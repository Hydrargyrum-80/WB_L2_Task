package main

import (
	"dev11/routers"
	"dev11/serverControl"
	"fmt"
	"net/http"
)

func main() {
	router := routers.NewEventControl()
	serverControlUnit := serverControl.NewController(router)
	err := http.ListenAndServe(":8080", serverControlUnit.GetRouter())
	if err != nil {
		fmt.Println(err.Error())
	}
}
