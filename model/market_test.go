package model

import (
	"fmt"
	"os"
	"testing"

	"github.com/takeshi/gorm-sample/util"

	"github.com/jinzhu/gorm"
)

var db *gorm.DB

func TestMarket(t *testing.T) {
	// db := util.CreateTestDb()

	market := &Market{Name: "サンプル市場"}

	asset := market.addAsset("株")
	asset2 := market.addAsset("債権")

	db.Save(market)
	db.Save(asset)
	db.Save(asset2)

}

func TestMain(m *testing.M) {

	fmt.Println("Test Start")
	db = util.CreateTestDb()
	db.LogMode(true)
	db.Exec("Drop Table markets")
	db.Exec("Drop Table financial_assets")
	db.AutoMigrate(&Market{})
	db.AutoMigrate(&FinancialAsset{})
	var code int
	defer func() {
		fmt.Println("db close")
		db.Close()
		os.Exit(code)
	}()
	code = m.Run()

	fmt.Println("Test End")

}
