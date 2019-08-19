package handler

import (
	"github.com/fearlesshyena/djtwebserver/platform/request"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)


func OwnerGet(getter request.GetOwner) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var contractreq request.ContractReq
		if err := ctx.ShouldBindUri(&contractreq); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"msg": err})
			log.Fatalln(err)
		} else {
			if  tokenOwner, err := getter.Get(contractreq); err != nil {
				ctx.JSON(http.StatusBadRequest, gin.H{"msg": err})
				log.Fatalln(err)
			} else {
				ctx.JSON(http.StatusOK, tokenOwner)
			}
		}
	}
}