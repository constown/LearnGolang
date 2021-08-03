package main

import (
	"fmt"
	"io/ioutil"
	http "net/http"
	"net/url"
)

// http client

func main() {
	//res, err := http.Get("127.0.0.1:9090/posts/xx?name=周玲玲&age=19")
	data := url.Values{}
	urlObj, _ := url.Parse("http://127.0.0.1:9090/posts/xx")
	data.Set("name", "周玲玲")
	data.Set("age", "19")
	urlStr := data.Encode()
	urlObj.RawQuery = urlStr
	req, err := http.NewRequest("GET", urlObj.String(), nil)
	if err != nil {
		return
	}
	res, err := http.DefaultClient.Do(req)
	resBody, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return
	}
	fmt.Println(string(resBody))
}
