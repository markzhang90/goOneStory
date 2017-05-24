package request

import (
	"net/http"
	"net/url"
	"io/ioutil"
	"fmt"
	"strings"
)

// 简单直接的GET请求
func HttpGet(urlStr string, queryList map[string]string) (string, error){
	var temp = make([]string, 0, len(queryList))

	for key, value := range queryList{
		stringQuery :=  key + "=" + value
		temp = append(temp, stringQuery)
	}

	queryVar := strings.Join(temp, "&")
	fullQuery := urlStr + "?" + queryVar
	resp, err := http.Get(fullQuery)
	if err != nil {
		return "", err
		// handle error
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
		// handle error
	}
	return string(body), nil
}

// POST请求 -- 使用http.Post()方法
//Tips：使用这个方法的话，第二个参数要设置成”application/x-www-form-urlencoded”，否则post参数无法传递。

func HttpPost(urlStr string) {
	resp, err := http.Post(urlStr,
		"application/x-www-form-urlencoded",
		strings.NewReader("name=cjb"))
	if err != nil {
		fmt.Println(err)
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		// handle error
	}

	fmt.Println(string(body))
}

// POST请求 -- 使用http.PostForm()方法
func HttpPostForm(urlStr string) {
	resp, err := http.PostForm(urlStr,
		url.Values{"key": {"Value"}, "id": {"123"}})

	if err != nil {
		// handle error
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		// handle error
	}

	fmt.Println(string(body))
}

// 复杂的请求（设置头参数、cookie之类的数据），可以使用http.Client的Do()方法</strong>
func HttpDo() {
	client := &http.Client{}

	req, err := http.NewRequest("POST", "http://www.baidu.com", strings.NewReader("name=cjb"))
	if err != nil {
		// handle error
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("Cookie", "name=anny")

	resp, err := client.Do(req)

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		// handle error
	}

	fmt.Println(string(body))
}