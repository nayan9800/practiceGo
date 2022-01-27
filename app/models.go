package app

import (
	"encoding/json"
	"io"
	"log"
	"os"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Car struct {
	gorm.Model
	CarMake      string `json:"car_make"`
	CarModel     string `json:"car_model"`
	CarModelYear int32  `json:"car_model_year"`
	Description  string `json:"description"`
}

func init() {

	_, err := os.Stat(DATABASE)
	if os.IsNotExist(err) {
		dbLog.Println("creating database")
		db, err := gorm.Open(sqlite.Open(DATABASE), &gorm.Config{})
		if err != nil {
			dbLog.Fatal(err.Error())
		}
		cars, err := loadMockData()
		if err != nil {
			dbLog.Println(err.Error())
			return
		}
		if err := db.AutoMigrate(&Car{}); err != nil {
			log.Println(err.Error())
			return
		}
		dbLog.Printf("inserting %d records", len(cars))
		db.CreateInBatches(cars, 50)
	} else {
		dbLog.Println(DATABASE, " already exists")
	}
}
func loadMockData() ([]Car, error) {
	carsData := []Car{}
	f, err := os.Open("./app/tmp/MOCK_DATA.json")
	if err != nil {
		return carsData, err
	}
	fdata, err := io.ReadAll(f)
	if err != nil {
		return carsData, err
	}
	err = json.Unmarshal(fdata, &carsData)
	if err != nil {
		return []Car{}, err
	}

	return carsData, nil
}
