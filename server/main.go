package main

import (
	"net/http"
	"strings"

	"github.com/getsentry/sentry-go"
	sentry_gin "github.com/getsentry/sentry-go/gin"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/wangxin688/narvis/server/core"
	"github.com/wangxin688/narvis/server/core/config"
	"github.com/wangxin688/narvis/server/global"
	"github.com/wangxin688/narvis/server/infra"
	"github.com/wangxin688/narvis/server/middleware"
	"github.com/wangxin688/narvis/server/tools/helpers"
	"go.uber.org/zap"
)

func main() {
	setupConfig()
	setupLogger()
	initializeSentry()
	infra.InitDB()
	// dal.SetDefault(infra.DB)
	router := gin.New()
	configureRouter(router)
	helpers.RegisterCustomValidator()
	// user, err := dal.Q.User.WithContext(context.Background()).Find()
	// if err != nil {
	// 	core.Logger.Fatal("Failed to get user", zap.Error(err))
	// }
	// core.Logger.Info("User: ", zap.Any("user", user))

	for _, route := range router.Routes() {
		if strings.HasPrefix(route.Path, "/api") {
			core.Logger.Info("route: " + route.Path)
			core.Logger.Info("method: " + route.Method)
		}
	}

	if err := router.Run(":8080"); err != nil {
		core.Logger.Fatal("Failed to run server", zap.Error(err))
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
			core.Logger.Fatal("Failed to initialize Sentry", zap.Error(err))
		}
	} else {
		core.Logger.Info("Sentry disabled because of environment", zap.String("environment", string(core.Environment)))
	}
}

func configureRouter(router *gin.Engine) {
	if core.Environment == config.Prod || core.Environment == config.Stage {
		router.Use(sentry_gin.New(sentry_gin.Options{}))
	}
	router.Use(middleware.ZapLoggerMiddleware(core.Logger), middleware.GinRecovery(core.Logger, true))
	router.Use(middleware.CORSByConfig())
	router.GET("/api/v1/health/:id", healthHandler)
	router.GET("/api/v1/test", testHandler)
}

func healthHandler(c *gin.Context) {
	global.OrganizationID.Set(uuid.New().String())
	core.Logger.Info("organization id: " + global.OrganizationID.Get())
	c.Set("id", c.Param("id"))
	core.Logger.Info("path: " + c.FullPath()) // full path test
	testThreadLocal()
	c.JSON(http.StatusOK, gin.H{
		"message": "ok",
	})
}

func testHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "ok",
	})
}

func testThreadLocal() {
	orgID := global.OrganizationID.Get()

	core.Logger.Info("organization id: " + orgID)
}
