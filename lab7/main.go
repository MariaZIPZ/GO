package main

import (
	"fmt"
	"lab7/packages/handlers"
	"lab7/packages/routes"
	"net/http"
)

func main() {
	run()
}

func run() {
	mux := http.NewServeMux()

	mux.HandleFunc(routes.QuadraticEquationRoute, handlers.QuadraticEquation)
	mux.HandleFunc(routes.SliceRoute, handlers.Slice)

	mux.HandleFunc("/", handlers.Index)

	fmt.Println("Server started on " + routes.ServerUrl + routes.SeverPort)

	err := http.ListenAndServe(routes.SeverPort, mux)
	if err != nil {
		fmt.Println("An error occurred while starting the server", "\n", err.Error())
		return
	}
}
