package main

import (
	"fmt"
	"github.com/ChimeraCoder/anaconda"
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
	anaconda.SetConsumerKey(apiKey)
	anaconda.SetConsumerSecret(apiSecretKey)
	api := anaconda.NewTwitterApi(accessToken, accessTokenSecret)
	res, _ := api.GetSearch("golang", nil)
	for _ , tweet := range res.Statuses {
		fmt.Print(tweet.Text)
	}

}