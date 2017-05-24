package request

import "errors"

func GetWeatherInfo(location string) (string, error) {
	if len(location) < 1 {
		return "", errors.New("")
	}
	var requestVars = make(map[string]string)
	requestVars["key"] = "77514aacee204dc697a27743f714d434";
	requestVars["cityname"] = "北京";
	stringRes, err := HttpGet("http://api.avatardata.cn/Weather/Query", requestVars)
	return stringRes, err
}
