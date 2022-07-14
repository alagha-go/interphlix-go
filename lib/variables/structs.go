package variables

import "go.mongodb.org/mongo-driver/bson/primitive"


type Response struct {
	Action								string						`json:"action,omitempty" bson:"action,omitempty"`
	Success								bool						`json:"success,omitempty" bson:"success,omitempty"`
	Failed								bool						`json:"failed,omitempty" bson:"failed,omitempty"`
	Data								interface{}					`json:"data,omitempty" bson:"data,omitempty"`
	Error								string						`json:"error,omitempty" bson:"error,omitempty"`
}

type Log struct {
	ID									primitive.ObjectID			`json:"_id,omitempty" bson:"_id,omitempty"`
	Error								string						`json:"error,omitempty" bson:"error,omitempty"`
	Package								string						`json:"package,omitempty" bson:"package,omitempty"`
	Function							string						`json:"function,omitempty" bson:"function,omitempty"`
}


type Secret struct {
	MongoDBUrl							string						`json:"mongodb_url,omitempty" bson:"mongodb_url,omitempty"`
	LocalUrl							string						`json:"local_url,omitempty" bson:"local_url,omitempty"`
	JwtKey								string						`json:"jwtkey,omitempty" bson:"jwtkey,omitempty"`
	Wallets								[]Wallet					`json:"wallets,omitempty" bson:"wallets,omitempty"`
}

type Wallet struct {
	APIKey   							string 						`json:"api_key,omitempty" bson:"api_key,omitempty"`
	Password 							string 						`json:"password,omitempty" bson:"password,omitempty"`
	Currency 							string 						`json:"currency,omitempty" bson:"currency,omitempty"`
	Coin     							string 						`json:"coin,omitempty" bson:"coin,omitempty"`
	CoinIcon 							string 						`json:"coin_icon,omitempty" bson:"coin_icon,omitempty"`
	Name     							string 						`json:"name,omitempty" bson:"name,omitempty"`
	Expiry								int							`json:"expiry,omitempty" bson:"expiry,omitempty"`
}