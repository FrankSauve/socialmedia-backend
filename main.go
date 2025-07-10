package main

import (
	"franksauve/socialmedia/db"
	"franksauve/socialmedia/user"

	"github.com/gin-gonic/gin"
)

func main() {
	db.ConnectPostgres()
	router := gin.Default()
	user.Routes(router)
	router.Run()
}
