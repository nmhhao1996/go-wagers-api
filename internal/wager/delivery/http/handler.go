package http

import "github.com/gin-gonic/gin"

// Handler is the interface for the wager http handler
type Handler interface {
	// Create creates a new wager
	Create(c *gin.Context)
	// List lists all wagers
	List(c *gin.Context)
	// Buy buys a wager
	Buy(c *gin.Context)
}
