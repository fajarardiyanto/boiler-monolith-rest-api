package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"net/http"
	"os"
	"os/signal"
	logger2 "privy/pkg/logger"
	"privy/pkg/parsing"
	"privy/routes"
	"syscall"
)

func main() {
	var log = logrus.New()
	log.Formatter = new(logrus.TextFormatter)
	log.SetReportCaller(parsing.StrToBoolEnv("APPLICATION_TRACE_REPORT"))
	log.Level = logrus.TraceLevel

	log.Info("service started")
	defer log.Info("service ended")

	r := mux.NewRouter()
	{
		newRoutes := routes.NewRoutes()
		newRoutes.Run(r)
	}

	errs := make(chan error)

	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
		errs <- fmt.Errorf("%s", <-c)
	}()

	loggingMiddleware := logger2.LoggingMiddleware(log)
	loggedRouter := loggingMiddleware(r)

	port := ":" + os.Getenv("APPLICATION_PORT")
	server := &http.Server{Addr: port, Handler: loggedRouter}
	go func() {
		log.Info("listing on port ", port)
		errs <- server.ListenAndServe()
	}()

	log.Error("exit ", <-errs)
}
