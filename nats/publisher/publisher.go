package publisher

import (
	"bytes"
	"encoding/gob"
	"fmt"
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
	fmt.Println(e)
	err := p.client.Publish(e.Subject, encodeToBytes(e.Data))
	if err != nil {
		log.Fatalf("Error during publish: %v\n", err)
	}
	log.Printf("Published [%s] : '%s'\n", e.Subject, e.Data)
}

func encodeToBytes(p interface{}) []byte {

	buf := bytes.Buffer{}
	enc := gob.NewEncoder(&buf)
	err := enc.Encode(p)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("uncompressed size (bytes): ", len(buf.Bytes()))
	return buf.Bytes()
}
