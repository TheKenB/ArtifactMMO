package main

import (
	parse "artifactMMO/logParsers"
	"io"
	"log"
	"net/http"
)

func main() {
	url := "https://api.artifactsmmo.com/my/DaGuile/action/fight"
	// yes yes this has been changed
	var bearer = "Bearer " + "insertToken"
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
