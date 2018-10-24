/*
 * Created by GoLand.
 * User: 小粽子
 * Date: 2018/10/19
 * Time: 16:27
 */
package router

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	. "github.com/omgzui/gin-demo/api/apis"
	"time"
)

func InitRouter() *gin.Engine {
	router := gin.Default()

	// 跨域
	router.Use(cors.New(cors.Config{
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "HEAD", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Length", "Content-Type", "X-Requested-With", "X-CSRF-TOKEN"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		AllowAllOrigins: true,
		MaxAge: 12 * time.Hour,
	}))

	router.GET("/api/user/list", List)
	router.GET("/api/statistics", Statistics)

	return router
}
