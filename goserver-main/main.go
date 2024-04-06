package main

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"goserver-main/db"
	"net/http"
)

var (
	ListenAddr = "178.253.32.178:8080"
	RedisAddr  = "178.253.32.178:6379"
)

func main() {

	database, err := db.NewDatabase(RedisAddr)
	if err != nil {
		fmt.Println("Failed to connect to redis %s", err.Error())
	}
	val, err := database.Client.Get(context.TODO(), "name2").Result()
	fmt.Println(val)
	router := initRouter(database)
	router.Run(ListenAddr)
}

func initRouter(database *db.Database) *gin.Engine {
	r := gin.Default()

	r.POST("/points", func(c *gin.Context) {
		var userJson db.User
		if err := c.ShouldBindJSON(&userJson); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		err := database.SaveUser(&userJson)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"user": userJson})
	})

	return r
}
