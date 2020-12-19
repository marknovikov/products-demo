package products

import (
	"context"
	"encoding/csv"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/shopspring/decimal"
)

type Client interface {
	List(ctx context.Context, url string) ([]Product, error)
}

type ClientConfig struct {
	HttpTimeout time.Duration
}

type httpClient struct {
	cli *http.Client
	cfg ClientConfig
}

func NewClient(cfg ClientConfig) Client {
	cli := &http.Client{
		Timeout: cfg.HttpTimeout,
	}

	return &httpClient{
		cli,
		cfg,
	}
}

func csvRowToProduct(row []string) (Product, error) {
	if len(row) < 2 {
		return Product{}, fmt.Errorf("csvRowToProduct: row: %v, expected exactly 2 cols per row, got: %d", row, len(row))
	}

	price, err := decimal.NewFromString(row[1])
	if err != nil {
		return Product{}, fmt.Errorf("csvRowToProduct: %w", err)
	}

	return Product{
		Name:  row[0],
		Price: price,
	}, nil
}

func (c *httpClient) List(ctx context.Context, url string) ([]Product, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, fmt.Errorf("List: %w", err)
	}
	req.Header.Set("Accept", "text/csv; charset=utf-8")
	req.Header.Set("Connection", "Keep-Alive")

	resp, err := c.cli.Do(req)
	if err != nil {
		return nil, fmt.Errorf("List: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("List: 200 http status expected, got: %s", resp.Status)
	}

	r := csv.NewReader(resp.Body)
	r.Comma = ';'

	// skip head
	_, _ = r.Read()

	var pp []Product
	for {
		row, err := r.Read()
		if err == io.EOF {
			return pp, nil
		}
		if err != nil {
			return pp, fmt.Errorf("List: %w", err)
		}

		p, err := csvRowToProduct(row)
		if err != nil {
			return pp, fmt.Errorf("List: %w", err)
		}

		pp = append(pp, p)
	}
}
