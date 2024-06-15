package routers

import (
	controllers "service-auth/Controllers"
	user_data_layer "service-auth/DataLayers/User"
	initializers "service-auth/Initializers"
	repositories "service-auth/Repositories"
	auth_services "service-auth/Services/AuthServices"

	"github.com/gin-gonic/gin"
)

func AuthRoutes(routes *gin.RouterGroup) {
	var userDataLayer = user_data_layer.NewUserDataLayer(initializers.DB)
	var userRepository = repositories.NewIntanceUserDataLayer(userDataLayer)
	var authService = auth_services.NewAuthService(userRepository)
	var authController = controllers.NewAuthController(authService)

	authGroup := routes.Group("/auth")
	{
		authGroup.POST("/sign-in")
		authGroup.POST("/sign-up", authController.SignUpController)
		authGroup.POST("/facebook")
		authGroup.POST("/google")

		authCallBackGroup := authGroup.Group("/callback")
		{
			authCallBackGroup.GET("/facebook")
			authCallBackGroup.GET("/google")
		}
	}
}
