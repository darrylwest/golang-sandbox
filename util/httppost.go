package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)

	var data map[string]interface{}

	if err := decoder.Decode(&data); err != nil {
		fmt.Println(err)

		w.Write([]byte("failed..."))
	} else {
		fmt.Println(data)

		w.Write([]byte("ok..."))
	}
}

func main() {
	port := ":9000"

	http.HandleFunc("/rest", handler)
	fmt.Println("listen on port: ", port)

	http.ListenAndServe(":9000", nil)
}
