package main

import (
	"github.com/gin-gonic/gin"
	"github.com/rwcoding/goback"
	_ "github.com/rwcoding/goback/pkg/router"
	"github.com/rwcoding/mrng/center"
	"github.com/rwcoding/mrng/config"
	_ "github.com/rwcoding/mrng/internal"
	"github.com/rwcoding/mrng/models"
	"github.com/rwcoding/mrng/services"
	"log"
)

func main() {

	goback.SetLang(config.GetLang())
	goback.SetDev(config.IsDev())
	goback.SetLogFile(config.GetLog())
	goback.SetOnlyGP(config.IsOnlyGP())
	goback.SetWriteHeader(config.IsWriteHeader())
	goback.SetDb(models.InitDB())

	services.InitRedis()

	gin.SetMode(config.GetMode())
	if goback.App().Console.HasCommand("cc") || config.IsOnlyCc() {
		//仅仅启动配置中心api
		e := gin.Default()
		e.POST("/cc", center.Handle)
		err := e.Run(config.GetAddr())
		if err != nil {
			log.Fatal(err)
		}
	} else {
		services.SyncTimer()
		e := gin.Default()
		e.POST("/api", func(context *gin.Context) {
			goback.Run(context)
		})
		e.NoRoute(staticHandler)
		err := e.Run(config.GetAddr())
		if err != nil {
			log.Fatal(err)
		}
	}
}