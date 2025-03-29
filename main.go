package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mariohalucyn/tus/controllers"
)

func main() {
	router := gin.Default()
	handler := controllers.Tus()
	router.Any("/files/*filepath", gin.WrapH(http.StripPrefix("/files", handler)))
	router.Run(":8080")
}
