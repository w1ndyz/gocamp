package data

import (
	"book/internal/conf"
	"book/internal/data/ent"
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"

	_ "github.com/go-sql-driver/mysql"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(NewData, NewGreeterRepo)

// Data .
type Data struct {
	// TODO warpped database client
	db *ent.Client
}

// NewData .
func NewData(c *conf.Data, logger log.Logger) (*Data, func(), error) {
	log := log.NewHelper("data", logger)
	client, err := ent.Open(
			c.Database.Driver,
			c.Database.Source,
		)
	if err != nil {
		log.Errorf("failed opening connection to mysql: %v", err)
		return nil, nil, err
	}
	// run migration
	if err := client.Schema.Create(context.Background()); err != nil {
		log.Errorf("failed creating schema resources: %v", err)
		return nil, nil, err
	}
	d := &Data{
		db: client,
	}
	cleanup := func() {
		logger.Log("msg", "closing the data resources")
	}
	return d, cleanup, nil
}
