package app

import (
	"github.com/gin-gonic/gin"
	"github.com/samyo0/go_micro/nats/publisher"
	"github.com/samyo0/go_micro/tickets/domain/ticket"
	http "github.com/samyo0/go_micro/tickets/http/ticket"
	"github.com/samyo0/go_micro/tickets/nats"
	"github.com/samyo0/go_micro/tickets/respository/db"
)

var (
	router = gin.Default()
)

func StartApplication() {
	ticketService := ticket.NewService(db.NewRepository(), publisher.NewPublisher(nats.NewNatsClient()))
	ticketHandler := http.NewHandler(ticketService)

	router.GET("/api/tickets", ticketHandler.GetAll)
	router.POST("/api/tickets", ticketHandler.CreateTicket)
	router.PUT("/api/tickets/:ticket_id", ticketHandler.UpdateById)
	router.GET("/api/tickets/:ticket_id", ticketHandler.FindById)

	router.Run(":3000")
}
