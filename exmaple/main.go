package main

import (
	"log"
	"net/http"

	"fmt"

	"github.com/kavehmz/weather"
)

func handler(w http.ResponseWriter, r *http.Request) {

	skyTmp, err := weather.DarkSkyTemp()
	checkErr(err)
	xuTmp, err := weather.XUTemp()
	checkErr(err)

	w.Write([]byte(fmt.Sprintf("[%f, %f]: %f", skyTmp, xuTmp, (skyTmp+xuTmp)/2)))
}

func main() {
	http.HandleFunc("/weather", handler)
	http.ListenAndServe(":8080", nil)
}

func checkErr(e error) {
	if e != nil {
		log.Fatal(e)
	}
}
