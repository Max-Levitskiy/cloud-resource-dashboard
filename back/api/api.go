package main

import (
	"fmt"
	"github.com/Max-Levitskiy/cloud-resource-dashboard/api/handlers"
	"net/http"
)

func main() {
	fmt.Println("qweqwe")
	http.HandleFunc("/resource", handlers.ResourceHandler)

	http.ListenAndServe(":8080", nil)
}
