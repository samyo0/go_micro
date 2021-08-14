package ticket

import (
	"bytes"
	"encoding/gob"
	"errors"
	"fmt"
	"log"

	"github.com/samyo0/go_micro/nats/constants"
	"github.com/samyo0/go_micro/nats/publisher"
)

type Repository interface {
	Create(Ticket) error
	GetAll() ([]*Ticket, error)
	FindById(string) (*Ticket, error)
	UpdateById(*Ticket) (*Ticket, error)
}

type Service interface {
	GetAll() ([]*Ticket, error)
	CreateTicket(Ticket) error
	FindById(string) (*Ticket, error)
	UpdateById(*Ticket) (*Ticket, error)
}

type service struct {
	repository Repository
	stan       publisher.Publisher
}

func NewService(repo Repository, stan publisher.Publisher) Service {
	return &service{
		repository: repo,
		stan:       stan,
	}
}

func (s *service) GetAll() ([]*Ticket, error) {
	tickets, err := s.repository.GetAll()
	if err != nil {
		return nil, err
	}

	return tickets, nil
}

func (s *service) CreateTicket(ticket Ticket) error {
	//TODO: validate ticket

	if err := s.repository.Create(ticket); err != nil {
		return err
	}

	e := constants.TicketEvent{
		Subject: constants.TicketCreated,
		Data: constants.Data{
			Title:  ticket.Title,
			Price:  ticket.Price,
			UserId: ticket.UserId,
		},
	}

	s.stan.Publish(e)

	return nil
}

func (s *service) FindById(id string) (*Ticket, error) {
	ticket, err := s.repository.FindById(id)
	if err != nil {
		return nil, err
	}
	return ticket, nil
}

func (s *service) UpdateById(ticket *Ticket) (*Ticket, error) {
	current, err := s.repository.FindById(ticket.Id.Hex())
	if err != nil {
		return nil, err
	}
	current.Price = ticket.Price
	current.Title = ticket.Title
	current.UserId = ticket.UserId

	s.repository.UpdateById(current)

	return nil, errors.New("implement me FindById")
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
