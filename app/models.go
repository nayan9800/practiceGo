package app

import (
	"encoding/json"
	"io"
	"log"
	"os"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

/*car model*/
type Car struct {
	gorm.Model
	CarMake      string `json:"car_make"`
	CarModel     string `json:"car_model"`
	CarModelYear int32  `json:"car_model_year"`
	Description  string `json:"description"`
}

/*car json model*/
type CarJson struct {
	ID           uint   `json:"id"`
	CarMake      string `json:"car_make"`
	CarModel     string `json:"car_model"`
	CarModelYear int32  `json:"car_model_year"`
	Description  string `json:"description"`
}

/*converts car model to carJson model*/
func (car Car) ToCarJson() CarJson {

	return CarJson{
		ID:           car.Model.ID,
		CarMake:      car.CarMake,
		CarModel:     car.CarModel,
		CarModelYear: car.CarModelYear,
		Description:  car.Description}

}

/*converts slice of car to slice of carjson*/
func carsTocarJson(cars []Car) (carJsons []CarJson) {
	for _, c := range cars {
		carJsons = append(carJsons, c.ToCarJson())
	}
	return
}

/*setup to create sqlite database*/
func SetupSqlite() {

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

/*loads test data form MOCK_DATA.json to slice car*/
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
