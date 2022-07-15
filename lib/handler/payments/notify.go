package payments

import (
	"encoding/json"
	"fmt"
	"interphlix/lib/payments"
	"interphlix/lib/variables"
	"io"
	"io/ioutil"
	"net/http"
	"reflect"
	"strings"
)

type Response struct {
	Flag   						int64  							`json:"flag"`  
	Msg    						string 							`json:"msg"`   
	Action 						string 							`json:"action"`
	Data   						payments.Invoice   				`json:"data"`
}

func Hook(res http.ResponseWriter, req *http.Request) {
	req.ParseMultipartForm(0)
	payment := payments.Payment{}
	t := reflect.ValueOf(payment)
	for index := 0; index < t.Type().NumField(); index++ {
		name := t.Type().Field(index).Tag.Get("json")
		name = strings.ReplaceAll(name, ",omitempty", "")
		field := t.Type().Field(index)
		value, err := payments.ChangeType(req.FormValue(name), field.Type)
		if err != nil {
			continue
		}
		reflect.ValueOf(&payment).Elem().FieldByName(field.Name).Set(value)
	}
	data, _ := json.Marshal(payment)
	ioutil.WriteFile("./payment.json", data, 0755)
}

func Notify(res http.ResponseWriter, req *http.Request) {
	var response Response
	response.Data = payments.Invoice{}
	req.ParseMultipartForm(0)
	form := req.Form
	id := ""
	for key, value := range form {
		if key == "invoice_id" && len(value) > 0{
			id = value[0]
			break
		}
	}
	if id == "" {
		return
	}
	wallet := GetWallet(id[:3])
	if wallet.Coin == "" {
		return
	}
	url := fmt.Sprintf("https://coinremitter.com/api/v3/%s/get-invoice", wallet.Coin)
	values := map[string]io.Reader{
		"api_key": strings.NewReader(wallet.APIKey),
		"password": strings.NewReader(wallet.Password),
		"invoice_id": strings.NewReader(id),
	}
	PostRequest(url, values, &response)
	data, _ := json.Marshal(response.Data)
	ioutil.WriteFile("./invoice.json", data, 0755)
}

func Success(res http.ResponseWriter, req *http.Request) {
	data, _ := ioutil.ReadAll(req.Body)
	ioutil.WriteFile("./success.json", data, 0755)
}

func Fail(res http.ResponseWriter, req *http.Request) {
	data, _ := ioutil.ReadAll(req.Body)
	ioutil.WriteFile("./fail.json", data, 0755)
}


func GetWallet(coin string) variables.Wallet {
	secret := variables.LoadSecret()
	for _, wallet := range secret.Wallets {
		if strings.HasPrefix(wallet.Coin, coin) {
			return wallet
		}
	}
	return variables.Wallet{}
}