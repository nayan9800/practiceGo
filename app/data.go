package app

import (
	"log"
	"os"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var (
	dbLog = log.New(os.Stdout, "[DBLOG] ", log.LstdFlags)
	//Default database for sqlite
	DATABASE = "./testdata/test.db"
	//logger for gorm
	looger = logger.New(
		dbLog,
		logger.Config{
			SlowThreshold:             time.Second,
			LogLevel:                  logger.Warn,
			Colorful:                  true,
			IgnoreRecordNotFoundError: false,
		},
	)
)

/*dbConn wrapper for *gorm.db*/
type dbConn struct {
	db *gorm.DB
}

/*creates new sqlite connection with givemn dsn*/
func NewSqliteDb(dsn string) (dbConn, error) {
	db, err := gorm.Open(sqlite.Open(dsn), &gorm.Config{
		Logger: looger,
	})
	return dbConn{db: db}, err
}

/*creates new mysql connection with givemn dsn*/
func NewMysqlDb(dsn string) (dbConn, error) {
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: looger,
	})
	return dbConn{db: db}, err
}

/*get call cars matching with given carmake*/
func (dbc dbConn) GetCarsByCarMake(carmake string) []Car {
	cars := []Car{}
	dbc.
		db.
		//Select([]string{"ID", "car_make", "car_model", "car_model_year", "description"}).
		Where(map[string]interface{}{"car_make": carmake}).
		Find(&cars)
	return cars
}

/*get car by given id*/
func (dbc dbConn) GetCarbyId(id string) (Car, error) {
	car := Car{}
	err := dbc.db.First(&car, id).Error

	return car, err
}

/*get car by car model*/
func (dbc dbConn) GetCarByCarModel(carmodel string) (Car, error) {
	car := Car{}
	err := dbc.db.Where("car_model=?", carmodel).First(&car).Error
	return car, err
}

/*Creates new car with given car object*/
func (dbc dbConn) AddNewCar(newcar Car) {
	dbc.db.Create(&newcar)
}

/*Update car by given car id*/
func (dbc dbConn) UpdateCarByID(id string, UpdatedCar Car) error {
	car, err := dbc.GetCarbyId(id)
	if err != nil {
		return err
	}
	return dbc.db.Model(&car).Updates(&UpdatedCar).Error

}

/*delete car by given id*/
func (dbc dbConn) DeleteCarByID(id string) error {
	return dbc.db.Delete(&Car{}, id).Error
}
