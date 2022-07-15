package payments

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Time struct {
	time.Time
}


type Payment struct {
	ID            					string 						`json:"id,omitempty" bson:"id,omitempty"`
	Txid          					string 						`json:"txid,omitempty" bson:"txid,omitempty"`
	ExplorerURL   					string 						`json:"explorer_url,omitempty" bson:"explorer_url,omitempty"`
	MerchantID    					string 						`json:"merchant_id,omitempty" bson:"merchant_id,omitempty"`
	Type          					string 						`json:"type,omitempty" bson:"type,omitempty"`
	CoinShortName 					string 						`json:"coin_short_name,omitempty" bson:"coin_short_name,omitempty"`
	WalletID      					string 						`json:"wallet_id,omitempty" bson:"wallet_id,omitempty"`
	WalletName    					string 						`json:"wallet_name,omitempty" bson:"wallet_name,omitempty"`
	Address       					string 						`json:"address,omitempty" bson:"address,omitempty"`
	Amount        					string 						`json:"amount,omitempty" bson:"amount,omitempty"`
	Confirmations 					int64  						`json:"confirmations,omitempty" bson:"confirmations,omitempty"`
	Date          					Time 						`json:"date,omitempty" bson:"date,omitempty"`
}

type Invoice struct {
	ID								*primitive.ObjectID			`json:"_id,omitempty" bson:"_id,omitempty"`
	Id              				string           			`json:"id,omitempty" bson:"id,omitempty"`
	InvoiceID       				string           			`json:"invoice_id,omitempty" bson:"invoice_id,omitempty"`
	MerchantID      				string           			`json:"merchant_id,omitempty" bson:"merchant_id,omitempty"`
	URL             				string           			`json:"url,omitempty" bson:"url,omitempty"`
	TotalAmount     				Amount           			`json:"total_amount,omitempty" bson:"total_amount,omitempty"`
	PaidAmount      				Amount           			`json:"paid_amount,omitempty" bson:"paid_amount,omitempty"`
	UsdAmount       				string           			`json:"usd_amount,omitempty" bson:"usd_amount,omitempty"`
	ConversionRate  				ConversionRate   			`json:"conversion_rate,omitempty" bson:"conversion_rate,omitempty"`
	BaseCurrency    				string           			`json:"base_currency,omitempty" bson:"base_currency,omitempty"`
	Coin            				string           			`json:"coin,omitempty" bson:"coin,omitempty"`
	Name            				string           			`json:"name,omitempty" bson:"name,omitempty"`
	Description     				string           			`json:"description,omitempty" bson:"description,omitempty"`
	WalletName      				string           			`json:"wallet_name,omitempty" bson:"wallet_name,omitempty"`
	Address         				string           			`json:"address,omitempty" bson:"address,omitempty"`
	PaymentHistory  				[]PaymentHistory 			`json:"payment_history,omitempty" bson:"payment_history,omitempty"`
	Status          				string           			`json:"status,omitempty" bson:"status,omitempty"`
	StatusCode      				int64            			`json:"status_code,omitempty" bson:"status_code,omitempty"`
	NotifyURL       				string           			`json:"notify_url,omitempty" bson:"notify_url,omitempty"`
	SuceessURL      				string           			`json:"suceess_url,omitempty" bson:"suceess_url,omitempty"`
	FailURL         				string           			`json:"fail_url,omitempty" bson:"fail_url,omitempty"`
	ExpireOn        				string           			`json:"expire_on,omitempty" bson:"expire_on,omitempty"`
	InvoiceDate     				Time           				`json:"invoice_date,omitempty" bson:"invoice_date,omitempty"`
	CustomData1     				string           			`json:"custom_data1,omitempty" bson:"custom_data1,omitempty"`
	CustomData2     				string           			`json:"custom_data2,omitempty" bson:"custom_data2,omitempty"`
	LastUpdatedDate 				Time           				`json:"last_updated_date,omitempty" bson:"last_updated_date,omitempty"`
}

type ConversionRate struct {
	FromCurrency 					float64 					`json:"from_currency,omitempty" bson:"from_currency,omitempty"`
	ToCurrency 						float64 					`json:"to_currency,omitempty" bson:"to_currency,omitempty"`
}

type Amount struct {
	Coin 							float64 					`json:"coin,omitempty" bson:"coin,omitempty"`
	Currency 						float64 					`json:"currency,omitempty" bson:"currency,omitempty"`
}

type PaymentHistory struct {
	Txid         					string 						`json:"txid,omitempty" bson:"txid,omitempty"`
	ExplorerURL  					string 						`json:"explorer_url,omitempty" bson:"explorer_url,omitempty"`
	Amount       					float64 					`json:"amount,omitempty" bson:"amount,omitempty"`
	Date         					Time 						`json:"date,omitempty" bson:"date,omitempty"`
	Confirmation 					int64  						`json:"confirmation,omitempty" bson:"confirmation,omitempty"`
}