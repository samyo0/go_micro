package ticket

import "go.mongodb.org/mongo-driver/bson/primitive"

type Ticket struct {
	Id     primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Title  string             `json:"title" bson:"title"`
	Price  int                `json:"price" bson:"price"`
	UserId string             `json:"user_id" bson:"user_id"`
}
