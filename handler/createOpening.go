package handler

import (
	"exemple/gooportunities/schemas"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateOpeningHandler(ctx *gin.Context) {
	request := CreateOpeningRequest{}
	ctx.BindJSON(&request)

	if err := request.validate(); err != nil {
		logger.Errorf("Validation error: %v", err.Error())
		sendError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	opening := schemas.Opening{
		Role:     request.Role,
		Company:  request.Company,
		Location: request.Location,
		Remote:   *request.Remote,
		Link:     request.Link,
		Salaray:  request.Salaray,
	}

	if err := db.Create(&opening); err != nil {
		logger.Errorf("error creating Opening: %v", err)
		return
	}

	sendSuccess(ctx, "create-opening", opening)
}
