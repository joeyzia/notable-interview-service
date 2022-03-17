package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"notable-interview-service/config"
	"notable-interview-service/service"
	"os"
	"time"

	"github.com/go-kit/log"
	"github.com/go-stack/stack"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func main() {
	var (
		cfg       config.Config
		fileBytes []byte
	)

	newLogger := log.NewLogfmtLogger(log.NewSyncWriter(os.Stdout))
	logger := newLogContext(newLogger, "API")
	serviceLogger := newLogContext(newLogger, "service")

	//Configuration file loading
	logger.Log("event", "loading config.json")
	filename := "config.json"

	if _, err := os.Stat(filename); err == nil {
		if fileBytes, err = ioutil.ReadFile(filename); err != nil {
			logger.Log("event", "exiting", "err", err)
			os.Exit(1)
		}

		if err = json.Unmarshal(fileBytes, &cfg); err != nil {
			logger.Log("event", "exiting", "err", err)
			os.Exit(1)
		}
	} else {
		logger.Log("event", "exiting", "err", err)
		os.Exit(1)
	}

	// Set up and start http server
	r := mux.NewRouter()

	apiRouter := r.PathPrefix("/api").Subrouter()

	var middleware = service.SomeMiddleware("anything")

	svc := service.NewService(logger)
	svc = service.NewLoggingService(svc, serviceLogger)
	service.MakeRoutes(apiRouter, svc, logger, middleware)


	methodsOk := handlers.AllowedMethods([]string{"GET", "HEAD", "POST", "PUT", "OPTIONS"})
	originsOk := handlers.AllowedOrigins([]string{"*"})


	// Check if the port is set
	port := cfg.Port;
	if (port == "0") {
		logger.Log("event", "exiting", "err", "port not set")
		port = "6000"
	}

	server := http.Server{
		Addr:         fmt.Sprintf(":%s", port),
		Handler:      handlers.CORS(originsOk, methodsOk)(r),
		WriteTimeout: 60 * time.Second,
		ReadTimeout:  60 * time.Second,
	}

	logger.Log("event", "server is listening", "port", port)
	if err := server.ListenAndServe(); err != nil {
		logger.Log("event", "exiting", "err", err)
		os.Exit(1)
	}

}

func newLogContext(logger log.Logger, app string) log.Logger {
	return log.With(logger,
		"time", log.DefaultTimestampUTC,
		"app", app,
	)
}

// pkgCaller wraps a stack.Call to make the default string output include the
// package path.
type pkgCaller struct {
	c stack.Call
}