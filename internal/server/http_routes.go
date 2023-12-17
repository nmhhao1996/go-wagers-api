package server

import (
	wagerHTTP "github.com/nmhhao1996/go-wagers-api/internal/wager/delivery/http"
	wagerRepository "github.com/nmhhao1996/go-wagers-api/internal/wager/repository"
	wagerUsecase "github.com/nmhhao1996/go-wagers-api/internal/wager/usecase"
)

func (srv HTTP) mapRoutes() {
	// Repositories
	wagerRepo := wagerRepository.New(srv.db)

	// Usecases
	wagerUC := wagerUsecase.New(wagerRepo)

	// Handlers
	wagerH := wagerHTTP.New(srv.l, wagerUC)

	// Routes
	srv.gin.POST("/wagers", wagerH.Create)
	srv.gin.GET("/wagers", wagerH.List)
	srv.gin.POST("/buy/:wager_id", wagerH.Buy)
}
