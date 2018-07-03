package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type joke struct {
	IconURL string `json:"icon_url"`
	ID      string `json:"id"`
	URL     string `json:"url"`
	Value   string `json:"value"`
}

type jokeList struct {
	Hits   int    `json:"hits"`
	Result []joke `json:"result"`
}

func jokeRequest(path, query string) ([]byte, error) {
	url := "https://api.chucknorris.io/jokes/"
	request, err := http.NewRequest("GET", url+path+"?"+query, nil)
	if err != nil {
		return nil, err
	}

	resp, err := http.DefaultClient.Do(request)
	if err != nil {
		return nil, err
	}

	return ioutil.ReadAll(resp.Body)
}

func jokeCategory(category string) string {
	body, err := jokeRequest("random", "category="+category)
	if err != nil {
		return err.Error()
	}

	var j joke
	json.Unmarshal(body, &j)
	return j.Value
}

func jokeSearch(query string) string {
	body, err := jokeRequest("search", "query="+query)
	if err != nil {
		return err.Error()
	}

	var j jokeList
	json.Unmarshal(body, &j)
	if len(j.Result) == 0 {
		return ""
	}
	return j.Result[0].Value
}
