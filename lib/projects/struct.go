package projects

import "go.mongodb.org/mongo-driver/bson/primitive"


type Project struct {
	ID							primitive.ObjectID							`json:"_id,omitempty" bson:"_id,omitempty"`
	Name						string										`json:"name,omitempty" bson:"name,omitempty"`
	AccountID					primitive.ObjectID							`json:"account_id,omitempty" bson:"account_id,omitempty"`
	ApiKeys						[]ApiKey									`json:"api_keys,omitempty" bson:"api_keys"`
	RequestsLimit				int64										`json:"requests_limit,omitempty" bson:"requests_limit,omitempty"`
}


type ApiKey struct {
	Name						string										`json:"name,omitempty" bson:"name,omitempty"`
	Key							string										`json:"key,omitempty" bson:"key,omitempty"`
}