package library

import (
	"encoding/json"
	"reflect"
	"net/http"
	"math/rand"
	"strings"
	"net/smtp"
	"qiniupkg.com/x/log.v7"
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

func SendToMail(to, subject, body, mailtype string) error {
	host := "smtpdm.aliyun.com:80"
	user := "service@mail.onestory.cn‍"
	password := "Zyy45612301Mark"
	hp := strings.Split(host, ":")
	auth := smtp.PlainAuth("", user, password, hp[0])
	var content_type string
	if mailtype == "html" {
		content_type = "Content-Type: text/" + mailtype + "; charset=UTF-8"
	} else {
		content_type = "Content-Type: text/plain" + "; charset=UTF-8"
	}
	to = to + ";" +user
	msg := []byte("To: " + to + "\r\nFrom: " + user + ">\r\nSubject: " + subject + "\r\n" + content_type + "\r\n\r\n" + body)
	send_to := strings.Split(to, ";")
	err := smtp.SendMail(host, auth, user, send_to, msg)
	log.Warn(err)
	return err
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
