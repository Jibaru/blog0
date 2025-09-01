package services

import (
	"context"
	"fmt"

	"blog0/internal/domain/dao"
)

type GetAuthorInfo struct {
	userDAO     dao.UserDAO
	postDAO     dao.PostDAO
	postLikeDAO dao.PostLikeDAO
}

type GetAuthorInfoReq struct {
	AuthorID string
}

type TopPostInfo struct {
	ID         string `json:"id"`
	Title      string `json:"title"`
	Slug       string `json:"slug"`
	LikesCount int    `json:"likes_count"`
}

type GetAuthorInfoResp struct {
	ID             string        `json:"id"`
	Name           string        `json:"name"`
	PostsCount     int           `json:"posts_count"`
	FollowersCount int           `json:"followers_count"`
	TopPosts       []TopPostInfo `json:"top_posts"`
}

func NewGetAuthorInfo(userDAO dao.UserDAO, postDAO dao.PostDAO, postLikeDAO dao.PostLikeDAO) *GetAuthorInfo {
	return &GetAuthorInfo{
		userDAO:     userDAO,
		postDAO:     postDAO,
		postLikeDAO: postLikeDAO,
	}
}

func (s *GetAuthorInfo) Exec(ctx context.Context, req *GetAuthorInfoReq) (*GetAuthorInfoResp, error) {
	author, err := s.userDAO.FindByPk(ctx, req.AuthorID)
	if err != nil {
		return nil, fmt.Errorf("author not found: %w", err)
	}

	postsCount, err := s.postDAO.Count(ctx, "author_id = $1 AND published_at IS NOT NULL", req.AuthorID)
	if err != nil {
		return nil, fmt.Errorf("failed to count author posts: %w", err)
	}

	// TODO: Implement followers count when FollowDAO is available
	followersCount := int64(0)

	topPosts, err := s.postDAO.FindPaginated(ctx, 5, 0, "author_id = $1 AND published_at IS NOT NULL", "published_at DESC", req.AuthorID)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch top posts: %w", err)
	}

	topPostInfos := make([]TopPostInfo, 0)
	for _, post := range topPosts {
		likesCount, err := s.postLikeDAO.Count(ctx, "post_id = $1", post.ID)
		if err != nil {
			return nil, fmt.Errorf("failed to count likes for post %s: %w", post.ID, err)
		}

		topPostInfos = append(topPostInfos, TopPostInfo{
			ID:         post.ID,
			Title:      post.Title,
			Slug:       post.Slug,
			LikesCount: int(likesCount),
		})
	}

	return &GetAuthorInfoResp{
		ID:             author.ID,
		Name:           author.Username,
		PostsCount:     int(postsCount),
		FollowersCount: int(followersCount),
		TopPosts:       topPostInfos,
	}, nil
}