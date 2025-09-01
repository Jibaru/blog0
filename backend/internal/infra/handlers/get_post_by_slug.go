package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"blog0/internal/services"
)

// GetPostBySlug godoc
// @Summary      Get post by slug
// @Description  Get post details by slug including comments
// @Accept       json
// @Produce      json
// @Param        slug path     string true "Post slug"
// @Success      200  {object} services.GetPostBySlugResp
// @Failure      404  {object} ErrorResp
// @Failure      500  {object} ErrorResp
// @Router       /api/v1/posts/{slug} [get]
func GetPostBySlug(getPostBySlug *services.GetPostBySlug) gin.HandlerFunc {
	return func(c *gin.Context) {
		slug := c.Param("slug")
		if slug == "" {
			c.JSON(http.StatusBadRequest, ErrorResp{Error: "slug is required"})
			return
		}

		req := &services.GetPostBySlugReq{
			Slug: slug,
		}

		resp, err := getPostBySlug.Exec(c, req)
		if err != nil {
			c.JSON(http.StatusInternalServerError, ErrorResp{Error: err.Error()})
			return
		}

		c.JSON(http.StatusOK, resp)
	}
}