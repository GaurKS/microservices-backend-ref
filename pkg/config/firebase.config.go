package config

import (
	"context"
	"fmt"
	"log"

	"cloud.google.com/go/firestore"
	cloud "cloud.google.com/go/storage"
	firebase "firebase.google.com/go"
	"google.golang.org/api/option"
)

type Storage struct {
	ctx     context.Context
	client  *firestore.Client
	storage *cloud.Client
}

func (route *Storage) Init() {
	route.ctx = context.Background()
	// config := &firebase.Config{
	// 	DatabaseURL: ,
	// }
	sa := option.WithCredentialsFile("serviceAccountKey.json")

	var err error
	app, err := firebase.NewApp(route.ctx, nil, sa)
	if err != nil {
		fmt.Println(err)
	}

	route.client, err = app.Firestore(route.ctx)
	if err != nil {
		log.Fatalln(err)
	}

	route.storage, err = cloud.NewClient(route.ctx, sa)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println("Successfully connected at port")
}
