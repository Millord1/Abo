package entities

import (
	"context"
	"fmt"
	"log"

	"github.com/edgedb/edgedb-go"
)

type Collection struct {
	id           edgedb.OptionalUUID
	CollectionId string `edgedb:"collectionId" form:"collectionId"`
	Name         string `edgedb:"name" form:"name"`
	ImgUrl       string `edgedb:"imgUrl" form:"imgUrl"`
	Description  string `edgedb:"descritpion" form:"description"`
}

func GetAllCollections() []Collection {
	ctx := context.Background()
	client, err := edgedb.CreateClient(ctx, edgedb.Options{})
	if err != nil {
		log.Fatal(err)
	}
	defer client.Close()

	var colls []Collection
	query := "select Collection{collectionId, name, imgUrl, description}"
	err = client.Query(ctx, query, &colls)
	if err != nil {
		log.Fatalln(err)
		fmt.Println(err)
	}
	return colls
}

func GetCollection(collectionId string) Collection {
	ctx := context.Background()
	client, err := edgedb.CreateClient(ctx, edgedb.Options{})
	if err != nil {
		log.Fatal(err)
	}
	defer client.Close()

	var coll Collection
	query := fmt.Sprintf("select Collection{collectionId, name, imgUrl, description} filter Collection.collectionId='%s'", collectionId)
	err = client.Query(ctx, query, &coll)
	if err != nil {
		log.Fatalln(err)
		fmt.Println(err)
	}
	return coll
}

func PushCollection(colls *[]Collection) []Collection {
	ctx := context.Background()
	client, err := edgedb.CreateClient(ctx, edgedb.Options{})
	if err != nil {
		log.Fatal(err)
	}
	defer client.Close()

	var query string
	for _, coll := range *colls {
		query += fmt.Sprintf("insert Collection{collectionId:='%s',name:='%s',imgUrl:='%s',description:='%s'}unless conflict on .collectionId;", coll.CollectionId, coll.Name, coll.ImgUrl, coll.Description)
	}

	var result []Collection
	if len(*colls) > 1 {
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
