package db

import (
	"context"
	"errors"
	"fmt"
	"log"
	"os"

	"github.com/samyo0/go_micro/tickets/domain/ticket"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func NewRepository() DbRepository {
	return &dbRepository{}
}

type DbRepository interface {
	Create(ticket.Ticket) error
	GetAll() ([]*ticket.Ticket, error)
	FindById(string) (*ticket.Ticket, error)
	UpdateById(*ticket.Ticket) (*ticket.Ticket, error)
}

type dbRepository struct{}

var (
	clientInstanceError error
	Client              *mongo.Client
	ticketDatabase      string = "ticketsDB"
	ticketCollection    string = "tickets"
	MONGO_URI           string = os.Getenv("MONGO_URI")
)

// "mongodb://root:example@localhost:27017"
func init() {
	// Set client options
	clientOptions := options.Client().ApplyURI(MONGO_URI)
	// Connect to MongoDB
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		clientInstanceError = err
	}
	// Check the connection
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		clientInstanceError = err
	}
	Client = client
}

func (db *dbRepository) Create(ticket ticket.Ticket) error {

	_, err := Client.Database(ticketDatabase).Collection(ticketCollection).InsertOne(context.TODO(), ticket)

	if err != nil {
		return errors.New("There was an error with db")
	}
	return nil
}

func (db *dbRepository) GetAll() ([]*ticket.Ticket, error) {
	cur, err := Client.Database(ticketDatabase).Collection(ticketCollection).Find(context.TODO(), bson.D{{}})
	if err != nil {
		return nil, errors.New("There was an error with db")
	}

	var results []*ticket.Ticket

	// Finding multiple documents returns a cursor
	// Iterating through the cursor allows us to decode documents one at a time
	for cur.Next(context.TODO()) {

		// create a value into which the single document can be decoded
		var elem ticket.Ticket
		err := cur.Decode(&elem)
		if err != nil {
			log.Fatal(err)
		}

		results = append(results, &elem)
	}

	if err := cur.Err(); err != nil {
		log.Fatal(err)
	}

	// Close the cursor once finished
	cur.Close(context.TODO())

	return results, nil
}

func (db *dbRepository) FindById(id string) (*ticket.Ticket, error) {
	objectId, _ := primitive.ObjectIDFromHex(id)
	filter := bson.M{"_id": objectId}
	var result ticket.Ticket

	dbErr := Client.Database(ticketDatabase).Collection(ticketCollection).FindOne(context.TODO(), filter).Decode(&result)
	if dbErr != nil {
		log.Fatal(dbErr)
	}

	return &result, nil
}

func (db *dbRepository) UpdateById(ticket *ticket.Ticket) (*ticket.Ticket, error) {

	result, dbErr := Client.Database(ticketDatabase).Collection(ticketCollection).UpdateOne(
		context.TODO(),
		bson.M{"_id": ticket.Id},
		bson.D{{
			Key: "$set",
			Value: bson.D{
				primitive.E{
					Key:   "price",
					Value: ticket.Price,
				},
				primitive.E{
					Key:   "title",
					Value: ticket.Title,
				},
				primitive.E{
					Key:   "user_id",
					Value: ticket.UserId,
				},
			}},
		})
	if dbErr != nil {
		log.Fatal(dbErr)
	}

	fmt.Println(result)
	return nil, nil
}
