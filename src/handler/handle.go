package handler

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	"github.com/terassyi/cyberrange-manager/src/database"
	"net/http"
)

type Handler struct {
	*gin.Engine
	DB *database.DB
}

func New(router *gin.Engine, db *sql.DB) *Handler {
	return &Handler{
		Engine: router,
		DB:     database.New(db),
	}
}

func (h *Handler) Index(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", nil)
}
