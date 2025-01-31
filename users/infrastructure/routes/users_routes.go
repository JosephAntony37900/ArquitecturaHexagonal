package routes

import (
    "github.com/JosephAntony37900/ArquitecturaHexagonal/users/infrastructure/controllers"
    "github.com/gin-gonic/gin"
)

func SetupUserRoutes(r *gin.Engine, createUserControlle *controllers.CreateUserController) {
	//rutas
	r.POST("/users", createUserControlle.Handle)
}