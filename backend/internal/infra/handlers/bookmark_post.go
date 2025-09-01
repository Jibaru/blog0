package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"blog0/internal/services"
)

// BookmarkPost godoc
// @Summary      Bookmark post
// @Description  Add post to user's bookmarks (requires authentication)
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        slug path     string true "Post slug"
// @Success      201  {object} services.BookmarkPostResp
// @Failure      401  {object} ErrorResp
// @Failure      404  {object} ErrorResp
// @Failure      409  {object} ErrorResp
// @Failure      500  {object} ErrorResp
// @Router       /api/v1/posts/{slug}/bookmarks [post]
func BookmarkPost(bookmarkPost *services.BookmarkPost) gin.HandlerFunc {
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

		req := &services.BookmarkPostReq{
			Slug:   slug,
			UserID: userID.(string),
		}

		resp, err := bookmarkPost.Exec(c, req)
		if err != nil {
			c.JSON(http.StatusInternalServerError, ErrorResp{Error: err.Error()})
			return
		}

		c.JSON(http.StatusCreated, resp)
	}
}

// UnbookmarkPost godoc
// @Summary      Remove bookmark
// @Description  Remove post from user's bookmarks (requires authentication)
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        slug path     string true "Post slug"
// @Success      200  {object} services.UnbookmarkPostResp
// @Failure      401  {object} ErrorResp
// @Failure      404  {object} ErrorResp
// @Failure      500  {object} ErrorResp
// @Router       /api/v1/posts/{slug}/bookmarks [delete]
func UnbookmarkPost(unbookmarkPost *services.UnbookmarkPost) gin.HandlerFunc {
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

		req := &services.UnbookmarkPostReq{
			Slug:   slug,
			UserID: userID.(string),
		}

		resp, err := unbookmarkPost.Exec(c, req)
		if err != nil {
			c.JSON(http.StatusInternalServerError, ErrorResp{Error: err.Error()})
			return
		}

		c.JSON(http.StatusOK, resp)
	}
}