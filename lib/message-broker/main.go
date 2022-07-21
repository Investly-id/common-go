package message_broker

type MessageBroker interface {
	Publish(topic string, message string) error
	Close() error
}

type ConsumerConfig struct {
	Topic   string
	Channel string
	Handler func(m *Message) error
}

type Message struct {
	Body      string
	Timestamp int64
}

const (
	DriverNSQ   = "nsq"
	DriverKafka = "kafka"
)

func New(driverName string, opt ...FnOpt) (MessageBroker, error) {
	switch driverName {
	case DriverNSQ:
		client, err := newNsq(opt...)
		if err != nil {
			return nil, err
		}
		return client, nil
	}

	return nil, nil
}
