/*
 * Created by GoLand.
 * User: 小粽子
 * Date: 2018/10/19
 * Time: 16:19
 */
package database

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

var Eloquent *gorm.DB

func init() {
	var err error

	Eloquent, err = gorm.Open(
		"mysql",
		"root:root@tcp(127.0.0.1:3307)/blog?charset=utf8mb4&parseTime=True&loc=Local&timeout=10ms",
	)
	if err != nil {
		fmt.Printf("mysql连接失败 %v", err)
	}

	if Eloquent.Error != nil {
		fmt.Printf("数据库错误 %v", Eloquent.Error)
	}
	// 禁用复数
	Eloquent.SingularTable(true)
}
