package main

import (
	"github.com/fearlesshyena/djtwebserver/httpd/handler"
	"github.com/fearlesshyena/djtwebserver/platform/cache"
	"github.com/fearlesshyena/djtwebserver/platform/etherium"
	"github.com/fearlesshyena/djtwebserver/platform/repo"
	"github.com/gin-gonic/gin"
	"log"
)

const Serveraddr = ":8080"
const Ethnetwork = "https://ropsten.infura.io"
const Dbconninfo = "user=postgres dbname=djthash sslmode=disable"

func main() {
	router := gin.Default()

	if eth, err := etherium.Connect(Ethnetwork); err != nil {
		log.Fatalln("Error while connecting to the repository", err)
	} else if pgcache, err := cache.NewPgClient(Dbconninfo);
		err != nil {
		log.Fatalln("Error while connecting to the cache db", err)
	} else {
		token := repo.Token{Repo: eth, Cache: pgcache}

		router.GET("/", handler.PingGet())
		router.GET("/contracts/:contract/:tokenid/owner", handler.OwnerGet(&token))
		router.GET("/contracts/:contract/:tokenid/transfers", handler.TransferGet(&token))

		log.Fatalln(router.Run(Serveraddr))
	}
}
