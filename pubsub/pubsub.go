package pubsub

type Pubsub struct {
	subscriberMap map[string]([]Subscriber)
}

type Event interface{}

type MessageQueue interface {
	Publish(eventType string, event Event) error
	Subscribe(eventType string, subscriber Subscriber) error
}

type Subscriber interface {
	Notify(event Event) error
}

func NewPubsub() *Pubsub {
	p := &Pubsub{
		subscriberMap: make(map[string]([]Subscriber)),
	}
	return p
}

func (p *Pubsub) Publish(eventType string, event Event) error {
	for _, subscriber := range p.subscriberMap[eventType] {
		go func(subscriber Subscriber) {
			subscriber.Notify(event)
		}(subscriber)
	}
	return nil
}

func (p *Pubsub) Subscribe(eventType string, subscriber Subscriber) error {
	p.subscriberMap[eventType] = append(p.subscriberMap[eventType], subscriber)
	return nil
}
