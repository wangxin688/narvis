package main

import (
	"github.com/getsentry/sentry-go"
	sentry_gin "github.com/getsentry/sentry-go/gin"
	"github.com/gin-gonic/gin"

	"github.com/wangxin688/narvis/server/core"
	"github.com/wangxin688/narvis/server/core/config"
	"github.com/wangxin688/narvis/server/middleware"
	"github.com/wangxin688/narvis/server/register"
	"github.com/wangxin688/narvis/server/tools/helpers"

	"github.com/wangxin688/narvis/server/dal/gen"
	"github.com/wangxin688/narvis/server/infra"

	"go.uber.org/zap"
)

// @title						Narvis API
// @version 					1.0
// @description 				Narvis OpenAPI for developers.
// @termsOfService 				http://swagger.io/terms/
// @contact.name 				Jeffry
// @securityDefinitions.apikey 	BearerAuth
// @in header
// @name Authorization
func main() {
	setupConfig()
	setupLogger()
	initializeSentry()
	infra.InitDB()
	gen.SetDefault(infra.DB)
	router := gin.New()
	configureRouter(router)
	helpers.RegisterCustomValidator()
	middleware.RegisterOpenAPI(router)
	if err := router.Run(":8080"); err != nil {
		core.Logger.Fatal("[mainStartHttpServer]: failed to run server", zap.Error(err))
	}
}

func setupConfig() {
	core.SetUpConfig()
}

func setupLogger() {
	core.SetUpLogger()
}

func initializeSentry() {
	if core.Environment == config.Prod || core.Environment == config.Stage {
		if err := sentry.Init(sentry.ClientOptions{
			Dsn:              core.Settings.Sentry.Dsn,
			EnableTracing:    core.Settings.Sentry.EnableTracing,
			TracesSampleRate: core.Settings.Sentry.TraceSampleRate,
			Release:          core.Settings.Sentry.Release,
		}); err != nil {
			core.Logger.Fatal("[mainStartHttpServer]: failed to initialize sentry", zap.Error(err))
		}
	} else {
		core.Logger.Info("[mainStartHttpServer]: sentry disabled because of environment", zap.String("environment", string(core.Environment)))
	}
}

func configureRouter(router *gin.Engine) {
	if core.Environment == config.Prod || core.Environment == config.Stage {
		router.Use(sentry_gin.New(sentry_gin.Options{}))
	}
	router.Use(
		middleware.ZapLoggerMiddleware(core.Logger),
		middleware.GinRecovery(core.Logger, true),
	)
	router.Use(middleware.CORSByConfig())
	register.RegisterRouter(router)
}
