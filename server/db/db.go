package db

import (
	"context"
	"fmt"
	"log"
	"time"
	"firebase.google.com/go/v4"
	"firebase.google.com/go/v4/db"
	"google.golang.org/api/option"
)

type TimestampData struct {
	Timestamp time.Time `json:"timestamp"`
}

var client *db.Client

/*
 * @brief Create a connection with database
 */
func Init() {
	opt := option.WithCredentialsFile("../cert/test-f0fe8-firebase-adminsdk-b1yqd-52792a7480.json")
    config := &firebase.Config{
        DatabaseURL: "https://test-f0fe8-default-rtdb.asia-southeast1.firebasedatabase.app/",
    }
    
	app, err := firebase.NewApp(context.Background(), config, opt)
	if err != nil {
		log.Fatalln("Error initializing app: %v", err)
	}

	client, err = app.Database(context.Background())
	if err != nil {
		log.Fatalln("Error initializing Firebase database client: %v", err)
	}
}

/*
 * @brief Read timestamp from database
 */
func ReadTime() TimestampData {
	messageRef := client.NewRef("/timestamp")
	var time TimestampData
	if err := messageRef.Get(context.Background(), &time); err != nil {
		log.Fatalln("Error reading message from database: %v", err)
	}

	fmt.Println("Timestamp:", time.Timestamp)
	return time
}

/*
 * @brief Write timestamp to database
 */
func WriteCurrentTime() {
	currentTime := time.Now()

	data := TimestampData{
		Timestamp: currentTime,
	}

	messageRef := client.NewRef("/timestamp")
	if err := messageRef.Set(context.Background(), &data); err != nil {
		log.Fatalln("Error writing timestamp to database: %v", err)
	}

	fmt.Println("Timestamp written to Firebase:", currentTime)
}
