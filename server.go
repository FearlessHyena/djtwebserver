package main

import (
	"github.com/fearlesshyena/djtwebserver/httpd/handler"
	"github.com/fearlesshyena/djtwebserver/platform/etherium"
	"github.com/fearlesshyena/djtwebserver/platform/request"
	"github.com/gin-gonic/gin"
	"log"
)

func main() {
	router := gin.Default()

	repo, err := etherium.Connect("https://ropsten.infura.io")
	if err != nil {
		log.Fatalln("Error while connecting to the repository", err)
	} else {
		router.GET("/", handler.PingGet())
		router.GET("/ping", handler.PingGet())
		router.GET("/contracts/:contract/:tokenid/owner", handler.OwnerGet(&request.TokenOwner{Repo: repo}))

		log.Fatalln(router.Run(":8080"))
	}
}
