package main

import (
	parse "artifactMMO/logParsers"
	menus "artifactMMO/menuMaps"
	"bufio"
	"bytes"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
)

func main() {

	Run()
}

func Run() {
	var running int = 1
	var character string = ""

	for running == 1 {

		// Print Character list
		var bearer = "Bearer " + "eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzI1NiJ9.eyJ1c2VybmFtZSI6ImRhZGVidHNpbXVsYXRvciIsInBhc3N3b3JkX2NoYW5nZWQiOiIyMDI0LTA4LTA2IDIzOjUyOjQ1LjI2Mjg2NiJ9.t9BDau2hFiQ1xRpWtGBk2wS9IDTL-zBqy7BIEfURmkc"
		fmt.Println(menus.GetCharactersMenu())
		character = HandleCharacterStep()
		keyPressed := HandleActionChoice(character)
		actionPost := menus.ActionOperator(keyPressed, character)

		client := &http.Client{}
		req := HttpReqBuilder(actionPost, bearer)
		resp, err := client.Do(req)
		fmt.Println("req: ", actionPost)
		fmt.Println("response: ", resp)
		if err != nil {
			log.Println("Error on response.\nERROR] -", err)
		}
		defer resp.Body.Close()
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			log.Println("Error while reading the response bytes:", err)
		}
		parse.ParseOperator(keyPressed, body, character)
	}
}

func HttpReqBuilder(action menus.ActionResponse, bearer string) *http.Request {
	var jsonStr = []byte(action.Data)
	byteStr := bytes.NewBuffer(jsonStr)
	if len(jsonStr) == 0 {
		req, err := http.NewRequest(action.PostGet, action.Url, nil)
		if err != nil {
			fmt.Println("Fail build request with json")
			fmt.Println(err)
		}
		req.Header.Add("Accept", "application/json")
		req.Header.Add("Authorization", bearer)
		return req
	} else {
		req, err := http.NewRequest(action.PostGet, action.Url, byteStr)
		if err != nil {
			fmt.Println("Fail build request with json")
			fmt.Println(err)
		}
		req.Header.Add("Content-Type", "application/json")
		req.Header.Add("Accept", "application/json")
		req.Header.Add("Authorization", bearer)
		return req
	}
}

func HandleCharacterStep() string {
	// Call the reader to read user's input
	reader := bufio.NewReader(os.Stdin)
	key, err := reader.ReadString('\n')
	if err != nil {
		panic(err)
	}
	if len(strings.TrimSpace(key)) > 0 {
		key = strings.TrimSpace(key)
		char := menus.CharacterOperator(key)
		return char
	} else {
		return HandleCharacterStep()
	}
}

func HandleActionChoice(char string) string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println(menus.ActionMenuOperator(char))
	key2, err := reader.ReadString('\n')
	if err != nil {
		panic(err)
	}
	return strings.TrimSpace(key2)
}
