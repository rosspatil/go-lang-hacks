package main

import (
	"context"
	"fmt"
	"os"

	"cloud.google.com/go/datastore"
)

type User struct {
	ID   string `json:"ID", datastore:"Id"`
	Name string `json:"name" datastore:"name"`
}

func main() {
	os.Setenv("GOOGLE_APPLICATION_CREDENTIALS", "/etc/config/prod/goodcop.json")
	client, err := datastore.NewClient(context.Background(), "goodcop-staging")
	if err != nil {
		panic(err)
	}
	// k := &datastore.Key{
	// 	Kind: "sample",
	// 	Name: "123",
	// }
	// user := User{Name: "Roshan"}
	// _, err = client.Put(context.Background(), k, &user)
	// if err != nil {
	// 	panic(err)
	// }
	users := []User{}
	q := datastore.NewQuery("sample")

	q = q.Filter("__key__=", datastore.NameKey("sample", "123", nil))

	keys, err := client.GetAll(context.Background(), q, &users)
	if err != nil {
		panic(err)
	}
	fmt.Println(keys)

	fmt.Println("------")
	for i := 0; i < len(keys); i++ {
		users[i].ID = keys[i].Name
	}
	fmt.Println(users)
}
