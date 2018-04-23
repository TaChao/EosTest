package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"Gcoin/dao/chain"
)

func GetInfoCtrl(c *gin.Context) {
	//var dat interface{}
	//jsoniter.Unmarshal([]byte("{\"tt\":123,\"eee\":222}"), &dat)

	res, err := chain.GetInfo()
	if err != nil {
		panic(err)
		c.JSON(http.StatusInternalServerError, gin.H{"message": err})
	} else {
		c.Data(http.StatusOK, "application/json", res)
	}
}

func GetAccountCtrl(c *gin.Context) {
	name := c.Param("name")
	res, err := chain.GetAccount(name)
	if err != nil {
		panic(err)
		c.JSON(http.StatusInternalServerError, gin.H{"message": err})
	} else {
		c.Data(http.StatusOK, "application/json", res)
	}
}

func GetCodeCtrl(c *gin.Context) {
	name := c.Param("name")
	res, err := chain.GetCode(name)
	if err != nil {
		panic(err)
		c.JSON(http.StatusInternalServerError, gin.H{"message": err})
	} else {
		c.Data(http.StatusOK, "application/json", res)
	}
}
func GetTableRowsCtrl(c *gin.Context) {
	code := c.Param("code")
	table := c.Param("table")
	scope := c.Param("scope")
	res, err := chain.GetTableRows(scope,code,table)
	if err != nil {
		panic(err)
		c.JSON(http.StatusInternalServerError, gin.H{"message": err})
	} else {
		c.Data(http.StatusOK, "application/json", res)
	}
}

func JsonToBin(c *gin.Context)  {
	var b chain.JsonToBin
	c.BindJSON(&b)
	res, err := chain.AbiJsonToBin(&b)
	if err != nil {
		panic(err)
		c.JSON(http.StatusInternalServerError, gin.H{"message": err})
	} else {
		c.Data(http.StatusOK, "application/json", res)
	}
}