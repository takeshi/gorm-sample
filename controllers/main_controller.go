package controllers

import "net/http"

import "app/util"

var fileName string

func init() {
	fileName = util.GetCurrentFile()
}

func mainControler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Sample Text " + fileName))
}
