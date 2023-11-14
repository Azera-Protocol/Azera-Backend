package main

import (
	"azera-backend/controllers"
	"azera-backend/initializer"
	"azera-backend/middleware"
	"github.com/gin-gonic/gin"
)

func init() {
	initializer.ConnectToDb()
	initializer.SyncDatabase()
}

func main() {
	r := gin.Default()
	r.Use(middleware.Logger())

	r.GET("/users/:walletID", controllers.GetUserByWalletID)
	r.POST("/users/create", controllers.CreateUserDetails)
	r.PUT("/users/edit/:walletID", controllers.UpdateUserSocials)
	r.GET("/users/customizations/:walletID", controllers.GetCustomizationByWalletID)
	r.POST("/users/customization/create/:walletID", controllers.CreateCustomizations)
	r.PUT("/users/customization/edit/:walletID", controllers.UpdateCustomizations)
	r.Run()
}
