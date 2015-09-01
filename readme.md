# goPubsub
![](https://travis-ci.org/anamewing/goPubsub.svg)  
goPubsub is a Publish–Subscribe component writen in Go 

## Usage
Using 
```
Publisher := NewPubsub()
```
to create a publisher.

Subscriber must implement interface `Notify(event Event) error`.

Using
```
Publisher.Subscribe(eventType string, subscriber Subscriber) error
```
to subscribe the publishter. `eventType` is a filter, and subscriber only get notified when the choosen `eventType`'s event is published.

Using
```
Publisher.Publish(eventType string, event Event) error
```
to publish an event.
