package main

import (
	"fmt"
	"net/http"
	"log"
)

func main() {
	http.HandleFunc("/", handler)
	
	// 憑證和私鑰的檔案路徑，請替換成你自己的路徑
	certFile := "./cert3/certificate.crt"
	keyFile := "./cert3/private.key"

	fmt.Println("Server is running on https://localhost:8080")

	// 啟動 HTTPS 伺服器
	err := http.ListenAndServeTLS(":8080", certFile, keyFile, nil)
	if err != nil {
		log.Fatal("ListenAndServeTLS: ", err)
	}
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, HTTPS!")
}
