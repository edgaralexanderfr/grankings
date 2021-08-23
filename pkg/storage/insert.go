package storage

import (
	"context"
	"time"
)

func Insert(tableName string, records ...interface{}) {
	SetTable(tableName)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if len(records) == 1 {
		collections[tableName].InsertOne(ctx, records[0])
	} else {
		collections[tableName].InsertMany(ctx, records)
	}
}
