package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	_ "github.com/joho/godotenv/autoload"
)

var (
	port = os.Getenv("PORT")
)

func main() {
	r := chi.NewRouter()

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello World"))
	})

	fmt.Println("Listening on port " + port)
	_ = http.ListenAndServe(":"+port, r)
}
