package core

import (
	"Ozinshe_restart/internal/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

type GenreController struct {
	GenreRepository models.GenreRepository
}

func (gc *GenreController) CreateGenre(c *gin.Context) {
	var request models.Genre

	if err := c.ShouldBind(&request); err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{
			Result: []models.ErrorDetail{
				{
					Code:    "ERROR_BIND_JSON",
					Message: "Datas dont match with struct of CreateGenre",
				},
			},
		})
		return
	}

	_, err := gc.GenreRepository.CreateGenre(c, request)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{
			Result: []models.ErrorDetail{
				{
					Code:    "ERROR_CREATE_GENRE",
					Message: "Couldn't create genre",
					Metadata: models.Properties{
						Properties1: err.Error(),
					},
				},
			},
		})
		return
	}

	c.JSON(http.StatusOK, models.SuccessResponse{Result: "Genre created!"})
}
