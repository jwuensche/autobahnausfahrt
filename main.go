package main

import (
	"net/http"
	"os"
	"os/signal"

	"github.com/jwuensche/autobahnausfahrt/handler"
	"github.com/jwuensche/autobahnausfahrt/util"
)

func main() {
	util.Init()
	util.Log.Info("Starting autobahnausfahrt")
	util.GetEnvVar()
	runHTTPEndPoint()

	// Wait for SIGINT (CTRL-c)
	shutdown := make(chan os.Signal, 1)
	signal.Notify(shutdown, os.Interrupt)
	<-shutdown
	util.Log.Info("SIGINT received, terminating")
	util.Log.Info("Goodbye")
}

func runTLSEndPoint() {
	util.Log.Info("Starting TLS Endpoint")
	util.Log.Debugf("Route is %s", util.Conf.Route)
	http.HandleFunc(util.Conf.Route, handler.Render)
	go http.ListenAndServeTLS(";"+util.Conf.Port, util.Conf.TLSCert, util.Conf.TLSKey, nil)
	util.Log.Infof("HTTP Enpoint listening at port %s and route %s", util.Conf.Port, util.Conf.Route)
	return
}

func runHTTPEndPoint() {
	util.Log.Info("Starting HTTP Endpoint")
	util.Log.Debugf("Route is %s", util.Conf.Route)
	http.HandleFunc(util.Conf.Route, handler.Render)
	go http.ListenAndServe(":"+util.Conf.Port, nil)
	util.Log.Infof("HTTP Enpoint listening at port %s and route %s", util.Conf.Port, util.Conf.Route)
	return
}
