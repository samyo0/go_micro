package nats

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"log"
	"os"

	nats "github.com/nats-io/nats.go"
	"github.com/nats-io/stan.go"
	"github.com/samyo0/go_micro/nats/constants"
)

const (
	Cluster_ID = "NATS_CLUSTER_ID"
	Client_ID  = "NATS_CLIENT_ID"
	URL        = "NATS_URL"
)

var (
	nc        *nats.Conn
	clusterid = os.Getenv(Cluster_ID)
	clientid  = os.Getenv(Client_ID)
	url       = os.Getenv(URL)
)

func NewNatsClient() stan.Conn {
	opts := []nats.Option{nats.Name("NATS Streaming Example Publisher")}

	// Connect to NATS
	var err error
	nc, err = nats.Connect(url, opts...)
	if err != nil {
		log.Fatal(url)
		log.Fatal(err)
	}
	defer nc.Close()

	client, err := stan.Connect(clusterid, clientid, stan.NatsConn(nc))
	if err != nil {
		log.Fatalf("Can't connect: %v.\nMake sure a NATS Streaming Server is running at: %s", err, url)
	}
	defer client.Close()

	e := constants.TicketEvent{
		Subject: constants.TicketCreated,
		Data: constants.Data{
			Title:  "matix",
			Price:  427,
			UserId: "123",
		},
	}

	err = client.Publish(e.Subject, encodeToBytes(e.Data))
	if err != nil {
		log.Fatalf("Error during publish: %v\n", err)
	}
	log.Printf("Published [%s] : '%v'\n", e.Subject, e.Data)

	log.Printf("Connected to NATS clusterid: [%s], clientId:[%s], url:[%s]", clusterid, clientid, url)

	return client
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
