package storage

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
)

type D []E

type E struct {
	Key   string
	Value interface{}
}

var (
	connectionCtx context.Context
	client        *mongo.Client
	collections   map[string]*mongo.Collection = make(map[string]*mongo.Collection)
)
