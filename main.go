package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httputil"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.Method, r.URL.Path)
	fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])
}

func invokeHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	var requestData struct {
		Method string `json:"method"`
		URL    string `json:"url"`
	}

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Unable to read request body", http.StatusBadRequest)
		return
	}

	err = json.Unmarshal(body, &requestData)
	if err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	req, err := http.NewRequest(requestData.Method, requestData.URL, nil)
	if err != nil {
		http.Error(w, "Failed to create request", http.StatusInternalServerError)
		return
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		http.Error(w, "Failed to execute request", http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	responseDump, err := httputil.DumpResponse(resp, true)
	if err != nil {
		http.Error(w, "Failed to read response", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(resp.StatusCode)
	w.Write(responseDump)
}

func main() {
	fmt.Println("Start web server")
	http.HandleFunc("/", handler)
	http.HandleFunc("/invoke", invokeHandler)
	log.Fatal(http.ListenAndServe(":9000", nil))
}
