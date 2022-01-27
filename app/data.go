package app

import (
	"log"
	"os"
	"time"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var (
	dbLog    = log.New(os.Stdout, "[DBLOG] ", log.LstdFlags)
	DATABASE = "./testdata/test.db"
	looger   = logger.New(
		dbLog,
		logger.Config{
			SlowThreshold:             time.Second,
			LogLevel:                  logger.Warn,
			Colorful:                  true,
			IgnoreRecordNotFoundError: false,
		},
	)
)

type dbConn struct {
	db *gorm.DB
}

func NewSqliteDb(dsn string) (dbConn, error) {
	db, err := gorm.Open(sqlite.Open(dsn), &gorm.Config{
		Logger: looger,
	})
	return dbConn{db: db}, err
}

func (dbc dbConn) GetCarsByCarMake(carmake string) []Car {
	cars := []Car{}
	dbc.
		db.
		//Select([]string{"ID", "car_make", "car_model", "car_model_year", "description"}).
		Where(map[string]interface{}{"car_make": carmake}).
		Find(&cars)
	return cars
}

func (dbc dbConn) GetCarbyId(id string) map[string]interface{} {
	car := map[string]interface{}{}
	dbc.db.Model(&Car{}).First(car, id)

	return car
}
