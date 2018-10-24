/*
 * Created by GoLand.
 * User: 小粽子
 * Date: 2018/10/19
 * Time: 16:13
 */
package apis

import (
	"github.com/gin-gonic/gin"
	"github.com/omgzui/gin-demo/api/models"
	"net/http"
)

// 所有用户
func List(c *gin.Context) {

	result, err := models.List()

	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":    400,
			"message": "没找到相关信息",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"data": result,
	})
}
