package entities

import (
	"context"
	"fmt"
	"log"

	"github.com/edgedb/edgedb-go"
)

type Contract struct {
	id         edgedb.OptionalUUID
	Address    string  `edgedb:"address"`
	Standard   string  `edgedb:"standard"`
	Blockchain string  `edgedb:"blockchain"`
	Assets     []Asset `edgedb:"assets"`
}

func GetAllContracts() []Contract {
	ctx := context.Background()
	client, err := edgedb.CreateClient(ctx, edgedb.Options{})
	if err != nil {
		log.Fatal(err)
	}
	defer client.Close()

	var contracts []Contract
	query := "select Contract{address, standard, blockchain}"
	err = client.Query(ctx, query, &contracts)
	if err != nil {
		log.Fatalln(err)
		fmt.Println(err)
	}
	return contracts
}

func GetContract(address string) Contract {
	ctx := context.Background()
	client, err := edgedb.CreateClient(ctx, edgedb.Options{})
	if err != nil {
		log.Fatal(err)
	}
	defer client.Close()

	var contract Contract
	query := fmt.Sprintf("select Contract{address, standard, blockchain} filter .address = %s", address)
	err = client.Query(ctx, query, &contract)
	if err != nil {
		log.Fatalln(err)
		fmt.Println(err)
	}
	return contract
}

func PushContracts(contracts *[]Contract) []Contract {
	ctx := context.Background()
	client, err := edgedb.CreateClient(ctx, edgedb.Options{})
	if err != nil {
		log.Fatal(err)
	}
	defer client.Close()

	var query string
	for _, contract := range *contracts {
		query += fmt.Sprintf("insert Contract{address:='%s',standard:='%s',blockchain:='%s'}unless conflict on .address;", contract.Address, contract.Standard, contract.Blockchain)
	}

	var result []Contract
	if len(*contracts) > 1 {
		err = client.Tx(ctx, func(ctx context.Context, tx *edgedb.Tx) error {
			e := tx.Execute(ctx, query)
			return e
		})
	} else {
		err = client.Query(ctx, query, &result)
	}

	if err != nil {
		log.Fatalln(err)
		fmt.Println(err)
	}

	return result
}
