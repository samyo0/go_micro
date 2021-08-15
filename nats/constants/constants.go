package constants

import "go.mongodb.org/mongo-driver/bson/primitive"

var TicketCreated = "ticket:created"
var TicketUpdated = "ticket:updated"

type Data struct {
	Id     primitive.ObjectID `json:"Id"`
	Title  string             `json:"Title"`
	Price  int                `json:"Price"`
	UserId string             `json:"UserId"`
}

type TicketEvent struct {
	Subject string `json:"subject"`
	Data    Data   `json:"data"`
}
