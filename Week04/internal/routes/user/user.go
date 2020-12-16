package user

import (
	"github.com/gin-gonic/gin"
	"jmolboy/week04/internal/biz"
	"jmolboy/week04/internal/dao"
	"net/http"
)

// 分页查询
func List(c *gin.Context) {

	userBiz := biz.User{}

	id, err := userBiz.Create(dao.User{})
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"error": 1,
			"msg":   "failed",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"error": 0,
		"msg":   id,
	})
	return
}
