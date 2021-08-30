package storage

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func Find(tableName string, start int, limit int) (results []D, err error) {
	SetTable(tableName)

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	options := options.Find()
	options.SetSkip(int64(start))
	options.SetLimit(int64(limit))

	cur, err := collections[tableName].Find(ctx, bson.D{}, options)

	if err != nil {
		log.Fatal(err)
	}

	defer cur.Close(ctx)

	for cur.Next(ctx) {
		var data primitive.D
		err = cur.Decode(&data)

		if err != nil {
			log.Fatal(err)
		}

		var result D

		for _, value := range data {
			var pair E = E{
				Key:   value.Key,
				Value: value.Value,
			}

			result = append(result, pair)
		}

		results = append(results, result)
	}

	if err := cur.Err(); err != nil {
		log.Fatal(err)
	}

	return
}
