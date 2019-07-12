package main

import (
	"context"
	"log"
	"time"

	firebase "firebase.google.com/go"
	"google.golang.org/api/option"
)

func main() {
	// g := gin.New()

	// g.Run(":4444")
	Init()
}

var ctx = context.Background()

func Init() {
	opt := option.WithCredentialsFile("./serviceaccount.json")
	app, err := firebase.NewApp(ctx, nil, opt)
	if err != nil {
		log.Println(err)
		return
	}
	client, err := app.Auth(ctx)
	if err != nil {
		log.Println(err)
		return
	}
	claim := map[string]interface{}{"rule": "abc"}
	token, err := client.CustomTokenWithClaims(ctx, "123456789abc", claim)
	if err != nil {
		log.Println(err)
		return
	}
	// _, err = client.VerifyIDToken(ctx, token)
	// if err != nil {
	// 	log.Println(err)
	// 	return
	// }
	_, err = client.SessionCookie(ctx, token, time.Duration(time.Hour*5))
	if err != nil {
		log.Println(err)
		return
	}
	// fmt.Println(s)
	// fstore, _ := app.Firestore(ctx)
}
