package rmq

import (
	"sync"

	"github.com/wagslane/go-rabbitmq"
	"github.com/wangxin688/narvis/intend/logger"
	"github.com/wangxin688/narvis/server/config"
	"go.uber.org/zap"
)

var once sync.Once
var oncePublish sync.Once
var conn *rabbitmq.Conn
var publisher *rabbitmq.Publisher

func GetMqConn() *rabbitmq.Conn {
	once.Do(func() {
		var err error
		conn, err = rabbitmq.NewConn(config.Settings.RabbitMQ.ProxyUrl)
		if err != nil {
			logger.Logger.Fatal("[rabbit-mq]: failed to connect rabbitmq", zap.Error(err))
		}
	})

	return conn
}

func GetPublisher() (*rabbitmq.Publisher, error) {
	var err error

	oncePublish.Do(func() {
		publisher, err = rabbitmq.NewPublisher(
			GetMqConn(),
			rabbitmq.WithPublisherOptionsLogging,
		)
	})
	return publisher, err
}

func PublishProxyMessage(message []byte, orgId string) error {
	var err error
	if publisher == nil {
		publisher, err = GetPublisher()
		if err != nil {
			return err
		}
	}
	// defer publisher.Close()
	err = publisher.Publish(
		message,
		[]string{orgId},
		// rabbitmq.WithPublishOptionsContentType("application/json"),
		rabbitmq.WithPublishOptionsExpiration("180000"), // 180s
	)
	return err
}
