package main

import (
	"fmt"
	"github.com/ChimeraCoder/anaconda"
	"net/url"
	"os"
)

//Tokens struct stores our access tokens and secret keys needed for
//authentication against Twitter REST API
type Tokens struct {
	ConsumerKey string
	ConsumerSecretKey string
	TokenKey string
	TokenSecretKey string
}

func getenv(name string) string {
	v := os.Getenv(name)
	if v == "" {
		panic("missing required environment variable " + name)
	}
	return v
}

func main() {
	tokens := Tokens{
		ConsumerKey:       getenv("TWITTERBOT_CONSUMER_KEY"),
		ConsumerSecretKey: getenv("TWITTERBOT_CONSUMER_SECRETKEY"),
		TokenKey:          getenv("TWITTERBOT_TOKEN_KEY"),
		TokenSecretKey:    getenv("TWITTERBOT_TOKEN_SECRETKEY"),
	}

	anaconda.SetConsumerKey(tokens.ConsumerKey)
	anaconda.SetConsumerSecret(tokens.ConsumerSecretKey)
	api := anaconda.NewTwitterApi(tokens.TokenKey, tokens.TokenSecretKey)
	//api.SetLogger(log)

	stream := api.PublicStreamFilter(url.Values{
		"track": []string{"#coronavirus"},
	})

	defer stream.Stop()

	api.Tweet
	for v := range stream.C {
		fmt.Printf("%T\n", v)
	}
}