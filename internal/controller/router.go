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

	ageController := &core.AgeController{
		AgeRepository: repository.NewAgeRepository(db),
	}

	screenshotController := &core.ScreenshotController{
		ScreenshotRepository: repository.NewScreenshotRepository(db),
	}

	contentController := &core.ContentController{
		ContentRepository: repository.NewContentRepository(db),
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

		coreRouter.POST("/ages", ageController.CreateAge)
		coreRouter.GET("/ages", ageController.GetAllAges)
		coreRouter.GET("/ages/:ageID", ageController.GetAgeByID)
		coreRouter.PATCH("/ages/:ageID", ageController.UpdateAge)
		coreRouter.DELETE("/ages/:ageID", ageController.DeleteAge)

		coreRouter.POST("/screenshot", screenshotController.CreateScreenshot)
		coreRouter.GET("/screenshots/:screenshotID", screenshotController.GetScreenshotByID)
		coreRouter.DELETE("/screenshots/:screenshotID", screenshotController.DeleteScreenshot)

		coreRouter.POST("/content", contentController.CreateContent)
		coreRouter.PUT("/content/:contentID", contentController.UpdateContent)
		coreRouter.DELETE("/content/:contentID", contentController.DeleteContent)
		coreRouter.GET("/content/:contentID", contentController.GetContentByID)
		coreRouter.GET("/content/byGenre/:genreID", contentController.GetContentByGenre)
		coreRouter.GET("/content/byTitle/:title", contentController.GetContentByTitle)
	}

}
