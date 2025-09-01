package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"blog0/internal/services"
)

// GetAuthorInfo godoc
// @Summary      Get author information
// @Description  Get public information about an author
// @Accept       json
// @Produce      json
// @Param        author_id path     string true "Author ID"
// @Success      200       {object} services.GetAuthorInfoResp
// @Failure      404       {object} ErrorResp
// @Failure      500       {object} ErrorResp
// @Router       /api/v1/users/{author_id} [get]
func GetAuthorInfo(getAuthorInfo *services.GetAuthorInfo) gin.HandlerFunc {
	return func(c *gin.Context) {
		authorID := c.Param("author_id")
		if authorID == "" {
			c.JSON(http.StatusBadRequest, ErrorResp{Error: "author_id is required"})
			return
		}

		req := &services.GetAuthorInfoReq{
			AuthorID: authorID,
		}

		resp, err := getAuthorInfo.Exec(c, req)
		if err != nil {
			c.JSON(http.StatusInternalServerError, ErrorResp{Error: err.Error()})
			return
		}

		c.JSON(http.StatusOK, resp)
	}
}