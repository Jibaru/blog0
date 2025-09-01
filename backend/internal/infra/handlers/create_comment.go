package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"blog0/internal/services"
)

type CreateCommentReq struct {
	Body     string  `json:"body" binding:"required"`
	ParentID *string `json:"parent_id"`
}

// CreateComment godoc
// @Summary      Create comment on post
// @Description  Create a new comment on a post (requires authentication)
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        slug path     string           true "Post slug"
// @Param        body body     CreateCommentReq true "Comment data"
// @Success      201  {object} services.CreateCommentResp
// @Failure      400  {object} ErrorResp
// @Failure      401  {object} ErrorResp
// @Failure      404  {object} ErrorResp
// @Failure      500  {object} ErrorResp
// @Router       /api/v1/posts/{slug}/comments [post]
func CreateComment(createComment *services.CreateComment) gin.HandlerFunc {
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

		var body CreateCommentReq
		if err := c.ShouldBindJSON(&body); err != nil {
			c.JSON(http.StatusBadRequest, ErrorResp{Error: err.Error()})
			return
		}

		req := &services.CreateCommentReq{
			Slug:     slug,
			Body:     body.Body,
			ParentID: body.ParentID,
			UserID:   userID.(string),
		}

		resp, err := createComment.Exec(c, req)
		if err != nil {
			c.JSON(http.StatusInternalServerError, ErrorResp{Error: err.Error()})
			return
		}

		c.JSON(http.StatusCreated, resp)
	}
}