package nats

import (
	"log"
	"os"

	nats "github.com/nats-io/nats.go"
	"github.com/nats-io/stan.go"
)

const (
	Cluster_ID = "NATS_CLUSTER_ID"
	Client_ID  = "NATS_CLIENT_ID"
	URL        = "NATS_URL"
)

var (
	Client    stan.Conn
	clusterid = os.Getenv(Cluster_ID)
	clientid  = os.Getenv(Client_ID)
	url       = os.Getenv(URL)
)

func init() {
	opts := []nats.Option{nats.Name("NATS Streaming Example Publisher")}

	// Connect to NATS
	nc, err := nats.Connect(url, opts...)
	if err != nil {
		log.Fatal(url)
		log.Fatal(err)
	}
	defer nc.Close()

	Client, err = stan.Connect(clusterid, clientid, stan.NatsConn(nc))
	if err != nil {
		log.Fatalf("Can't connect: %v.\nMake sure a NATS Streaming Server is running at: %s", err, url)
	}
	defer Client.Close()

	log.Printf("Connected to NATS")
}
