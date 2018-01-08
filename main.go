package main

import (
	"log"
	"net/http"

	"github.com/takeshi/gorm-sample/controllers"
	"github.com/takeshi/gorm-sample/util"
)

func main() {

	db := util.CreateTestDb()
	defer db.Close()

	r := controllers.CreateRouter()

	log.Fatal(http.ListenAndServe(":3001", r))

}
