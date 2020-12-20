package products

import (
	"context"
	goErrors "errors"
	"fmt"
	"reflect"
	"strings"
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

func (p mongoProduct) get(field string) (v interface{}, err error) {
	field = strings.Title(field)

	val := reflect.ValueOf(&p).Elem()
	fieldVal := val.FieldByName(field)
	if !fieldVal.IsValid() {
		return nil, fmt.Errorf("get: product has no field: %s", field)
	}

	return fieldVal.Interface(), nil
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
		{"$set", bson.D{{"lastModified", primitive.NewDateTimeFromTime(time.Now().UTC())}}},
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

func mongoSorting(sorting Sorting) bson.D {
	sortOrder := 1
	if !sorting.Ascending {
		sortOrder = -1
	}

	return bson.D{{sorting.SortBy, sortOrder}, {"_id", sortOrder}}
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

	ctx, cancel = context.WithTimeout(context.Background(), cfg.ConnTimeout)
	defer cancel()

	coll := cli.Database("products").Collection("products")
	nameUniqueIdx := mongo.IndexModel{
		Keys:    bson.D{{"name", 1}},
		Options: options.Index().SetUnique(true),
	}

	if _, err := coll.Indexes().CreateOne(ctx, nameUniqueIdx); err != nil {
		_ = closeMongoCli(cli, cfg.ConnTimeout)

		return nil, nil, fmt.Errorf("NewMongoConn: Index: %w", err)
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
			SetFilter(p.updateFilter()).
			SetUpdate(p.updateQuery()).
			SetUpsert(true)
	}

	opts := options.BulkWrite().
		SetOrdered(false)

	_, err := coll.BulkWrite(ctx, writeModel, opts)
	if err != nil && !isErrDuplicateKey(err) {
		return fmt.Errorf("UpdateProducts: %w", err)
	}

	return nil
}

func isErrDuplicateKey(err error) bool {
	var e mongo.BulkWriteException

	if !goErrors.As(err, &e) {
		return false
	}

	for _, we := range e.WriteErrors {
		if we.Code == 11000 {
			return true
		}
	}
	return false
}

func mongoFindFilterOpts(opts ...option) (filter bson.D, mongoOpts *options.FindOptions, err error) {
	filter = bson.D{}
	mongoOpts = options.Find()

	optsHolder := applyOptions(opts)
	if optsHolder.sorting != nil {
		mongoOpts.SetSort(mongoSorting(*optsHolder.sorting))
	}

	if optsHolder.paging != nil {
		mongoOpts.SetLimit(int64(optsHolder.paging.Limit))
	}

	seekPageFilter, err := resolveSeekPageFilter(optsHolder)
	if err != nil {
		return nil, nil, fmt.Errorf("mongoFindFilterOpts: %w", err)
	}
	if seekPageFilter != nil {
		filter = append(filter, *seekPageFilter)
	}

	return filter, mongoOpts, nil
}

func resolveSeekPageFilter(opts *optsHolder) (*bson.E, error) {
	if opts.paging == nil || opts.paging.Last == nil {
		return nil, nil
	}

	if opts.paging.Last != nil && opts.sorting != nil {
		return optsToSeekPageFilterStrategy[seekPageFilterStrategyWithSorting](opts)
	}

	if opts.paging.Last != nil {
		return optsToSeekPageFilterStrategy[seekPageFilterStrategyNextPage](opts)
	}

	return nil, fmt.Errorf("seekPageFilter: no possible paging strategy for provided paging and sorting params")
}

type seekPageFilterStrategy func(opts *optsHolder) (*bson.E, error)

const (
	seekPageFilterStrategyNextPage    = "nextPage"
	seekPageFilterStrategyWithSorting = "withSorting"
)

var optsToSeekPageFilterStrategy = map[string]seekPageFilterStrategy{
	seekPageFilterStrategyNextPage: func(opts *optsHolder) (*bson.E, error) {
		id, err := primitive.ObjectIDFromHex(opts.paging.Last.ID)
		if err != nil && err != primitive.ErrInvalidHex {
			return nil, fmt.Errorf("seekPageFilterStrategyNextPage: %w", err)
		}

		return &bson.E{"_id", bson.D{{"$gt", id}}}, nil
	},
	seekPageFilterStrategyWithSorting: func(opts *optsHolder) (*bson.E, error) {
		sortByCond := "$gte"
		idCond := "$gt"

		if !opts.sorting.Ascending {
			sortByCond = "$lte"
			idCond = "$lt"
		}

		last, err := newMongoProduct(*opts.paging.Last)
		if err != nil {
			return nil, fmt.Errorf("seekPageFilterStrategyWithSorting: %w", err)
		}

		lastSortBy, err := last.get(opts.sorting.SortBy)
		if err != nil {
			return nil, fmt.Errorf("seekPageFilterStrategyWithSorting: %w", err)
		}

		lastID := last.ID

		return &bson.E{
			"$and", bson.A{
				bson.D{{
					opts.sorting.SortBy,
					bson.D{{sortByCond, lastSortBy}},
				}},
				bson.D{{
					"$or", bson.A{
						bson.D{{
							opts.sorting.SortBy,
							bson.D{{"$not", bson.D{{"$eq", lastSortBy}}}},
						}},
						bson.D{{
							"_id",
							bson.D{{idCond, lastID}},
						}},
					},
				}},
			},
		}, nil
	},
}

func (s *mongodb) FindProducts(ctx context.Context, opts ...option) ([]Product, error) {
	coll := s.cli.Database(s.cfg.Database).Collection("products")

	filter, mongoOpts, err := mongoFindFilterOpts(opts...)
	if err != nil {
		return nil, fmt.Errorf("FindProducts: %w", err)
	}

	curs, err := coll.Find(ctx, filter, mongoOpts)
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
