package main

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

var(
	server *gin.Engine
)
func init(){
	server = gin.Default()
}
func main() {
	fmt.Println("Welcome to 3iOj")
	err := godotenv.Load()
	if err != nil{
		log.Fatal("Error loading .env file")
	}
	server.SetTrustedProxies(nil)
	// basepath := server.Group("/api")
	log.Fatal(server.Run(":8080"))
}
