package main

import (
	"log"
	"os"

	"github.com/urfave/cli"

	"github.com/marknovikov/products-demo/internal/server"
)

func main() {
	app := &cli.App{
		Name:   "products-demo",
		Action: server.Server,
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:   "appName",
				EnvVar: "APP_NAME",
			},
			&cli.IntFlag{
				Name:   "appPort",
				EnvVar: "APP_PORT",
			},
			&cli.DurationFlag{
				Name:   "httpTimeout",
				EnvVar: "HTTP_TIMEOUT",
			},
			&cli.StringFlag{
				Name:   "mongoHost",
				EnvVar: "MONGO_HOST",
			},
			&cli.IntFlag{
				Name:   "mongoPort",
				EnvVar: "MONGO_PORT",
			},
			&cli.StringFlag{
				Name:   "mongoUser",
				EnvVar: "MONGO_USER",
			},
			&cli.StringFlag{
				Name:   "mongoPassword",
				EnvVar: "MONGO_PASSWORD",
			},
			&cli.StringFlag{
				Name:   "mongoDatabase",
				EnvVar: "MONGO_DATABASE",
			},
			&cli.DurationFlag{
				Name:   "mongoConnTimeout",
				EnvVar: "MONGO_CONN_TIMEOUT",
			},
			&cli.DurationFlag{
				Name:   "mongoQueryTimeout",
				EnvVar: "MONGO_QUERY_TIMEOUT",
			},
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
