package main

import (
	"encoding/json"
	"flighttracker/engine"
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/calculate", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case "POST":
			w.Header().Set("Content-Type", "application/json")
			decoder := json.NewDecoder(r.Body)
			var input [][]string
			err := decoder.Decode(&input)
			if err != nil {
				log.Println("Decoder: " + err.Error())
				http.ResponseWriter.WriteHeader(w, http.StatusBadRequest)
				return
			}

			result, err := engine.Track(input)
			if err != nil {
				log.Println("Tracker: " + err.Error())
				http.ResponseWriter.WriteHeader(w, http.StatusBadRequest)
				return
			}

			http.ResponseWriter.WriteHeader(w, http.StatusOK)
			http.ResponseWriter.Write(w, []byte(fmt.Sprintf("[\"%s\",\"%s\"]\n", result[0], result[1])))
		default:
			http.ResponseWriter.WriteHeader(w, http.StatusMethodNotAllowed)
			return
		}
	})

	fmt.Printf("Starting the flight tracker server ... \n")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
