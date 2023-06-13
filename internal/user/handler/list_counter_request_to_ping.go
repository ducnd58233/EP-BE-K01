package handler

import (
	"github.com/EngineerProOrg/BE-K01/internal/user/model"
	"github.com/gin-gonic/gin"
)

func (hdl *userHandler) GetCounter() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var username model.LoginPing
		if err := ctx.ShouldBindJSON(&username); err != nil {
			ctx.JSON(400, gin.H{"error": err.Error()})
			return
		}
		count := hdl.redis.Get(ctx, username.Username).Val()
		ctx.JSON(200, gin.H{"count": count})
	}
}
