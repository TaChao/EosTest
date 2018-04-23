package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"Gcoin/dao/wallet"
)

func GetPublicKeysCtrl(c *gin.Context) {

	res, err := wallet.GetPublickKeys()
	if err != nil {
		panic(err)
		c.JSON(http.StatusInternalServerError, gin.H{"message": err})
	} else {
		c.Data(http.StatusOK, "application/json", res)
	}
}
