package main

import (
	"fmt"
	"net/http"
)

func WebC1Handle(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(writer, "Hello World %s!", request.URL.Path[1:])
}

func UseWebC1Server() {
	http.HandleFunc("/", WebC1Handle)
	http.ListenAndServe(":8080", nil)
}
