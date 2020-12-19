package config

import (
	"time"

	"github.com/urfave/cli"
)

type Config struct {
	AppName           string
	AppPort           int
	HTTPTimeout       time.Duration
	MongoHost         string
	MongoPort         int
	MongoUser         string
	MongoPassword     string
	MongoDatabase     string
	MongoConnTimeout  time.Duration
	MongoQueryTimeout time.Duration
}

func New(c *cli.Context) Config {
	return Config{
		AppName:           c.String("appName"),
		AppPort:           c.Int("appPort"),
		HTTPTimeout:       c.Duration("httpTimeout"),
		MongoHost:         c.String("mongoHost"),
		MongoPort:         c.Int("mongoPort"),
		MongoUser:         c.String("mongoUser"),
		MongoPassword:     c.String("mongoPassword"),
		MongoDatabase:     c.String("mongoDatabase"),
		MongoConnTimeout:  c.Duration("mongoConnTimeout"),
		MongoQueryTimeout: c.Duration("mongoQueryTimeout"),
	}
}
