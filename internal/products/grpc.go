package products

import (
	"context"
	goErrors "errors"
	"fmt"

	"github.com/marknovikov/products-demo/internal/errors"
	"github.com/marknovikov/products-demo/pkg/productspb"
	"github.com/shopspring/decimal"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type grpcServer struct {
	// embedding required by google.golang.org/protobuf
	productspb.UnimplementedProductsServer

	s Service
}

func NewGrpcServer(s Service) productspb.ProductsServer {
	return &grpcServer{
		s: s,
	}
}

func (srv *grpcServer) Fetch(ctx context.Context, req *productspb.FetchRequest) (*productspb.FetchResponse, error) {
	resp := &productspb.FetchResponse{}

	err := srv.s.Fetch(ctx, req.Url)

	var invalidInput errors.ErrInvalidInput
	if goErrors.As(err, &invalidInput) {
		return resp, status.Errorf(codes.InvalidArgument, "Fetch: %w", err)
	}
	if err != nil {
		return resp, status.Errorf(codes.Internal, "Fetch: %w", err)
	}

	return resp, nil
}

func (srv *grpcServer) List(ctx context.Context, req *productspb.ListRequest) (*productspb.ListResponse, error) {
	resp := &productspb.ListResponse{}

	var opts []option
	if err := applyPaging(&opts, req); err != nil {
		return resp, status.Errorf(codes.InvalidArgument, "List: %w", err)
	}
	if err := applySorting(&opts, req); err != nil {
		return resp, status.Errorf(codes.InvalidArgument, "List: %w", err)
	}

	pp, err := srv.s.List(ctx, opts...)

	var invalidInput errors.ErrInvalidInput
	if goErrors.As(err, &invalidInput) {
		return resp, status.Errorf(codes.InvalidArgument, "List: %w", err)
	}
	if err != nil {
		return resp, status.Errorf(codes.Internal, "List: %w", err)
	}

	resp.Products = make([]*productspb.Product, 0, len(pp))

	resp.Products = make([]*productspb.Product, len(pp))
	for i, p := range pp {
		resp.Products[i] = &productspb.Product{
			Id:               p.ID,
			Name:             p.Name,
			Price:            p.Price.StringFixed(2),
			PriceUpdateCount: p.PriceUpdateCount,
			LastModified:     timestamppb.New(p.LastModified),
		}
	}

	return resp, nil
}

func toProduct(pb *productspb.Product) (*Product, error) {
	if pb == nil {
		return nil, nil
	}

	p := &Product{
		ID:               pb.Id,
		Name:             pb.Name,
		PriceUpdateCount: pb.PriceUpdateCount,
	}

	if pb.Price != "" {
		price, err := decimal.NewFromString(pb.Price)
		if err != nil {
			return nil, fmt.Errorf("applyPaging: %w", err)
		}
		p.Price = price
	}

	if pb.LastModified.IsValid() {
		p.LastModified = pb.LastModified.AsTime()
	}

	return p, nil
}

func applyPaging(opts *[]option, req *productspb.ListRequest) error {
	if opts == nil {
		return fmt.Errorf("opts is nil")
	}

	if req == nil {
		return nil
	}

	if req.Paging == nil {
		return nil
	}

	paging := Paging{
		Limit: req.Paging.Limit,
	}

	last, err := toProduct(req.Paging.Last)
	if err != nil {
		return fmt.Errorf("applyPaging: %w", err)
	}
	paging.Last = last

	optsVal := *opts
	optsVal = append(optsVal, Options().WithPaging(paging))

	*opts = optsVal

	return nil
}

func applySorting(opts *[]option, req *productspb.ListRequest) error {
	if opts == nil {
		return fmt.Errorf("opts is nil")
	}

	if req == nil {
		return nil
	}

	if req.Sorting == nil {
		return nil
	}

	sorting, err := Options().WithSorting(Sorting{
		Ascending: req.Sorting.Ascending,
		SortBy:    req.Sorting.SortBy,
	})
	if err != nil {
		return fmt.Errorf("applySorting: %w", err)
	}

	optsVal := *opts
	optsVal = append(optsVal, sorting)

	*opts = optsVal

	return nil
}
