package casts

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// get cast info from tmdb's api
func GetCastInfo(name string) Cast {
	var Cast Cast
	var response Response
	url := fmt.Sprintf("https://api.themoviedb.org/3/search/person?api_key=27cc73002943ca37c5422425af69c720&query=%s&include_adult=true", name)
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
	req.Header.Set("User-Agent", "Mozilla/5.0 (X11; Linux x86_64; rv:101.0) Gecko/20100101 Firefox/101.0")
	req.Header.Set("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,*/*;q=0.8")
	req.Header.Set("Accept-Language", "en-US,en;q=0.5")
	req.Header.Set("Accept-Encoding", "gzip, deflate, br")
	req.Header.Set("Connection", "keep-alive")
	req.Header.Set("Upgrade-Insecure-Requests", "1")
	req.Header.Set("Sec-Fetch-Dest", "document")
	req.Header.Set("Sec-Fetch-Mode", "navigate")
	req.Header.Set("Sec-Fetch-Site", "none")
	req.Header.Set("Sec-Fetch-User", "?1")
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