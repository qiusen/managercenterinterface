/*
*	错误工具
 */
package util

import (
	"log"
)

/*
* 检查是否有ERROR出现，打印错语信息，并停止程序继续进行
 */
func CheckErr(err error, mes string) {
	if err != nil {
		log.Println("because of : ", mes)
		panic(err.Error())
	}
}
