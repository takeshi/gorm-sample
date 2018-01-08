package model

import "testing"
import "github.com/jinzhu/gorm"
import "encoding/json"
import _ "github.com/mattn/go-sqlite3"
import _ "github.com/jinzhu/gorm/dialects/sqlite"

func TestPet(t *testing.T) {
	db, err := gorm.Open("sqlite3", "test.db")

	if err != nil {
		panic("fail to open db")
	}

	defer db.Close()

	db.LogMode(true)

	db.AutoMigrate(&Cat{})
	db.AutoMigrate(&Dog{})
	db.AutoMigrate(&Toy{})

	db.Delete(&Cat{}, true)
	db.Delete(&Toy{}, true)

	db.Save(&Cat{
		// Named{
		Named: Named{Name: "Tama", LastName: "多摩"},
		// },
		Toy: Toy{
			Name: "猫じゃらし",
		},
	})
	var cat Cat
	db.Preload("Toy").First(&cat)

	b, _ := json.Marshal(&cat)

	t.Logf("json %s", string(b))
	t.Logf("named %s", cat.name())

}
