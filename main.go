package main

import (
	"fmt"
	"github.com/ChimeraCoder/anaconda"
	"github.com/sirupsen/logrus"
	"net/url"
	"os"
)

//Tokens struct stores our access tokens and secret keys needed for
//authentication against Twitter REST API
type Credentials struct {
	ConsumerKey string
	ConsumerSecretKey string
	TokenKey string
	TokenSecretKey string
	NameAcc string
}

func getenv(name string) string {
	v := os.Getenv(name)
	if v == "" {
		panic("missing required environment variable " + name)
	}
	return v
}

func findUsermentions(tweet anaconda.Tweet, name string) bool {
	e := false
	users_mentions := tweet.Entities.User_mentions
	for _, v := range users_mentions {
		if v.Screen_name == name {
			e = true
		}
	}
	return e
}

//func generateQuestion(username string) string {
//
//}

func main() {
	tokens := Credentials{
		ConsumerKey:       getenv("TWITTERBOT_CONSUMER_KEY"),
		ConsumerSecretKey: getenv("TWITTERBOT_CONSUMER_SECRETKEY"),
		TokenKey:          getenv("TWITTERBOT_TOKEN_KEY"),
		TokenSecretKey:    getenv("TWITTERBOT_TOKEN_SECRETKEY"),
		NameAcc:		   getenv("TWITTERBOT_NAMEACC"),
	}

	anaconda.SetConsumerKey(tokens.ConsumerKey)
	anaconda.SetConsumerSecret(tokens.ConsumerSecretKey)
	api := anaconda.NewTwitterApi(tokens.TokenKey, tokens.TokenSecretKey)
	//api.SetLogger(log)

	stream := api.PublicStreamFilter(url.Values{
		"track": []string{tokens.NameAcc},
	})

	defer stream.Stop()

	for v := range stream.C {
		t, ok := v.(anaconda.Tweet)
		if !ok {
			logrus.Warningf("received unexpected value of type %T", v)
			continue
		}
		isUserMentioned := findUsermentions(t, tokens.NameAcc)

		if isUserMentioned == false {
			continue
		}

		if t.InReplyToStatusID != 0 {
			continue
		}

		b, err := api.PostTweet("@"+t.User.ScreenName+" hi this is your question", url.Values{
			//TODO:why here IdStr is working (its a string) and string(t.Id which is int64) is not? check it
			"in_reply_to_status_id" : []string{t.IdStr},
		})
		fmt.Printf("%+v\n", b.InReplyToStatusID)
		if err != nil {
			logrus.Errorf("could not replied %d: %v", t.Id, err)
			continue
		}
		logrus.Infof("replied succesfully %d", t.Id)
	}

}