package main

import (
	"log"
	"net/http"

	"app/controllers"
)

func main() {
	r := controllers.CreateRouter()

	log.Fatal(http.ListenAndServe(":3001", r))

}
