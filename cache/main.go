package main

import (
	"cache/controller"
	"fmt"
	"net/http"
	_ "cache/scheduler"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

const (
	port string = "5556"
)

// Expose endpoitns for the cache
func main() {
	start()
}

func start() {
	fmt.Println("Starting caching service...")
	router := mux.NewRouter()
	// Stores the search values
	// Check if a search has been done before
	router.HandleFunc("/cache/v1/getSearch", controller.GetSearch).Methods("POST")
	c := cors.New(cors.Options{
		AllowedOrigins: []string{
			"http://localhost:3000",
			"http://127.0.0.1:3000",
			"http://localhost:5555",
			"http://127.0.0.1:5555",
		},
		AllowCredentials: true,
	})
	handler := c.Handler(router)
	fmt.Printf("Cache routes exposed on port %s\n", port)
	err := http.ListenAndServe(":"+port, handler) //Launch the app, visit localhost:5555/api
	if err != nil {
		fmt.Print(err)
	}
}
