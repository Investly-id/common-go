package message_broker

import (
	"time"

	nsqGo "github.com/nsqio/go-nsq"
)

type nsq struct {
	producer     *nsqGo.Producer
	consumerList []*nsqGo.Consumer
}

type consumerHandler struct {
	handler func(m *Message) error
}

func (c *consumerHandler) HandleMessage(m *nsqGo.Message) error {
	err := c.handler(&Message{
		Timestamp: m.Timestamp,
		Body:      string(m.Body),
	})

	if err != nil {
		m.Requeue(10 * time.Second)
	}

	return nil
}

func newNsq(fnOpt ...FnOpt) (*nsq, error) {

	options := new(Options)
	for _, o := range fnOpt {
		o(options)
	}

	config := nsqGo.NewConfig()

	producer, err := nsqGo.NewProducer(options.NSQd, config)
	if err != nil {
		return nil, err
	}

	consumerList := make([]*nsqGo.Consumer, 0)
	for _, optConsumer := range options.ConsumerList {
		consumer, err := nsqGo.NewConsumer(optConsumer.Topic, optConsumer.Channel, config)
		if err != nil {
			return nil, err
		}

		consumer.AddHandler(&consumerHandler{
			handler: optConsumer.Handler,
		})

		err = consumer.ConnectToNSQLookupd(options.NSQLookupd)
		if err != nil {
			return nil, err
		}

		consumerList = append(consumerList, consumer)
	}

	return &nsq{
		producer:     producer,
		consumerList: consumerList,
	}, nil
}

func (n *nsq) Publish(topic string, message string) error {
	err := n.producer.Publish(topic, []byte(message))
	if err != nil {
		return err
	}
	return nil
}

func (n *nsq) Close() error {
	if n.producer != nil {
		n.producer.Stop()
	}

	for _, consumer := range n.consumerList {
		consumer.Stop()
	}

	return nil
}
