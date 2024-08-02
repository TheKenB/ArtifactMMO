package main

import (
	parse "artifactMMO/logParsers"
	menus "artifactMMO/menuMaps"
	"bufio"
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
	actionPost := menus.ActionResponse{
		Url:     "",
		PostGet: "",
		Data:    "",
	}
	reader := bufio.NewReader(os.Stdin)

	for running == 1 {
		// Print Character list
		fmt.Println(menus.GetCharacterMenu())

		// Call the reader to read user's input
		key, err := reader.ReadString('\n')
		if err != nil {
			panic(err)
		}
		key = strings.TrimSpace(key)
		character = menus.GetCharacter(key)
		fmt.Println(menus.GetActionsMenu())

		// Call the reader to read user's input
		key2, err := reader.ReadString('\n')
		if err != nil {
			panic(err)
		}
		key2 = strings.TrimSpace(key2)
		actionPost = menus.GetAction(key2, character)

		var bearer = "Bearer " + "PlaceHolder"
		if err != nil {
			fmt.Println("Can't marshal jsonvalue: ", actionPost.Data)
		}

		client := &http.Client{}
		req := HttpReqBuilder(actionPost, bearer)
		resp, err := client.Do(req)
		if err != nil {
			log.Println("Error on response.\nERROR] -", err)
		}
		defer resp.Body.Close()
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			log.Println("Error while reading the response bytes:", err)
		}
		parse.ParseOperator(key2, body)
	}
}

func HttpReqBuilder(action menus.ActionResponse, bearer string) *http.Request {
	payload := strings.NewReader("{\n  \"x\": 0,\n  \"y\": 0\n}")
	req, err := http.NewRequest(action.PostGet, action.Url, payload)
	fmt.Println(action)
	if err != nil {
		fmt.Println("Fail build request with json")
		fmt.Println(err)
	}
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Authorization", bearer)
	return req
}
