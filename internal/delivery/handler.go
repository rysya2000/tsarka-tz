package delivery

import (
	"github.com/gin-gonic/gin"
	"rest/internal/usecase"
)
type Handler struct {
	rest usecase.Rester
}

func New(rest usecase.Rester) *Handler {
	return &Handler{
		rest: rest,
	}
}

func (h *Handler) SubStrHandler(c *gin.Context) {

}

func (h *Handler) EmailHandler(c *gin.Context) {

}

func (h *Handler) SelfHandler(c *gin.Context) {

}

