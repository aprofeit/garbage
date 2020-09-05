package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/", hello)
	fmt.Println("listening...")
	err := http.ListenAndServe(":"+os.Getenv("PORT"), nil)
	if err != nil {
		panic(err)
	}
}

var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func hello(res http.ResponseWriter, req *http.Request) {
	for i := 0; i < 250000000; i++ {
		fmt.Fprintln(res, string(letters[rand.Intn(len(letters))]))
	}
}
