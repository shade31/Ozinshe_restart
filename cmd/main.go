package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/username/GitRepoName/internal/controller"
	"github.com/username/GitRepoName/pkg"
)

func main() {
	app, err := pkg.App()

	if err != nil {
		log.Fatal(err)
	}
	defer app.CloseDBConnection()

	ginRouter := gin.Default()
	
	controller.Setup(app, ginRouter)

	ginRouter.Run(fmt.Sprintf(":%s", app.Env.PORT))
}
