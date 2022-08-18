package main

import (
	"api/src/config"
	"api/src/router"
	"fmt"
	"log"
	"net/http"
)

type Person struct {
	Name string
	Age  int
}

func main() {
	config.Charge()
	r := router.Generate()

	fmt.Println("Listening port: 5000")
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", config.Port), r))
}
