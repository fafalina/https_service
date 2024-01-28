package main

import (
	"encoding/json"
	"crypto/tls"
	"fmt"
	"net/http"
	"server/utils"
	"server/db"
)

type Response struct {
	Time	db.TimestampData `json:"timestamp"`
}

func main() {
	http.HandleFunc("/", handler)

	tlsConfig := &tls.Config{
		ClientCAs:  utils.LoadCA("../cert/ca.crt"),
		ClientAuth: tls.RequestClientCert,
	}

	// Enable https server
	server := &http.Server{
		Addr:      ":8080",
		TLSConfig: tlsConfig,
	}
	fmt.Println("HTTPS server listening on :8080")
	err := server.ListenAndServeTLS("../cert/server.crt", "../cert/server.key")
	if err != nil {
		fmt.Println("Error starting HTTPS server:", err)
	}
}

func handler(w http.ResponseWriter, r *http.Request) {
	db.Init()
	db.WriteCurrentTime()
    time := db.ReadTime()

    response := Response{
        Time:    time,
    }

    jsonBytes, err := json.Marshal(response)
    if err != nil {
        http.Error(w, "Error encoding JSON", http.StatusInternalServerError)
        return
    }

    w.Header().Set("Access-Control-Allow-Origin", "*")
    w.Header().Set("Content-Type", "application/json")
    w.Write(jsonBytes)
}
