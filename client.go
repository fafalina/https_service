package main

import (
	"crypto/tls"
	"fmt"
	"io"
	"os"
	"net/http"
	"https_service/utils"
)

func main() {
	clientCert, err := tls.LoadX509KeyPair("./cert/client.crt", "./cert/client.key")
	if err != nil {
		fmt.Println("LoadX509KeyPair err:", err)
		return
	}

	tlsConfig := &tls.Config{
		RootCAs:      utils.LoadCA("./cert/ca.crt"),
		Certificates: []tls.Certificate{clientCert},
	}

	client := &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: tlsConfig,
		},
	}

	response, err := client.Get("https://localhost:8080")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer response.Body.Close()

	fmt.Println("Status Code:", response.Status)
	io.Copy(os.Stdout, response.Body)
}
