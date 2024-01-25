package main

import (
	"fmt"
	"log"

	"around/handler"
	"net/http"
)

func main() {
	fmt.Println("started-service")
	log.Fatal(http.ListenAndServe(":8080", handler.InitRouter()))
}
