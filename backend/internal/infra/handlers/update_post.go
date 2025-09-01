package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"blog0/internal/services"
)

type UpdatePostReq struct {
	Title       string `json:"title"`
	Slug        string `json:"slug"`
	RawMarkdown string `json:"raw_markdown"`
	Publish     *bool  `json:"publish"`
}

// UpdatePost godoc
// @Summary      Update a post
// @Description  Update an existing post (requires authentication and ownership)
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        slug path     string        true "Post slug"
// @Param        body body     UpdatePostReq true "Post data"
// @Success      200  {object} services.UpdatePostResp
// @Failure      400  {object} ErrorResp
// @Failure      401  {object} ErrorResp
// @Failure      403  {object} ErrorResp
// @Failure      404  {object} ErrorResp
// @Failure      500  {object} ErrorResp
// @Router       /api/v1/me/posts/{slug} [put]
func UpdatePost(updatePost *services.UpdatePost) gin.HandlerFunc {
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

		var body UpdatePostReq
		if err := c.ShouldBindJSON(&body); err != nil {
			c.JSON(http.StatusBadRequest, ErrorResp{Error: err.Error()})
			return
		}

		req := &services.UpdatePostReq{
			Slug:        slug,
			Title:       body.Title,
			NewSlug:     body.Slug,
			RawMarkdown: body.RawMarkdown,
			UserID:      userID.(string),
			Publish:     body.Publish,
		}

		resp, err := updatePost.Exec(c, req)
		if err != nil {
			if err.Error() == "unauthorized: you can only update your own posts" {
				c.JSON(http.StatusForbidden, ErrorResp{Error: err.Error()})
				return
			}
			c.JSON(http.StatusInternalServerError, ErrorResp{Error: err.Error()})
			return
		}

		c.JSON(http.StatusOK, resp)
	}
}