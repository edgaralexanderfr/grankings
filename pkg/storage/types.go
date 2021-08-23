package storage

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
)

var (
	connectionCtx context.Context
	client        *mongo.Client
	collections   map[string]*mongo.Collection = make(map[string]*mongo.Collection)
)
