package utils

import (
	"context"
	"fmt"
	"firebase.google.com/go/v4"
	"google.golang.org/api/option"
)

type Data struct {
	Message string `json:"message"`
	Value   int    `json:"value"`
}

func ReadData() Data {
	opt := option.WithCredentialsFile("./cert/test-f0fe8-firebase-adminsdk-b1yqd-52792a7480.json")
    config := &firebase.Config{
        DatabaseURL: "https://test-f0fe8-default-rtdb.asia-southeast1.firebasedatabase.app/",
    }
    
	app, err := firebase.NewApp(context.Background(), config, opt)
	if err != nil {
		log.Fatalln("Error initializing app: %v", err)
	}

	// 初始化 Firebase Realtime Database 客戶端
	client, err := app.Database(context.Background())
	if err != nil {
		log.Fatalln("Error initializing Firebase database client: %v", err)
	}

	// 讀取資料庫中的訊息
	messageRef := client.NewRef("/")
	var data Data
	if err := messageRef.Get(context.Background(), &data); err != nil {
		log.Fatalln("Error reading message from database: %v", err)
	}

	// 顯示訊息
	fmt.Println("Message:", data.Message)
	fmt.Println("Value:", data.Value)
	return data
}
