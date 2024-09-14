package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/wagslane/go-rabbitmq"
	"github.com/wangxin688/narvis/client/config"
)

func main() {
	config.SetupConfig()

	conn, err := rabbitmq.NewConn(
		config.Settings.AMQP_URL,
		rabbitmq.WithConnectionOptionsLogging,
	)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	consumer, err := rabbitmq.NewConsumer(
		conn,
		config.Settings.ORGANIZATION_ID,
	)
	if err != nil {
		log.Fatal(err)
	}
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		fmt.Println(" [*] Waiting for logs. To exit press CTRL+C")
		sig := <-sigs
		fmt.Println()
		fmt.Println(" [x] Received signal:", sig)
		consumer.Close()
		os.Exit(0)
	}()

	// block main thread - wait for shutdown signal
	err = consumer.Run(func(delivery rabbitmq.Delivery) rabbitmq.Action {
		log.Printf(" [x] %s", delivery.Body)
		return rabbitmq.Ack
	})
	if err != nil {
		log.Fatal(err)
	}
}
