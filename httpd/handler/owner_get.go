package handler

import (
	"github.com/fearlesshyena/djtwebserver/platform"
	"github.com/fearlesshyena/djtwebserver/platform/request"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)


func OwnerGet(getter request.GetTokenOwner) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var contractreq platform.ContractToken
		if err := ctx.ShouldBindUri(&contractreq); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"msg": err})
			log.Fatalln(err)
		} else if !getter.TokenExists(contractreq) {
			ctx.Status(http.StatusNotFound)
		} else {
			if  tokenOwner, err := getter.GetOwner(contractreq); err != nil {
				ctx.JSON(http.StatusBadRequest, gin.H{"msg": err})
				log.Fatalln(err)
			} else {
				ctx.JSON(http.StatusOK, tokenOwner)
			}
		}
	}
}