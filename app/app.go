package app

import (
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var (
	netLog = log.New(os.Stdout, "NETLOG: ", log.LstdFlags|log.Lshortfile)
)

func Run() {
	log.Println("Running app")
	/*Create databsase connection*/
	db, err := NewSqliteDb(DATABASE)
	if err != nil {
		netLog.Fatal(err.Error())
	}

	/*gin router*/
	route := gin.Default()

	/*gives the status of api*/
	route.GET("api/v1/status", func(c *gin.Context) {
		c.JSON(http.StatusAccepted, gin.H{
			"status": "up",
		})
	})

	/*gives cars by carmake*/
	route.GET("api/v1/cars/:carmake", func(c *gin.Context) {
		carmake := c.Params.ByName("carmake")
		cars := db.GetCarsByCarMake(carmake)
		c.JSON(http.StatusAccepted, carsTocarJson(cars))
	})

	/*gives car by carmodel*/
	route.GET("api/v1/car/:carmodel", func(c *gin.Context) {
		carmodel := c.Params.ByName("carmodel")
		car, err := db.GetCarByCarModel(carmodel)
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{})
			return
		}
		c.JSON(http.StatusAccepted, car.ToCarJson())
	})

	/*gives car by car id*/
	route.GET("api/v1/car/get/:id", func(c *gin.Context) {
		id := c.Params.ByName("id")
		car, err := db.GetCarbyId(id)
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{})
			return
		}
		c.JSON(http.StatusAccepted, car.ToCarJson())

	})

	/*post request add new car*/
	route.POST("api/v1/car/add", func(c *gin.Context) {
		var newCar Car
		if err := c.BindJSON(&newCar); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		db.AddNewCar(newCar)
		c.JSON(http.StatusAccepted, gin.H{})
	})

	/*path request to udate car by car id*/
	route.PATCH("api/v1/car/:id", func(c *gin.Context) {
		id := c.Params.ByName("id")
		var updatedCar Car
		if err := c.ShouldBindJSON(&updatedCar); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		if err := db.UpdateCarByID(id, updatedCar); err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": gorm.ErrRecordNotFound})
		} else if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		} else {
			c.JSON(http.StatusAccepted, gin.H{"id": id})
		}
	})

	/*delete car by car id*/
	route.DELETE("api/v1/car/delete/:id", func(c *gin.Context) {
		id := c.Params.ByName("id")
		if err := db.DeleteCarByID(id); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{})

	})

	/*run app at given port*/
	if err := route.Run(":8080"); err != nil {
		netLog.Fatal(err.Error())
	}

}
