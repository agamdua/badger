package main

import (
	"io"
	"log"
	"net/http"
	"os"
)

// these arguments will be data-driven in the future
// just creates a new request object from data
func createRequest(verb, url string, body io.Reader, headers map[string]string) (http.Request, error) {
	req, err := http.NewRequest(verb, url, body)
	for k, v := range headers {
		req.Header.Add(k, v)
	}
	return *req, err
}

// Returns the API Key from the environment
func getAPIKey() string {
	return os.Getenv("BADGER_API")
}

func main() {
	// TODO:  hookup to db
	url := "https://api.projectoxford.ai/vision/v1.0/model"

	// TODO:  hookup to db
	var headers = map[string]string{
		"Ocp-Apim-Subscription-Key": getAPIKey(),
	}

	// TODO: add error handling & defaults for
	//	- timeouts
	//	- connection errors
	//	- non-OK response codes
	httpClient := &http.Client{}

	// TODO:  hookup to db
	req, err := createRequest("GET", url, nil, headers)

	if err != nil {
		log.Fatal(err)
	}

	resp, err := httpClient.Do(&req)

	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()
}
