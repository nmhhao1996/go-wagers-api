package server

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/gin-gonic/gin"
	"github.com/nmhhao1996/go-wagers-api/config"
	"github.com/nmhhao1996/go-wagers-api/pkg/log"
	"gorm.io/gorm"
)

const (
	defaultGinMode = gin.DebugMode
	productionMode = "production"
)

// HTTP is the HTTP server
type HTTP struct {
	port int
	mode string
	gin  *gin.Engine
	l    log.Logger
	db   *gorm.DB
}

// NewHTTP creates a new HTTP server
func NewHTTP(
	cfg config.HTTPServerConfig,
	l log.Logger,
	db *gorm.DB,
) HTTP {
	if cfg.Mode == productionMode {
		gin.SetMode(gin.ReleaseMode)
	} else {
		gin.SetMode(defaultGinMode)
	}

	return HTTP{
		gin:  gin.Default(),
		port: cfg.Port,
		mode: cfg.Mode,
		l:    l,
		db:   db,
	}
}

// Start starts the HTTP server
func (srv HTTP) Start() error {

	srv.mapRoutes()

	ctx := context.Background()
	go func() {
		srv.gin.Run(fmt.Sprintf(":%d", srv.port))
	}()

	srv.l.Infof(ctx, "Started server on :%d", srv.port)
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, syscall.SIGINT, syscall.SIGTERM)
	srv.l.Info(ctx, <-ch)
	srv.l.Info(ctx, "Stopping API server.")

	return nil
}
