package products

import (
	"context"
	goErrors "errors"

	"github.com/marknovikov/products-demo/internal/errors"
	"github.com/marknovikov/products-demo/pkg/productspb"
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
	if req != nil {
		if req.Paging != nil {
			opts = append(opts, Options().WithPaging(Paging{
				Offset: req.Paging.Offset,
				Limit:  req.Paging.Limit,
			}))
		}

		if req.Sorting != nil {
			sorting, err := Options().WithSorting(Sorting{
				Ascending: req.Sorting.Ascending,
				SortBy:    req.Sorting.SortBy,
			})
			if err != nil {
				return resp, status.Errorf(codes.InvalidArgument, "List: %w", err)
			}

			opts = append(opts, sorting)
		}
	}

	pp, err := srv.s.List(ctx, opts...)

	var invalidInput errors.ErrInvalidInput
	if goErrors.As(err, &invalidInput) {
		return resp, status.Errorf(codes.InvalidArgument, "List: %w", err)
	}
	if err != nil {
		return resp, status.Errorf(codes.Internal, "List: %w", err)
	}

	resp.Products = make([]*productspb.ListResponse_Product, 0, len(pp))

	resp.Products = make([]*productspb.ListResponse_Product, len(pp))
	for i, p := range pp {
		resp.Products[i] = &productspb.ListResponse_Product{
			Id:               p.ID,
			Name:             p.Name,
			Price:            p.Price.StringFixed(2),
			PriceUpdateCount: p.PriceUpdateCount,
			LastModified:     timestamppb.New(p.LastModified),
		}
	}

	return resp, nil
}
