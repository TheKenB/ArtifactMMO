package api

import (
	response "artifactMMO/responseType"
	"bytes"
	"fmt"
	"io"
	"log"
	"net/http"
)

var bearer = "Bearer " + "eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzI1NiJ9.eyJ1c2VybmFtZSI6ImRhZGVidHNpbXVsYXRvciIsInBhc3N3b3JkX2NoYW5nZWQiOiIyMDI0LTA4LTA2IDIzOjUyOjQ1LjI2Mjg2NiJ9.t9BDau2hFiQ1xRpWtGBk2wS9IDTL-zBqy7BIEfURmkc"

func HttpReqBuilder(action response.ActionResponse, bearer string) *http.Request {
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

func ExecQuery(actionPost response.ActionResponse) []byte {
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
	return body
}
