package main

import (
	"shabo_edge/api"
	"shabo_edge/controllers"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.Use(cors.Default())

	public := router.Group("/api")

	// assets
	public.GET("asset/get/all", api.GetAllAssets)
	public.GET("asset/get/:assetId", api.GetAsset)
	public.POST("asset/add", api.AddAsset)

	// collections
	public.GET("collection/get/all", api.GetAllCollections)
	public.GET("collection/get/:collectionId", api.GetCollection)
	public.POST("collection/add", api.AddCollection)

	// Bind asset and collection
	public.POST("asset/bind/:assetId/:collectionId", api.BindAssetToColl)

	public.POST("/register", controllers.Register)
	public.POST("/login", controllers.Login)

	router.Run(":8080")
}
