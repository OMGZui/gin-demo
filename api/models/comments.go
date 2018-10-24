/*
 * Created by GoLand.
 * User: 小粽子
 * Date: 2018/10/23
 * Time: 17:53
 */
package models

import (
	db "github.com/omgzui/gin-demo/api/database"
)

type Comments struct {
	ID int `json:"id"`
}

func CommentsNumber() int {
	var num int

	db.Eloquent.Table("comments").Count(&num)

	return num
}
