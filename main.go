package main

import (
	"log"
	"net/http"

	"app/controllers"
	"app/util"
)

func main() {

	db := util.CreateTestDb()
	defer db.Close()

	r := controllers.CreateRouter()

	log.Fatal(http.ListenAndServe(":3001", r))

}
