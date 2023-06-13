package handler

import (
	"time"

	"github.com/EngineerProOrg/BE-K01/internal/user/model"
	"github.com/gin-gonic/gin"
)

// rate limit mỗi người chỉ được gọi API /ping 2 lần trong 60s
func (hdl *userHandler) RateLimitPing() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var username model.LoginPing
		if err := ctx.ShouldBindJSON(&username); err != nil {
			ctx.JSON(400, gin.H{"error": err.Error()})
			return
		}

		exists, _ := hdl.redis.Exists(ctx, username.Username).Result()
		if exists == 0 {
			hdl.redis.Set(ctx, username.Username, 0, time.Duration(60*time.Second))
		} else {
			count := hdl.redis.Get(ctx, username.Username).Val()
			if count == "2" {
				ctx.JSON(400, gin.H{"error": "limit 2 requests per 5s"})
				return
			}
		}
		hdl.redis.Incr(ctx, username.Username)

		ctx.JSON(200, gin.H{"message": "pong 2"})
	}
}