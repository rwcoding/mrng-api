package main

import (
	"github.com/gin-gonic/gin"
	"github.com/rwcoding/goback"
	"github.com/rwcoding/mrng/center"
	"github.com/rwcoding/mrng/config"
	"github.com/rwcoding/mrng/models"
	"github.com/rwcoding/mrng/services"
	"log"

	_ "github.com/rwcoding/goback/pkg/router"
	_ "github.com/rwcoding/mrng/internal"
)

func main() {

	goback.SetLang(config.GetLang())
	goback.SetDev(config.IsDev())
	goback.SetLogFile(config.GetLog())
	goback.SetOnlyGP(config.IsOnlyGP())
	goback.SetWriteHeader(config.IsWriteHeader())
	goback.SetDb(models.GetDB())

	services.InitRedis()

	gin.SetMode(config.GetMode())
	e := gin.Default()
	e.POST("/cc", center.Handle)

	if !goback.App().Console.HasCommand("cc") && !config.IsOnlyCc() {
		services.SyncTimer()
		e.POST("/api", func(context *gin.Context) {
			goback.Run(context)
		})
		e.OPTIONS("/api", func(context *gin.Context) {
			goback.Run(context)
		})
		e.NoRoute(staticHandler)
	}

	err := e.Run(config.GetAddr())
	if err != nil {
		log.Fatal(err)
	}
}
