package main

import (
	"log"

	"github.com/nguyenbt456/blocklist/app"
)

func main() {
	router := app.InitRouter()

	if err := router.Run(":8000"); err != nil {
		log.Println("Fail!!!!!")
	}
}
