package routes

import (
    "github.com/JosephAntony37900/ArquitecturaHexagonal/users/infrastructure/controllers"
    "github.com/gin-gonic/gin"
)

func SetupUserRoutes(r *gin.Engine, createUserControlle *controllers.CreateUserController, getUserControlle *controllers.GetUsersController, deleteUserControlle *controllers.DeleteUserController, updateUserController *controllers.UpdateUserController) {
	//rutas
	r.POST("/users", createUserControlle.Handle)
	r.GET("/users", getUserControlle.Handle)
	r.DELETE("/delete/users/:id", deleteUserControlle.Handle)
	r.PUT("/users/:id", updateUserController.Handle)
}