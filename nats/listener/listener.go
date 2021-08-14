package listener

import (
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/nats-io/stan.go"
	// "github.com/nats-io/stan.go/pb"
	"github.com/samyo0/go_micro/nats/constants"
)

type Listener interface {
	Listen()
}

type listener struct {
	subject        string
	queueGroupName string
	client         stan.Conn
}

func NewListener(subject string, queueGroupName string, client stan.Conn) Listener {
	log.Printf("New Listener")
	return &listener{
		subject,
		queueGroupName,
		client,
	}
}

func printMsg(m *stan.Msg) {
	fmt.Println(m.Sequence)
	fmt.Println(m.Data)
	log.Printf("Received: %s\n", m)
}

func (l *listener) Listen() {
	mcb := func(msg *stan.Msg) {
		var data constants.TicketEvent
		fmt.Println(msg.Data)
		json.Unmarshal(msg.Data, &data)

		fmt.Println("------")
		fmt.Println(data)
		msg.Ack()
	}

	aw, _ := time.ParseDuration("60s")
	_, err := l.client.QueueSubscribe(l.subject, l.queueGroupName, mcb,
		stan.DeliverAllAvailable(),
		stan.DurableName(l.queueGroupName),
		stan.SetManualAckMode(),
		stan.AckWait(aw),
	)

	if err != nil {
		l.client.Close()
		log.Fatal(err)
	}

	log.Printf("Listening on [%s], qgroup=[%s], durable=[%s]\n", l.subject, l.queueGroupName, l.queueGroupName)
}

// type TicketCreatedData struct {
// 	id     string
// 	title  string
// 	price  int
// 	userId string
// }

// type TicketCreatedEvent struct {
// 	subject string
// 	data    TicketCreatedData
// }

// type listener struct {
// 	subject        string
// 	queueGroupName string
// }

// func (l listener) onMessage(data TicketCreatedData, msg *stan.Msg) {
// 	fmt.Print("Event data!", data)

// 	fmt.Print("data price", data.price)
// 	fmt.Print("data title", data.title)
// 	fmt.Print("data id", data.id)

// 	msg.Ack()
// }
