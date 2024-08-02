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
	url := ""
	reader := bufio.NewReader(os.Stdin)

	for running == 1 {
		// Print the instruction to the reader in the console

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
		url = menus.GetAction(key2, character)
		// yes yes this has been changed
		var bearer = "Bearer " + "eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzI1NiJ9.eyJ1c2VybmFtZSI6ImRhZGVidHNpbXVsYXRvciIsInBhc3N3b3JkX2NoYW5nZWQiOiIyMDI0LTA4LTAyIDE1OjAzOjIxLjM3MzUzNCJ9.gcnIvBLjqesqXVBQ6zSr78ELK6B-trsjDWRtCr-DgPg"
		req, err := http.NewRequest("POST", url, nil)
		if err != nil {
			log.Println("Error trying to make post request")
		}
		req.Header.Add("Accept", "application/json")
		req.Header.Add("Content-Type", "application/json")
		req.Header.Add("Authorization", bearer)

		client := &http.Client{}
		resp, err := client.Do(req)
		if err != nil {
			log.Println("Error on response.\nERROR] -", err)
		}
		defer resp.Body.Close()

		body, err := io.ReadAll(resp.Body)
		if err != nil {
			log.Println("Error while reading the response bytes:", err)
		}
		parse.ParseCombatLog(body)
	}
}
