package tserver

import (
	"fmt"
	"managercentertserver/managercenter"
	//"git.apache.org/thrift.git/lib/go/thrift"
	"application"
	"database/sql"
	"github.com/apache/thrift/lib/go/thrift"
	"os"
	"strconv"
	"strings"
	"util"
)

//const (
//	NetworkAddr = "127.0.0.1:19090"
//)

type ServiceImpl struct { //通用struct
}

func (this *ServiceImpl) FunCallManager(callTime int64, funCode string, paramMap map[string]string) (r []*managercenter.Manager, err error) {
	fmt.Println("-->FunCall:", callTime, funCode, paramMap)

	var db *sql.DB
	db = application.GetDBConn()              //获取连接
	sqlmapData := application.GetSqlmapData() //获取所有sql语句配置

	//执行
	tName := paramMap["tName"]
	sName := paramMap["sName"]

	var sqlstr string = sqlmapData[tName][sName] //sql语句

	pValue := paramMap["pValue"] //参数之间用“,”分隔，有序
	var ps []interface{}
	if pValue != "" {
		ss := strings.Split(pValue, ",")

		for _, p := range ss {
			//sqlstr = strings.Replace(sqlstr, "?", p, 1)

			if i, err := strconv.Atoi(p); err == nil {
				ps = append(ps, i)
			} else if f, err := strconv.ParseFloat(p, -1); err == nil {
				ps = append(ps, f)
			} else {
				ps = append(ps, p)
			}

		}

	}

	rows, err := db.Query(sqlstr, ps...)

	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(rows)
	for rows.Next() {
		var id, status, department_id, position_id, role_id int32
		var createtime, logintime string

		var email, password, nickname, ename, logo, creator, loginip string
		err = rows.Scan(&id, &email, &password, &nickname, &ename, &logo, &status, &department_id, &position_id, &role_id, &creator, &createtime, &logintime, &loginip)
		if err != nil {
			fmt.Println(err.Error())
		}
		fmt.Println(id, "----", email, "----", password, "---", nickname, "----", ename, "----", logo, "---", status, "---", department_id, "----", position_id, "----", role_id, "---", creator, "--xx--", createtime, "--xx--", logintime, "---", loginip)

		var temp = managercenter.Manager{id, email, password, nickname, ename, logo, status, department_id, position_id, role_id, creator, createtime, logintime, loginip}
		r = append(r, &temp)
	}

	return
}

func (this *ServiceImpl) FunCallManagerInfo(callTime int64, funCode string, paramMap map[string]string) (r []*managercenter.ManagerInfo, err error) {
	fmt.Println("-->FunCall:", callTime, funCode, paramMap)

	var db *sql.DB
	db = application.GetDBConn()              //获取连接
	sqlmapData := application.GetSqlmapData() //获取所有sql语句配置

	//执行
	tName := paramMap["tName"]
	sName := paramMap["sName"]

	var sqlstr string = sqlmapData[tName][sName] //sql语句

	pValue := paramMap["pValue"] //参数之间用“,”分隔，有序
	var ps []interface{}
	if pValue != "" {
		ss := strings.Split(pValue, ",")

		for _, p := range ss {
			//sqlstr = strings.Replace(sqlstr, "?", p, 1)

			if i, err := strconv.Atoi(p); err == nil {
				ps = append(ps, i)
			} else if f, err := strconv.ParseFloat(p, -1); err == nil {
				ps = append(ps, f)
			} else {
				ps = append(ps, p)
			}

		}

	}

	rows, err := db.Query(sqlstr, ps...)

	if err != nil {
		fmt.Println(err.Error())
	}
	for rows.Next() {
		var id, sex int32
		var updatetime string
		var email, tel, mobile, address, qq string
		err = rows.Scan(&id, &email, &tel, &mobile, &sex, &address, &qq, &updatetime)
		if err != nil {
			fmt.Println(err.Error())
		}
		fmt.Println(id, "----", email, "----", tel, "---", mobile, "----", sex, "----", address, "---", qq, "---", updatetime)
		var temp = managercenter.ManagerInfo{id, email, tel, mobile, sex, address, qq, updatetime}
		r = append(r, &temp)
	}

	return
}

func (this *ServiceImpl) FunCallDepartment(callTime int64, funCode string, paramMap map[string]string) (r []*managercenter.Department, err error) {
	fmt.Println("-->FunCall:", callTime, funCode, paramMap)

	var db *sql.DB
	db = application.GetDBConn()              //获取连接
	sqlmapData := application.GetSqlmapData() //获取所有sql语句配置

	//执行
	tName := paramMap["tName"]
	sName := paramMap["sName"]

	var sqlstr string = sqlmapData[tName][sName] //sql语句

	pValue := paramMap["pValue"] //参数之间用“,”分隔，有序

	fmt.Printf("tName: %s; sName: %s; pValue: %s\n", tName, sName, pValue)
	fmt.Println("sqlstr::::::::::::::", sqlstr)

	var ps []interface{}
	if pValue != "" {
		ss := strings.Split(pValue, ",")

		fmt.Println("len(ss)::::::::::::::", len(ss))

		for _, p := range ss {
			//sqlstr = strings.Replace(sqlstr, "?", p, 1)

			if i, err := strconv.Atoi(p); err == nil {
				ps = append(ps, i)
			} else if f, err := strconv.ParseFloat(p, -1); err == nil {
				ps = append(ps, f)
			} else {
				ps = append(ps, p)
			}

		}
		fmt.Println("len(ps)::::::::::::::", len(ps))

	}

	rows, err := db.Query(sqlstr, ps...)

	//fmt.Println("rows.Columns()::::::::::::::", rows.Columns())

	if err != nil {
		fmt.Println(err.Error())
	}

	for rows.Next() {
		var id, dtype, super_id int32
		var createtime, updatetime string
		var name, code, leader_email, creator, updator string
		err = rows.Scan(&id, &name, &code, &dtype, &leader_email, &super_id, &creator, &createtime, &updator, &updatetime)
		if err != nil {
			fmt.Println(err.Error())
		}
		fmt.Println(id, "----", name, "----", code, "---", dtype, "----", leader_email, "----", super_id, "---", creator, "---", createtime, "---", updator, "---", updatetime)

		var temp = managercenter.Department{id, name, code, dtype, leader_email, super_id, creator, createtime, updator, updatetime}
		r = append(r, &temp)
	}

	return
}

func StartTserver() {
	tconfig := application.GetThrfitConfig()
	fmt.Println("thrift tconfig.ip", tconfig.Ip)
	fmt.Println("thrift tconfig.port", tconfig.Port)

	NetworkAddr := tconfig.Ip + ":" + strconv.Itoa(tconfig.Port)

	transportFactory := thrift.NewTFramedTransportFactory(thrift.NewTTransportFactory())
	protocolFactory := thrift.NewTBinaryProtocolFactoryDefault()
	//protocolFactory := thrift.NewTCompactProtocolFactory()

	serverTransport, err := thrift.NewTServerSocket(NetworkAddr)
	if err != nil {
		fmt.Println("Error!", err)
		os.Exit(1)
	}

	handler := &ServiceImpl{}
	processor := managercenter.NewManagerCenterServiceProcessor(handler)

	server := thrift.NewTSimpleServer4(processor, serverTransport, transportFactory, protocolFactory)
	fmt.Println("thrift server in", NetworkAddr)
	err = server.Serve()
	util.CheckErr(err, "服务启动失败... ")

}
