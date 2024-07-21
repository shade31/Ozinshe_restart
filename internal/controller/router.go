package controller

import (
	"Ozinshe_restart/middleware"
	"Ozinshe_restart/pkg"
	"os"

	"github.com/gin-gonic/gin"

	"Ozinshe_restart/internal/controller/auth"
	"Ozinshe_restart/internal/controller/core"
	"Ozinshe_restart/internal/controller/user"
	"Ozinshe_restart/internal/repository"
)

var AccessTokenSecret string

const AccessTokenExpiryHour = 24

func Setup(app pkg.Application, router *gin.Engine) {
	AccessTokenSecret := os.Getenv("ACCESS_TOKEN_SECRET")

	db := app.DB

	loginController := &auth.AuthController{
		UserRepository:        repository.NewUserRepository(db),
		AccessTokenSecret:     AccessTokenSecret,
		AccessTokenExpiryHour: AccessTokenExpiryHour,
	}

	userController := &user.UserController{
		UserRepository: repository.NewUserRepository(db),
	}

	genreController := &core.GenreController{
		GenreRepository: repository.NewGenreRepository(db),
	}

	router.POST("/signup", loginController.Signup)
	router.POST("/signin", loginController.Signin)

	router.Use(middleware.JWTAuth(AccessTokenSecret))

	userRouter := router.Group("/user")
	{
		userRouter.GET("/profile", userController.GetProfile)
		userRouter.PATCH("/updateProfile", userController.UpdateProfile)
		userRouter.PATCH("/changePassword", userController.ChangePassword)
	}

	coreRouter := router.Group("/core")
	{
		coreRouter.POST("/genres", genreController.CreateGenre)
		coreRouter.GET("/genres", genreController.GetAllGenres)
		coreRouter.GET("/genres/:genreID", genreController.GetGenreByID)
		coreRouter.PATCH("/genres/:genreID", genreController.UpdateGenre)
		coreRouter.DELETE("/genres/:genreID", genreController.DeleteGenre)
	}

}
