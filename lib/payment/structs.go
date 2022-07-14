package payment


type Payment struct {
	ID            					string 						`json:"id,omitempty" json:"id,omitempty"`
	Txid          					string 						`json:"txid,omitempty" json:"txid,omitempty"`
	ExplorerURL   					string 						`json:"explorer_url,omitempty" json:"explorer_url,omitempty"`
	MerchantID    					string 						`json:"merchant_id,omitempty" json:"merchant_id,omitempty"`
	Type          					string 						`json:"type,omitempty" json:"type,omitempty"`
	CoinShortName 					string 						`json:"coin_short_name,omitempty" json:"coin_short_name,omitempty"`
	WalletID      					string 						`json:"wallet_id,omitempty" json:"wallet_id,omitempty"`
	WalletName    					string 						`json:"wallet_name,omitempty" json:"wallet_name,omitempty"`
	Address       					string 						`json:"address,omitempty" json:"address,omitempty"`
	Amount        					string 						`json:"amount,omitempty" json:"amount,omitempty"`
	Confirmations 					int64  						`json:"confirmations,omitempty" json:"confirmations,omitempty"`
	Date          					string 						`json:"date,omitempty" json:"date,omitempty"`
}

type Invoice struct {
	ID              				string           			`json:"id,omitempty" json:"id,omitempty"`
	InvoiceID       				string           			`json:"invoice_id,omitempty" json:"invoice_id,omitempty"`
	MerchantID      				string           			`json:"merchant_id,omitempty" json:"merchant_id,omitempty"`
	URL             				string           			`json:"url,omitempty" json:"url,omitempty"`
	TotalAmount     				Amount           			`json:"total_amount,omitempty" json:"total_amount,omitempty"`
	PaidAmount      				Amount           			`json:"paid_amount,omitempty" json:"paid_amount,omitempty"`
	UsdAmount       				string           			`json:"usd_amount,omitempty" json:"usd_amount,omitempty"`
	ConversionRate  				ConversionRate   			`json:"conversion_rate,omitempty" json:"conversion_rate,omitempty"`
	BaseCurrency    				string           			`json:"base_currency,omitempty" json:"base_currency,omitempty"`
	Coin            				string           			`json:"coin,omitempty" json:"coin,omitempty"`
	Name            				string           			`json:"name,omitempty" json:"name,omitempty"`
	Description     				string           			`json:"description,omitempty" json:"description,omitempty"`
	WalletName      				string           			`json:"wallet_name,omitempty" json:"wallet_name,omitempty"`
	Address         				string           			`json:"address,omitempty" json:"address,omitempty"`
	PaymentHistory  				[]PaymentHistory 			`json:"payment_history,omitempty" json:"payment_history,omitempty"`
	Status          				string           			`json:"status,omitempty" json:"status,omitempty"`
	StatusCode      				int64            			`json:"status_code,omitempty" json:"status_code,omitempty"`
	NotifyURL       				string           			`json:"notify_url,omitempty" json:"notify_url,omitempty"`
	SuceessURL      				string           			`json:"suceess_url,omitempty" json:"suceess_url,omitempty"`
	FailURL         				string           			`json:"fail_url,omitempty" json:"fail_url,omitempty"`
	ExpireOn        				string           			`json:"expire_on,omitempty" json:"expire_on,omitempty"`
	InvoiceDate     				string           			`json:"invoice_date,omitempty" json:"invoice_date,omitempty"`
	CustomData1     				string           			`json:"custom_data1,omitempty" json:"custom_data1,omitempty"`
	CustomData2     				string           			`json:"custom_data2,omitempty" json:"custom_data2,omitempty"`
	LastUpdatedDate 				string           			`json:"last_updated_date,omitempty" json:"last_updated_date,omitempty"`
}

type ConversionRate struct {
	FromCurrency 						string 						`json:"from_currency,omitempty"`
	ToCurrency 							string 						`json:"to_currency,omitempty"`
}

type Amount struct {
	Coin 							string 						`json:"coin,omitempty"`
	Currency 						string 						`json:"currency,omitempty"`
}

type PaymentHistory struct {
	Txid         					string 						`json:"txid,omitempty"`
	ExplorerURL  					string 						`json:"explorer_url,omitempty"`
	Amount       					string 						`json:"amount,omitempty"`
	Date         					string 						`json:"date,omitempty"`
	Confirmation 					int64  						`json:"confirmation,omitempty"`
}