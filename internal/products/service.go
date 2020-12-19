package products

import (
	"context"
	"fmt"
	"net/url"

	"github.com/marknovikov/products-demo/internal/errors"
)

type Service interface {
	Fetch(ctx context.Context, path string) error
	List(ctx context.Context, opts ...option) ([]Product, error)
}

type service struct {
	client  Client
	storage Storage
}

func NewService(client Client, storage Storage) Service {
	return &service{
		client:  client,
		storage: storage,
	}
}

func (s *service) Fetch(ctx context.Context, path string) error {
	_, err := url.Parse(path)
	if err != nil {
		return errors.NewErrInvalidInput(fmt.Errorf("Fetch: %w", err))
	}

	pp, err := s.client.List(ctx, path)
	if err != nil {
		return fmt.Errorf("Fetch: %w", err)
	}

	err = s.storage.UpdateProducts(ctx, pp)
	if err != nil {
		return fmt.Errorf("Fetch: %w", err)
	}

	return nil
}

func (s *service) List(ctx context.Context, opts ...option) ([]Product, error) {
	pp, err := s.storage.FindProducts(ctx, opts...)
	if err != nil {
		return pp, fmt.Errorf("List: %w", err)
	}

	return pp, nil
}
