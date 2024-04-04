package main

import (
	"fmt"
	"github.com/go-chi/chi"
	"github.com/kx0101/api-demo/internal/handlers"
	log "github.com/sirupsen/logrus"
	"net/http"
)

func main() {
	log.SetReportCaller(true)
	r := chi.NewRouter()

	handlers.Handler(r)

	fmt.Println("Starting GO API service...")
	fmt.Println("Listening on PORT 8080")

	err := http.ListenAndServe("localhost:8080", r)

	if err != nil {
		log.Error(err)
	}
}
