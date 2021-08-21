package storage

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	// "go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var (
	connectionCtx context.Context
	client        *mongo.Client
	collections   map[string]*mongo.Collection = make(map[string]*mongo.Collection)
)

func Init() {
	var ctx context.Context
	var cancel context.CancelFunc
	var err error

	connectionCtx, cancel = context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	uri := os.Getenv("MONGO_DB_CONNECTION_URI")
	fmt.Println("Connecting to " + uri + "...")
	client, err = mongo.Connect(ctx, options.Client().ApplyURI(uri).SetAuth(options.Credential{
		Username: os.Getenv("MONGO_DB_CONNECTION_USER"),
		Password: os.Getenv("MONGO_DB_CONNECTION_PASS"),
	}))

	if err != nil {
		panic(err)
	}

	ctx, cancel = context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	err = client.Ping(ctx, readpref.Primary())

	if err != nil {
		log.Fatal("Couldn't connect to " + uri)
	}

	// SetTable("games")
	// collections["games"].InsertOne(ctx, bson.D{{Key: "name", Value: "Snipe Intruders"}, {Key: "key", Value: "SECRET"}})
}

func SetTable(name string) {
	collections[name] = client.Database(os.Getenv("MONGO_DB_CONNECTION_DB")).Collection(name)
}

func Finish() {
	if err := client.Disconnect(connectionCtx); err != nil {
		panic(err)
	}

	fmt.Println("Disconnected from " + os.Getenv("MONGO_DB_CONNECTION_URI"))
}
