package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/wagslane/go-rabbitmq"
	"github.com/wangxin688/narvis/client/config"
	"github.com/wangxin688/narvis/client/tasks"
	"github.com/wangxin688/narvis/client/utils/helpers"
	"github.com/wangxin688/narvis/client/utils/logger"
)

func main() {
	config.SetupConfig()

	conn, err := rabbitmq.NewConn(
		config.Settings.AMQP_URL,
		rabbitmq.WithConnectionOptionsLogging,
	)
	if err != nil {
		logger.Logger.Error(err)
	}
	defer conn.Close()

	consumer, err := rabbitmq.NewConsumer(
		conn,
		config.Settings.ORGANIZATION_ID,
	)
	if err != nil {
		logger.Logger.Error(err)
	}
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		logger.Logger.Info(" [startConsumer] Waiting for logs. To exit press CTRL+C")
		sig := <-sigs
		logger.Logger.Info()
		logger.Logger.Info(" [consumer] Received signal:", sig)
		consumer.Close()
		os.Exit(0)
	}()

	// block main thread - wait for shutdown signal
	err = consumer.Run(func(delivery rabbitmq.Delivery) rabbitmq.Action {
		logger.Logger.Info(fmt.Sprintf(" [consumer] received new message %s", delivery.Body))
		helpers.BackgroundTask(func() {
			tasks.TaskDispatcher(delivery.Body)
		})
		return rabbitmq.Ack
	})
	if err != nil {
		logger.Logger.Error(err)
	}
}
