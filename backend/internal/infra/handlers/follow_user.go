package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"blog0/internal/services"
)

// FollowUser godoc
// @Summary      Follow user
// @Description  Follow an author (requires authentication)
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        author_id path     string true "Author ID"
// @Success      200       {object} services.FollowUserResp
// @Failure      400       {object} ErrorResp
// @Failure      401       {object} ErrorResp
// @Failure      404       {object} ErrorResp
// @Failure      500       {object} ErrorResp
// @Router       /api/v1/users/{author_id}/follow [post]
func FollowUser(followUser *services.FollowUser) gin.HandlerFunc {
	return func(c *gin.Context) {
		authorID := c.Param("author_id")
		if authorID == "" {
			c.JSON(http.StatusBadRequest, ErrorResp{Error: "author_id is required"})
			return
		}

		userID, exists := c.Get("user_id")
		if !exists {
			c.JSON(http.StatusUnauthorized, ErrorResp{Error: "user not authenticated"})
			return
		}

		userIDStr := userID.(string)
		if userIDStr == authorID {
			c.JSON(http.StatusBadRequest, ErrorResp{Error: "cannot follow yourself"})
			return
		}

		req := &services.FollowUserReq{
			AuthorID: authorID,
			UserID:   userIDStr,
		}

		resp, err := followUser.Exec(c, req)
		if err != nil {
			c.JSON(http.StatusInternalServerError, ErrorResp{Error: err.Error()})
			return
		}

		c.JSON(http.StatusOK, resp)
	}
}

// UnfollowUser godoc
// @Summary      Unfollow user
// @Description  Unfollow an author (requires authentication)
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        author_id path     string true "Author ID"
// @Success      200       {object} services.UnfollowUserResp
// @Failure      400       {object} ErrorResp
// @Failure      401       {object} ErrorResp
// @Failure      404       {object} ErrorResp
// @Failure      500       {object} ErrorResp
// @Router       /api/v1/users/{author_id}/follow [delete]
func UnfollowUser(unfollowUser *services.UnfollowUser) gin.HandlerFunc {
	return func(c *gin.Context) {
		authorID := c.Param("author_id")
		if authorID == "" {
			c.JSON(http.StatusBadRequest, ErrorResp{Error: "author_id is required"})
			return
		}

		userID, exists := c.Get("user_id")
		if !exists {
			c.JSON(http.StatusUnauthorized, ErrorResp{Error: "user not authenticated"})
			return
		}

		userIDStr := userID.(string)
		if userIDStr == authorID {
			c.JSON(http.StatusBadRequest, ErrorResp{Error: "cannot unfollow yourself"})
			return
		}

		req := &services.UnfollowUserReq{
			AuthorID: authorID,
			UserID:   userIDStr,
		}

		resp, err := unfollowUser.Exec(c, req)
		if err != nil {
			c.JSON(http.StatusInternalServerError, ErrorResp{Error: err.Error()})
			return
		}

		c.JSON(http.StatusOK, resp)
	}
}