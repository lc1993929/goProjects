package main

import (
	"fmt"
	"io"
	"net/http"
)

func main121() {
	http.HandleFunc("/health", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Println(request.URL)

		all, _ := io.ReadAll(request.Body)
		fmt.Println(string(all))

		_, err := writer.Write([]byte("test"))
		if err != nil {
			fmt.Println(err)
			return
		}
	})

	err := http.ListenAndServe("127.0.0.1:8080", nil)
	if err != nil {
		return
	}

}
