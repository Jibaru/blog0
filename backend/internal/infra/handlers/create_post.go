package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"blog0/internal/services"
)

type CreatePostReq struct {
	Title       string `json:"title" binding:"required"`
	Slug        string `json:"slug" binding:"required"`
	RawMarkdown string `json:"raw_markdown" binding:"required"`
	Publish     bool   `json:"publish"`
}

// CreatePost godoc
// @Summary      Create a new post
// @Description  Create a new post (requires authentication)
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        body body     CreatePostReq true "Post data"
// @Success      201  {object} services.CreatePostResp
// @Failure      400  {object} ErrorResp
// @Failure      401  {object} ErrorResp
// @Failure      500  {object} ErrorResp
// @Router       /api/v1/me/posts [post]
func CreatePost(createPost *services.CreatePost) gin.HandlerFunc {
	return func(c *gin.Context) {
		userID, exists := c.Get("user_id")
		if !exists {
			c.JSON(http.StatusUnauthorized, ErrorResp{Error: "user not authenticated"})
			return
		}

		var body CreatePostReq
		if err := c.ShouldBindJSON(&body); err != nil {
			c.JSON(http.StatusBadRequest, ErrorResp{Error: err.Error()})
			return
		}

		req := &services.CreatePostReq{
			Title:       body.Title,
			Slug:        body.Slug,
			RawMarkdown: body.RawMarkdown,
			UserID:      userID.(string),
			Publish:     body.Publish,
		}

		resp, err := createPost.Exec(c, req)
		if err != nil {
			c.JSON(http.StatusInternalServerError, ErrorResp{Error: err.Error()})
			return
		}

		c.JSON(http.StatusCreated, resp)
	}
}