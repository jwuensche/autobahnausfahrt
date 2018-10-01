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
	// TLS certficate, if specified TLS will always be used
	UseTLS  bool
	TLSCert string
	TLSKey  string
	//
	// other things i currently do not think of like tls, websocket etc.
}

// Conf is an instance of the Global Config
var Conf Config

// GetEnvVar initializes the config with all applicable environment variables
func GetEnvVar() {
	Log.Info("Reading in environment Variables")
	if Conf.Port = os.Getenv("AUSFAHRT_PORT"); Conf.Port == "" {
		Log.Critical("No Port specified, please define AUSFAHRT_PORT")
		os.Exit(1)
	}
	if Conf.Route = os.Getenv("AUSFAHRT_ROUTE"); Conf.Route == "" {
		Log.Critical("No Route specifiy, please define AUSFAHRT_ROUTE")
		os.Exit(1)
	}
	Conf.InterconnectPort = os.Getenv("INTERCONNECT_PORT")
	Conf.InterconnectAddress = os.Getenv("INTERCONNECT_ADDRESS")
	if Conf.TLSCert = os.Getenv("AUSFAHRT_CERT"); Conf.TLSCert != "" {
		Conf.UseTLS = true
		if Conf.TLSKey = os.Getenv("AUSFAHRT_KEY"); Conf.TLSKey != "" {
			Log.Critical("No key specified. please define AUSFAHRT_KEY")
			os.Exit(1)
		}
	}
}
