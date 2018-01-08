package controllers

import "github.com/gorilla/mux"
import "fmt"
import "github.com/takeshi/gorm-sample/util"

func init() {
	fmt.Println(util.GetCurrentFile())
}

func CreateRouter() *mux.Router {
	r := mux.NewRouter()

	r.HandleFunc("/", mainControler).Methods("Get")

	return r
}
