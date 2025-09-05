package services

import (
	"context"
	"fmt"
	"strings"
	"time"

	"blog0/internal/domain/dao"
)

type GetPostBySlug struct {
	postDAO     dao.PostDAO
	userDAO     dao.UserDAO
	commentDAO  dao.CommentDAO
	postLikeDAO dao.PostLikeDAO
}

type GetPostBySlugReq struct {
	Slug string
}

type AuthorInfo struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type CommentInfo struct {
	ID        string     `json:"id"`
	Author    AuthorInfo `json:"author"`
	ParentID  *string    `json:"parent_id"`
	Body      string     `json:"body"`
	CreatedAt time.Time  `json:"created_at"`
}

type GetPostBySlugResp struct {
	ID                  string        `json:"id"`
	Title               string        `json:"title"`
	Summary             string        `json:"summary"`
	Tags                []string      `json:"tags"`
	Slug                string        `json:"slug"`
	RawMarkdown         string        `json:"raw_markdown"`
	Author              AuthorInfo    `json:"author"`
	PublishedAt         time.Time     `json:"published_at"`
	LikesCount          int           `json:"likes_count"`
	Comments            []CommentInfo `json:"comments"`
	RawMarkdownAudioURL *string       `json:"raw_markdown_audio_url"`
	SummaryAudioURL     *string       `json:"summary_audio_url"`
}

func NewGetPostBySlug(postDAO dao.PostDAO, userDAO dao.UserDAO, commentDAO dao.CommentDAO, postLikeDAO dao.PostLikeDAO) *GetPostBySlug {
	return &GetPostBySlug{
		postDAO:     postDAO,
		userDAO:     userDAO,
		commentDAO:  commentDAO,
		postLikeDAO: postLikeDAO,
	}
}

func (s *GetPostBySlug) Exec(ctx context.Context, req *GetPostBySlugReq) (*GetPostBySlugResp, error) {
	post, err := s.postDAO.FindOne(ctx, "slug = $1", "", req.Slug)
	if err != nil {
		return nil, fmt.Errorf("post not found: %w", err)
	}

	author, err := s.userDAO.FindByPk(ctx, post.AuthorID)
	if err != nil {
		return nil, fmt.Errorf("author not found: %w", err)
	}

	comments, err := s.commentDAO.FindAll(ctx, "post_id = $1", "created_at ASC", post.ID)
	if err != nil {
		return nil, fmt.Errorf("failed to load comments: %w", err)
	}

	commentAuthorsIDs := make([]any, 0)
	placeholders := make([]string, 0)
	for i, comment := range comments {
		commentAuthorsIDs = append(commentAuthorsIDs, comment.AuthorID)
		placeholders = append(placeholders, fmt.Sprintf("$%d", i+1))
	}

	var commentAuthors []*dao.User
	if len(commentAuthorsIDs) > 0 {
		commentAuthors, err = s.userDAO.FindAll(ctx, "id IN ("+strings.Join(placeholders, ",")+")", "", commentAuthorsIDs...)
		if err != nil {
			return nil, fmt.Errorf("failed to load comment authors: %w", err)
		}
	}

	commentAuthorsMap := make(map[string]*dao.User)
	for _, author := range commentAuthors {
		commentAuthorsMap[author.ID] = author
	}

	likesCount, err := s.postLikeDAO.Count(ctx, "post_id = $1", post.ID)
	if err != nil {
		return nil, fmt.Errorf("failed to count likes: %w", err)
	}

	commentInfos := make([]CommentInfo, 0)
	for _, comment := range comments {
		commentAuthor, ok := commentAuthorsMap[comment.AuthorID]
		if !ok {
			return nil, fmt.Errorf("comment author %s not found", comment.AuthorID)
		}

		commentInfos = append(commentInfos, CommentInfo{
			ID:        comment.ID,
			Author:    AuthorInfo{ID: commentAuthor.ID, Name: commentAuthor.Username},
			ParentID:  comment.ParentID,
			Body:      comment.Body,
			CreatedAt: comment.CreatedAt,
		})
	}

	return &GetPostBySlugResp{
		ID:                  post.ID,
		Title:               post.Title,
		Summary:             post.Summary,
		Tags:                post.ItsTags(),
		Slug:                post.Slug,
		RawMarkdown:         post.RawMarkdown,
		Author:              AuthorInfo{ID: author.ID, Name: author.Username},
		PublishedAt:         *post.PublishedAt,
		LikesCount:          int(likesCount),
		Comments:            commentInfos,
		RawMarkdownAudioURL: post.RawMarkdownAudioURL,
		SummaryAudioURL:     post.SummaryAudioURL,
	}, nil
}
