package main

import (
	parse "artifactMMO/logParsers"
	"io"
	"log"
	"net/http"
)

func main() {
	url := "https://api.artifactsmmo.com/my/DaGuile/action/fight"
	var bearer = "Bearer " + "eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzI1NiJ9.eyJ1c2VybmFtZSI6ImRhZGVidHNpbXVsYXRvciIsInBhc3N3b3JkX2NoYW5nZWQiOiIifQ.c7hgfF8UzfyIZ-ppQyYH-AMMlJw5uUIrBOhWSMDMjCw"
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
