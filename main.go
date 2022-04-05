package main

import (
	"go-api-starter-kit/routes"
	"go-api-starter-kit/utils"
	"go-api-starter-kit/utils/logger"
	"log"
	"net/http"

	"github.com/joho/godotenv"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Print(".env file does not exist, will get the variables from the environment")
	}
}

func main() {
	var err error

	logger, err := logger.New()
	if err != nil {
		panic(err)
	}

	db, err := utils.InitSQL()
	if err != nil {
		panic(err)
	}

	err = utils.InitMigration(db)
	if err != nil {
		panic(err)
	}

	audit, err := utils.InitAudit()
	if err != nil {
		panic(err)
	}

	r, err := utils.InitRouter()
	if err != nil {
		panic(err)
	}

	routes.AddRoutes(r, db, logger, audit)

	apiPort, prometheusPort := utils.InitPorts()
	go func() { panic(http.ListenAndServe(":"+prometheusPort, promhttp.Handler())) }()
	r.Run(":" + apiPort)
}
