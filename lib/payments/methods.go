package payments

import (
	"encoding/json"
	"fmt"
	"interphlix/lib/variables"
	"reflect"
	"strconv"
	"strings"
	"time"
)

func (T *Time) UnmarshalJSON(data []byte) error {
	date := string(data)
	date = strings.ReplaceAll(date, "\"", "")
	var err error
	T.Time, err = time.Parse(time.RFC3339, date)
	if err != nil {
		return err
	}
	return nil
}


func (ph *PaymentHistory) UnmarshalJSON(data []byte) error {
	Map := make(map[string]any)
	err := json.Unmarshal(data, &Map)
	if err != nil {
		return err
	}
	t := reflect.ValueOf(*ph)
	for key, value := range Map {
		var Value reflect.Value
		Int64 := reflect.TypeOf(int64(0))
		index := GetFieldIndex(t, key)
		field := t.Type().Field(index)
		if field.Type == Int64 {
			i, _ := value.(int64)
			Value = reflect.ValueOf(i)
		}else {
			Value, err = ChangeType(fmt.Sprintf("%v", value), field.Type)
			if err != nil {
				continue
			}
		}
		reflect.ValueOf(ph).Elem().FieldByName(field.Name).Set(Value)
	}
	return nil
}


func (cr *ConversionRate) UnmarshalJSON(data []byte) error {
	Map := make(map[string]float64)
	err := json.Unmarshal(data, &Map)
	if err != nil {
		return err
	}
	for key, value := range Map {
		if IsFrom(key) {
			cr.FromCurrency = value
		}else {
			cr.ToCurrency = value
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