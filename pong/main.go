package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"
)

func main() {
	fmt.Println("Starting Pong Server")
	http.HandleFunc("/pong", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Sending Pong")
	})

	go func() {
		if err := http.ListenAndServe(":81", nil); err != nil {
			log.Fatal("ListenAndServe: ", err)
		}
	}()

	for {
		time.Sleep(5 * time.Second)
		response, err := http.Get("http://localhost:80/ping")
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		responseData, err := ioutil.ReadAll(response.Body)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(string(responseData))
	}

	fmt.Print("Pong Server Started")

}
