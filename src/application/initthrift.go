package application

import (
	//"fmt"
	"io/ioutil"
	"strconv"
	"strings"
	"util"
)

func loadThriftConfig(tcstring string) ThriftConfig {
	content, err := ioutil.ReadFile(tcstring)
	util.CheckErr(err, "读取thrift.properties失败")
	//fmt.Println(string(content))

	lines := strings.Split(string(content), "\n")

	var tc ThriftConfig
	for _, line := range lines {
		l := strings.TrimSpace(line)
		if len(l) > 0 && !strings.HasPrefix(l, "#") { //不为空，并且不以#开头 string(l[0]) != "#"

			strs := strings.Split(l, "=")
			if len(strs) == 2 {
				//fmt.Println(strs[0], strs[1])
				if strs[0] == "ip" {
					tc.Ip = strs[1]
				}
				if strs[0] == "port" {
					p, err := strconv.Atoi(strs[1])
					util.CheckErr(err, "thrift 端口号配置错语")
					tc.Port = p
				}

			}
		}

	}
	return tc
}
