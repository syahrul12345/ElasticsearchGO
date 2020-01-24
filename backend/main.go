package main

import (
	"backend/controller"
	"fmt"
	"net/http"
)

const (
	port string = "5555"
)

func main() {
	start()
}

// Start the webserver
func start() {
	fmt.Println("Starting webserver...")
	// router := mux.NewRouter()
	// //API routes

	// router.HandleFunc("/api/v1/createAccount", controller.CreateAccount).Methods("POST")
	// router.HandleFunc("/api/v1/changePassword", controller.ChangePassword).Methods("POST")
	// router.HandleFunc("/api/v1/login", controller.Login).Methods("POST")
	// router.HandleFunc("/api/v1/search", controller.Search).Methods("POST")

	// // Routes to serve the webpage
	// router.PathPrefix("/").Handler(http.HandlerFunc(controller.Serve))

	// // Use the JWT middle ware
	// router.Use(auth.JwtAuthentication)
	// // Set the cors
	// c := cors.New(cors.Options{
	// 	AllowedOrigins:   []string{"http://localhost:3000", "http://127.0.0.1:3000"},
	// 	AllowCredentials: true,
	// })
	// handler := c.Handler(router)
	// fmt.Printf("Serving frontend on http://127.0.0.1:%s\n", port)
	// fmt.Printf("Api end routes exposed on port %s\n", port)
	// err := http.ListenAndServe(":"+port, handler) //Launch the app, visit localhost:5555/api
	// if err != nil {
	// 	fmt.Print(err)
	// }

	fmt.Printf("Api end routes exposed on port %s\n", port)
	fmt.Printf("Serving frontend on http://127.0.0.1:%s\n", port)
	http.HandleFunc("/api/v1/createAccount", controller.CreateAccount)
	http.HandleFunc("/api/v1/changePassword", controller.ChangePassword)
	http.HandleFunc("/api/v1/login", controller.Login)
	http.HandleFunc("/api/v1/search", controller.SearchAuth)
	http.HandleFunc("/", controller.Serve)
	http.ListenAndServe(":"+port, nil)
}
