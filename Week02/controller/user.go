package controller

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/yann0917/Go-000/Week02/internal/dao"
)

func GetUser(c *gin.Context) {

	id := c.Param("id")
	var u dao.User
	u.ID, _ = strconv.Atoi(id)
	err := u.Detail()
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 1,
			"msg":  err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code": 0,
		"msg":  "ok",
		"data": u})
}
