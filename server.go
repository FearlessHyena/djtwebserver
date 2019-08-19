package main

import (
	"github.com/fearlesshyena/djtwebserver/httpd/handler"
	"github.com/fearlesshyena/djtwebserver/platform/owner"
	"github.com/gin-gonic/gin"
	"log"
)

func main() {
	router := gin.Default()

	repo, err := owner.New()
	if err != nil {
		log.Fatalln("Error while connecting to the repository", err)
	} else {
		router.GET("/", handler.PingGet())
		router.GET("/ping", handler.PingGet())
		router.GET("/contracts/:contract/:tokenid/owner", handler.OwnerGet(repo))

		log.Fatalln(router.Run(":8080"))
	}
}
