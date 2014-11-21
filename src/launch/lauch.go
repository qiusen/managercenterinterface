package main

import (
	"application"
	"fmt"
)

func init() {
	application.InitConfig()
}

func main() {

	db := application.GetDBConn()             //获取连接
	sqlmapData := application.GetSqlmapData() //获取所有sql语句配置

	//执行示例
	rows, err := db.Query(sqlmapData["city"]["selectCityAll"])
	if err != nil {
		fmt.Println(err.Error())
	}
	for rows.Next() {
		var id int
		var code, name, cityCode string
		err = rows.Scan(&id, &code, &name, &cityCode)
		if err != nil {
			fmt.Println(err.Error())
		}
		//fmt.Println(id, "----", code, "----", name, "---", cityCode)
	}
}
