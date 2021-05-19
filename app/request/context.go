package request

import (
	"context"

	"github.com/zengqiang96/mattermostclone/model"
)

type Context struct {
	session model.Session

	context context.Context
}

func (c *Context) Session() *model.Session {
	return &c.session
}

func (c *Context) SetSession(s *model.Session) {
	c.session = *s
}
