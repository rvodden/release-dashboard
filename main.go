//go:generate go run -tags=dev generateClient.go
package main

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"log"
	"os"
	"os/signal"
	"path/filepath"
	"release-dashboard/app/route"
	"release-dashboard/app/server"
	"runtime"
	//"app/shared/database"
	//"app/shared/email"
	//"app/shared/jsonconfig"
	//"app/shared/recaptcha"
	//"app/shared/server"
	//"app/shared/session"
	//"app/shared/view"
	//"app/shared/view/plugin"
)

func init() {
	log.SetFlags(log.Lshortfile | log.Lmicroseconds | log.Ldate | log.LUTC)
	runtime.GOMAXPROCS(runtime.NumCPU())
}

func main() {
	// Load the configuration file
	load("config"+string(os.PathSeparator)+"config.json", config)

	// Start the server
	go func() {
		server.Run(*route.LoadHTTP(), *route.LoadHTTPS(), config.ServerConfiguration)
	}()

	channel := make(chan os.Signal, 1)
	// We'll accept graceful shutdowns when quit via SIGINT (Ctrl+C)
	// SIGKILL, SIGQUIT or SIGTERM (Ctrl+/) will not be caught.
	signal.Notify(channel, os.Interrupt)
	<-channel

	// Implement graceful shutdown here

	log.Print("Shutting Down")
}

/**
 * Application Settings
 */

// config the settings variable
var config = &configuration{}

// configuration contains the application settings
type configuration struct {
	ServerConfiguration server.ServerConfiguration `json:"Server"`
}

// ParseJSON unmarshals bytes to structs
func (configuration *configuration) ParseJSON(bytes []byte) error {
	return json.Unmarshal(bytes, &configuration)
}

// Parser must implement ParseJSON
type Parser interface {
	ParseJSON([]byte) error
}

// Load the JSON config file
func load(configFile string, p Parser) {
	var err error
	var absPath string
	var input = io.ReadCloser(os.Stdin)
	if absPath, err = filepath.Abs(configFile); err != nil {
		log.Fatalln(err)
	}

	if input, err = os.Open(absPath); err != nil {
		log.Fatalln(err)
	}

	// Read the config file
	jsonBytes, err := ioutil.ReadAll(input)
	input.Close()
	if err != nil {
		log.Fatalln(err)
	}

	// Parse the config
	if err := p.ParseJSON(jsonBytes); err != nil {
		log.Fatalln("Could not parse %q: %v", configFile, err)
	}
}
