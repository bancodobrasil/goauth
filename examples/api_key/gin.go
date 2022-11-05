package api_key

import (
	"github.com/bancodobrasil/goauth"
	goauthgin "github.com/bancodobrasil/goauth-gin"
	"github.com/bancodobrasil/goauth/handler"
	"github.com/bancodobrasil/goauth/log"
	"github.com/gin-gonic/gin"
)

// ApiKeyGin runs the example using Gin
func ApiKeyGin(logger log.Logger) {
	log.SetLogger(logger)

	cfg := handler.VerifyAPIKeyConfig{
		Header: "X-API-Key",
		Key:    "123456",
	}
	h := []goauth.AuthHandler{
		handler.NewVerifyAPIKey(cfg),
	}
	goauth.SetHandlers(h)

	r := gin.Default()
	r.Use(goauthgin.AuthenticateGin(nil))
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.Run()
}
