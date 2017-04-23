package library

import (
	"encoding/json"
	"github.com/astaxie/beego/logs"
)

type (
	returnFormat struct {
		ErrNo  int
		ErrMsg string
		Data   interface{}
	}
)

func ReturnJsonWithError(errNo int, errMsg string, data interface{}) (res string, err error) {

	if data == nil{
		data = ""
	}
	if errMsg == "ref" {
		errMsg = CodeString(errNo)
	}
	formatter := new(returnFormat)
	formatter.ErrNo = errNo
	formatter.ErrMsg = errMsg
	formatter.Data = data

	result, err := json.Marshal(formatter)
	if err != nil {
		logs.Warn(err)
		return "", err
	}

	return string(result), nil
}
