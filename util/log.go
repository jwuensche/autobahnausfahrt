package util

import (
	"os"

	logging "github.com/op/go-logging"
)

// Log for the autobahnausfahrt
var Log = logging.MustGetLogger("autobahnkreuz")
var format = logging.MustStringFormatter(
	`%{color}[%{level:-8s}] %{time:15:04:05.000} %{longpkg}@%{shortfile}%{color:reset} -- %{message}`,
)

// Init sets all required information for go-logging to be as pretty as possible
func Init() {

	backend := logging.NewLogBackend(os.Stderr, "", 0)

	backendFormatter := logging.NewBackendFormatter(backend, format)

	backendLeveled := logging.AddModuleLevel(backend)
	backendLeveled.SetLevel(logging.ERROR, "")

	logging.SetBackend(backendLeveled, backendFormatter)
}
