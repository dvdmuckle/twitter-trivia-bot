package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
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
var url string = "http://www.jservice.io/api"

func Jget(jfill interface{}) error {
	res, err := getClient.Get(url + "/random")
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

func GetQ() (string, string, error) {
	isValid := false
	Jfill := []Trivia{}
	for !isValid {
		err := Jget(&Jfill)
		if err != nil {
			fmt.Println(err)
			return "", "", err
		}
		if Jfill[0].InvalidCount == nil {
			isValid = true
		}
	}
	return fmt.Sprintf("Under \"%s\": %s.", Jfill[0].Category.Title, Jfill[0].Question), fmt.Sprintf("The answer? %s.", Jfill[0].Answer), nil
}
func test() {
	Jfill := []Trivia{}
	Jget(&Jfill)
	question, answer, err := GetQ()
	if err != nil {
		os.Exit(1)
	}
	fmt.Printf("%s %s", question, answer)
}
