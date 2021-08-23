package storage

import (
	"fmt"
	"os"
)

func Finish() {
	if err := client.Disconnect(connectionCtx); err != nil {
		panic(err)
	}

	fmt.Println("Disconnected from " + os.Getenv("MONGO_DB_CONNECTION_URI"))
}
