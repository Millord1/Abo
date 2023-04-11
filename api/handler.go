package api

import (
	"encoding/json"
	"log"
	"net/http"
	"shabo_edge/entities"

	"github.com/gin-gonic/gin"
)

// Assets
func GetAllAssets(c *gin.Context) {
	allAssets := entities.GetAllAssets()
	result, err := json.Marshal(allAssets)
	if err != nil {
		log.Fatalln(err)
	}
	c.String(http.StatusOK, string(result))
}

func GetAsset(c *gin.Context) {
	assetId := c.Query("assetId")
	if len(assetId) == 0 {
		log.Fatalln("Arg assetId not found")
	}

	asset := entities.GetAsset(assetId)
	result, err := json.Marshal(asset)
	if err != nil {
		log.Fatalln(err)
	}
	c.String(http.StatusOK, string(result))
}

func AddAsset(c *gin.Context) {
	var asset entities.Asset
	err := c.ShouldBind(&asset)
	if err != nil {
		log.Println(asset.AssetId)
	}

	var assets []entities.Asset
	assets = append(assets, asset)
	result := entities.PushAssets(&assets)

	if len(result) == 0 {
		c.String(http.StatusConflict, "Already existing")
	}

	c.String(http.StatusOK, "Success")
}

// Collections
func GetAllCollections(c *gin.Context) {
	allColls := entities.GetAllCollections()
	result, err := json.Marshal(allColls)
	if err != nil {
		log.Fatalln(err)
	}
	c.String(http.StatusOK, string(result))
}

func GetCollection(c *gin.Context) {
	collectionId := c.Query("collectionId")
	if len(collectionId) == 0 {
		log.Fatalln("Arg collectionId not found")
	}

	coll := entities.GetCollection(collectionId)
	result, err := json.Marshal(coll)
	if err != nil {
		log.Fatalln(err)
	}
	c.String(http.StatusOK, string(result))
}

func AddCollection(c *gin.Context) {
	var coll entities.Collection
	err := c.ShouldBind(&coll)
	if err != nil {
		log.Println(coll.CollectionId)
	}

	var collections []entities.Collection
	collections = append(collections, coll)
	result := entities.PushCollection(&collections)

	if len(result) == 0 {
		c.String(http.StatusConflict, "Already existing")
	}

	c.String(http.StatusOK, "Success")
}

// Bind asset and collection
func BindAssetToColl(c *gin.Context) {
	collectionId := c.Param("collectionId")
	assetId := c.Param("assetId")

	err := entities.BindAssetToCollection(assetId, collectionId)

	if err != nil {
		c.String(http.StatusBadRequest, "No bind")
	}

	c.String(http.StatusOK, "Success")
}

// Contracts
func GetAllContracts(c *gin.Context) {
	allCtts := entities.GetAllContracts()
	result, err := json.Marshal(allCtts)
	if err != nil {
		log.Fatalln(err)
	}
	c.String(http.StatusOK, string(result))
}

func GetContract(c *gin.Context) {
	address := c.Query("address")
	if len(address) == 0 {
		log.Fatalln("Arg address not found")
	}

	ctt := entities.GetCollection(address)
	result, err := json.Marshal(ctt)
	if err != nil {
		log.Fatalln(err)
	}
	c.String(http.StatusOK, string(result))
}

func AddContract(c *gin.Context) {
	var ctt entities.Contract
	err := c.ShouldBind(&ctt)
	if err != nil {
		log.Println(ctt.Address)
	}

	var contracts []entities.Contract
	contracts = append(contracts, ctt)
	result := entities.PushContracts(&contracts)

	if len(result) == 0 {
		c.String(http.StatusConflict, "Already existing")
	}

	c.String(http.StatusOK, "Success")
}

// Subscriptions
func GetAllSubscriptions(c *gin.Context) {
	allSubs := entities.GetAllSubscriptions()
	result, err := json.Marshal(allSubs)
	if err != nil {
		log.Fatalln(err)
	}
	c.String(http.StatusOK, string(result))
}

func GetSubscription(c *gin.Context) {
	name := c.Query("name")
	if len(name) == 0 {
		log.Fatalln("Arg name not found")
	}

	subscriptions := entities.GetSubscription(name)
	result, err := json.Marshal(subscriptions)
	if err != nil {
		log.Fatalln(err)
	}
	c.String(http.StatusOK, string(result))
}
