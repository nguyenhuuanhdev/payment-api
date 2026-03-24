package main

import (
	"payment-api/controllers"

	"github.com/gin-gonic/gin"
)
//ngay1
func main() {
	r := gin.Default()

	r.POST("/pay", controllers.CreatePayment)
	r.GET("/pay/:id", controllers.Pay)
	r.GET("/order/:id", controllers.GetOrder)
	r.PUT("/order/:id", controllers.UpdatePayment)
	r.DELETE("/order/:id", controllers.DeleteOrder)
	r.Run(":8080")
}
