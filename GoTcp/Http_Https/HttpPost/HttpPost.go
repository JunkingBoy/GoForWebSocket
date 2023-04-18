package HttpPost

import (
	"net/http"
	"strings"
)

type RequestJson struct {
	Version string `json:"version"`
}

func (rj *RequestJson) SetVersion(str string) {
	if str != "" {
		rj.Version = str
	}
}

func (rj *RequestJson) GetVersion() string {
	return rj.Version
}

func HttpPost(url string, args []string) *http.Response {
	var targetUrl string
	if url != "" {
		targetUrl = url
	}
	// 构造json
	payload := strings.NewReader("")

	req, _ := http.NewRequest("POST", targetUrl, payload)

	for i, j := range args {
		req.Header.Add(args[i], j)
	}

	res, err := http.DefaultClient.Do(req)

	if err == nil {
		return res
	}

	return nil
}
