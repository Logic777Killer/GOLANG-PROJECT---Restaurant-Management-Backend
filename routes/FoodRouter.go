package routes

import (
	controller "Restaurant_Management_Backend/controllers"
	"github.com/gin-gonic/gin"
)

func FoodRouter(incomingRoutes *gin.Engine) {
	incomingRoutes.GET("/foods", controller.GetFoods())
	incomingRoutes.GET("/foods/:food_id", controller.GetFoods())
	incomingRoutes.POST("/foods", controller.CreateFood())
	incomingRoutes.PATCH("/foods:food_id", controller.UpdateFood())
}
