package message_broker

type Options struct {
	Host         string
	NSQLookupd   string
	NSQd         string
	Port         string
	ConsumerList []*ConsumerConfig
}

type FnOpt func(o *Options) (err error)

func WithHost(host string) FnOpt {
	return func(o *Options) (err error) {
		o.Host = host
		return
	}
}

func WithNSQLookupd(NSQLookupd string) FnOpt {
	return func(o *Options) (err error) {
		o.NSQLookupd = NSQLookupd
		return
	}
}

func WithNSQd(NSQd string) FnOpt {
	return func(o *Options) (err error) {
		o.NSQd = NSQd
		return
	}
}

func WithPort(port string) FnOpt {
	return func(o *Options) (err error) {
		o.Port = port
		return
	}
}

func AddConsumer(config *ConsumerConfig) FnOpt {
	return func(o *Options) (err error) {
		o.ConsumerList = append(o.ConsumerList, config)
		return
	}
}
