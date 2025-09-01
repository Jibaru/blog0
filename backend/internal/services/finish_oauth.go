package services

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt"
	"golang.org/x/oauth2"

	"blog0/config"
	"blog0/internal/domain"
	"blog0/internal/domain/dao"
)

type FinishOAuthReq struct {
	Code string `json:"code"`
}

type FinishOAuthResp struct {
	URL string
}

type FinishOAuth struct {
	userDAO            dao.UserDAO
	oauthCfg           *oauth2.Config
	oauthInfoExtractor domain.InfoExtractor
	nextID             domain.NextID
	cfg                config.Config
}

func NewFinishOAuth(
	userDAO dao.UserDAO,
	oauthCfg *oauth2.Config,
	oauthInfoExtractor domain.InfoExtractor,
	nextID domain.NextID,
	cfg config.Config,
) *FinishOAuth {
	return &FinishOAuth{
		userDAO:            userDAO,
		oauthCfg:           oauthCfg,
		oauthInfoExtractor: oauthInfoExtractor,
		nextID:             nextID,
		cfg:                cfg,
	}
}

func (s *FinishOAuth) Exec(ctx context.Context, req FinishOAuthReq) (*FinishOAuthResp, error) {
	token, err := s.oauthCfg.Exchange(ctx, req.Code)
	if err != nil {
		return nil, err
	}

	username, email, err := s.oauthInfoExtractor(token.AccessToken)
	if err != nil {
		return nil, err
	}

	user, err := s.userDAO.FindOne(ctx, "email = $1", email)
	if err != nil && errors.Is(err, sql.ErrNoRows) {
		newUser, err := domain.NewUser(s.nextID(), email, username)
		if err != nil {
			return nil, err
		}

		err = s.userDAO.Create(ctx, newUser)
		if err != nil {
			return nil, err
		}

		user = newUser
	} else if err != nil {
		return nil, err
	}

	tokenString, err := generateToken(user.ID, user.Email, []byte(s.cfg.JWTSecret))
	if err != nil {
		return nil, err
	}

	url := fmt.Sprintf("%s/login?token=%s&id=%s&email=%s&username=%s",
		s.cfg.WebBaseURI,
		tokenString,
		user.ID,
		user.Email,
		user.Username,
	)

	return &FinishOAuthResp{
		URL: url,
	}, nil
}

func generateToken(userID string, userEmail string, jwtSecret []byte) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": userID,
		"email":   userEmail,
		"exp":     time.Now().Add(24 * time.Hour).Unix(),
	})
	tokenString, err := token.SignedString(jwtSecret)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
