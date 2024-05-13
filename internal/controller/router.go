package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/username/GitRepoName/middleware"
	"github.com/username/GitRepoName/pkg"

	"github.com/username/GitRepoName/internal/controller/auth"
	"github.com/username/GitRepoName/internal/controller/user"
	"github.com/username/GitRepoName/internal/repository"
)

func Setup(app pkg.Application, router *gin.Engine) {
	db := app.Pql

	loginController := &auth.AuthController{
		UserRepository: repository.NewUserRepository(db),
	}

	userController := &user.UserController{
		UserRepository: repository.NewUserRepository(db),
	}

	router.POST("/signup", loginController.Signup)
	router.POST("/signin", loginController.Signin)

	router.Use(middleware.JWTAuth(env.AccessTokenSecret))

	userRouter := router.Group("/user")
	{
		userRouter.GET("/profile", userController.GetProfile)
	}

}
