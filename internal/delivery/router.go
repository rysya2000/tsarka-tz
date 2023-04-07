package delivery

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func NewRouter(handler *gin.Engine, r *Handler) {

	handler.Use(gin.Recovery())

	handler.GET("/health", func(c *gin.Context) { c.Status(http.StatusOK) })

	h := handler.Group("/rest")
	{
		h.POST("/substr/find", r.SubStrHandler)
		h.POST("/email/check", r.EmailHandler)
		h.GET("/self/find/:str", r.SelfHandler)
	}
	
}

