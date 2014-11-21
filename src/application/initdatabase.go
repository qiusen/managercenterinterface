/*
	1、解析数据库配置文件
	2、初始化数据库连接
*/
package application

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"io/ioutil"
	"strconv"
	"strings"
	"util"
)

func loadDatabaseConfig(database Database) map[string]string {
	var dbmap = make(map[string]string)

	fmt.Println("database.Config loading...", database.Config)
	content, err := ioutil.ReadFile(database.Config)
	util.CheckErr(err, "读取database.Config失败")
	//fmt.Println(string(content))

	lines := strings.Split(string(content), "\n")
	for _, line := range lines {
		l := strings.TrimSpace(line)
		if len(l) > 0 && !strings.HasPrefix(l, "#") { //不为空，并且不以#开头 string(l[0]) != "#"

			strs := strings.Split(l, "=")
			if len(strs) == 2 {
				//fmt.Println(strs[0], strs[1])
				dbmap[strs[0]] = strs[1]
			}
		}

	}
	return dbmap
}

func initConnection(dbmap map[string]string) *sql.DB {
	db, err := sql.Open("mysql", dbmap["db.username"]+":"+
		dbmap["db.password"]+"@"+
		dbmap["db.protocol"]+"("+
		dbmap["db.host"]+":"+
		dbmap["db.port"]+")/"+
		dbmap["db.dbname"]+"?charset="+
		dbmap["db.charset"])
	//db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/acomp?charset=utf8")

	util.CheckErr(err, "连接数据库失败")

	if dbmap["db.maxopenconns"] != "" {
		m, err := strconv.Atoi(dbmap["db.maxopenconns"])
		util.CheckErr(err, "最大连接数请使用数字")
		db.SetMaxOpenConns(m)
	}

	if dbmap["db.maxidleconns"] != "" {
		i, err := strconv.Atoi(dbmap["db.maxopenconns"])
		util.CheckErr(err, "最大连接数请使用数字")
		db.SetMaxIdleConns(i)
	}

	//rows, err := db.Query("select * from area")
	//if err != nil {
	//	fmt.Println(err.Error())
	//}
	//for rows.Next() {
	//	var id int
	//	var code, name, cityCode string
	//	err = rows.Scan(&id, &code, &name, &cityCode)
	//	if err != nil {
	//		fmt.Println(err.Error())
	//	}
	//	fmt.Println(id, "----", code, "----", name, "---", cityCode)
	//}

	//defer db.Close()
	//defer rows.Close()

	return db
}
