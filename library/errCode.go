package library

const CodeErrCommen = 1
const GetUserFail = 10001
const AddPostFail = 20001
const InternalError = 500
const HttpError = 30001
const CodeSucc = 0

func CodeString(errorNo int) string {
	switch errorNo {
	case CodeSucc:
		return ""
	case CodeErrCommen:
		return "发生错误"
	case InternalError:
		return "内部错误"
	//user related
	case GetUserFail:
		return "登录失败，用户名或者密码错误"
	case HttpError:
		return "http请求失败"
	//post related
	case AddPostFail:
		return "发布日记失败"
	default:
		return "error"
	}
}
