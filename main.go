package main

import (
	"context"
	"log"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	
	db "gin/database"
	handler "gin/handlers"
	services "gin/services"
	
)

var rdb *redis.Client
var ctx = context.Background()

func init_services() {
	rdb = db.Redis_init()

	services.SetClient(rdb)

}

func setupRouter() *gin.Engine {
	router := gin.Default()

	// Load HTML templates
	router.LoadHTMLGlob("templates/*")

	// Define the route using the handler from another file
	router.GET("/", handler.Home)
	router.GET("/about", handler.About)
	router.GET("/login", handler.Login)
	router.GET("/signup", handler.SignUp)

	router.NoRoute(handler.NotFound404)

	return router
}

func main() {

	// //this is  a comment
	// router := setupRouter()

	// // Start serving the application
	// println("localhost:8080")
	// router.Run(":8080") // listens on :8080 by default

	init_services()

	err := rdb.Set(ctx, "key_name", "value_of Rajveer", 0).Err()
	if err != nil {
		log.Fatal(err)
	}

	// val, err := rdb.Get(ctx, "key_name").Result()
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Println("key:", val)

	s, err := services.FetchData("key_name")
	println(s)


}
