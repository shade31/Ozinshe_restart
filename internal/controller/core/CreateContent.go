package core

import (
	"Ozinshe_restart/internal/models"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ContentController struct {
	ContentRepository models.ContentRepository
}

func (cc *ContentController) CreateContent(c *gin.Context) {
	var request models.Content
	fmt.Println(request)
	if err := c.ShouldBind(&request); err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{
			Result: []models.ErrorDetail{
				{
					Code:    "ERROR_BIND_JSON",
					Message: "Datas dont match with struct of CreateContent",
				},
			},
		})
		return
	}

	contentID, err := cc.ContentRepository.CreateContent(c, request)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{
			Result: []models.ErrorDetail{
				{
					Code:    "ERROR_CREATE_CONTENT",
					Message: "Couldn't create content",
					Metadata: models.Properties{
						Properties1: err.Error(),
					},
				},
			},
		})
		return
	}

	c.JSON(http.StatusOK, models.SuccessResponse{Result: contentID})
}
