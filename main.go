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

w.Header().Set("Content-Disposition", "attachment; filename=garbage.txt")
w.Header().Set("Content-Type", req.Header.Get("Content-Type"))

func hello(res http.ResponseWriter, req *http.Request) {
	for i := 0; i < 1000; i++ {
		fmt.Fprintln(res, string(letters[rand.Intn(len(letters))]))
	}
}
