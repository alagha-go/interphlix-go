package payment

import (
	"encoding/json"
	"interphlix/lib/variables"
	"strconv"
	"strings"
)


func (cr *ConversionRate) UnmarshalJSON(data []byte) error {
	Map := make(map[string]string)
	err := json.Unmarshal(data, &Map)
	if err != nil {
		return err
	}
	for key, value := range Map {
		answer, _ := strconv.ParseFloat(value, 64)
		if IsFrom(key) {
			cr.FromCurrency = answer
		}else {
			cr.ToCurrency = answer
		}
	}
	return nil
}

func (amount *Amount) UnmarshalJSON(data []byte) error {
	Map := make(map[string]string)
	err := json.Unmarshal(data, &Map)
	if err != nil {
		return err
	}
	for key, value := range Map {
		answer, _ := strconv.ParseFloat(value, 64)
		if IsFrom(key) {
			amount.Currency = answer
		}else {
			amount.Coin = answer
		}
	}
	return nil
}

// check if string starts with a currency name
func IsFrom(s string) bool {
	secret := variables.LoadSecret()
	for _, wallet := range secret.Wallets {
		if strings.HasPrefix(s, wallet.Currency) {
			return true
		}
	}
	return false
}