// Code generated by goctl. DO NOT EDIT.
// goctl 1.7.2

package types

type AlarmLogReqest struct {
	Current          int64  `form:"current,default=1,optional"`      // 页码
	PageSize         int64  `form:"page_size,default=10,optional"`   // 数据量
	CreatedStartTime int64  `form:"created_start_time,optional"`     // 创建开始时间
	CreatedEndTime   int64  `form:"created_end_time,optional"`       // 创建结束时间
	AlarmCategory    int64  `form:"alarm_category,options=1|2|3"`    // 告警类别 1:告警 2:yx提醒 3:网关下线提醒
	Id               string `form:"id,optional"`                     // 告警ID
	Mid              string `form:"mid,optional"`                    // mid 预警设备ID
	AlarmType        int64  `form:"alarm_type,default=99,optional"`  // 类型：1 越上限/2 越下限/ 3 变位 / 4 网关下线
	AlarmGrade       int64  `form:"alarm_grade,default=99,optional"` // 等级：1 预警/2 告警 /3 提醒
	AssetCode        string `form:"asset_code,optional"`             // 柜号
	AlarmState       int64  `form:"alarm_state,default=99,optional"` // 状态 0 未读  1已读  2已确认
}

type AlarmUpStateReqest struct {
	Id string `form:"id"` // 告警ID
}

type AppLogReqest struct {
	Current          int64  `json:"current,default=1,optional"`     // 页码
	PageSize         int64  `json:"page_size,default=10,optional"`  // 数据量
	CreatedStartTime int64  `json:"created_start_time,optional"`    // 创建开始时间
	CreatedEndTime   int64  `json:"created_end_time,optional"`      // 创建结束时间
	Uid              string `json:"uid,optional"`                   // 用户Uid
	Ip               string `json:"ip,optional"`                    // 请求IP
	InterfaceType    string `json:"interface_type,optional"`        // 请求类型
	InterfaceAddress string `json:"interface_address,optional"`     // 请求路由
	IsRequest        int64  `json:"is_request,default=99,optional"` // 请求结果
}

type Response struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

type ScheduledTasksLogReqest struct {
	Current          int64  `json:"current,default=1,optional"`     // 页码
	PageSize         int64  `json:"page_size,default=10,optional"`  // 数据量
	CreatedStartTime int64  `json:"created_start_time,optional"`    // 创建开始时间
	CreatedEndTime   int64  `json:"created_end_time,optional"`      // 创建结束时间
	ScheduledTasksId string `json:"scheduled_tasks_id"`             // 任务ID
	IsRequest        int64  `json:"is_request,default=99,optional"` // 是否成功
}
