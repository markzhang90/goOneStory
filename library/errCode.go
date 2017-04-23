package library

const CodeErrCommen = 1
const GetUserFail = 10001
const InternalError = 500
const CodeSucc = 0

func CodeString(errorNo int) string {
	switch errorNo {
	case CodeSucc:
		return ""
	case CodeErrCommen:
		return "发生错误"
	case InternalError:
		return "内部错误"
	case GetUserFail:
		return "登录失败，用户名或者密码错误"
	default:
		return "error"
	}
}
