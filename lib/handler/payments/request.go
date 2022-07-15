package payments

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
)



func PostRequest(url string, values map[string]io.Reader, pointer any) {
	client := &http.Client{}
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	for key, value := range values {
		part, _ := writer.CreateFormField(key)
		io.Copy(part, value)
	}
	writer.Close()
	req, err := http.NewRequest("POST", url, body)
	req.Header.Add("Content-Type", writer.FormDataContentType())
	if err != nil {
		return
	}
	res, err := client.Do(req)
	if err != nil {
		return
	}
	defer res.Body.Close()
	data, _ := ioutil.ReadAll(res.Body)
	err = json.Unmarshal(data, pointer)
	if err != nil {
		fmt.Println(err.Error())
	}
}