package main

import (
	"flag"
	"github.com/fearlesshyena/djtwebserver/httpd/handler"
	"github.com/fearlesshyena/djtwebserver/platform/cache"
	"github.com/fearlesshyena/djtwebserver/platform/etherium"
	"github.com/fearlesshyena/djtwebserver/platform/repo"
	"github.com/gin-gonic/gin"
	"log"
)

const DefServeraddr = ":8080"
const DefEthnetwork = "https://ropsten.infura.io"
const DefDbconninfo = "user=postgres dbname=djthash sslmode=disable"

func main() {
	var serveraddr string
	var ethnetwork string
	var dbconninfo string

	flag.StringVar(&serveraddr, "addr", DefServeraddr, "The address to run the webserver at")
	flag.StringVar(&ethnetwork, "eth", DefEthnetwork, "The Etherium network address")
	flag.StringVar(&dbconninfo, "dbconn", DefDbconninfo, "The DB connection options")
	flag.Parse()

	router := gin.Default()

	if eth, err := etherium.Connect(ethnetwork); err != nil {
		log.Fatalln("Error while connecting to the repository", err)
	} else if pgcache, err := cache.NewPgClient(dbconninfo);
		err != nil {
		log.Fatalln("Error while connecting to the cache db", err)
	} else {
		token := repo.Token{Repo: eth, Cache: pgcache}

		router.GET("/", handler.PingGet())
		router.GET("/contracts/:contract/:tokenid/owner", handler.OwnerGet(&token))
		router.GET("/contracts/:contract/:tokenid/transfers", handler.TransferGet(&token))

		log.Fatalln(router.Run(serveraddr))
	}
}
