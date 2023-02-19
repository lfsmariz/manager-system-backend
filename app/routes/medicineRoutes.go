package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/lfsmariz/manager-system-backend/app/handlers"
)

func addMedicineRoutes(rg *gin.RouterGroup) {
	medicine := rg.Group("/medicine")

	a := handlers.NewMedicineHandler()

	medicine.GET("/", a.GetAll)
	medicine.POST("/", a.Save)
	medicine.OPTIONS("/")
}
