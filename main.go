package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/edgaralexanderfr/grankings/pkg/router"
	"github.com/edgaralexanderfr/grankings/pkg/storage"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatal("Couldn't open the .env file.")
	}

	appAddr := os.Getenv("APP_ADDR")
	r := router.CreateRoutes()

	storage.Init()
	defer storage.Finish()

	fmt.Println("App running at " + appAddr + ".")
	http.ListenAndServe(appAddr, r)
}
