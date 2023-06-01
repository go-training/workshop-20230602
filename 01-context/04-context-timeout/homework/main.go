package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type response struct {
	UserID    int    `json:"userId"`
	ID        int    `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

type callResponse struct {
	Resp *response
	Err  error
}

func helper(ctx context.Context) *callResponse {
	resp, err := http.Get("https://jsonplaceholder.typicode.com/todos/1")
	if err != nil {
		return &callResponse{nil, fmt.Errorf("error in http call")}
	}

	defer resp.Body.Close()
	byteResp, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return &callResponse{nil, fmt.Errorf("error in reading response")}
	}

	structResp := &response{}
	err = json.Unmarshal(byteResp, structResp)

	if err != nil {
		return &callResponse{nil, fmt.Errorf("error in unmarshalling response")}
	}

	return &callResponse{structResp, nil}
}

func getHTTPResponse(ctx context.Context) (*response, error) {
	return nil, nil
}

func main() {
	res, err := getHTTPResponse(context.Background())
	if err != nil {
		fmt.Printf("err %v", err)
	}

	fmt.Printf("res %v", res)
}
