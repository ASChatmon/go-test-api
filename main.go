package main

import (
	"go-test-api/config"
	"go-test-api/handlers"
	"go-test-api/handlers/middleware"
	"go-test-api/storage"
	"github.com/goji/httpauth"
	"github.com/zenazn/goji"
	"github.com/zenazn/goji/web"
	gojiMiddleware "github.com/zenazn/goji/web/middleware"
	"net/http"
	"time"
	"os"
)

func main() {

	// App Config
	conf := config.New()
	// add db
	database, err := storage.NewConnection(conf.DatabaseUrl)
	if err != nil {
		conf.Log.Log.Fatalf("Failed to connect to DB! [%s]", err)
		return
	}

	conf.Connection = database

	authOpts := httpauth.AuthOptions{
		AuthFunc: conf.IsInUserMap,
	}

	go startWorker(conf)

	auth := web.New()

	goji.Abandon(gojiMiddleware.Logger)

	goji.Handle("/api/*", auth)

	// Basic Auth
	auth.Use(httpauth.BasicAuth(authOpts))
	auth.Use(middleware.LogWithTransactionId(&conf))

	// host documentation
	docsPath := os.Getenv("GOPATH") + "/"
	goji.Get("/docs/*", http.FileServer(http.Dir(docsPath)))

	// Health Check endpoint
	goji.Get("/healthcheck", func(c web.C, w http.ResponseWriter, r *http.Request) {
		handlers.GetHealthHandler(&conf, c, w, r)
	})

	auth.Get("/api/metrics", func(c web.C, w http.ResponseWriter, r *http.Request) {
		handlers.GetCurrentMetics(&conf, c, w, r)
	})

	auth.Get("/api/metrics/datetime/:timestamp", func(c web.C, w http.ResponseWriter, r *http.Request) {
		handlers.GetMeticsByTimestamp(&conf, c, w, r)
	})

	auth.Get("/api/metrics/aggregates", func(c web.C, w http.ResponseWriter, r *http.Request) {
		handlers.GetAggregatedMetrics(&conf, c, w, r)
	})

	auth.Get("/api/metrics/averages", func(c web.C, w http.ResponseWriter, r *http.Request) {
		handlers.GetAverageMetrics(&conf, c, w, r)
	})

	// Endpoints for manipulating sites and users.
	goji.Serve()
}

func work(c config.Config) {
	c.Log.LogInfo("work","getting latest metrics","")
	handlers.WorkerMetrics(&c)
}

func startWorker(c config.Config) {
	c.Log.LogInfo("Worker","Starting Worker","")
	duration, err := time.ParseDuration(c.Interval)
	if err != nil {
		return
	}
	for {
		time.Sleep(duration)
		go work(c)
	}
}