package http

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/samyo0/go_micro/tickets/domain/ticket"
)

type TicketHandler interface {
	GetAll(*gin.Context)
	CreateTicket(*gin.Context)
	FindById(*gin.Context)
	UpdateById(*gin.Context)
}

type ticketHandler struct {
	service ticket.Service
}

func NewHandler(service ticket.Service) TicketHandler {
	return &ticketHandler{
		service: service,
	}
}

func getTicketId(idParam string) (int64, error) {
	id, err := strconv.ParseInt(idParam, 10, 64)
	if err != nil {
		return 0, errors.New("invalid id")

	}
	return id, nil
}

func (h *ticketHandler) GetAll(c *gin.Context) {
	tickets, err := h.service.GetAll()

	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, tickets)
}

func (h *ticketHandler) CreateTicket(c *gin.Context) {
	var ticket ticket.Ticket

	if err := c.ShouldBindJSON(&ticket); err != nil {
		restErr := errors.New("invalid json body")
		c.JSON(http.StatusInternalServerError, restErr.Error())
		return
	}

	saveErr := h.service.CreateTicket(ticket)
	if saveErr != nil {
		c.JSON(http.StatusInternalServerError, saveErr.Error())
		return
	}

	c.JSON(http.StatusCreated, "Created")
}

func (h *ticketHandler) FindById(c *gin.Context) {

	ticket, ticketErr := h.service.FindById(c.Param("ticket_id"))
	if ticketErr != nil {
		c.JSON(http.StatusInternalServerError, ticketErr)
		return
	}

	c.JSON(http.StatusOK, ticket)

}

func (h *ticketHandler) UpdateById(c *gin.Context) {

	var ticket *ticket.Ticket

	if err := c.ShouldBindJSON(&ticket); err != nil {
		restErr := errors.New("invalid json body")
		c.JSON(http.StatusInternalServerError, restErr.Error())
		return
	}

	fmt.Println("handler", ticket)

	result, updateErr := h.service.UpdateById(ticket)
	if updateErr != nil {
		c.JSON(http.StatusInternalServerError, updateErr)
		return
	}

	c.JSON(http.StatusOK, result)
}
