package database

import (
	"context"
	"fmt"
	"time"

	"ylanzinhoy-operator-management/config"

	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
	"go.mongodb.org/mongo-driver/v2/mongo/readpref"
)

type DatabaseAdapter struct {
	cfg    config.Config
	client *mongo.Client
}

func NewDatabaseConnectionAdapter(cfg config.Config) *DatabaseAdapter {
	return &DatabaseAdapter{
		cfg: cfg,
	}
}

func (d *DatabaseAdapter) Connect(ctx context.Context) error {
	uri := d.cfg.Database.Uri
	if uri == "" {
		uri = "mongodb://localhost:27017"
	}

	opts := options.Client().ApplyURI(uri)

	opts.SetAuth(options.Credential{
		Username:   "root",
		Password:   "foo",
		AuthSource: "admin",
	})

	client, err := mongo.Connect(opts)
	if err != nil {
		return fmt.Errorf("falha ao inicializar client mongo: %w", err)
	}

	pingCtx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	if err := client.Ping(pingCtx, readpref.Primary()); err != nil {
		return fmt.Errorf("falha ao pingar o banco de dados: %w", err)
	}

	d.client = client
	return nil
}

func (d *DatabaseAdapter) Close(ctx context.Context) error {
	if d.client == nil {
		return nil
	}
	if err := d.client.Disconnect(ctx); err != nil {
		return fmt.Errorf("erro ao desconectar do banco: %w", err)
	}
	return nil
}

func (d *DatabaseAdapter) CreateCollection(ctx context.Context, dbName, collectionName string) error {
	err := d.client.Database(dbName).CreateCollection(ctx, collectionName)
	if err != nil {
		return err
	}

	return nil
}

func (d *DatabaseAdapter) GetCollection(dbName, collectionName string) *mongo.Collection {
	if d.client == nil {
		return nil
	}
	return d.client.Database(dbName).Collection(collectionName)
}
