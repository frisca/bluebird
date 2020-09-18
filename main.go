package main

import (
	_ "bluebird/routers"
	"github.com/astaxie/beego"
	"github.com/gin-gonic/gin"
	routers "bluebird/routers"
)

func main() {
	runMode := beego.AppConfig.DefaultString("gin.mode", "debug")
	serverPort := beego.AppConfig.DefaultString("httpport", "8080")

	gin.SetMode(runMode)
	routersInit := routers.InitRouter()
	routersInit.Run(":" + serverPort)
}

