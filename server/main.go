package main

import (
	"fmt"

	"github.com/getsentry/sentry-go"
	sentry_gin "github.com/getsentry/sentry-go/gin"
	"github.com/gin-gonic/gin"

	"github.com/wangxin688/narvis/intend/logger"
	"github.com/wangxin688/narvis/server/config"
	"github.com/wangxin688/narvis/server/middleware"

	// "github.com/wangxin688/narvis/server/pkg/rmq"
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
	config.InitConfig()
	config.InitLogger()
	initializeSentry()
	err := infra.InitDB()
	if err != nil {
		logger.Logger.Fatal("[mainStartHttpServer]: failed to initialize database", zap.Error(err))
		panic(err)
	}
	err = infra.InitClickHouseDB()
	if err != nil {
		logger.Logger.Fatal("[mainStartHttpServer]: failed to initialize clickhouse database", zap.Error(err))
		panic(err)
	}
	gen.SetDefault(infra.DB)
	logger.Logger.Info("[mainStartHttpServer]: server started to run on port", zap.Int("port", config.Settings.System.ServerPort))
	if config.Settings.Env == config.Prod || config.Settings.Env == config.OnPrem {
		gin.SetMode(gin.ReleaseMode)
	}
	router := gin.New()
	configureRouter(router)
	helpers.RegisterCustomValidator()
	middleware.RegisterOpenAPI(router)
	go register.RegisterScheduler()
	// rmq.GetMqConn()
	if err := router.Run(fmt.Sprintf(":%d", config.Settings.System.ServerPort)); err != nil {
		logger.Logger.Fatal("[mainStartHttpServer]: failed to run server", zap.Error(err))
	}
}

func initializeSentry() {
	if config.Settings.Env == config.Prod {
		if err := sentry.Init(sentry.ClientOptions{
			Dsn:              config.Settings.Sentry.Dsn,
			EnableTracing:    config.Settings.Sentry.EnableTracing,
			TracesSampleRate: config.Settings.Sentry.TraceSampleRate,
			Release:          config.Settings.Sentry.Release,
		}); err != nil {
			logger.Logger.Fatal("[mainStartHttpServer]: failed to initialize sentry", zap.Error(err))
		}
	} else {
		logger.Logger.Info("[mainStartHttpServer]: sentry disabled because of environment", zap.String("environment", string(config.Settings.Env)))
	}
}

func configureRouter(router *gin.Engine) {
	if config.Settings.Env == config.Prod {
		router.Use(sentry_gin.New(sentry_gin.Options{}))
	}
	router.Use(
		middleware.ZapLoggerMiddleware(logger.Logger),
		middleware.GinRecovery(logger.Logger, true),
	)
	router.Use(middleware.CorsMiddleware())
	register.RegisterRouter(router)
}
