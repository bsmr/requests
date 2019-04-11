package main

import (
	"fmt"
	"io/ioutil"
	"net/url"

	"github.com/xuanbo/requests"
)

func main() {
	getRow()
	getText()
	postForm()
	postJson()
	handler()
	save()
	getJson()
}

func getRow() {
	raw, err := requests.Get("http://127.0.0.1:8080/ping").
		Params(url.Values{
			"param1": {"value1"},
			"param2": {"123"},
		}).
		Send().
		Raw()
	if err != nil {
		panic(err)
	}
	fmt.Println(raw)
}

func getText() {
	text, err := requests.Get("http://127.0.0.1:8080/ping").
		Params(url.Values{
			"param1": {"value1"},
			"param2": {"123"},
		}).
		Send().
		Text()
	if err != nil {
		panic(err)
	}
	fmt.Println(text)
}

func postForm() {
	text, err := requests.Post("http://127.0.0.1:8080/ping").
		Params(url.Values{
			"param1": {"value1"},
			"param2": {"123"},
		}).
		Form(url.Values{
			"form1": {"value1"},
			"form2": {"123"},
		}).
		Send().
		Text()
	if err != nil {
		panic(err)
	}
	fmt.Println(text)
}

func postJson() {
	text, err := requests.Post("http://127.0.0.1:8080/ping").
		Params(url.Values{
			"param1": {"value1"},
			"param2": {"123"},
		}).
		Json(map[string]interface{}{
			"json1": "value1",
			"json2": 2,
		}).
		Send().
		Text()
	if err != nil {
		panic(err)
	}
	fmt.Println(text)
}

func save() {
	err := requests.Get("https://www.cnblogs.com/bener/p/10683404.html").
		Send().
		Save("./10683404.html")
	if err != nil {
		panic(err)
	}
}

func getJson() {
	var v map[string]interface{}
	err := requests.Post("http://127.0.0.1:8080/ping").
		Params(url.Values{
			"param1": {"value1"},
			"param2": {"123"},
		}).
		Json(map[string]interface{}{
			"json1": "value1",
			"json2": 2,
		}).
		Send().
		Json(&v)
	if err != nil {
		panic(err)
	}
	fmt.Println(v)
}

func handler() {
	result := requests.Post("http://127.0.0.1:8080/ping").
		Params(url.Values{
			"param1": {"value1"},
			"param2": {"123"},
		}).
		Json(map[string]interface{}{
			"json1": "value1",
			"json2": 2,
		}).
		Send()
	if result.Err != nil {
		panic(result.Err)
	}

	b, err := ioutil.ReadAll(result.Resp.Body)
	if err != nil {
		panic(err)
	}
	defer result.Resp.Body.Close()

	fmt.Println(string(b))
}
