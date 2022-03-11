package main

import (
	"context"
	"go-api-starter-kit/router"
	"go-api-starter-kit/utils"
	"log"
	"net/http"

	"github.com/joho/godotenv"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"go.uber.org/zap"
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Print(".env file does not exist, will get the variables from the environment")
	}
}

func main() {
	var err error

	ctx := context.Background()

	logger, err := utils.InitLogger()
	if err != nil {
		log.Fatal(err)
	}

	db, err := utils.InitSQL()
	if err != nil {
		logger.Panic("Cannot init sql connection..", zap.Error(err))
	}

	r, err := utils.InitRouter()
	if err != nil {
		logger.Panic("Cannot init gin router..", zap.Error(err))
	}

	router.AddRoutes(r, db, logger, ctx)

	apiPort, prometheusPort := utils.InitPorts()
	go func() { log.Fatal(http.ListenAndServe(":"+prometheusPort, promhttp.Handler())) }()
	r.Run(":" + apiPort)
}
