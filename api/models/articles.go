/*
 * Created by GoLand.
 * User: 小粽子
 * Date: 2018/10/23
 * Time: 16:35
 */
package models

import (
	db "github.com/omgzui/gin-demo/api/database"
)

type Articles struct {
	ID int `json:"id"`
}

func ArticlesNumber() int {
	var num int

	db.Eloquent.Table("articles").Count(&num)

	return num
}
