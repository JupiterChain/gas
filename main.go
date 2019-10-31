package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"
)

const (
	gasPriceEndpoint = "https://etherscan.io/chart/gasprice?output=csv"
)

func main() {
	port := ":" + os.Getenv("PORT")

	http.HandleFunc("/", logMw(home))
	http.HandleFunc("/history", logMw(history))

	log.Println("Server running on port:", port)
	log.Fatal(http.ListenAndServe(port, nil))
}

func logMw(handler http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		from := time.Now()
		handler(w, r)
		elapse := time.Since(from)
		log.Printf("%s %s - 200 %.2fs", r.Method, r.URL.Path, elapse.Seconds())
	}
}

func home(w http.ResponseWriter, r *http.Request) {
	http.Error(w, "ok", http.StatusOK)
}

func history(w http.ResponseWriter, r *http.Request) {
	body := ""

	resp, err := http.Get(gasPriceEndpoint)
	if err == nil {
		chunks, _ := ioutil.ReadAll(resp.Body)
		body = string(chunks)
	}

	http.Error(w, body, http.StatusOK)
}
