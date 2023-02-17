package main

import (
	"github.com/sam80719/echoBot/router"
)

// @host localhost:8088
func main() {
	router := router.SetRouterPublic()
	err := router.Run(":8088")
	if err != nil {
		panic(err)
	}
}
