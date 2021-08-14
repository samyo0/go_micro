package constants

import "go.mongodb.org/mongo-driver/bson/primitive"

var TicketCreated = "ticket:created"
var TicketUpdated = "ticket:updated"

type Data struct {
	Id     primitive.ObjectID `json:"id"`
	Title  string             `json:"title"`
	Price  int                `json:"price"`
	UserId string             `json:"userId"`
}

type TicketEvent struct {
	Subject string
	Data    Data
}
