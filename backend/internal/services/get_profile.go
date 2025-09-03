package services

import (
	"context"
	"fmt"

	"blog0/internal/domain/dao"
)

type GetProfile struct {
	userDAO     dao.UserDAO
	followDAO   dao.FollowDAO
	bookmarkDAO dao.BookmarkDAO
	postLikeDAO dao.PostLikeDAO
	postDAO     dao.PostDAO
}

type GetProfileReq struct {
	UserID string `json:"-"`
}

type ProfileUser struct {
	ID       string `json:"id"`
	Username string `json:"username"`
}

type ProfilePost struct {
	ID    string `json:"id"`
	Title string `json:"title"`
}

type GetProfileResp struct {
	Following  []ProfileUser `json:"following"`
	Bookmarks  []ProfilePost `json:"bookmarks"`
	LikedPosts []ProfilePost `json:"liked_posts"`
}

func NewGetProfile(userDAO dao.UserDAO, followDAO dao.FollowDAO, bookmarkDAO dao.BookmarkDAO, postLikeDAO dao.PostLikeDAO, postDAO dao.PostDAO) *GetProfile {
	return &GetProfile{
		userDAO:     userDAO,
		followDAO:   followDAO,
		bookmarkDAO: bookmarkDAO,
		postLikeDAO: postLikeDAO,
		postDAO:     postDAO,
	}
}

func (s *GetProfile) Exec(ctx context.Context, req *GetProfileReq) (*GetProfileResp, error) {
	// Get following users
	follows, err := s.followDAO.FindAll(ctx, "follower_id = $1", "created_at DESC", req.UserID)
	if err != nil {
		return nil, fmt.Errorf("failed to get follows: %w", err)
	}

	following := make([]ProfileUser, 0, len(follows))
	for _, follow := range follows {
		user, err := s.userDAO.FindByPk(ctx, follow.FolloweeID)
		if err != nil {
			continue // Skip if user not found
		}
		following = append(following, ProfileUser{
			ID:       user.ID,
			Username: user.Username,
		})
	}

	// Get bookmarked posts
	bookmarks, err := s.bookmarkDAO.FindAll(ctx, "user_id = $1", "created_at DESC", req.UserID)
	if err != nil {
		return nil, fmt.Errorf("failed to get bookmarks: %w", err)
	}

	bookmarkedPosts := make([]ProfilePost, 0, len(bookmarks))
	for _, bookmark := range bookmarks {
		post, err := s.postDAO.FindByPk(ctx, bookmark.PostID)
		if err != nil {
			continue // Skip if post not found
		}
		bookmarkedPosts = append(bookmarkedPosts, ProfilePost{
			ID:    post.ID,
			Title: post.Title,
		})
	}

	// Get liked posts
	likes, err := s.postLikeDAO.FindAll(ctx, "user_id = $1", "created_at DESC", req.UserID)
	if err != nil {
		return nil, fmt.Errorf("failed to get likes: %w", err)
	}

	likedPosts := make([]ProfilePost, 0, len(likes))
	for _, like := range likes {
		post, err := s.postDAO.FindByPk(ctx, like.PostID)
		if err != nil {
			continue // Skip if post not found
		}
		likedPosts = append(likedPosts, ProfilePost{
			ID:    post.ID,
			Title: post.Title,
		})
	}

	return &GetProfileResp{
		Following:  following,
		Bookmarks:  bookmarkedPosts,
		LikedPosts: likedPosts,
	}, nil
}
