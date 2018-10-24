package main

import (
	orm "github.com/omgzui/gin-demo/api/database"
	router2 "github.com/omgzui/gin-demo/api/router"
)

func main() {
	defer orm.Eloquent.Close()
	router := router2.InitRouter()
	router.Run(":8081")
}