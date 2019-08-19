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

	if repo, err := etherium.Connect("https://ropsten.infura.io"); err != nil {
		log.Fatalln("Error while connecting to the repository", err)
	} else {
		router.GET("/", handler.PingGet())
		router.GET("/ping", handler.PingGet())
		token := request.Token{Repo: repo}
		router.GET("/contracts/:contract/:tokenid/owner", handler.OwnerGet(&token))
		router.GET("/contracts/:contract/:tokenid/transfers", handler.TransferGet(&token))

		log.Fatalln(router.Run(":8080"))
	}
}
