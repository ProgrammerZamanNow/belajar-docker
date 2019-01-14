package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		resp, err := http.Get("http://target-server/")
		defer resp.Body.Close()
		if err != nil {
			fmt.Fprintf(w, "Ups, Error!")
		} else {
			body, _ := ioutil.ReadAll(resp.Body)
			fmt.Fprintf(w, string(body))
		}
	})

	http.ListenAndServe(":8080", nil)
}
