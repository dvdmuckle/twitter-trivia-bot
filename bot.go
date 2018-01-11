package main

import (
	"fmt"
	"github.com/dghubble/go-twitter/twitter"
	"github.com/dghubble/oauth1"
	"gopkg.in/robfig/cron.v2"
	"net/http"
	"os"
	"time"
)

var httpClient http.Client
var client *twitter.Client

func Config() {
	config := oauth1.NewConfig(os.Getenv("CONSUMERKEY"), os.Getenv("CONSUMERSECRET"))
	token := oauth1.NewToken(os.Getenv("ACCESSTOKEN"), os.Getenv("ACCESSSECRET"))
	httpClient := config.Client(oauth1.NoContext, token)
	client = twitter.NewClient(httpClient)
}
func TweetQuestion(question string) (int64, error) {
	tweet, _, err := client.Statuses.Update(question, nil)
	if err != nil {
		fmt.Println(err)
		return 0, err
	}
	return tweet.ID, nil
}
func TweetAnswer(answer string, tweetid int64) error {
	replyTo := new(twitter.StatusUpdateParams)
	replyTo.InReplyToStatusID = tweetid
	_, _, err := client.Statuses.Update(answer, replyTo)
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}
func TweetThread() error {
	question, answer, err := GetQ()
	if err != nil {
		fmt.Println(err)
		return err
	}
	tweetid, err := TweetQuestion(question)
	if err != nil {
		fmt.Println(err)
		return err
	}
	time.Sleep(10 * time.Hour)
	err = TweetAnswer(answer, tweetid)
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil

}
func main() {
	Config()
	c := cron.New()
	c.AddFunc("0 0 8 * * *", func() {
		question, answer, err := GetQ()
		if err != nil {
			fmt.Println(err)
			return
		}
		tweetid, err := TweetQuestion(question)
		if err != nil {
			fmt.Println(err)
			return
		}
		time.Sleep(10 * time.Hour)
		err = TweetAnswer(answer, tweetid)
		if err != nil {
			fmt.Println(err)
			return
		}
	})
	c.Start()
}
