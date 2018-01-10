package main

import (
	"fmt"
	"github.com/dghubble/go-twitter/twitter"
	"github.com/dghubble/oauth1"
	"net/http"
	"os"
)

var httpClient http.Client
var client *twitter.Client

func Config() {
	config := oauth1.NewConfig(os.Getenv("CONSUMERKEY"), os.Getenv("CONSUMERSECRET"))
	token := oauth1.NewToken(os.Getenv("ACCESSTOKEN"), os.Getenv("ACCESSSECRET"))
	httpClient := config.Client(oauth1.NoContext, token)
	client = twitter.NewClient(httpClient)
	fmt.Println(os.Getenv("CONSUMERKEY"))
	fmt.Println(os.Getenv("CONSUMERSECRET"))
	fmt.Println(os.Getenv("ACCESSTOKEN"))
	fmt.Println(os.Getenv("ACCESSSECRET"))
}
func CreateThread() {
	client.Statuses.Update("This is another test!", nil)
}
func main() {
	Config()
	CreateThread()
}
