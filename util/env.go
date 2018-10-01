package util

import (
	"os"
)

// Config includes all general information required for the endpoint to be opened
type Config struct {
	// to which port the endpoint should exposed at
	Port string
	// the route the metrics will be exposed at, default for prometheus is `/metric`
	Route string
	// Interconnect specific values
	InterconnectPort string
	// Address for Interconnect, leave empty for localhost
	InterconnectAddress string
	// which scraper will be used to collect data from `autobahnausfahrt`
	Scraper string
	//
	// other things i currently do not think of like tls, websocket etc.
}

// Conf is an instance of the Global Config
var Conf Config

// GetEnvVar initializes the config with all applicable environment variables
func GetEnvVar() {
	Log.Info("Reading in environmennt Variables")
	Conf.Port = os.Getenv("EXPORT_PORT")
	Conf.Route = os.Getenv("EXPORT_ROUTE")
	Conf.InterconnectPort = os.Getenv("INTERCONNECT_PORT")
	Conf.InterconnectAddress = os.Getenv("INTERCONNECT_ADDRESS")
}
