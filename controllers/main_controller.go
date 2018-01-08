package controllers

import "net/http"
import "fmt"
import "app/util"

func init() {
	fmt.Println(util.GetCurrentFile())
}

func mainControler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Sample Text"))
}
