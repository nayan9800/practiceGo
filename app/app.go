package app

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
)

var (
	netLog = log.New(os.Stdout, "NETLOG: ", log.LstdFlags|log.Lshortfile)
)

func Run() {
	log.Println("Running app")
	db, err := NewSqliteDb(DATABASE)
	if err != nil {
		netLog.Fatal(err.Error())
	}
	route := gin.Default()
	route.GET("/status", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status": "up",
		})
	})
	route.GET("/cars/:carmake", func(c *gin.Context) {
		carmake := c.Params.ByName("carmake")
		cars := db.GetCarsByCarMake(carmake)
		c.JSON(200, cars)
	})

	route.GET("/car/:id", func(c *gin.Context) {
		id := c.Params.ByName("id")
		car := db.GetCarbyId(id)
		c.JSON(200, car)
	})

	if err := route.Run(":8080"); err != nil {
		netLog.Fatal(err.Error())
	}

}
