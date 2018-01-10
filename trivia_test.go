package main

import (
	"fmt"
	"testing"
)

func TestJservFmt(t *testing.T) {
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
func TestJservGet(t *testing.T) {
	for i := 0; i < 10; i++ {
		Jfill := []Trivia{}
		err := Jget(&Jfill)
		if err != nil {
			t.Error(err)
		}
	}
}
