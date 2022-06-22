package variables


type Response struct {
	Action								string						`json:"action,omitempty" bson:"action,omitempty"`
	Success								bool						`json:"success,omitempty" bson:"success,omitempty"`
	Failed								bool						`json:"failed,omitempty" bson:"failed,omitempty"`
	Data								interface{}					`json:"data,omitempty" bson:"data,omitempty"`
	Error								string						`json:"error,omitempty" bson:"error,omitempty"`
}


type Secret struct {
	MongoDBUrl							string						`json:"mongodb_url,omitempty" bson:"mongodb_url,omitempty"`
	JwtKey								string						`json:"jwtkey,omitempty" bson:"jwtkey,omitempty"`
}