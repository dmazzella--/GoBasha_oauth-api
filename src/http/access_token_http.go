package http

import (
	"github.com/dmazzella--/GoBasha_oauth-api/src/domain/access_token"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

type AccessTokenHandler interface {
	GetById(ctx *gin.Context)
}

type accessTokenHandler struct {
	service access_token.Service
}

func NewHandler(service access_token.Service) AccessTokenHandler {
	return &accessTokenHandler{
		service: service,
	}
}

func (a *accessTokenHandler) GetById(ctx *gin.Context) {
	accessToken, err := a.service.GetById(strings.TrimSpace(ctx.Param("access_token_id")))
	if err != nil {
		ctx.JSON(err.Status, err)
		return
	}
	ctx.JSON(http.StatusOK, accessToken)
}
