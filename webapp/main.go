package main

import (
	"fmt"
	"log"
	"net/http"
	"webapp/src/routes"
)

func main() {
	fmt.Println("Running webapp!")

	r := routes.Generate()
	log.Fatal(http.ListenAndServe(":3000", r))

}
