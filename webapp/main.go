package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	fmt.Println("Running webapp!")

	r := routes.Generate()
	log.Fatal(http.ListenAndServe(":3000", r))

}
