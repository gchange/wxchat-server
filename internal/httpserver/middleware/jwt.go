package middleware

import (
	"context"
	"errors"
	"net/http"
	"wx-server/internal/httpserver"

	"github.com/gin-gonic/gin"
)

const (
	XAccessToken = "X-Access-Token"
	AccessToken  = "ACCESS_TOKEN"
	UID          = "UID"
	Username     = "Username"
)

func WrapJwtTokenParser() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader(XAccessToken)
		// get token from cookie if not found in header
		var err error
		if token == "" {
			token, err = c.Cookie(AccessToken)
		}
		if err != nil && !errors.Is(err, http.ErrNoCookie) {
			httpserver.ResponseUnauthorized(c, err.Error())
			c.Abort()
			return
		}
		if err != nil {
			httpserver.ResponseUnauthorized(c, "empty token")
			c.Abort()
			return
		}

		// TODO: authorize token

		// set user id and name
		ctx := c.Request.Context()
		// if cliams.Username != "" {
		// 	ctx = WithUsername(ctx, cliams.Username)
		// }
		c.Request = c.Request.WithContext(ctx)
		c.Next()
	}
}

func WithUID(ctx context.Context, uid uint) context.Context {
	return context.WithValue(ctx, UID, uid)
}

func WithUsername(ctx context.Context, username string) context.Context {
	return context.WithValue(ctx, Username, username)
}

func GetUID(ctx context.Context) uint {
	if uid, ok := ctx.Value(UID).(uint); ok {
		return uid
	}
	return 0
}

func GetUsername(ctx context.Context) string {
	if username, ok := ctx.Value(Username).(string); ok {
		return username
	}
	return ""
}
