package handler

import (
	"github.com/fearlesshyena/djtwebserver/platform/owner"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)


func OwnerGet(getter owner.GetOwner) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var contractreq owner.ContractRequest
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