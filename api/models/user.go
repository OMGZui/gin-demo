/*
 * Created by GoLand.
 * User: 小粽子
 * Date: 2018/10/19
 * Time: 16:15
 */
package models

import (
	db "github.com/omgzui/gin-demo/api/database"
)

type Users struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Password string `json:"password"`
}

func List() (*[]Users, error) {
	var list []Users

	err := db.Eloquent.Table("users").Find(&list).Error

	return &list, err
}

func UsersNumber() int {
	var num int

	db.Eloquent.Table("users").Count(&num)

	return num
}
