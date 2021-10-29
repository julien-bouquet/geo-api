package infrastructure

import (
	"context"
	"fmt"
	"os"

	"github.com/geo-api/interfaces"
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
	clientOptions := options.Client().ApplyURI(
		fmt.Sprintf("mongodb://%s:%s@db:%s", os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_PORT")),
	)
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		return nil, err
	}
	err = client.Ping(ctx, nil)
	if err != nil {
		return nil, err
	}
	collection := client.Database("geo-app-db").Collection("point")
	noSqlHandler.Collection = collection

	return noSqlHandler, nil
}

func (noSQLHander *Collection) Get(args map[string]interface{}) (interfaces.Documents, error) {
	filter := bson.D{{}}
	if args != nil {
		for key, value := range args {
			filter = append(filter, bson.E{Key: key, Value: value})
		}
	}

	cursor, err := noSQLHander.Collection.Find(ctx, filter)

	if err != nil {
		return nil, err
	}
	row := &Cursor{}
	row.Results = cursor
	return row, nil
}

func (noSQLHander *Collection) Add(args interface{}) error {
	_, err := noSQLHander.Collection.InsertOne(ctx, args)
	if err != nil {
		return err
	}
	return nil
}

func (noSQLHander *Collection) Delete(args interface{}) error {
	_, err := noSQLHander.Collection.DeleteOne(ctx, args)
	if err != nil {
		return err
	}
	return nil
}

func (noSQLHander *Collection) Update(filterValue string, args interface{}) error {
	filter := bson.D{{"name", filterValue}}
	update := bson.D{{"$set", args}}

	_ = noSQLHander.Collection.FindOneAndUpdate(ctx, filter, update)
	return nil
}

func (r *Cursor) Read(args interface{}) error {
	return r.Results.Decode(args)
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
