package main

import (
	"crypto/tls"
	"fmt"
	"net/http"
	"utils/utils"
)

func main() {
	http.HandleFunc("/", handler)

	tlsConfig := &tls.Config{
		ClientCAs:  utils.LoadCA("./cert/ca.crt"),
		ClientAuth: tls.RequestClientCert, // for Chrome
		// ClientAuth: tls.RequestClientCert, // for client go
	}

	// Enable https server
	server := &http.Server{
		Addr:      ":8080",
		TLSConfig: tlsConfig,
	}
	fmt.Println("HTTPS server listening on :8080")
	err := server.ListenAndServeTLS("./cert/server.crt", "./cert/server.key")
	if err != nil {
		fmt.Println("Error starting HTTPS server:", err)
	}
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, HTTPS service!")
}
