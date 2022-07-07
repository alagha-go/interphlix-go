package casts

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

// get cast info from tmdb's api
func GetCastInfo(name string) Cast {
	Name := strings.ReplaceAll(name, " ", "%20")
	var Cast Cast
	var response Response
	url := fmt.Sprintf("https://api.themoviedb.org/3/search/person?api_key=27cc73002943ca37c5422425af69c720&query=%s&include_adult=true", Name)
	err := json.Unmarshal(GetRequest(url), &response)
	if err != nil {
		return Cast
	}
	for index := range response.Results {
		if response.Results[index].Name == name {
			url := fmt.Sprintf("https://api.themoviedb.org/3/person/%d?api_key=27cc73002943ca37c5422425af69c720", response.Results[index].ID)
			json.Unmarshal(GetRequest(url), &Cast)
			Cast.ProfilePath = "https://image.tmdb.org/t/p/w300_and_h450_bestv2" + Cast.ProfilePath
			return Cast
		}
	}

	return Cast
}


/// seand http get request
func GetRequest(url string, headers ...map[string]string) []byte {
	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return []byte("")
	}
	for _, header := range headers {
		for key, value := range header {
			req.Header.Set(key, value)
		}
	}
	resp, err := client.Do(req)
	if err != nil {
		return []byte("")
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return []byte("")
	}
	return body
}