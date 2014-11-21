/*
	加载所有sqlmap文件
*/
package application

import (
	"bytes"
	"container/list"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"path/filepath"
	"strings"
	"util"
)

func loadSqlmap(database Database) *list.List {
	fmt.Println("loadSqlmap ", database.Sqlmap)
	absFilePath, _ := filepath.Abs(database.Sqlmap)
	//fmt.Println(absFilePath)
	//absFilePath = strings.Replace(absFilePath, ".", "\\.", -1) //第四个参数小于0，表示所有的都替换
	//fmt.Println(absFilePath)

	i := strings.Index(absFilePath, "*")
	absFolder := util.Substr(absFilePath, 0, i)
	extName := util.Substr(absFilePath, i+1, len(absFilePath))
	fmt.Println(absFolder)
	fmt.Println(extName)

	fileList := util.GetAllFileListByFilePath(absFolder)
	fmt.Println(fileList.Len())

	for fl := fileList.Front(); fl != nil; fl = fl.Next() {

		flstr, ok := fl.Value.(string)
		if ok && strings.HasSuffix(flstr, extName) {
			fmt.Println(fl.Value)
		} else {
			fileList.Remove(fl)
		}
	}

	return fileList

}

func initSqlmapData(fileList *list.List) map[string]map[string]string {
	//var sqlmapData = make(map[string]map[string]string)

	sqlmapData := map[string]map[string]string{}
	fmt.Println(fileList.Len())

	var t xml.Token

	for fl := fileList.Front(); fl != nil; fl = fl.Next() {

		flstr, _ := fl.Value.(string)

		content, err := ioutil.ReadFile(flstr)
		util.CheckErr(err, "读取"+flstr+"失败")

		decoder := xml.NewDecoder(bytes.NewBuffer(content))
		var c = 0
		var tName, sName, sValue string
		var sMap map[string]string
		var ok bool = false

		for t, err = decoder.Token(); err == nil; t, err = decoder.Token() {
			switch token := t.(type) {
			case xml.StartElement:
				name := token.Name.Local

				if name == "sqlList" {
					for _, attr := range token.Attr {
						attrName := attr.Name.Local
						attrValue := attr.Value
						if attrName == "id" {
							tName = attrValue
							//fmt.Println("tName::::::::::", tName)
							sMap, ok = sqlmapData[tName]
							if !ok {
								sMap = make(map[string]string)
								sqlmapData[tName] = sMap
								//fmt.Println("new sMap")
							}
						}
						//fmt.Printf("An attribute is : %s %s\n", attrName, attrValue)
					}
				}
				if name == "sql" {
					c = 1
					for _, attr := range token.Attr {
						attrName := attr.Name.Local
						attrValue := attr.Value
						if attrName == "id" {
							sName = attrValue
							//fmt.Println("sName::::::::::", tName)

						}
						//fmt.Printf("An attribute is : %s %s\n", attrName, attrValue)
					}

				}

			case xml.EndElement:

				c = 0
				//fmt.Println("end", token.Name.Local)
			case xml.CharData:
				content := string([]byte(token))
				sValue = strings.TrimSpace(content)

				if c == 1 {
					sMap[sName] = sValue
					sName = ""
					sValue = ""
				}

				//fmt.Println("content", strings.TrimSpace(content))
			}
		}

		//sqlmapData[tName] = sMap

		//sMap = nil

		fmt.Printf("-----解析文件 %s 完成-----\n", flstr)
	}

	fmt.Println("selectAreaAll: ", sqlmapData["area"]["selectAreaAll"])
	fmt.Println("selectCityAll: ", sqlmapData["city"]["selectCityAll"])
	return sqlmapData
}
