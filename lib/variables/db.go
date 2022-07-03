package variables

import (
	"bytes"
	"context"
	"encoding/json"
	"io/ioutil"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	Client *mongo.Client
	Local *mongo.Client
)


func ConnectToDB() {
	ConnectToLocl()
	var err error
	ctx := context.Background()
	secret := LoadSecret()
	clientOptions := options.Client().ApplyURI(secret.MongoDBUrl)
	Client, err = mongo.Connect(ctx, clientOptions)
	HandleError(err)
}

func ConnectToLocl() {
	var err error
	ctx := context.Background()
	secret := LoadSecret()
	clientOptions := options.Client().ApplyURI(secret.LocalUrl)
	Local, err = mongo.Connect(ctx, clientOptions)
	HandleError(err)
}

func LoadSecret() Secret {
	var secret Secret
	data, err := ioutil.ReadFile("./secret.json")
	HandleError(err)
	err = json.Unmarshal(data, &secret)
	HandleError(err)
	return secret
}

func HandleError(err error) {
	if err != nil {
		log.Panic(err)
	}
}


func JsonMarshal(data interface{}) []byte {
	var buff bytes.Buffer
	encoder := json.NewEncoder(&buff)
	encoder.SetEscapeHTML(false)
	encoder.Encode(data)
	return buff.Bytes()
}