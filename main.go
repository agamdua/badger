package main

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

type Message struct {
	Action string
	Data   string
}

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

func inboundHandler(w http.ResponseWriter, r *http.Request) {
	var message Message

	if r.Body == nil {
		http.Error(w, "No request body found", 400)
		return
	}

	err := json.NewDecoder(r.Body).Decode(&message)

	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}

	var config OutboundConfig

	config = GetConfig(message)

	// TODO: add error handling & defaults for newtwork call
	httpClient := &http.Client{}

	// TODO: make headers more configurable
	var headers = map[string]string{
		config.APIKeyHeader: config.APIKey,
	}
	req, err := createRequest(config.Verb, config.URL, nil, headers)

	if err != nil {
		log.Fatal(err)
	}

	resp, err := httpClient.Do(&req)

	log.Print(resp.StatusCode)

	if err != nil {
		log.Fatal(err)
	}

	w.Header().Set("Content-Type", "application/json")

	body, _ := ioutil.ReadAll(resp.Body)

	log.Print(string(body))
	fmt.Fprintln(w, string(body))

	defer resp.Body.Close()

}

func main() {
	http.HandleFunc("/", inboundHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
