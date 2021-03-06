package controllers

import "net/http"

import "github.com/takeshi/gorm-sample/util"
import "github.com/takeshi/gorm-sample/model"
import "encoding/json"

// var fileName string

func init() {
	// fileName = util.GetCurrentFile()
}

func mainControler(w http.ResponseWriter, r *http.Request) {
	db := util.GetDb()

	var market model.Market
	db.First(&market)

	b, err := json.Marshal(&market)
	if err != nil {
		w.Write([]byte("unexpected Error"))
	}

	w.Write(b)

}
