package middlewares

import (
	"github.com/gin-gonic/gin"
	"im/helper"
	"net/http"
)

func AuthCheck() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("token")
		UserClaims, err := helper.AnalyseToken(token)
		if err != nil {
			c.Abort()
			c.JSON(http.StatusOK, gin.H{
				"code": -1,
				"mes":  "用户认证不通过",
			})
			return
		}
		c.Set("user_claims", UserClaims)
		c.Next()
	}

}
