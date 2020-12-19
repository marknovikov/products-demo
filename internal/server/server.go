package server

import (
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"

	"github.com/urfave/cli"
	"google.golang.org/grpc"

	"github.com/marknovikov/products-demo/internal/config"
	"github.com/marknovikov/products-demo/internal/products"
	"github.com/marknovikov/products-demo/pkg/productspb"
)

func Server(c *cli.Context) error {
	cfg := config.New(c)

	storageConfig := products.StorageConfig{
		Host:         cfg.MongoHost,
		Port:         cfg.MongoPort,
		User:         cfg.MongoUser,
		Password:     cfg.MongoPassword,
		Database:     cfg.MongoDatabase,
		ConnTimeout:  cfg.MongoConnTimeout,
		QueryTimeout: cfg.MongoQueryTimeout,
	}

	mongoConn, closeMongo, err := products.NewMongoConn(storageConfig)
	if err != nil {
		return fmt.Errorf("server: %w", err)
	}
	defer closeMongo()

	storage, err := products.NewMongoStorage(mongoConn, storageConfig)
	if err != nil {
		log.Fatal(err)
	}

	httpCli := products.NewClient(products.ClientConfig{
		HttpTimeout: cfg.HTTPTimeout,
	})

	productsSvc := products.NewService(httpCli, storage)
	productsGrpcServer := products.NewGrpcServer(productsSvc)

	lis, err := net.Listen("tcp", ":8080")
	if err != nil {
		return fmt.Errorf("server: %w", err)
	}
	var opts []grpc.ServerOption

	// TODO FILL GRPC SERVER OPTIONS

	grpcServer := grpc.NewServer(opts...)
	productspb.RegisterProductsServer(grpcServer, productsGrpcServer)

	signalCh := make(chan os.Signal, 1)
	signal.Notify(signalCh, os.Interrupt)
	go func() {
		<-signalCh
		grpcServer.GracefulStop()
	}()

	fmt.Printf("grpc server is starting on port %d...\n", cfg.AppPort)
	if err := grpcServer.Serve(lis); err != nil {
		return fmt.Errorf("server: %w", err)
	}

	return nil
}
