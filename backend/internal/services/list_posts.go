package services

import (
	"context"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

type ListPosts struct{}

type ListPostsReq struct {
	Page    int
	PerPage int
	Order   string
}

type PostItem struct {
	Title       string    `json:"title"`
	Author      string    `json:"author"`
	PublishedAt time.Time `json:"published_at"`
	Slug        string    `json:"slug"`
}

type ListPostsResp struct {
	Page    int        `json:"page"`
	PerPage int        `json:"per_page"`
	Total   int        `json:"total"`
	Items   []PostItem `json:"items"`
}

func NewListPosts() *ListPosts {
	return &ListPosts{}
}

func (s *ListPosts) Exec(ctx context.Context, req *ListPostsReq) (*ListPostsResp, error) {
	return &ListPostsResp{
		Page:    req.Page,
		PerPage: req.PerPage,
		Total:   0,
		Items:   []PostItem{},
	}, nil
}

func (s *ListPosts) ParseRequest(c *gin.Context) (*ListPostsReq, error) {
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

	order := "published_at_desc"
	if o := c.Query("order"); o != "" {
		order = o
	}

	return &ListPostsReq{
		Page:    page,
		PerPage: perPage,
		Order:   order,
	}, nil
}