package pubsub

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

func TestPubsub(t *testing.T) {
	//	compile time assertion
	Publisher := NewPubsub()
	eventTypes := []string{"notes", "images", "videos"}
	suberPerType := 20
	suberList := make([]*Suber, 0, len(eventTypes)*suberPerType)

	eventNameCount := 20
	eventMaxPubTimes := 30
	chanMaxCount := len(eventTypes) * eventNameCount * eventMaxPubTimes * suberPerType

	bufChan := make(chan bool, chanMaxCount)

	//	make subscribers
	for _, eventType := range eventTypes {
		for i := 0; i < suberPerType; i++ {
			sub := &Suber{tester: t, eventT: eventType, eventCounter: make(map[string]int), bufChan: bufChan}
			suberList = append(suberList, sub)
			Publisher.Subscribe(eventType, sub)
		}
	}

	//	start publish
	eventMap := make(map[string](map[string]int))
	for _, eventType := range eventTypes {
		eventMap[eventType] = make(map[string]int)
	}

	realPubTimes := 0
	pubStart := time.Now()
	for i := 0; i < len(eventTypes); i++ {
		for j := 0; j < 2; j++ {
			pubTimes := rand.Intn(30)
			eventName := fmt.Sprintf("Name%d", j)
			eventMap[eventTypes[i]][eventName] += pubTimes
			for k := 0; k < pubTimes; k++ {
				Publisher.Publish(eventTypes[i], eventName)
			}
			realPubTimes += pubTimes
		}
	}
	t.Logf("pub runs %v", time.Now().Sub(pubStart))

	//	blocking
	for i := 0; i < realPubTimes*suberPerType; i++ {
		<-bufChan
	}
	//	check subers
	for _, suber := range suberList {
		for eventName, counts := range eventMap[suber.eventT] {
			// t.Logf("pub %s on %s for %d times", suber.eventT, eventName, counts)
			if suber.eventCounter[eventName] != counts {
				t.Errorf("suber %p count %s event %s: %d, expect %d", suber, suber.eventT, eventName, suber.eventCounter[eventName], counts)
			}
		}
	}
}

type Suber struct {
	tester       *testing.T
	eventT       string
	eventCounter map[string]int
	bufChan      chan bool
}

func (s *Suber) Notify(event Event) error {
	// s.tester.Logf("suber %p listen on %s get %s", s, s.eventT, event.(string))
	s.eventCounter[event.(string)]++
	time.Sleep(100 * time.Millisecond)
	s.bufChan <- true
	return nil
}
