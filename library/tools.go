package library

import (
	"encoding/json"
	"reflect"
	"net/http"
)

func Json2Map(input string) (map[string]interface{}, error) {
	var result map[string]interface{}
	if err := json.Unmarshal([]byte(input), &result); err != nil {
		return nil, err
	}
	return result, nil
}

func Struct2Map(obj interface{}) map[string]interface{} {
	t := reflect.TypeOf(obj)
	v := reflect.ValueOf(obj)

	var data = make(map[string]interface{})
	for i := 0; i < t.NumField(); i++ {
		data[t.Field(i).Name] = v.Field(i).Interface()
	}
	return data
}

func GetClientIp(r *http.Request) string {
	ip := r.Header.Get("remote_addr")
	if ip == "" {
		ip = r.RemoteAddr
	}
	return ip
}
