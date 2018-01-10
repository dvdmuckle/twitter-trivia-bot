package main

import (
	"fmt"
	"testing"
)

func TestJserv(t *testing.T) {
	for i := 0; i < 10; i++ {
		question, answer, err := GetQ()
		if err != nil {
			t.Error(err)
		}
		if question == "" {
			t.Error("Question string empty!")
		}
		if answer == "" {
			t.Error("Answer string empty!")
		}
		fmt.Println(question)
		fmt.Println(answer)
	}
}
