package services

import (
	"context"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"

	"blog0/internal/domain/dao"
)

type ListPosts struct {
	postDAO     dao.PostDAO
	userDAO     dao.UserDAO
	postLikeDAO dao.PostLikeDAO
	commentDAO  dao.CommentDAO
}

type ListPostsReq struct {
	Page    int
	PerPage int
	Order   string
}

type PostItem struct {
	Title        string    `json:"title"`
	Author       string    `json:"author"`
	AuthorID     string    `json:"author_id"`
	PublishedAt  time.Time `json:"published_at"`
	Slug         string    `json:"slug"`
	LikeCount    int       `json:"like_count"`
	CommentCount int       `json:"comment_count"`
}

type ListPostsResp struct {
	Page    int        `json:"page"`
	PerPage int        `json:"per_page"`
	Total   int        `json:"total"`
	Items   []PostItem `json:"items"`
}

func NewListPosts(postDAO dao.PostDAO, userDAO dao.UserDAO, postLikeDAO dao.PostLikeDAO, commentDAO dao.CommentDAO) *ListPosts {
	return &ListPosts{
		postDAO:     postDAO,
		userDAO:     userDAO,
		postLikeDAO: postLikeDAO,
		commentDAO:  commentDAO,
	}
}

func (s *ListPosts) Exec(ctx context.Context, req *ListPostsReq) (*ListPostsResp, error) {
	limit := req.PerPage
	offset := (req.Page - 1) * req.PerPage

	posts, err := s.postDAO.FindPaginated(ctx, limit, offset, "", "published_at "+req.Order)
	if err != nil {
		return nil, err
	}

	totalPosts, err := s.postDAO.Count(ctx, "")
	if err != nil {
		return nil, err
	}

	if len(posts) == 0 {
		return &ListPostsResp{
			Page:    req.Page,
			PerPage: req.PerPage,
			Total:   int(totalPosts),
			Items:   make([]PostItem, 0),
		}, nil
	}

	authorsIDs := make([]any, 0)
	placeholders := make([]string, 0)
	for i, post := range posts {
		authorsIDs = append(authorsIDs, post.AuthorID)
		placeholders = append(placeholders, fmt.Sprintf("$%d", i+1))
	}

	authors, err := s.userDAO.FindAll(ctx, "id IN ("+strings.Join(placeholders, ",")+")", "", authorsIDs...)
	if err != nil {
		return nil, err
	}

	authorsMap := make(map[string]*dao.User)
	for _, author := range authors {
		authorsMap[author.ID] = author
	}

	// Get like counts for all posts
	likeCounts := make(map[string]int)
	for _, post := range posts {
		count, err := s.postLikeDAO.Count(ctx, "post_id = $1", post.ID)
		if err != nil {
			return nil, err
		}
		likeCounts[post.ID] = int(count)
	}

	// Get comment counts for all posts
	commentCounts := make(map[string]int)
	for _, post := range posts {
		count, err := s.commentDAO.Count(ctx, "post_id = $1", post.ID)
		if err != nil {
			return nil, err
		}
		commentCounts[post.ID] = int(count)
	}

	items := make([]PostItem, 0)
	for _, post := range posts {
		author, ok := authorsMap[post.AuthorID]
		if !ok {
			return nil, fmt.Errorf("author for post %s not found", post.ID)
		}

		items = append(items, PostItem{
			Title:        post.Title,
			Author:       author.Username,
			AuthorID:     author.ID,
			PublishedAt:  *post.PublishedAt,
			Slug:         post.Slug,
			LikeCount:    likeCounts[post.ID],
			CommentCount: commentCounts[post.ID],
		})
	}

	return &ListPostsResp{
		Page:    req.Page,
		PerPage: req.PerPage,
		Total:   int(totalPosts),
		Items:   items,
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

	order := "desc"
	if o := c.Query("order"); o != "" {
		order = o
	}

	return &ListPostsReq{
		Page:    page,
		PerPage: perPage,
		Order:   order,
	}, nil
}
