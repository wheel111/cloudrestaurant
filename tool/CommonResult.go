package tool

import "github.com/gin-gonic/gin"

const (
	SUCCESS = 0 // 操作成功
	FAILED  = 1 // 操作失败
)

// 普通成功返回
func Success(c *gin.Context, v interface{}) {
	c.JSON(200, map[string]interface{}{
		"code": SUCCESS,
		"msg":  "成功",
		"data": v,
	})
}

func Fail(c *gin.Context, v interface{}) {
	c.JSON(200, map[string]interface{}{
		"code": FAILED,
		"msg":  "失败",
		"data": v,
	})
}
