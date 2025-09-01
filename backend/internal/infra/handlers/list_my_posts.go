package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"blog0/internal/services"
)

// ListMyPosts godoc
// @Summary      List my posts
// @Description  List all posts created by the authenticated user with pagination and ordering
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        page     query    int    false  "Page number" default(1)
// @Param        per_page query    int    false  "Items per page" default(20)
// @Param        order    query    string false  "Order by" Enums(ASC, DESC) default(DESC)
// @Success      200      {object} services.ListMyPostsResp
// @Failure      400      {object} ErrorResp
// @Failure      401      {object} ErrorResp
// @Failure      500      {object} ErrorResp
// @Router       /api/v1/me/posts [get]
func ListMyPosts(listMyPosts *services.ListMyPosts) gin.HandlerFunc {
	return func(c *gin.Context) {
		userID, exists := c.Get("user_id")
		if !exists {
			c.JSON(http.StatusUnauthorized, ErrorResp{Error: "user not authenticated"})
			return
		}

		req, err := listMyPosts.ParseRequest(c, userID.(string))
		if err != nil {
			c.JSON(http.StatusBadRequest, ErrorResp{Error: err.Error()})
			return
		}

		resp, err := listMyPosts.Exec(c, req)
		if err != nil {
			c.JSON(http.StatusInternalServerError, ErrorResp{Error: err.Error()})
			return
		}

		c.JSON(http.StatusOK, resp)
	}
}