package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"blog0/internal/services"
)

// GetProfile godoc
// @Summary      Get user profile
// @Description  Get user's following, bookmarks, and liked posts (requires authentication)
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Success      200 {object} services.GetProfileResp
// @Failure      401 {object} ErrorResp
// @Failure      500 {object} ErrorResp
// @Router       /api/v1/me/profile [get]
func GetProfile(getProfile *services.GetProfile) gin.HandlerFunc {
	return func(c *gin.Context) {
		userID, exists := c.Get("user_id")
		if !exists {
			c.JSON(http.StatusUnauthorized, ErrorResp{Error: "user not authenticated"})
			return
		}

		userIDStr := userID.(string)

		req := &services.GetProfileReq{
			UserID: userIDStr,
		}

		resp, err := getProfile.Exec(c, req)
		if err != nil {
			c.JSON(http.StatusInternalServerError, ErrorResp{Error: err.Error()})
			return
		}

		c.JSON(http.StatusOK, resp)
	}
}
