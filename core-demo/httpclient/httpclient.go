package httpclient

import (
	"net/http"
	"time"
)

//Send request
func Send(req *http.Request) (*http.Response, error) {
	timeout := time.Duration(1 * time.Second)
	client := http.Client{
		Timeout: timeout,
	}
	return client.Do(req)
}
