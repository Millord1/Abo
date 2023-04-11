package entities

import (
	"context"
	"fmt"
	"log"

	"github.com/edgedb/edgedb-go"
)

type Asset struct {
	id          edgedb.OptionalUUID
	AssetId     string     `edgedb:"assetId" form:"assetId"`
	Name        string     `edgedb:"name" form:"name"`
	ImgUrl      string     `edgedb:"imgUrl" form:"imgUrl"`
	Metadata    string     `edgedb:"metadata" form:"metadata"`
	Description string     `edgedb:"description" form:"description"`
	Contracts   []Contract `edhedb:"contracts"`
}

func GetAllAssets() []Asset {
	ctx := context.Background()
	client, err := edgedb.CreateClient(ctx, edgedb.Options{})
	if err != nil {
		log.Fatal(err)
	}
	defer client.Close()

	var assets []Asset
	query := "select Asset{assetId, name, imgUrl, metadata, description}"
	err = client.Query(ctx, query, &assets)
	if err != nil {
		log.Fatalln(err)
		fmt.Println(err)
	}
	return assets
}

func GetAsset(assetId string) Asset {
	ctx := context.Background()
	client, err := edgedb.CreateClient(ctx, edgedb.Options{})
	if err != nil {
		log.Fatal(err)
	}
	defer client.Close()

	var asset Asset
	query := fmt.Sprintf("select Asset{assetId, name, imgUrl, metadata, description} filter Asset.assetId = %s", assetId)
	err = client.Query(ctx, query, &asset)
	if err != nil {
		log.Fatalln(err)
		fmt.Println(err)
	}
	return asset
}

func GetAssetsByContract(contractId string) []Asset {
	ctx := context.Background()
	client, err := edgedb.CreateClient(ctx, edgedb.Options{})
	if err != nil {
		log.Fatal(err)
	}
	defer client.Close()

	var assets []Asset
	query := fmt.Sprintf(
		"select Asset{assetId, name, imgUrl, metadata, description, contracts{address, blockchain, standard} filter address=%s} ", contractId)
	err = client.Query(ctx, query, &assets)
	if err != nil {
		log.Fatalln(err)
		fmt.Println(err)
	}
	return assets
}

func PushAssets(assets *[]Asset) []Asset {
	ctx := context.Background()
	client, err := edgedb.CreateClient(ctx, edgedb.Options{})
	if err != nil {
		log.Fatal(err)
	}
	defer client.Close()

	var query string
	for _, asset := range *assets {
		query += fmt.Sprintf("insert Asset{assetId:='%s',name:='%s',imgUrl:='%s',metadata:='%s',description:='%s'}unless conflict on .assetId;", asset.AssetId, asset.Name, asset.ImgUrl, asset.Metadata, asset.Description)
	}

	var result []Asset
	if len(*assets) > 1 {
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

func BindAssetToCollection(assetId string, collectionId string) error {
	ctx := context.Background()
	client, err := edgedb.CreateClient(ctx, edgedb.Options{})
	if err != nil {
		log.Fatal(err)
	}
	defer client.Close()

	query := fmt.Sprintf("update Asset filter .assetId='%s' set collection=(select distinct Collection filter .collectionId='%s')", assetId, collectionId)

	var result []Asset
	err = client.Query(ctx, query, &result)

	if err != nil {
		fmt.Println(err)
	}

	return err
}

// func BindAsset(subType Type, assetId string, contractId string, collectionId string) Asset {
// 	ctx := context.Background()
// 	client, err := edgedb.CreateClient(ctx, edgedb.Options{})
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	defer client.Close()

// }
