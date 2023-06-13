package handler

import (
	"math"
	"time"

	"github.com/EngineerProOrg/BE-K01/internal/user/model"
	"github.com/gin-gonic/gin"
)

func (hdl *userHandler) Ping() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var username model.LoginPing

		if err := ctx.ShouldBind(&username); err != nil {
			ctx.JSON(400, gin.H{"error": err.Error()})
			return
		}

		lockKey := "distributed-lock"
		lockValue := "distributed-lock-value"

		setNXResult := hdl.redis.SetNX(ctx, lockKey, lockValue, time.Duration(int(math.Pow(10, 9))))
		if setNXResult.Err() != nil || !setNXResult.Val() {
			ctx.JSON(400, gin.H{"error": "lock is being used"})
			return
		}

		hdl.redis.Incr(ctx, username.Username) // serve as a counter for the number of requests of each user

		hdl.redis.ZIncrBy(ctx, "leaderboard", 1, username.Username) // serve for api top 10

		hdl.redis.PFAdd(ctx, "membersCallPingAPI", username.Username) // serve for api count unique users call api /ping use hyperloglog

		ctx.JSON(200, gin.H{"message": "pong"})
	}
}