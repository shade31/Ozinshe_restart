package core

import (
	"Ozinshe_restart/internal/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (cc *ContentController) GetContentByTitle(c *gin.Context) {
	title := c.Param("title")
	// title, err := strconv.Atoi(titleStr)
	// if err != nil {
	// 	c.JSON(http.StatusBadRequest, models.ErrorResponse{
	// 		Result: []models.ErrorDetail{
	// 			{
	// 				Code:    "ERROR_INVALID_TITLE",
	// 				Message: "Invalid title",
	// 				Metadata: models.Properties{
	// 					Properties1: err.Error(),
	// 				},
	// 			},
	// 		},
	// 	})
	// }
	contents, err := cc.ContentRepository.GetContentByTitle(c, title)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{
			Result: []models.ErrorDetail{
				{
					Code:    "ERROR_GET_GENRE",
					Message: "Can't get content from db",
					Metadata: models.Properties{
						Properties1: err.Error(),
					},
				},
			},
		})
		return
	}

	c.JSON(http.StatusOK, models.SuccessResponse{Result: contents})
}
