package main

import (
	"log"

	"github.com/nayan9800/practiceGo/app"
	//to import the local package use module name with path of package
	//<module name>/<path of package folder>
)

func init() {

	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
}
func main() {
	app.Run()
}
