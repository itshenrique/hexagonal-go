package middlewares

import (
	"nito/api/internals/config"
	"nito/api/internals/core/domain"
	"nito/api/internals/core/ports"
	"nito/api/internals/utils/crypto_util"
	"nito/api/internals/utils/http_error"

	"github.com/gin-gonic/gin"
)

// DI

type IAuthRequiredMiddleware interface {
	Initialize() gin.HandlerFunc
}

type AuthRequiredMiddleware struct {
	sessionService ports.ISessionService
}

func NewAuthRequiredMiddleware(sessionService ports.ISessionService) IAuthRequiredMiddleware {
	return &AuthRequiredMiddleware{
		sessionService: sessionService,
	}
}

// Functions

func (m *AuthRequiredMiddleware) Initialize() gin.HandlerFunc {
	return func(c *gin.Context) {
		config := config.LoadConfig(".")
		sessionID, err := c.Cookie("session-id")

		if err != nil {
			error := http_error.NewForbiddenError("user not authenticated")
			c.AbortWithStatusJSON(error.Status(), error)
			return
		}

		decriptedSessionId, err := crypto_util.Decrypt(sessionID, config.SessionSecret)

		if err != nil {
			println(err.Error())
		}

		data, err := m.sessionService.Get(*decriptedSessionId)

		if err != nil {
			error := http_error.NewForbiddenError("user not authenticated")
			c.AbortWithStatusJSON(error.Status(), error)
			return
		}

		m.sessionService.Delete(*decriptedSessionId)
		m.sessionService.Set(*decriptedSessionId, domain.Session{
			ID:       sessionID,
			Username: data.Username,
		})

		c.Next()
	}
}
