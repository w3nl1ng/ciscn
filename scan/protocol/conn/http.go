package conn

import (
	"crypto/tls"
	"net/http"
	"time"
)

func ConnHttp(request *http.Request, timeout int) (*http.Response, error) {
	var tr *http.Transport

	tr = &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}

	client := &http.Client{
		Timeout:   time.Duration(timeout) * time.Second,
		Transport: tr,
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			return http.ErrUseLastResponse
		},
	}
	response, err := client.Do(request)
	return response, err
}
