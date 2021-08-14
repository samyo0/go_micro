package main

import (
	"bytes"
	"encoding/gob"
	"flag"
	"fmt"
	"log"
	"math/rand"
	"os"

	nats "github.com/nats-io/nats.go"
	"github.com/nats-io/stan.go"
	"github.com/samyo0/go_micro/nats/constants"
	"github.com/samyo0/go_micro/nats/publisher"
)

var usageStr = `
Usage: stan-pub [options] <subject> <message>
Options:
	-s,  --server   <url>            NATS Streaming server URL(s)
	-c,  --cluster  <cluster name>   NATS Streaming cluster name
`

// NOTE: Use tls scheme for TLS, e.g. stan-pub -s tls://demo.nats.io:4443 foo hello
func usage() {
	fmt.Printf("%s\n", usageStr)
	os.Exit(0)
}

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func RandStringBytes(n int) string {
	b := make([]byte, n)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return string(b)
}

func main() {
	var (
		clusterID string
		URL       string
	)

	flag.StringVar(&URL, "s", stan.DefaultNatsURL, "The nats server URLs (separated by comma)")
	flag.StringVar(&URL, "server", stan.DefaultNatsURL, "The nats server URLs (separated by comma)")
	flag.StringVar(&clusterID, "c", "test-cluster", "The NATS Streaming cluster ID")
	flag.StringVar(&clusterID, "cluster", "test-cluster", "The NATS Streaming cluster ID")

	log.SetFlags(0)
	flag.Usage = usage
	flag.Parse()

	// Connect Options.
	opts := []nats.Option{nats.Name("NATS Streaming Example Publisher")}

	// Connect to NATS
	nc, err := nats.Connect(URL, opts...)
	if err != nil {
		log.Fatal(err)
	}
	defer nc.Close()

	sc, err := stan.Connect(clusterID, RandStringBytes(4), stan.NatsConn(nc))
	if err != nil {
		log.Fatalf("Can't connect: %v.\nMake sure a NATS Streaming Server is running at: %s", err, URL)
	}
	defer sc.Close()

	// subj, msg := args[0], []byte(args[1])
	e := constants.TicketEvent{
		Subject: constants.TicketCreated,
		Data: constants.Data{
			Title:  "matix",
			Price:  427,
			UserId: "123",
		},
	}

	publisher.NewPublisher(sc).Publish(e)
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
