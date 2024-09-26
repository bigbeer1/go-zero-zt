package common

import (
	"strings"
)

const DefaultCode = 1001 // 默认错误

const VideoErrorCode = 1002 // 视频流错误

const VideoDoErrorCode = 1003 // 摄像头操作错误

const CameraDoErrorCode = 1004 // 摄像头信息获取失败

const VideoNotOnlineCode = 1005 // 摄像头设备不在线

const FileErrorCode = 1200      // 文件找不到
const ImageErrorCode = 1201     // 图片文件找不到
const ImageTypeErrorCode = 1202 // 图片类型错误

const RobotErrorCode = 1301 // 机器人错误

const TokenErrorCode = 1401 // token错误

const TokenNewErrorCode = 1402 // token生成错误

const ReqNotAllCode = 1402 // 请求方式错误

const ReqRoutesErrorCode = 1403 // 没有该请求路由

const ParamErrorCode = 1100 // 请求参数错误

const AuthErrorCode = 1500 // 鉴权错误

const ProgrammeIsNilCode = 9999 // 鉴权错误

type CodeError struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

func NewCodeError(code int, msg string, data interface{}) error {
	return &CodeError{Code: code, Msg: msg, Data: data}
}

func NewDefaultError(msg string) error {
	var errMsg string
	errorMsgs := strings.Split(msg, "=")
	if len(errorMsgs) > 2 {
		errMsg = strings.TrimSpace(errorMsgs[len(errorMsgs)-1])
	} else {
		errMsg = msg
	}
	return NewCodeError(DefaultCode, errMsg, "")
}

func (e *CodeError) Error() string {
	return e.Msg
}
