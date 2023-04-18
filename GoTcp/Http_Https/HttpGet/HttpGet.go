package main

import (
	"fmt"
	"net/http"
)

type GetRes struct {
	Version string `json:"version"`
}

func (gr *GetRes) SetVersion(str string) {
	if str != "" {
		gr.Version = str
	}
}

func (gr *GetRes) GetVersion() string {
	return gr.Version
}

func HttpGetNotHeader(str string) *http.Response {
	if str != "" {
		res, err := http.Get(str)
		if err == nil {
			return res
		}
	}
	return nil
}

func HttpGetHeader(str string, args []string) *http.Response {
	var url string
	if str != "" {
		url = str
	}
	req, _ := http.NewRequest("GET", url, nil)
	for i, j := range args {
		req.Header.Add(args[i], j)
	}

	res, err := http.DefaultClient.Do(req)
	if err == nil {
		return res
	}
	return nil
}

func main() {
	login := "https://console-pre.raylink.live/login"
	// strings := []string{"1", "2"}
	fmt.Print(HttpGetNotHeader(login))
	// fmt.Print(HttpGetHeader(login, strings))
}
