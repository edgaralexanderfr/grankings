package storage

import (
	"context"
	"time"
)

func Insert(tableName string, records ...interface{}) (err error) {
	SetTable(tableName)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if len(records) == 1 {
		_, err = collections[tableName].InsertOne(ctx, records[0])
	} else {
		_, err = collections[tableName].InsertMany(ctx, records)
	}

	return
}
