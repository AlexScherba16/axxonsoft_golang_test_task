package main

import (
	"axxonsoft_golang_test_task/config"
	"axxonsoft_golang_test_task/internal/api"
	"axxonsoft_golang_test_task/internal/middleware"
	"axxonsoft_golang_test_task/internal/server"
	_map "axxonsoft_golang_test_task/pkg/map"
	"context"
	"github.com/joho/godotenv"
	log "github.com/sirupsen/logrus"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	appConfig := config.NewApplicationSettings()
	reqRespMap := _map.NewRequestResponseMap()
	proxyMiddleware := middleware.NewProxyServerMiddleware(reqRespMap)
	appHandler := api.NewApplicationHandler(appConfig, proxyMiddleware)
	proxyServer := server.NewProxyServer(appConfig)

	go func() {
		if err := proxyServer.RunServer(appHandler.NewRouters()); err != nil {
			log.Fatalf("error while start server %s", err.Error())
		}
	}()
	log.Infof("Server started on port %d...", appConfig.Server.Port)

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit

	log.Info("Server shutting down")
	if err := proxyServer.ShutdownServer(context.Background()); err != nil {
		log.Fatalf("error occurred on server shutting down: %s", err.Error())
	}
}

func init() {
	if err := godotenv.Load("./common.env"); err != nil {
		log.Fatalf("No .env file found: %s", err.Error())
	}
}
