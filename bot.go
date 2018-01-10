package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

type Trivia []struct {
	Airdate  time.Time `json:"airdate"`
	Answer   string    `json:"answer"`
	Category struct {
		CluesCount int       `json:"clues_count"`
		CreatedAt  time.Time `json:"created_at"`
		ID         int       `json:"id"`
		Title      string    `json:"title"`
		UpdatedAt  time.Time `json:"updated_at"`
	} `json:"category"`
	CategoryID   int         `json:"category_id"`
	CreatedAt    time.Time   `json:"created_at"`
	GameID       interface{} `json:"game_id"`
	ID           int         `json:"id"`
	InvalidCount interface{} `json:"invalid_count"`
	Question     string      `json:"question"`
	UpdatedAt    time.Time   `json:"updated_at"`
	Value        int         `json:"value"`
}

var getClient = &http.Client{Timeout: 10 * time.Second}
var url string = "http://www.jservice.io/api/random"

func Jget() ([]byte, error) {
	res, err := getClient.Get(url)
	if err != nil {
		return nil, err
	}
	body, err := ioutil.ReadAll(res.Body)

	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	return body, nil
}

func Jparse(body []byte) (*Trivia, error) {
	var t = new(Trivia)
	err := json.Unmarshal(body, &t)
	if err != nil {
		return nil, err
	}
	return t, err
}

func main() {
	res, err := Jget()
	if err != nil {
		fmt.Println(err)
	}
	NewTriv, err := Jparse(res)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(NewTriv.Question)

}
