package main

import (
	"Gcoin/enum"
	"github.com/gin-gonic/gin"
	"Gcoin/controller"
	"net/http"
)

var (
	Passworld = "PW5JSb4Dsx22TbMLK2Pch9oj8p4At6GdLNRPDUDG77wsUq2ZuEm9x"
)

func main() {
	//NODEOS API URL
	enum.URL = "http://127.0.0.1:8888"

	r := gin.Default()

	//html files
	r.LoadHTMLGlob("templates/*")

	//other static files
	r.Static("/static", "static")

	//index
	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})

	//v1 group api
	v1 := r.Group("/v1")

	//chain
	v1.GET("/info", controller.GetInfoCtrl)
	v1.GET("/account/:name", controller.GetAccountCtrl)
	v1.GET("/code/:name", controller.GetCodeCtrl)
	v1.GET("/table/:code/:table/:scope", controller.GetTableRowsCtrl)

	v1.POST("/bin", controller.JsonToBin)

	//wallet
	v1.GET("/wallet/publicKeys", controller.GetPublicKeysCtrl)

	//custom
	v1.POST("/currency/bin", controller.CurrencyToBin)
	v1.GET("/issue/:code/:name/:q", controller.Issue)
	v1.GET("/transfer/:code/:name/:receiver/:q", controller.Transfer)

	//ajax cors
	//r.Use(cors.Default())

	r.Run() // listen and serve on 0.0.0.0:8080

	//params := service.CurrencyCmdParams{
	//	Action:     "issue",
	//	Code:       "awcoin",
	//	Permission: "awcoin",
	//	Context:    "{\"to\":\"awcoin\",\"quantity\":\"1000.0000 CUR\",\"memo\":\"first\"}",
	//}
	//out, err := service.CurrencyCmd(params)
	//fmt.Println(out, err)
	//fmt.Println(strings.Trim(string(ip), "\n"))
}
