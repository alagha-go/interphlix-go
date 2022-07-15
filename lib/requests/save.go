package requests

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	// "go.mongodb.org/mongo-driver/bson/primitive"
)


func SaveRequest(req *http.Request, action string) {
	// key := 
	// request := Request{ID: primitive.NewObjectID(), Url: req.URL.String(), Action: action, Parameters: req.URL.Query(), IPData: GetIpData(req)}
}

func GetIpData(req *http.Request) IPData {
	var ip string
	var ipdata IPData
	client := &http.Client{}
	ips := strings.Split(req.Header.Get("X-Forwarded-For"), ",")
	if len(ips) > 0 {
		ip = ips[0]
	}
	url := fmt.Sprintf("http://ip-api.com/json/%s?fields=66846719", ip)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return ipdata
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
	res, err := client.Do(req)
	if err != nil {
		return ipdata
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return ipdata
	}
	err = json.Unmarshal(body, &ipdata)
	if err != nil {
		return ipdata
	}
	return ipdata
}