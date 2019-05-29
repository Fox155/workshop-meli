package main

import (
	"net/http"
	"os"
	"workshop-meli/config"
	"workshop-meli/controllers"

	"github.com/gin-gonic/gin"
)

var router = gin.Default()

func main() {
	routes()
	// var db = db.DBPurchases{}
	// for k, val := range db.GetAll() {
	// 	fmt.Printf("Key: %v - Value: %#v \n", k, val)
	// }

	if port := os.Getenv("PORT"); port != "" {
		router.Run(":" + port)
	} else {
		router.Run(":8080")
	}
}

func routes() {
	//Purchases
	router.POST("/purchases", onlyAdmin, controllers.CreatePurchase)
	router.GET("/purchases", controllers.GetPurchases)
	router.GET("/purchases/:id", controllers.ReadPurchases)
	router.PUT("/purchases/:id", onlyAdmin, controllers.UpdatePurchase)
	router.DELETE("/purchases/:id", onlyAdmin, controllers.DeletePurchase)

	//Usuarios
	router.POST("/users", controllers.AltaUsuario)
	router.GET("/users", controllers.MostrarUsuarios)
	router.GET("/users/:id", controllers.MostrarUsuario)
	router.PUT("/users/:id", controllers.ActualizarUsuario)
	router.DELETE("/users/:id", controllers.EliminarUsuario)
}

func checkQueryParams(c *gin.Context) {
	if userID := c.Query("user_id"); userID == "" {
		c.AbortWithStatusJSON(http.StatusBadRequest, "Query params user_id required.")
	}
}

func onlyAdmin(c *gin.Context) {
	if role := c.GetHeader("role"); role != "admin" {
		c.AbortWithStatusJSON(http.StatusUnauthorized, "Unauthorized!")
	}
}

func challenge(c *gin.Context) {
	if config.IsProduction() {
		c.AbortWithStatusJSON(http.StatusMethodNotAllowed, "Only gets allowed, for challenge purpose! - visit github.com/seansa/Workshop_go-UTNFRT")
		return
	}
}
