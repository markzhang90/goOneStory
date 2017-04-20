package library

const CodeErrCommen = 1
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
	default:
		return ""
	}
}
