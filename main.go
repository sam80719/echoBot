package main

import (
	"github.com/sam80719/echoBot/router"
	"log"
)

// @host localhost:8088
func main() {
	router := router.SetRouterPublic()
	err := router.Run(":8080")
	if err != nil {
		log.Fatalf("%+v", err)
	}
}
