package main

import (
	"github.com/gin-gonic/gin"
	"github.com/ivanramirez-git/DoublePartners/my_ticket_api_golang/controllers"
	"github.com/ivanramirez-git/DoublePartners/my_ticket_api_golang/initializers"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}

func main() {
	r := gin.Default()

	// POST routes
	r.POST("/user", controllers.UserCreate)

	r.POST("/ticket", controllers.TicketCreate)

	// GET routes
	r.GET("/tickets", controllers.TicketIndex)
	r.GET("/ticket/:id", controllers.TicketShow)

	r.GET("/users", controllers.UserIndex)
	r.GET("/user/:id", controllers.UserShow)

	// PUT routes
	r.PUT("/ticket/:id", controllers.TicketUpdate)

	r.PUT("/user/:id", controllers.UserUpdate)

	// DELETE routes
	r.DELETE("/ticket/:id", controllers.TicketDelete)

	r.DELETE("/user/:id", controllers.UserDelete)

	r.Run()
}
