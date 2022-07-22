package message_broker

import (
	"time"

	nsqGo "github.com/nsqio/go-nsq"
)

type nsqWriter struct {
	producer *nsqGo.Producer
}

type nsqReader struct {
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

func newNsqWriter(fnOpt ...FnOpt) (*nsqWriter, error) {

	options := new(Options)
	for _, o := range fnOpt {
		o(options)
	}

	config := nsqGo.NewConfig()

	producer, err := nsqGo.NewProducer(options.NSQd, config)
	if err != nil {
		return nil, err
	}

	return &nsqWriter{
		producer: producer,
	}, nil
}

func (n *nsqWriter) Publish(topic string, message string) error {
	err := n.producer.Publish(topic, []byte(message))
	if err != nil {
		return err
	}
	return nil
}

func (n *nsqWriter) Close() error {
	if n.producer != nil {
		n.producer.Stop()
	}

	return nil
}

func newNsqReader(fnOpt ...FnOpt) (*nsqReader, error) {

	options := new(Options)
	for _, o := range fnOpt {
		o(options)
	}

	config := nsqGo.NewConfig()

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

	return &nsqReader{
		consumerList: consumerList,
	}, nil
}

func (n *nsqReader) Close() error {

	for _, consumer := range n.consumerList {
		consumer.Stop()
	}

	return nil
}
