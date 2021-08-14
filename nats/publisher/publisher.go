package publisher

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"log"
	"os"

	"github.com/nats-io/nats.go"
	"github.com/nats-io/stan.go"
	"github.com/samyo0/go_micro/nats/constants"
)

const (
	Cluster_ID = "NATS_CLUSTER_ID"
	Client_ID  = "NATS_CLIENT_ID"
	URL        = "NATS_URL"
)

var (
	clusterid = os.Getenv(Cluster_ID)
	clientid  = os.Getenv(Client_ID)
	url       = os.Getenv(URL)
)

type Publisher interface {
	Publish(constants.TicketEvent)
}

type publisher struct {
	nats *nats.Conn
}

func NewPublisher(nats *nats.Conn) Publisher {
	return &publisher{
		nats,
	}
}

func (p *publisher) Publish(e constants.TicketEvent) {
	client, err := stan.Connect(clusterid, clientid, stan.NatsConn(p.nats))
	if err != nil {
		log.Fatalf("Can't connect: %v.\nMake sure a NATS Streaming Server is running at: %s", err, url)
	}
	defer client.Close()

	err = client.Publish(e.Subject, encodeToBytes(e.Data))
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
