package app

import (
	"net/http"

	"github.com/zengqiang96/mattermostclone/app/request"
	"github.com/zengqiang96/mattermostclone/model"
)

type AppIface interface {
	AttachSessionCookies(c *request.Context, w http.ResponseWriter, r *http.Request)
	AuthenticateUserForLogin(c *request.Context, loginId, password string) (*model.User, *model.AppError)
	CreatePostAsUser(c *request.Context, post *model.Post, currentSessionId string) (*model.Post, *model.AppError)
	CreateUserFromSignup(c *request.Context, user *model.User) (*model.User, *model.AppError)
	Config() *model.Config
	DoLogin(c *request.Context, w http.ResponseWriter, r *http.Request, user *model.User) *model.AppError
	HubRegister(webConn *WebConn)
	NewWebConn(cfg *WebConnConfig) *WebConn
	OriginChecker() func(*http.Request) bool
	GetSession(token string) (*model.Session, *model.AppError)
}
