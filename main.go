package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
)

func handler(w http.ResponseWriter, r *http.Request) {
	message := fmt.Sprintln(r.Method, r.URL.Path)
	fmt.Println(message)

	if err := os.Mkdir("./logs", 0777); err != nil {
		if !strings.Contains(err.Error(), "file exists") {
			fmt.Println("[ERROR]Fail to create directory.", err)
			fmt.Fprintf(w, "[%s]ERROR: %s", r.URL.Path[1:], err.Error())
			return
		}
	}
	f, err := os.OpenFile("./logs/info", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
	if err != nil {
		fmt.Println("[ERROR]Fail to open file.", err)
		fmt.Fprintf(w, "[%s]ERROR: %s", r.URL.Path[1:], err.Error())
		return
	}

	if _, err := f.WriteString(message); err != nil {
		fmt.Println("[ERROR]", err)
	}

	fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])
}

func main() {
	fmt.Println("Start web server")
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":80", nil))
}
