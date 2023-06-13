package handler

import (
	"crypto/rand"
	"encoding/base64"
	"math"
	"time"

	"github.com/EngineerProOrg/BE-K01/internal/user/model"
	"github.com/gin-gonic/gin"
)

func (hdl *userHandler) Login() gin.HandlerFunc {
	return func(ctx *gin.Context) {
			var user model.User
			if err := ctx.ShouldBindJSON(&user); err != nil {
				ctx.JSON(400, gin.H{"error": err.Error()})
				return
			}

			sessionId, error := generateSessionID()
			if error != nil {
				ctx.JSON(400, gin.H{"error": error.Error()})
				return
			}
			if err := hdl.redis.Set(ctx, sessionId, user.Username, time.Duration(int(math.Pow(10, 9)))).Err(); err != nil {
				ctx.JSON(400, gin.H{"error": err.Error()})
				return
			}
			ctx.JSON(200, gin.H{"sessionId": sessionId})
	}
}

func generateSessionID() (string, error) {
	randomBytes := make([]byte, 32)
	_, err := rand.Read(randomBytes)

	if err != nil {
		return "", err
	}

	sessionID := base64.URLEncoding.EncodeToString(randomBytes)

	return sessionID, nil
}