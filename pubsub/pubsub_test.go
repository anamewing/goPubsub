package pubsub

import (
	"testing"
)

func TestPubsub(t *testing.T) {
	//	compile time assertion
	//	var _ pubsub.Publisher = (*pubsub.Pubsub)(nil)
	//	run time check, both is ok
	var p interface{} = &Pubsub{}
	if _, ok := p.(MessageQueue); !ok {
		t.Fatal("Pubsub does not implement MessageQueue")
	}

type Suber struct {
	tester *testing.T
	eventT string
}

func (s *Suber) Notify(event Event) error {
	s.tester.Logf(event.(string))
	return nil
}
