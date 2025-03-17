package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	mux := http.NewServeMux()

	mux.Handle("/", http.HandlerFunc(handler))

	http.ListenAndServe(":8080", mux)
}

func handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "text/event-stream")
	w.Header().Set("Cache-Control", "no-cache")
	w.Header().Set("Connection", "keep-alive")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	for i := 0; i < 10; i++ {
		time.Sleep(time.Second)
		w.Write([]byte(fmt.Sprintf("data: %s%d\n\n", "Event #", i)))
		w.(http.Flusher).Flush()
	}

	<-r.Context().Done()
	return
}
