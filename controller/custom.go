package controller

import (
	"github.com/gin-gonic/gin"
	"Gcoin/dao/chain"
	"net/http"
	"time"
	"Gcoin/service"
	"github.com/gin-gonic/gin/json"
)

func CurrencyToBin(c *gin.Context) {
	b := chain.JsonToBin{
		Action: "transfer",
		Code:   "currency",
	}
	c.BindJSON(&b.Args)
	res, err := chain.AbiJsonToBin(&b)
	if err != nil {
		panic(err)
		c.JSON(http.StatusInternalServerError, gin.H{"message": err})
	} else {
		c.Data(http.StatusOK, "application/json", res)
	}
}

func Issue(c *gin.Context) {
	now := time.Now().Format("2006-01-02T15:04:05")
	context := map[string]string{
		"to":       c.Param("name"),
		"quantity": c.Param("q") + ".0000 CUR",
		"memo":     now,
	}
	json, _ := json.Marshal(context)
	contextStr := string(json)

	params := service.CurrencyCmdParams{
		Action:     "issue",
		Code:       c.Param("code"),
		Context:    contextStr,
		Permission: c.Param("code"),
	}
	res, err := service.CurrencyCmd(params)
	if err != nil {
		//panic(err)
		c.JSON(http.StatusInternalServerError, gin.H{"message": err})
	} else {
		c.Data(http.StatusOK, "application/json", res)
	}
}
func Transfer(c *gin.Context) {
	now := time.Now().Format("2006-01-02T15:04:05")
	context := map[string]string{
		"to":       c.Param("receiver"),
		"from":       c.Param("name"),
		"quantity": c.Param("q") + ".0000 CUR",
		"memo":     now,
	}
	json, _ := json.Marshal(context)
	contextStr := string(json)

	params := service.CurrencyCmdParams{
		Action:     "transfer",
		Code:       c.Param("code"),
		Context:    contextStr,
		Permission: c.Param("name"),
	}
	res, err := service.CurrencyCmd(params)
	if err != nil {
		//panic(err)
		c.JSON(http.StatusInternalServerError, gin.H{"message": err})
	} else {
		c.Data(http.StatusOK, "application/json", res)
	}
}