package session

import (
	"context"
	"fmt"
	"log"

	firebase "firebase.google.com/go"
	"firebase.google.com/go/auth"
)

type FirebaseAuth struct {
	client *auth.Client
}

func NewFirebaseAuth() *FirebaseAuth {
	ctx := context.Background()

	app, err := firebase.NewApp(context.Background(), nil)

	if err != nil {
		log.Panic(fmt.Errorf("error initializing firebase app: %v", err))
	}

	client, err := app.Auth(ctx)

	if err != nil {
		log.Fatalf("error getting firebase auth client: %v\n", err)
	}

	return &FirebaseAuth{
		client: client,
	}
}

func (auth *FirebaseAuth) ValidateIdToken(idToken string) bool {
	ctx := context.Background()
	if auth.client == nil {
		log.Fatalf("error not setting auth client")
		return false
	}

	token, err := auth.client.VerifyIDToken(ctx, idToken)

	if err != nil {
		log.Fatalf("error verifing Id Token: %v\n", err)
		return false
	}

	log.Printf("Verified ID token: %v\n", token)

	return true
}
