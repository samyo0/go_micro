package nats

import (
	"log"
	"os"

	nats "github.com/nats-io/nats.go"
)

const (
	Cluster_ID = "NATS_CLUSTER_ID"
	Client_ID  = "NATS_CLIENT_ID"
	URL        = "NATS_URL"
)

var (
	NC  *nats.Conn
	url = os.Getenv(URL)
)

func init() {
	opts := []nats.Option{nats.Name("NATS Streaming Example Publisher")}

	// Connect to NATS
	var err error
	NC, err = nats.Connect(url, opts...)
	if err != nil {
		log.Fatal(url)
		log.Fatal(err)
	}
	defer NC.Close()

	// e := constants.TicketEvent{
	// 	Subject: constants.TicketCreated,
	// 	Data: constants.Data{
	// 		Title:  "matix",
	// 		Price:  427,
	// 		UserId: "123",
	// 	},
	// }

	// err = Client.Publish(e.Subject, encodeToBytes(e.Data))
	// if err != nil {
	// 	log.Fatalf("Error during publish: %v\n", err)
	// }
	// log.Printf("Published [%s] : '%s'\n", e.Subject, e.Data)

	// log.Printf("Connected to NATS clusterid: [%s], clientId:[%s], url:[%s]", clusterid, clientid, url)
}
