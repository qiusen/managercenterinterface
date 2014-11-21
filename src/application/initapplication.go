/*
	应用启始配置
	1、解析系统XML文件
	2、提取并解析数据库配置

*/
package application

import (
	"encoding/xml"
	"fmt"
	//"strings"
	"bytes"
	"errors"
	"io/ioutil"
	//"log"
	"database/sql"
	"util"
)

var DBconn *sql.DB

var SqlmapData map[string]map[string]string

type Database struct {
	Config string
	Sqlmap string
}

func InitConfig() {
	var t xml.Token
	var err error

	//input := "<?xml version=\"1.0\" encoding=\"UTF-8\"?><application><database><dbconfig>./db.properties</dbconfig><sqlList>./sqlmap/*.xml</sqlList></database></application>"
	//inputrReader := strings.NewReader(input)

	//decoder := xml.NewDecoder(inputrReader)

	content, err := ioutil.ReadFile("application.xml")
	util.CheckErr(err, "读取application.xml失败")

	decoder := xml.NewDecoder(bytes.NewBuffer(content))

	var database Database

	var c = 0

	for t, err = decoder.Token(); err == nil; t, err = decoder.Token() {
		switch token := t.(type) {
		case xml.StartElement:
			name := token.Name.Local

			if name == "dbconfig" {
				c = 1
			}
			if name == "sqlmap" {
				c = 2
			}

		case xml.EndElement:

			c = 0
		case xml.CharData:
			content := string([]byte(token))

			if c == 1 {
				database.Config = content
			}
			if c == 2 {
				database.Sqlmap = content
			}
		}
	}

	fmt.Printf("database.Config = %s, database.Sqlmap= %s \n", database.Config, database.Sqlmap)
	if database.Config == "" {
		util.CheckErr(errors.New("database.Config is nil"), "database.Config is nil")

	}
	if database.Sqlmap == "" {
		util.CheckErr(errors.New("database.SqlList is nil"), "database.Sqlmap is nil")
	}

	dbmap := loadDatabaseConfig(database)

	for i, v := range dbmap {
		fmt.Println(i, v)
	}
	//fmt.Println("sssssss", dbmap["db.username"])
	DBconn = initConnection(dbmap)

	fileList := loadSqlmap(database)
	SqlmapData = initSqlmapData(fileList)
	fmt.Println(len(SqlmapData))

	//for t, err = decoder.Token(); err == nil; t, err = decoder.Token() {
	//	switch token := t.(type) {
	//	case xml.StartElement:
	//		name := token.Name.Local
	//		fmt.Printf("Token name: %s\n", name)
	//		for _, attr := range token.Attr {
	//			attrName := attr.Name.Local
	//			attrValue := attr.Value
	//			fmt.Printf("An attribute is : %s %s\n", attrName, attrValue)
	//		}
	//	case xml.EndElement:
	//		fmt.Printf("Token of '%s' end\n", token.Name.Local)
	//	case xml.CharData:
	//		content := string([]byte(token))
	//		fmt.Printf("This is the content: %v\n", content)
	//	default:
	//		fmt.Println("-------")
	//	}
	//}
}

func GetDBConn() *sql.DB {
	return DBconn
}

func GetSqlmapData() map[string]map[string]string {
	return SqlmapData
}
