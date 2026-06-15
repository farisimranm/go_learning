package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi"
	log "github.com/sirupsen/logrus" // import library as "log"

	// our custom packages
	"tutorial_11/internal/handlers"
)

func main() {
	// enable log to file
	log.SetReportCaller(true)

	var r *chi.Mux = chi.NewRouter()
	handlers.Handler(r)

	fmt.Println("Starting GO API service...")
	fmt.Println(`
  ________  _________      _____ __________.___ 
 /  _____/ /         \    /  _  \\______   \   |
/   \  ___ |    |    |   /  /_\  \|     ___/   |
\    \_\  \|    |    |  /    |    \    |   |   |
 \______  /\_________/  \____|__  /____|   |___|
        \/                     \/               `)
	fmt.Println("\nGO API service started!\n")

	err := http.ListenAndServe("localhost:8000", r)
	if err != nil {
		log.Error(err)
	}
}
