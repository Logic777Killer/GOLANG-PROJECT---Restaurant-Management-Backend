package routes

import (
	controller "Restaurant_Management_Backend/controllers"
	"github.com/gin-gonic/gin"
)

func OrderItemRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.GET("/orderItems", controller.GetOrderItems())
	incomingRoutes.GET("/orderItems/:orderItem_id", controller.GetOrderItem())
	incomingRoutes.GET("/orderItems-order/:order_id", controller.GetOrderItemsByOrder())
	incomingRoutes.POST("/orderItems", controller.CreateOrderItems())
	incomingRoutes.PATCH("/orderItems:orderItem_id", controller.UpdateOrderItems())
}
