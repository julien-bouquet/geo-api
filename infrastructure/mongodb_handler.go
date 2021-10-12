package infrastructure

import (
	"context"
	"fmt"
	"os"

	"github.com/coast-nav-api/interfaces"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Collection struct {
	Collection *mongo.Collection
}

type Cursor struct {
	Results *mongo.Cursor
}

type ResultInsert struct {
	Result *mongo.InsertOneResult
}

var ctx = context.TODO()

func NewMongoDBHandler() (interfaces.NoSQLHandler, error) {
	noSqlHandler := &Collection{}
	clientOptions := options.Client().ApplyURI(fmt.Sprintf("mongodb://db:%s", os.Getenv("DB_PORT")))
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		return nil, err
	}
	err = client.Ping(ctx, nil)
	if err != nil {
		return nil, err
	}
	collection := client.Database("point-app").Collection("point")
	noSqlHandler.Collection = collection

	return noSqlHandler, nil
}

func (noSQLHander *Collection) Get(args ...interface{}) (interfaces.Documents, error) {
	filter := bson.D{{}}
	cursor, err := noSQLHander.Collection.Find(ctx, filter)
	if err != nil {
		return nil, err
	}
	row := &Cursor{}
	row.Results = cursor
	return row, nil
}

func (noSQLHander *Collection) Add(args ...interface{}) error {
	_, err := noSQLHander.Collection.InsertOne(ctx, args)
	if err != nil {
		return err
	}
	return nil
}

func (r *Cursor) Read(args ...interface{}) error {
	return r.Results.Decode(ctx)
}

func (r *Cursor) Next(args ...interface{}) bool {
	return r.Results.Next(ctx)
}

func (r *Cursor) Close() error {
	return r.Results.Close(ctx)
}

func (r *Cursor) Err() error {
	return r.Results.Err()
}
