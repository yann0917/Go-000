package main

import (
	"fmt"
	"log"
	"net/http"
)

func nullSelect() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "hello")
	})

	go func() {
		if err := http.ListenAndServe(":8080", nil); err != nil {
			log.Fatal(err)
		}
	}()

	select {}
}

func main() {
	nullSelect()
}
