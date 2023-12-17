package main

import (
	"context"

	pkgErrors "github.com/go-errors/errors"
	"github.com/nmhhao1996/go-wagers-api/config"
	"github.com/nmhhao1996/go-wagers-api/internal/core/mysql"
	"github.com/nmhhao1996/go-wagers-api/internal/server"
	"github.com/nmhhao1996/go-wagers-api/pkg/log"
)

func main() {
	ctx := context.Background()

	cfg, err := config.Load()
	if err != nil {
		panic(err)
	}

	l := log.NewLogger(log.Config(cfg.Logger))
	l.InitLogger()

	db, err := mysql.Connect(cfg.MySQL)
	if err != nil {
		panic(err)
	}
	l.Info(ctx, "Connected to MySQL database")

	pkgErrors.MaxStackDepth = 5

	if err := server.NewHTTP(
		cfg.HTTPServer,
		l,
		db,
	).Start(); err != nil {
		panic(err)
	}
}
