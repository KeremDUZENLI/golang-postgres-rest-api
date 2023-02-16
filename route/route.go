package route

import (
	"net/http"
	"validation/configs"
	"validation/model"

	"github.com/gin-gonic/gin"

	"validation/controller"
)

func URL() {
	// config.go
	configs.DatabaseDB()
	configs.Database.AutoMigrate(&model.MyStruct{})

	// router.go
	ginRouter := gin.Default()

	// controller.go
	ginRouter.GET("/begin", controller.BeginDatabase)
	ginRouter.POST("/create", controller.CreateDatabase)
	ginRouter.GET("/read", controller.ReadDatabase)
	ginRouter.PUT("/update/:id", controller.UpdateDatabase)
	ginRouter.DELETE("/delete/:id", controller.DeleteDatabase)

	// adress
	http.ListenAndServe(":6000", ginRouter)
}
