package server

import (
	"database/sql"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"

	"blog0/config"
	"blog0/internal/infra/handlers"
	"blog0/internal/infra/middlewares"
	"blog0/internal/infra/persistence/postgres"
	infraServices "blog0/internal/infra/services"
	"blog0/internal/services"
)

func New(cfg config.Config, db *sql.DB) *gin.Engine {
	router := gin.Default()
	router.Use(middlewares.UseCORS())

	googleOAuthConfig := &oauth2.Config{
		RedirectURL:  fmt.Sprintf("%s/api/v1/auth/google/callback", cfg.APIBaseURI),
		ClientID:     cfg.GoogleClientID,
		ClientSecret: cfg.GoogleClientSecret,
		Scopes:       []string{"https://www.googleapis.com/auth/userinfo.email", "https://www.googleapis.com/auth/userinfo.profile"},
		Endpoint:     google.Endpoint,
	}

	userDAO := postgres.NewUserDAO(db)

	nextIDFunc := uuid.NewString

	startOAuthServ := services.NewStartOAuth(googleOAuthConfig)
	finishOAuthServ := services.NewFinishOAuth(userDAO, googleOAuthConfig, infraServices.GoogleInfoExtractor, nextIDFunc, cfg)
	listPostsServ := services.NewListPosts()
	getPostBySlugServ := services.NewGetPostBySlug()
	createCommentServ := services.NewCreateComment()
	toggleLikeServ := services.NewToggleLike()
	bookmarkPostServ := services.NewBookmarkPost()
	unbookmarkPostServ := services.NewUnbookmarkPost()
	followUserServ := services.NewFollowUser()
	unfollowUserServ := services.NewUnfollowUser()
	getAuthorInfoServ := services.NewGetAuthorInfo()

	api := router.Group("/api/v1")
	{
		api.GET("/auth/google", handlers.StartOAuth(startOAuthServ))
		api.GET("/auth/google/callback", handlers.OAuthCallback(finishOAuthServ))

		api.GET("/posts", handlers.ListPosts(listPostsServ))
		api.GET("/posts/:slug", handlers.GetPostBySlug(getPostBySlugServ))
		api.GET("/users/:author_id", handlers.GetAuthorInfo(getAuthorInfoServ))

		api.Use(middlewares.HasAuthorization(cfg.JWTSecret))
		{
			api.POST("/posts/:slug/comments", handlers.CreateComment(createCommentServ))
			api.POST("/posts/:slug/likes", handlers.ToggleLike(toggleLikeServ))
			api.POST("/posts/:slug/bookmarks", handlers.BookmarkPost(bookmarkPostServ))
			api.DELETE("/posts/:slug/bookmarks", handlers.UnbookmarkPost(unbookmarkPostServ))
			api.POST("/users/:author_id/follow", handlers.FollowUser(followUserServ))
			api.DELETE("/users/:author_id/follow", handlers.UnfollowUser(unfollowUserServ))
		}
	}

	router.GET("/api/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	return router
}
