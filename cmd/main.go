package main

import (
	"fmt"
	"log"
	"os"

	"Ozinshe_restart/internal/controller"
	"Ozinshe_restart/pkg"

	"github.com/gin-gonic/gin"
)

func main() {
	app, err := pkg.App()

	if err != nil {
		log.Fatal(err)
	}
	db, err := pkg.ConnectDatabase()
	if err != nil {
		panic(err)
	}
	defer db.Close()

	ginRouter := gin.Default()

	controller.Setup(app, ginRouter)

	PORT := os.Getenv("PORT")
	ginRouter.Run(fmt.Sprintf(":%s", PORT))
}
