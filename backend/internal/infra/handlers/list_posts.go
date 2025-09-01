package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"blog0/internal/services"
)

// ListPosts godoc
// @Summary      List all posts
// @Description  List all posts with pagination and ordering
// @Accept       json
// @Produce      json
// @Param        page     query    int    false  "Page number" default(1)
// @Param        per_page query    int    false  "Items per page" default(20)
// @Param        order    query    string false  "Order by" default(published_at_desc)
// @Success      200      {object} services.ListPostsResp
// @Failure      400      {object} ErrorResp
// @Failure      500      {object} ErrorResp
// @Router       /api/v1/posts [get]
func ListPosts(listPosts *services.ListPosts) gin.HandlerFunc {
	return func(c *gin.Context) {
		req, err := listPosts.ParseRequest(c)
		if err != nil {
			c.JSON(http.StatusBadRequest, ErrorResp{Error: err.Error()})
			return
		}

		resp, err := listPosts.Exec(c, req)
		if err != nil {
			c.JSON(http.StatusInternalServerError, ErrorResp{Error: err.Error()})
			return
		}

		c.JSON(http.StatusOK, resp)
	}
}