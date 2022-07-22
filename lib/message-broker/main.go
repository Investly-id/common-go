package message_broker

type Writer interface {
	Publish(topic string, message string) error
	Close() error
}

type Reader interface {
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

func NewWriter(driverName string, opt ...FnOpt) (Writer, error) {
	switch driverName {
	case DriverNSQ:
		client, err := newNsqWriter(opt...)
		if err != nil {
			return nil, err
		}
		return client, nil
	}

	return nil, nil
}

func NewReader(driverName string, opt ...FnOpt) (Reader, error) {
	switch driverName {
	case DriverNSQ:
		client, err := newNsqReader(opt...)
		if err != nil {
			return nil, err
		}
		return client, nil
	}

	return nil, nil
}
