package main

import (
	"io"
	"net/http"
)

func hello(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "text/html")
	name := req.URL.Query().Get("name")
	if name == "" {
		name = "World"
	}

	io.WriteString(res,
		`<doctype html>
<html>
	<head>
    	<title>Hello World!</title>
	</head>
	<body>
    	Hello, `+name+`!
	</body>
</html>`)
}

func main() {
	fs := http.FileServer(http.Dir("static"))
	http.Handle("/", fs)
	http.HandleFunc("/hello", hello)

	http.ListenAndServe(":8081", nil)
}