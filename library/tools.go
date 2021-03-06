package library

import (
	"encoding/json"
	"reflect"
	"net/http"
	"math/rand"
	"time"
	"onesteam/library"
	"strconv"
	"fmt"
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
	ip := r.Header.Get("X-Real-Ip")
	if ip == "" {
		ip = r.RemoteAddr
	}
	return ip
}

var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
var mix = []rune("1234567890abcdefghijklmnopqrstuvwxyz")
var numbers = []rune("1234567890")

func RandSeq(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

func RandNum(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = numbers[rand.Intn(len(numbers))]
	}
	return string(b)
}

func RandMix(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = mix[rand.Intn(len(mix))]
	}
	return string(b)
}

func Substr(str string, start int, end int) string {
	rs := []rune(str)
	length := len(rs)

	if start < 0 || start > length {
		panic("start is wrong")
	}

	if end < 0 || end > length {
		panic("end is wrong")
	}

	return string(rs[start:end])
}

func CreateRandId(firstInt int) int64 {
	nowTimer := time.Now().Unix()
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	randNum := r.Intn(1000)
	return nowTimer*100000 + int64(randNum*10) + int64(firstInt%10)
}

func PanicFunc(recoverErr interface{}) (string, error) {
	var output string
	var errorFlag error
	if recoverErr != nil {
		switch err := recoverErr.(type) {
		case int:
			output, _ = library.ReturnJsonWithError(library.CodeErrApi, "ref", "")
			break
		case string:
			output, _ = library.ReturnJsonWithError(library.CodeErrApi, err, "")
			break
		default:
			output, _ = library.ReturnJsonWithError(library.CodeErrApi, "ref", "")
		}
		errorFlag = fmt.Errorf("error alert")
	}
	return output, errorFlag
}
