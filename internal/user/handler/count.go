package handler

import "github.com/gin-gonic/gin"

// hyperloglog lưu xấp sỉ số người gọi api /ping, trả về trong api /count
func (hdl *userHandler) CountHyperLogLog() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		count := hdl.redis.PFCount(ctx, "memberCallPingAPI").Val()
		ctx.JSON(200, gin.H{"count": count})
	}
}