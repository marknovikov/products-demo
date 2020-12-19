package products

import (
	"context"
	"fmt"
	"time"

	"github.com/shopspring/decimal"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Storage interface {
	UpdateProducts(ctx context.Context, pp []Product) error
	FindProducts(ctx context.Context, opts ...option) ([]Product, error)
}

type StorageConfig struct {
	Host            string
	Port            int
	User            string
	Password        string
	Database        string
	ConnTimeout     time.Duration
	QueryTimeout    time.Duration
	MaxConns        uint32
	IdleConnTimeout time.Duration
}

func (cfg StorageConfig) connString() string {
	return fmt.Sprintf("mongodb://%s:%s@%s:%d", cfg.User, cfg.Password, cfg.Host, cfg.Port)
}

type mongodb struct {
	cli *mongo.Client
	cfg StorageConfig
}

type mongoProduct struct {
	ID               primitive.ObjectID   `bson:"_id,omitempty"`
	Name             string               `bson:"name,omitempty"`
	Price            primitive.Decimal128 `bson:"price,omitempty"`
	PriceUpdateCount uint32               `bson:"priceUpdateCount,omitempty"`
	LastModified     time.Time            `bson:"lastModified,omitempty"`
}

func newMongoProduct(p Product) (mongoProduct, error) {
	id, err := primitive.ObjectIDFromHex(p.ID)
	if err != nil && err != primitive.ErrInvalidHex {
		return mongoProduct{}, fmt.Errorf("newMongoProduct: %w", err)
	}
	if err == primitive.ErrInvalidHex {
		id = [12]byte{}
	}

	price, err := primitive.ParseDecimal128(p.Price.StringFixed(2))
	if err != nil {
		return mongoProduct{}, fmt.Errorf("newMongoProduct: %w", err)
	}

	return mongoProduct{
		ID:               id,
		Name:             p.Name,
		Price:            price,
		PriceUpdateCount: p.PriceUpdateCount,
		LastModified:     p.LastModified,
	}, nil
}

func (p mongoProduct) updateFilter() bson.D {
	return bson.D{{"$and", bson.A{
		bson.D{{"name", p.Name}},
		bson.D{{"price", bson.D{{"$not", bson.D{{"$eq", p.Price}}}}}}}}}
}

func (p mongoProduct) updateQuery() bson.D {
	return bson.D{
		{"$setOnInsert", bson.D{{"name", p.Name}}},
		{"$set", bson.D{{"price", p.Price}}},
		{"$inc", bson.D{{"priceUpdateCount", 1}}},
		{"$currentDate", bson.D{
			{"lastModified", bson.D{{"$type", "timestamp"}}},
		}},
	}
}

func (p mongoProduct) toProduct() (Product, error) {
	price, err := decimal.NewFromString(p.Price.String())
	if err != nil {
		return Product{}, fmt.Errorf("toProduct: %w", err)
	}

	return Product{
		ID:               p.ID.Hex(),
		Name:             p.Name,
		Price:            price,
		PriceUpdateCount: p.PriceUpdateCount,
		LastModified:     p.LastModified,
	}, nil
}

func mongoSorting(sorting Sorting) (bson.D, error) {
	sortOrder := 1
	if !sorting.Ascending {
		sortOrder = -1
	}

	return bson.D{{sorting.SortBy, sortOrder}}, nil
}

func NewMongoConn(cfg StorageConfig) (cli *mongo.Client, close func() error, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), cfg.ConnTimeout)
	defer cancel()

	cli, err = mongo.Connect(ctx, options.Client().ApplyURI(cfg.connString()))
	if err != nil {
		return nil, nil, fmt.Errorf("NewMongoConn: Connect: %w", err)
	}

	ctx, cancel = context.WithTimeout(context.Background(), cfg.ConnTimeout)
	defer cancel()

	if err := cli.Ping(ctx, nil); err != nil {
		_ = closeMongoCli(cli, cfg.ConnTimeout)

		return nil, nil, fmt.Errorf("NewMongoConn: Ping: %w", err)
	}

	closer := func() error {
		return closeMongoCli(cli, cfg.ConnTimeout)
	}

	return cli, closer, nil
}

func NewMongoStorage(cli *mongo.Client, cfg StorageConfig) (Storage, error) {
	return &mongodb{
		cli: cli,
		cfg: cfg,
	}, nil
}

func (s *mongodb) UpdateProducts(ctx context.Context, pp []Product) error {
	coll := s.cli.Database(s.cfg.Database).Collection("products")

	writeModel := make([]mongo.WriteModel, len(pp))
	for i := range pp {
		p, err := newMongoProduct(pp[i])
		if err != nil {
			return fmt.Errorf("UpdateProducts %w", err)
		}

		writeModel[i] = mongo.NewUpdateOneModel().
			SetUpsert(true).
			SetFilter(p.updateFilter()).
			SetUpdate(p.updateQuery())
	}

	_, err := coll.BulkWrite(ctx, writeModel)
	if err != nil {
		return fmt.Errorf("UpdateProducts: %w", err)
	}

	return nil
}

func (s *mongodb) FindProducts(ctx context.Context, opts ...option) ([]Product, error) {
	coll := s.cli.Database(s.cfg.Database).Collection("products")

	mongoOpts := options.Find()
	optsHolder := applyOptions(opts)
	if optsHolder.sorting != nil {
		mSorting, err := mongoSorting(*optsHolder.sorting)
		if err != nil {
			return nil, fmt.Errorf("FindProducts: %w", err)
		}
		mongoOpts.SetSort(mSorting)
	}

	if optsHolder.paging != nil {
		mongoOpts.SetSkip(int64(optsHolder.paging.Offset))
		mongoOpts.SetLimit(int64(optsHolder.paging.Limit))
	}

	curs, err := coll.Find(ctx, bson.M{}, mongoOpts)
	if err != nil {
		return nil, fmt.Errorf("FindProducts: %w", err)
	}
	defer curs.Close(ctx)

	var pp []Product

	for curs.Next(ctx) {
		var p mongoProduct
		if err := curs.Decode(&p); err != nil {
			return pp, fmt.Errorf("FindProducts: %w", err)
		}

		product, err := p.toProduct()
		if err != nil {
			return pp, fmt.Errorf("FindProducts: %w", err)
		}
		pp = append(pp, product)
	}
	if curs.Err() != nil {
		return pp, fmt.Errorf("FindProducts: %w", err)
	}
	return pp, nil
}

func closeMongoCli(cli *mongo.Client, connTimeout time.Duration) error {
	ctx, cancel := context.WithTimeout(context.Background(), connTimeout)
	defer cancel()

	return cli.Disconnect(ctx)
}
