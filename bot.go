package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

type Trivia struct {
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
var url string = "http://www.jservice.io/api/random?count=10"

func Jget(jfill interface{}) error {
	res, err := getClient.Get(url)
	if err != nil {
		fmt.Println(err)
		return err
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return err
	}
	return json.Unmarshal(body, jfill)
}

func main() {
	Jfill := []Trivia{}
	Jget(&Jfill)
	fmt.Printf("Under \"%s\": %s.\n", Jfill[0].Category.Title, Jfill[0].Question)
}
