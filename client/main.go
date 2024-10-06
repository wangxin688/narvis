package main

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/wagslane/go-rabbitmq"
	"github.com/wangxin688/narvis/client/config"
	"github.com/wangxin688/narvis/client/tasks"
	"github.com/wangxin688/narvis/client/utils/helpers"
	"github.com/wangxin688/narvis/client/utils/logger"
	"go.uber.org/zap"
)

func main() {
	if err := config.SetupConfig(); err != nil {
		logger.Logger.Error("[setupConfig]: failed to setup config ", zap.Error(err))
		os.Exit(1)
	}
	conn, err := rabbitmq.NewConn(config.Settings.AMQP_URL, rabbitmq.WithConnectionOptionsLogging)
	if err != nil {
		logger.Logger.Error("[proxyConsumer]: failed to create rabbitmq connection ", zap.Error(err))
		os.Exit(1)
	}
	defer conn.Close()

	consumer, err := rabbitmq.NewConsumer(conn, config.Settings.ORGANIZATION_ID)
	if err != nil {
		logger.Logger.Error("[proxyConsumer]: failed to create rabbitmq consumer ", zap.Error(err))
		os.Exit(1)
	}

	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		sig := <-sigs
		logger.Logger.Info("[proxyConsumer]: Received signal:", zap.Any("signal", sig))
		consumer.Close()
		os.Exit(0)
	}()

	// block main thread - wait for shutdown signal
	if err := consumer.Run(func(delivery rabbitmq.Delivery) rabbitmq.Action {
		logger.Logger.Info("[proxyConsumer]: Received message:", zap.String("message", string(delivery.Body)))
		helpers.BackgroundTask(func() {
			tasks.TaskDispatcher(delivery.Body)
		})
		return rabbitmq.Ack
	}); err != nil {
		logger.Logger.Error("[proxyConsumer]: failed to run consumer ", zap.Error(err))
	}
}
