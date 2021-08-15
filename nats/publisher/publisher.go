package publisher

import (
	"encoding/json"
	"log"

	"github.com/nats-io/stan.go"
	"github.com/samyo0/go_micro/nats/constants"
)

type Publisher interface {
	Publish(constants.TicketEvent)
}

type publisher struct {
	client stan.Conn
}

func NewPublisher(client stan.Conn) Publisher {
	return &publisher{
		client,
	}
}

func (p *publisher) Publish(e constants.TicketEvent) {
	err := p.client.Publish(e.Subject, json.Marshal(e.Data))
	if err != nil {
		log.Fatalf("Error during publish: %v\n", err)
	}
	log.Printf("Published [%s] : '%s'\n", e.Subject, e.Data)
}
