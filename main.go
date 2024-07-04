package main

import (
	"time"

	route "trial/api/route"
	"trial/bootstrap"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {

	app := bootstrap.App()

	env := app.Env
	r := gin.Default()

	db := app.Mongo.Database(env.DBName)
	defer app.CloseDBConnection()

	cld := app.Cloudinary

	timeout := time.Duration(env.ContextTimeout) * time.Second

	gin := gin.Default()
	// Global CORS configuration
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"*"}
	config.AllowMethods = []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"}
	config.AllowHeaders = []string{"Origin", "Authorization", "Content-Type"}
	gin.Use(cors.New(config))

	route.Setup(env, timeout, db, cld, gin)

	gin.Run()
}
