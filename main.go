package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/", inception)
	log.Fatal(http.ListenAndServe(":"+os.Getenv("PORT"), nil))
}

func inception(res http.ResponseWriter, req *http.Request) {
	url := fmt.Sprintf("%s://%s%s", "http", req.Host, req.RequestURI)
	log.Printf("Bwong!: %s\n", url)
	resp, err := http.Get(url)
	if err != nil {
		http.Error(res, err.Error(), http.StatusInternalServerError)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		http.Error(res, err.Error(), http.StatusInternalServerError)
	}
	fmt.Fprintln(res, body)
}
