/*
 * Created by GoLand.
 * User: 小粽子
 * Date: 2018/10/23
 * Time: 16:33
 */
package apis

import (
	"github.com/gin-gonic/gin"
	"github.com/omgzui/gin-demo/api/models"
	"net/http"
)

// 统计
func Statistics(c *gin.Context) {
	users := models.UsersNumber()
	articles := models.ArticlesNumber()
	visitors := models.VisitorsNumber().Num
	comments := models.CommentsNumber()

	c.JSON(http.StatusOK, gin.H{
		"users":    users,
		"articles": articles,
		"visitors": visitors,
		"comments": comments,
	})
}
