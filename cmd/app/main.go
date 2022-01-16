package main

import (
	"blog/internal/config"
	"blog/internal/router"
	"blog/internal/store"
	"fmt"
	"net/http"
	"os"

	"github.com/joeshaw/envdecode"
	"github.com/joho/godotenv"
	//"github.com/joho/godotenv"
)

var TestVar int

func init() {
	err := godotenv.Load(".env")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to load env: %s", err)
	}
}

func main() {

	var config config.Config
	if err := envdecode.Decode(&config); err != nil {
		fmt.Fprintf(os.Stderr, "Failed to decode %s", err)
	}

	store.NewConnection(config.SqlStore.DatabaseURL)

	http.ListenAndServe(":333", router.ConfigureRoutes(routes))

}
