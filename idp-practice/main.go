package main

import (
	"context"
	"log"
	"os"

	firebase "firebase.google.com/go/v4"
	"firebase.google.com/go/v4/auth"
)

func main() {
	ctx := context.Background()

	gcpProjectID := os.Getenv("GCP_PROJECT_ID")

	app, err := firebase.NewApp(ctx, &firebase.Config{
		ProjectID: gcpProjectID,
	})
	if err != nil {
		log.Fatalf("error initializing app: %v\n", err)
	}

	client, err := app.Auth(ctx)
	if err != nil {
		log.Fatalf("error getting Auth client: %v\n", err)
	}

	users, err := client.GetUsers(ctx, []auth.UserIdentifier{
		auth.EmailIdentifier{Email: "ride.or.die.2215@icloud.com"},
	})
	if err != nil {
		log.Fatalf("error getting users: %+v\n", err)
	}

	log.Printf("Successfully fetched user data")
	for _, u := range users.Users {
		log.Printf("%+v", u.UID)
	}
}
