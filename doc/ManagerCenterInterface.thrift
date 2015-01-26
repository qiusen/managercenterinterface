namespace go tserver.managercenter
namespace java com.dihaitech.tserver.managercenter

//管理员，使用string处理时间类型
struct Manager{
	1:i32 id,
	2:string email,
	3:string password,
	4:string nickname,
	5:string ename,
	6:string logo,
	7:i32 status,
	8:i32 department_id,
	9:i32 position_id,
	10:i32 role_id,
	11:string creator,
	12:string createtime,
	13:string logintime,
	14:string loginip,
}

//管理员详细信息
struct ManagerInfo{
	1:i32 id,
	2:string email,
	3:string tel,
	4:string mobile,
	5:i32 sex,
	6:string address,
	7:string qq,
	8:string updatetime,
}

struct Department{
	1:i32 id,
	2:string name,
	3:string code,
	4:i32 type,
	5:string leader_email,
	6:i32 super_id,
	7:string creator,
	8:string createtime,
	9:string updator,
	10:string updatetime,
	
}

// 城市服务
service ManagerCenterService {

  // 发起远程调用
 
  //查询
  list<Manager> funCallManager(1:i64 callTime, 2:string funCode, 3:map<string, string> paramMap),
  
  list<ManagerInfo> funCallManagerInfo(1:i64 callTime, 2:string funCode, 3:map<string, string> paramMap),
  
  list<Department> funCallDepartment(1:i64 callTime, 2:string funCode, 3:map<string, string> paramMap),

}