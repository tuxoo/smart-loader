package client

import (
	"fmt"
	"github.com/nats-io/nats.go"
	"github.com/tuxoo/smart-loader/loader-service/internal/domain/model/config"
)

type NatsClient struct {
	url     string
	options []nats.Option
	Conn    *nats.Conn
}

func NewNatsClient(cfg *config.NatsConfig) *NatsClient {
	return &NatsClient{
		url:     fmt.Sprintf("%s:%s", cfg.HOST, cfg.Port),
		options: nil,
	}
}

func (c *NatsClient) Connect() error {
	if connect, err := nats.Connect(c.url, c.options...); err != nil {
		return err
	} else {
		c.Conn = connect
	}

	return nil
}

func (c *NatsClient) Disconnect() {
	c.Conn.Close()
}
