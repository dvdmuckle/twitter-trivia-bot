package main

import (
	"fmt"
	"github.com/dghubble/go-twitter/twitter"
	"github.com/dghubble/oauth1"
	"github.com/jasonlvhit/gocron"
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
	fmt.Printf("Sent question %s at %s\n", question, time.Now())
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
	fmt.Printf("Sent answer %s at %s\n", answer, time.Now())
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
	time.Sleep(8 * time.Hour)
	err = TweetAnswer(answer, tweetid)
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil

}
func main() {
	Config()
	gocron.Every(1).Day().At("18:00").Do(TweetThread)
	<-gocron.Start()
}
