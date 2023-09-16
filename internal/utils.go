package internal

import (
	"fmt"
	"log"
	"net/http"
)

type RequestClient struct {
}

func (r *RequestClient) Head(url string) *http.Response {
	return r.respProcessor(url, http.Head)
}

func (r *RequestClient) Get(url string) *http.Response {
	return r.respProcessor(url, http.Get)
}

func (r *RequestClient) respProcessor(url string, httpCallBack func(string) (*http.Response, error)) *http.Response {
	resp, err := httpCallBack(url)
	if err != nil || (resp.StatusCode < 200 || resp.StatusCode > 299) {
		msg := fmt.Sprintf("Something went wrong with url: %s, error: %s, status: %s", url, err, resp.Status)
		log.Println(msg)
		return nil
	}
	return resp
}
