package nats

import (
	"log"
	"math/rand"
	"os"

	nats "github.com/nats-io/nats.go"
	"github.com/nats-io/stan.go"
)

const (
	Cluster_ID = "CLUSTER_ID"
	Client_ID  = "CLIENT_ID"
	URL        = "URL"
)

var (
	Client    stan.Conn
	clusterid = os.Getenv(Cluster_ID)
	//clientid  = os.Getenv(Client_ID)
	//url = os.Getenv(URL)
)

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func RandStringBytes(n int) string {
	b := make([]byte, n)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return string(b)
}

func init() {
	opts := []nats.Option{nats.Name("NATS Streaming Example Publisher")}

	// Connect to NATS
	nc, err := nats.Connect(stan.DefaultNatsURL, opts...)
	if err != nil {
		log.Fatal(err)
	}
	defer nc.Close()

	Client, err = stan.Connect(clusterid, RandStringBytes(4), stan.NatsConn(nc))
	if err != nil {
		log.Fatalf("Can't connect: %v.\nMake sure a NATS Streaming Server is running at: %s", err, stan.DefaultNatsURL)
	}
	defer Client.Close()

	log.Printf("Connected to NATS")
}
