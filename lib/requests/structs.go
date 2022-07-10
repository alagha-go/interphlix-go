package requests

import (
	"net/url"

	"go.mongodb.org/mongo-driver/bson/primitive"
)


type Request struct {
	ID						primitive.ObjectID						`json:"_id,omitempty" bson:"_id,omitempty"`
	Url						string									`json:"url,omitempty" bson:"url,omitempty"`
	Action					string									`json:"action,omitempty" bson:"action,omitempty"`
	Parameters				url.Values								`json:"parameters,omitempty" bson:"parameters,omitempty"`
	IPData					IPData									`json:"ip_data,omitempty" bson:"ip_data,omitempty"`
}


type IPData struct {
	Query         			string  								`json:"query,omitempty" bson:"query,omitempty"`
	Status        			string  								`json:"status,omitempty" bson:"status,omitempty"`
	Continent     			string  								`json:"continent,omitempty" bson:"continent,omitempty"`
	ContinentCode 			string  								`json:"continentCode,omitempty" bson:"continentCode,omitempty"`
	Country       			string  								`json:"country,omitempty" bson:"country,omitempty"`
	CountryCode   			string  								`json:"countryCode,omitempty" bson:"countryCode,omitempty"`
	Region        			string  								`json:"region,omitempty" bson:"region,omitempty"`
	RegionName    			string  								`json:"regionName,omitempty" bson:"regionName,omitempty"`
	City          			string  								`json:"city,omitempty" bson:"city,omitempty"`
	District      			string  								`json:"district,omitempty" bson:"district,omitempty"`
	Zip           			string  								`json:"zip,omitempty" bson:"zip,omitempty"`
	Lat           			float64 								`json:"lat,omitempty" bson:"lat,omitempty"`
	Lon           			float64 								`json:"lon,omitempty" bson:"lon,omitempty"`
	Timezone      			string  								`json:"timezone,omitempty" bson:"timezone,omitempty"`
	Offset        			int64   								`json:"offset,omitempty" bson:"offset,omitempty"`
	Currency      			string  								`json:"currency,omitempty" bson:"currency,omitempty"`
	ISP           			string  								`json:"isp,omitempty" bson:"isp,omitempty"`
	Org           			string  								`json:"org,omitempty" bson:"org,omitempty"`
	As            			string  								`json:"as,omitempty" bson:"as,omitempty"`
	Asname        			string  								`json:"asname,omitempty" bson:"asname,omitempty"`
	Reverse       			string  								`json:"reverse,omitempty" bson:"reverse,omitempty"`
	Mobile        			bool    								`json:"mobile,omitempty" bson:"mobile,omitempty"`
	Proxy         			bool    								`json:"proxy,omitempty" bson:"proxy,omitempty"`
	Hosting       			bool    								`json:"hosting,omitempty" bson:"hosting,omitempty"`
}