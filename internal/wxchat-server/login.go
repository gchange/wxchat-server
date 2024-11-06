package wxchatserver

import (
	"time"
	"wx-server/internal/httpserver"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

type LoginRequest struct {
	Code string `json:"code"`
}

type LoginResponse struct {
	Token string `json:"token"`
}

type user struct {
	Nmae string `json:"name"`
	jwt.RegisteredClaims
}

func (s *Server) login(c *gin.Context) {
	var req LoginRequest
	err := c.Bind(&req)
	if err != nil {
		httpserver.ResponseBadRequest(c, err)
		return
	}

	openid, err := s.wx.GetOpenId(req.Code)
	if err != nil {
		httpserver.ResponseUnauthorized(c, err)
		return
	}

	u := user{
		openid,
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
		},
	}
	t := jwt.NewWithClaims(jwt.SigningMethodHS256, u)
	token, err := t.SignedString(s.config.SecretKey)
	if err != nil {
		httpserver.ResponseServerError(c, err)
		return
	}
	httpserver.ResponseOK(c, LoginResponse{token})
}
