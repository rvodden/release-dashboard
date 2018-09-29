package server

import (
    "fmt"
    "log"
    "net/http"
)

// Server stores the hostname and port number
type ServerConfiguration struct {
	Hostname  string `json:"Hostname"`  // Server name
	UseHTTP   bool   `json:"UseHTTP"`   // Listen on HTTP
	UseHTTPS  bool   `json:"UseHTTPS"`  // Listen on HTTPS
	HTTPPort  int    `json:"HTTPPort"`  // HTTP port
	HTTPSPort int    `json:"HTTPSPort"` // HTTPS port
	CertFile  string `json:"CertFile"`  // HTTPS certificate
	KeyFile   string `json:"KeyFile"`   // HTTPS private key
}

// Run starts the HTTP and/or HTTPS listener
func Run(httpHandlers http.Handler, httpsHandlers http.Handler, server ServerConfiguration) {
	if server.UseHTTP && server.UseHTTPS {
		go func() {
			startHTTPS(httpsHandlers, server)
		}()
		startHTTP(httpHandlers, server)
	} else if server.UseHTTP {
		startHTTP(httpHandlers, server)
	} else if server.UseHTTPS {
		startHTTPS(httpsHandlers, server)
	} else {
		log.Println("Config file does not specify a listener to start")
	}
}

// startHTTP starts the HTTP listener
func startHTTP(handlers http.Handler, serverConfiguration ServerConfiguration) {
	log.Println("Running HTTP "+httpAddress(serverConfiguration))

	// Start the HTTP listener
	log.Fatal(http.ListenAndServe(httpAddress(serverConfiguration), handlers))
}

// startHTTPs starts the HTTPS listener
func startHTTPS(handlers http.Handler, serverConfiguration ServerConfiguration) {
	log.Println("Running HTTPS "+httpsAddress(serverConfiguration))

	// Start the HTTPS listener
	log.Fatal(http.ListenAndServeTLS(httpsAddress(serverConfiguration), serverConfiguration.CertFile, serverConfiguration.KeyFile, handlers))
}

// httpAddress returns the HTTP address
func httpAddress(serverConfiguration ServerConfiguration) string {
	return serverConfiguration.Hostname + ":" + fmt.Sprintf("%d", serverConfiguration.HTTPPort)
}

// httpsAddress returns the HTTPS address
func httpsAddress(serverConfiguration ServerConfiguration) string {
	return serverConfiguration.Hostname + ":" + fmt.Sprintf("%d", serverConfiguration.HTTPSPort)
}

