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

	postDAO := postgres.NewPostDAO(db)
	userDAO := postgres.NewUserDAO(db)
	commentDAO := postgres.NewCommentDAO(db)
	postLikeDAO := postgres.NewPostLikeDAO(db)
	bookmarkDAO := postgres.NewBookmarkDAO(db)
	followDAO := postgres.NewFollowDAO(db)

	postContentGenerator := infraServices.NewOpenAIGenerator(cfg.OpenAIApiKey, "gpt-4o")
	nextIDFunc := uuid.NewString
	triggerDev := infraServices.NewTriggerDev(cfg.TriggerSecretKey)
	eventBus := infraServices.NewTriggerDevEventBus(triggerDev)

	startOAuthServ := services.NewStartOAuth(googleOAuthConfig)
	finishOAuthServ := services.NewFinishOAuth(userDAO, googleOAuthConfig, infraServices.GoogleInfoExtractor, nextIDFunc, cfg)
	listPostsServ := services.NewListPosts(postDAO, userDAO, postLikeDAO, commentDAO)
	getPostBySlugServ := services.NewGetPostBySlug(postDAO, userDAO, commentDAO, postLikeDAO)
	createCommentServ := services.NewCreateComment(postDAO, userDAO, commentDAO, nextIDFunc)
	toggleLikeServ := services.NewToggleLike(postDAO, postLikeDAO, nextIDFunc)
	bookmarkPostServ := services.NewBookmarkPost(postDAO, bookmarkDAO, nextIDFunc)
	unbookmarkPostServ := services.NewUnbookmarkPost(postDAO, bookmarkDAO)
	createPostServ := services.NewCreatePost(postDAO, nextIDFunc, postContentGenerator, eventBus)
	updatePostServ := services.NewUpdatePost(postDAO, postContentGenerator, eventBus)
	deletePostServ := services.NewDeletePost(postDAO)
	listMyPostsServ := services.NewListMyPosts(postDAO, userDAO)
	getAuthorInfoServ := services.NewGetAuthorInfo(userDAO, postDAO, postLikeDAO)
	followUserServ := services.NewFollowUser(userDAO, followDAO, nextIDFunc)
	unfollowUserServ := services.NewUnfollowUser(userDAO, followDAO)
	getProfileServ := services.NewGetProfile(userDAO, followDAO, bookmarkDAO, postLikeDAO, postDAO)

	api := router.Group("/api/v1")
	{
		api.GET("/auth/google", handlers.StartOAuth(startOAuthServ))
		api.GET("/auth/google/callback", handlers.OAuthCallback(finishOAuthServ))

		api.GET("/posts", handlers.ListPosts(listPostsServ))
		api.GET("/posts/:slug", handlers.GetPostBySlug(getPostBySlugServ))
		api.GET("/users/:author_id", handlers.GetAuthorInfo(getAuthorInfoServ))

		api.Use(middlewares.HasAuthorization(cfg.JWTSecret))
		{
			// User-specific endpoints (my content)
			api.GET("/me/profile", handlers.GetProfile(getProfileServ))
			api.POST("/me/posts", handlers.CreatePost(createPostServ))
			api.PUT("/me/posts/:slug", handlers.UpdatePost(updatePostServ))
			api.DELETE("/me/posts/:slug", handlers.DeletePost(deletePostServ))
			api.GET("/me/posts", handlers.ListMyPosts(listMyPostsServ))

			// Post interactions
			api.POST("/posts/:slug/comments", handlers.CreateComment(createCommentServ))
			api.POST("/posts/:slug/likes", handlers.ToggleLike(toggleLikeServ))
			api.POST("/posts/:slug/bookmarks", handlers.BookmarkPost(bookmarkPostServ))
			api.DELETE("/posts/:slug/bookmarks", handlers.UnbookmarkPost(unbookmarkPostServ))

			// User interactions
			api.POST("/users/:author_id/follow", handlers.FollowUser(followUserServ))
			api.DELETE("/users/:author_id/follow", handlers.UnfollowUser(unfollowUserServ))
		}
	}

	router.GET("/api/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	return router
}
