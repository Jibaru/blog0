package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"blog0/internal/services"
)

// DeletePost godoc
// @Summary      Delete a post
// @Description  Delete an existing post (requires authentication and ownership)
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        slug path     string true "Post slug"
// @Success      200  {object} services.DeletePostResp
// @Failure      400  {object} ErrorResp
// @Failure      401  {object} ErrorResp
// @Failure      403  {object} ErrorResp
// @Failure      404  {object} ErrorResp
// @Failure      500  {object} ErrorResp
// @Router       /api/v1/me/posts/{slug} [delete]
func DeletePost(deletePost *services.DeletePost) gin.HandlerFunc {
	return func(c *gin.Context) {
		slug := c.Param("slug")
		if slug == "" {
			c.JSON(http.StatusBadRequest, ErrorResp{Error: "slug is required"})
			return
		}

		userID, exists := c.Get("user_id")
		if !exists {
			c.JSON(http.StatusUnauthorized, ErrorResp{Error: "user not authenticated"})
			return
		}

		req := &services.DeletePostReq{
			Slug:   slug,
			UserID: userID.(string),
		}

		resp, err := deletePost.Exec(c, req)
		if err != nil {
			if err.Error() == "unauthorized: you can only delete your own posts" {
				c.JSON(http.StatusForbidden, ErrorResp{Error: err.Error()})
				return
			}
			c.JSON(http.StatusInternalServerError, ErrorResp{Error: err.Error()})
			return
		}

		c.JSON(http.StatusOK, resp)
	}
}