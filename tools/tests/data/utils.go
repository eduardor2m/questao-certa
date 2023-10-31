package data

import (
	"fmt"
	"os"
	"path"
	"runtime"
)

const (
	base      = "../jsons"
	questions = "/questions"
	request   = "/request"
	user      = "/user"
	filter    = "/filter"
)

func GetQuestionMock(filename string) []byte {
	_, pwd, _, _ := runtime.Caller(0)
	filepath := path.Join(pwd, base, questions, request, filename)
	content, err := os.ReadFile(filepath)
	if err != nil {
		fmt.Println("[ERROR] Error while reading a JSON file: " + err.Error())
		panic(err)
	}

	return content
}

func GetSignInMock(filename string) []byte {
	_, pwd, _, _ := runtime.Caller(0)
	filepath := path.Join(pwd, base, user, request, filename)
	content, err := os.ReadFile(filepath)
	if err != nil {
		fmt.Println("[ERROR] Error while reading a JSON file: " + err.Error())
		panic(err)
	}

	return content
}

func GetSignUpMock(filename string) []byte {
	_, pwd, _, _ := runtime.Caller(0)
	filepath := path.Join(pwd, base, user, request, filename)
	content, err := os.ReadFile(filepath)
	if err != nil {
		fmt.Println("[ERROR] Error while reading a JSON file: " + err.Error())
		panic(err)
	}

	return content
}

func GetFilterMock(filename string) []byte {
	_, pwd, _, _ := runtime.Caller(0)
	filepath := path.Join(pwd, base, questions, filter, filename)
	content, err := os.ReadFile(filepath)
	if err != nil {
		fmt.Println("[ERROR] Error while reading a JSON file: " + err.Error())
		panic(err)
	}

	return content
}
