package variables


type Secret struct {
	MongoDBUrl							string						`json:"mongodb_url,omitempty" bson:"mongodb_url,omitempty"`
}