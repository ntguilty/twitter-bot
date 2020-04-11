package main

import (
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

func main() {
	tokens := Tokens{
		ConsumerKey:       os.Getenv("TWITTERBOT_CONSUMER_KEY"),
		ConsumerSecretKey: os.Getenv("TWITTERBOT_CONSUMER_SECRETKEY"),
		TokenKey:       os.Getenv("TWITTERBOT_TOKEN_KEY"),
		TokenSecretKey:    os.Getenv("TWITTERBOT_TOKEN_SECRETKEY"),
	}
	anaconda.SetConsumerKey(tokens.ConsumerKey)
	anaconda.SetConsumerSecret(tokens.ConsumerSecretKey)
	api := anaconda.NewTwitterApi(tokens.TokenKey, tokens.TokenSecretKey)

	stream := api.PublicStreamFilter(url.Values{
		"": []string{"#ntguiltydevbot-test", "@ntguiltydevbot"},
	})

	defer stream.Stop()

}