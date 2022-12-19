package client

import (
	"fmt"
	"github.com/nats-io/nats.go"
	"github.com/tuxoo/smart-loader/loader-service/internal/config"
)

//var e = nats.Connect()

//nc, err := nats.Connect("nats://host.docker.internal:4222")
//if err != nil {
//	logrus.Fatalf("error initializing nats: %s", err.Error())
//}

//err = nc.Publish("foo", []byte("Hello World"))

type NatsClient struct {
	url     string
	options []nats.Option
	Conn    *nats.Conn
}

func NewNatsClient(cfg *config.NatsConfig) *NatsClient {
	return &NatsClient{
		url:     fmt.Sprintf("%s:%s", cfg.URL, cfg.Port),
		options: nil,
	}
}

func (c *NatsClient) Connect() error {
	if connect, err := nats.Connect(c.url, c.options...); err != nil {
		return err
	} else {
		c.Conn = connect
	}

	//ch := make(chan *nats.Msg, 64)
	//_, err := c.conn.ChanSubscribe("foo", ch)
	//if err != nil {
	//	return err
	//}
	//
	//msg := <-ch
	//
	//fmt.Println(string(msg.Data))

	return nil
}

func (c *NatsClient) Disconnect() {
	c.Conn.Close()
}
