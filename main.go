package main

import (
	"errors"
	"fmt"
	"goPlayground/basic"
	"goPlayground/serv"
	"net/http"
	"os"
)

func main() {
	basic.BasicSyntax()
	http.HandleFunc("/", serv.GetRoot)
	http.HandleFunc("/hello", serv.GetHello)

	err := http.ListenAndServe(":3333", nil)
	if errors.Is(err, http.ErrServerClosed) {
		fmt.Printf("server closed\n")
	} else if err != nil {
		fmt.Printf("error starting server: %s\n", err)
		os.Exit(1)
	}
}
