package api

import (
	db "github.com/Nazoviya/basketService/db/sqlc"
	"github.com/gin-gonic/gin"
)

// create Server struct to hold db and gin pointers.
type Server struct {
	store  *db.Store
	router *gin.Engine
}

// create a NewServer function that returns Server pointer.
func NewServer(store *db.Store) *Server {
	// gin.SetMode(gin.ReleaseMode)
	server := &Server{store: store}
	router := gin.Default()

	// GET requests from db to execute following functions.
	router.GET("/products", server.getProducts)                      // returns all products.
	router.GET("/basket", server.calculateBasket, server.showBasket) // returns all items in basket.

	// POST request from db to execute following functions.
	router.POST("/basket/:product_id", server.addToBasket)          // posts specified product with it's id.
	router.POST("/basket/del/:product_id", server.deleteFromBasket) // deletes specified product with it's id.
	router.POST("/complete", server.completeOrder)                  // completes order, clears basket, returns message.

	server.router = router
	return server
}

// Create a method Start, to run http server on given address.
func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

// returns an error if any, that belongs to gin functions.
func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
