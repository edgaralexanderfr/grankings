package storage

import (
	"os"
)

func SetTable(name string) {
	_, ok := collections[name]

	if !ok {
		collections[name] = client.Database(os.Getenv("MONGO_DB_CONNECTION_DB")).Collection(name)
	}
}
