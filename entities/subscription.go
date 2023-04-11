package entities

import (
	"context"
	"fmt"
	"log"

	"github.com/edgedb/edgedb-go"
)

type Subscription struct {
	id         edgedb.OptionalUUID
	Name       string `edgedb:"name"`
	Blockchain string `edgedb:"blockchain"`
	LastMint   string `edgedb:"lastMint"`
	Types      []Type `edgedb:"types"`
}

func GetAllSubscriptions() []Subscription {
	ctx := context.Background()
	client, err := edgedb.CreateClient(ctx, edgedb.Options{})
	if err != nil {
		log.Fatal(err)
	}
	defer client.Close()

	var subscriptions []Subscription
	query := "select Subscription{name, blockchain, lastMint, types{name}}"
	err = client.Query(ctx, query, &subscriptions)
	if err != nil {
		log.Fatalln(err)
		fmt.Println(err)
	}
	return subscriptions
}

func GetSubscription(name string) Subscription {
	ctx := context.Background()
	client, err := edgedb.CreateClient(ctx, edgedb.Options{})
	if err != nil {
		log.Fatal(err)
	}
	defer client.Close()

	var subscription Subscription
	query := fmt.Sprintf("select Subscription{name, blockchain, lastMint, types{name}}filter name =%s", name)
	err = client.Query(ctx, query, &subscription)
	if err != nil {
		log.Fatalln(err)
		fmt.Println(err)
	}
	return subscription
}

func CreateSubscription(name string, blockchain string) Subscription {
	ctx := context.Background()
	client, err := edgedb.CreateClient(ctx, edgedb.Options{})
	if err != nil {
		log.Fatal(err)
	}
	defer client.Close()

	query := fmt.Sprintf("with allTypes := (select Type{name}) insert Subscription name='%s', blockchain='%s', lastMint='never', types:={allTypes}}", name, blockchain)
	var result Subscription
	err = client.Query(ctx, query, &result)

	if err != nil {
		fmt.Println(err)
	}
	return result
}
