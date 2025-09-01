package services

import (
	"context"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"

	"blog0/internal/domain/dao"
)

type ListMyPosts struct {
	postDAO dao.PostDAO
	userDAO dao.UserDAO
}

type ListMyPostsReq struct {
	Page    int
	PerPage int
	Order   string
	UserID  string
}

type MyPostItem struct {
	ID          string     `json:"id"`
	Title       string     `json:"title"`
	Slug        string     `json:"slug"`
	PublishedAt *time.Time `json:"published_at"`
	CreatedAt   time.Time  `json:"created_at"`
	UpdatedAt   time.Time  `json:"updated_at"`
	Status      string     `json:"status"`
}

type ListMyPostsResp struct {
	Page    int          `json:"page"`
	PerPage int          `json:"per_page"`
	Total   int          `json:"total"`
	Items   []MyPostItem `json:"items"`
}

func NewListMyPosts(postDAO dao.PostDAO, userDAO dao.UserDAO) *ListMyPosts {
	return &ListMyPosts{
		postDAO: postDAO,
		userDAO: userDAO,
	}
}

func (s *ListMyPosts) Exec(ctx context.Context, req *ListMyPostsReq) (*ListMyPostsResp, error) {
	limit := req.PerPage
	offset := (req.Page - 1) * req.PerPage

	posts, err := s.postDAO.FindPaginated(ctx, limit, offset, "author_id = $1", "created_at "+req.Order, req.UserID)
	if err != nil {
		return nil, err
	}

	totalPosts, err := s.postDAO.Count(ctx, "author_id = $1", req.UserID)
	if err != nil {
		return nil, err
	}

	items := make([]MyPostItem, 0)
	for _, post := range posts {
		status := "draft"
		if post.PublishedAt != nil {
			status = "published"
		}

		items = append(items, MyPostItem{
			ID:          post.ID,
			Title:       post.Title,
			Slug:        post.Slug,
			PublishedAt: post.PublishedAt,
			CreatedAt:   post.CreatedAt,
			UpdatedAt:   post.UpdatedAt,
			Status:      status,
		})
	}

	return &ListMyPostsResp{
		Page:    req.Page,
		PerPage: req.PerPage,
		Total:   int(totalPosts),
		Items:   items,
	}, nil
}

func (s *ListMyPosts) ParseRequest(c *gin.Context, userID string) (*ListMyPostsReq, error) {
	page := 1
	if p := c.Query("page"); p != "" {
		if parsed, err := strconv.Atoi(p); err == nil && parsed > 0 {
			page = parsed
		}
	}

	perPage := 20
	if pp := c.Query("per_page"); pp != "" {
		if parsed, err := strconv.Atoi(pp); err == nil && parsed > 0 && parsed <= 100 {
			perPage = parsed
		}
	}

	order := "DESC"
	if o := c.Query("order"); o == "ASC" || o == "DESC" {
		order = o
	}

	return &ListMyPostsReq{
		Page:    page,
		PerPage: perPage,
		Order:   order,
		UserID:  userID,
	}, nil
}