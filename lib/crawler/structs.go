package crawler

import "interphlix/lib/movies"

type LinkResponse struct {
	Type                                        string                                   `json:"type,omitempty" bson:"type,omitempty"`   
	Link                                        string                                   `json:"link,omitempty" bson:"link,omitempty"`   
	Sources                                     []interface{}                            `json:"sources,omitempty" bson:"sources,omitempty"`
	Tracks                                      []interface{}                            `json:"tracks,omitempty" bson:"tracks,omitempty"` 
	Title                                       string                                   `json:"title,omitempty" bson:"title,omitempty"`  
}

type Setter struct {
	Url				string
	Server			*movies.Server
	Servers			[]movies.Server
	Available		bool
}