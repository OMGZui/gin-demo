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

type Visitors struct {
	ID  int `json:"id"`
	Num int `json:"num"` //总的点击量
}

func VisitorsNumber() Visitors {
	var num Visitors

	db.Eloquent.Table("visitors").Select("sum(clicks) as num").Scan(&num)

	return num
}
