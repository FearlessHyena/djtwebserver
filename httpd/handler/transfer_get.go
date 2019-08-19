package handler

import (
	"github.com/fearlesshyena/djtwebserver/platform"
	"github.com/fearlesshyena/djtwebserver/platform/repo"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)


func TransferGet(getter repo.GetTokenTransfers) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var contractreq platform.ContractToken
		if err := ctx.ShouldBindUri(&contractreq); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"msg": err})
			log.Fatalln(err)
		} else if !getter.TokenExists(contractreq) {
			ctx.Status(http.StatusNotFound)
		} else {
			if  transfers, err := getter.GetTransfers(contractreq); err != nil {
				ctx.JSON(http.StatusBadRequest, gin.H{"msg": err})
				log.Fatalln(err)
			} else {
				ctx.JSON(http.StatusOK, transfers)
			}
		}
	}
}