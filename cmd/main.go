package main

import (
	"github.com/spf13/viper"
	"os"
	"tinkapi/v2/internal/server"
	"tinkapi/v2/pkg/config"
	"tinkapi/v2/pkg/handler"
	"tinkapi/v2/pkg/logger"
	"tinkapi/v2/pkg/repository"
	"tinkapi/v2/service"
)

func main() {
	// READ CONFIG
	cfg := config.MustLoad()

	// INIT LOGGER
	l := logger.NewLogger(cfg.Env)

	// INIT DB
	db, err := repository.NewSqlite(os.Getenv("SQLITE_PATH"))
	if err != nil {
		l.Panic(err.Error())
	}

	// INIT REPO & SERVICES
	repos := repository.NewRepository(db)
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)

	srv := new(server.Server)
	if err := srv.Run(viper.GetString("http_server.address"), handlers.InitRouter(),
		cfg.HTTPServer.ReadTimeout, cfg.HTTPServer.WriteTimeout); err != nil {
		l.Error("error occurred while running http server: %s", err.Error())
	}
}
