package entities

import (
	"context"
	"fmt"
	"log"

	"github.com/edgedb/edgedb-go"
)

type Type struct {
	id     edgedb.OptionalUUID
	Name   string  `edgedb:"name"`
	Assets []Asset `edgedb:"assets"`
}

func GetAllTypes() []Type {
	ctx := context.Background()
	client, err := edgedb.CreateClient(ctx, edgedb.Options{})
	if err != nil {
		log.Fatal(err)
	}
	defer client.Close()

	var types []Type
	query := "select Type{name}"
	err = client.Query(ctx, query, &types)
	if err != nil {
		log.Fatalln(err)
		fmt.Println(err)
	}
	return types
}

func GetType(name string) Type {
	ctx := context.Background()
	client, err := edgedb.CreateClient(ctx, edgedb.Options{})
	if err != nil {
		log.Fatal(err)
	}
	defer client.Close()

	var aboType Type
	query := fmt.Sprintf("select distinct Type{name} filter Type.name = '%s'", name)
	err = client.Query(ctx, query, &aboType)
	if err != nil {
		log.Fatalln(err)
		fmt.Println(err)
	}
	return aboType
}

func FillTypes() bool {
	ctx := context.Background()
	client, err := edgedb.CreateClient(ctx, edgedb.Options{})
	if err != nil {
		log.Fatal(err)
	}
	defer client.Close()

	var query string
	var allTypes = [3]string{"ogType", "curatorType", "essentialType"}
	for _, strType := range allTypes {
		query += fmt.Sprintf("insert Type{name='%s'}unless conflict on .name else(select Type{name});", strType)
	}

	err = client.Tx(ctx, func(ctx context.Context, tx *edgedb.Tx) error {
		e := tx.Execute(ctx, query)
		return e
	})

	if err != nil {
		log.Fatalln(err)
	}
	return true
}
