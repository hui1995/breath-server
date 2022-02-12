package utils

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"strings"
)

//http://img.cssmoban.com/UploadFiles/2021/17/2021092419140967590.jpg
func DoGet(url string) io.ReadCloser {
	client := &http.Client{}

	request, _ := http.NewRequest("GET", url, nil)
	request.Header.Add("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/83.0.4103.61 Safari/537.36")
	response, _ := client.Do(request)
	defer response.Body.Close()
	if response.StatusCode == 200 {
		return response.Body

	}
	return nil
}
func DoPOST(url string, data map[string]interface{}) io.ReadCloser {
	client := &http.Client{}
	bytesData, _ := json.Marshal(data)
	reader := bytes.NewReader(bytesData)

	request, _ := http.NewRequest("POST", url, reader)
	request.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/83.0.4103.61 Safari/537.36")
	request.Header.Set("Content-Type", "application/json")
	response, _ := client.Do(request)
	defer response.Body.Close()
	if response.StatusCode == 200 {
		return response.Body

	}
	return nil
}
func DoPostFrom(url string, payload string) []byte {
	method := "POST"

	payload2 := strings.NewReader(payload)

	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload2)

	if err != nil {
		fmt.Println(err)
		return nil
	}
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	fmt.Println(string(body))
	return body
}
