package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

func handler(w http.ResponseWriter, r *http.Request) {
	if !strings.HasPrefix(r.URL.Path, "/call") {
		fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])
		return
	}
	fmt.Println(r.URL.Path[6:])
	resp, err := http.Get("http://" + r.URL.Path[6:])
	if err != nil {
		fmt.Fprintf(w, "%s[GET ERROR]%s", r.URL.Path[6:], err.Error())
		return
	}

	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Fprintf(w, "%s[READ ERROR]%s", r.URL.Path[6:], err.Error())
		return
	}

	fmt.Fprintf(w, "[URL]%s\n\n%s", r.URL.Path[6:], string(b))

}

func main() {
	fmt.Println("Start web server")
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":9203", nil))
}
