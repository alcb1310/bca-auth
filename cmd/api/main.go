package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/alcb1310/bca-auth/internal/server"
	_ "github.com/joho/godotenv/autoload"
)

var (
	port = os.Getenv("PORT")
)

func main() {
	server := server.NewServer()

	if err := http.ListenAndServe(fmt.Sprintf(":%s", port), server.Router); err != nil {
		panic(err)
	}
}
