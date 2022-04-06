package main

import (
	"go-api-starter-kit/routes"
	"go-api-starter-kit/utils/audit"
	"go-api-starter-kit/utils/config"
	"go-api-starter-kit/utils/db"
	"go-api-starter-kit/utils/logger"
	"go-api-starter-kit/utils/router"
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

	l, err := logger.New()
	if err != nil {
		panic(err)
	}

	d, err := db.New()
	if err != nil {
		l.Panic(err.Error())
	}
	if err := db.Migrate(d, "file://migrations"); err != nil {
		l.Panic(err.Error())
	}

	a, err := audit.New()
	if err != nil {
		l.Panic(err.Error())
	}

	r, err := router.New("./assets")
	if err != nil {
		l.Panic(err.Error())
	}

	routes.AddRoutes(r, d, l, a)

	apiPort, prometheusPort := config.SetPorts()
	go func() { l.Panic(http.ListenAndServe(":"+prometheusPort, promhttp.Handler()).Error()) }()
	r.Run(":" + apiPort)
}
