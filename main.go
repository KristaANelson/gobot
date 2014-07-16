package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/arjunsharma/gobot/carlton"
	"github.com/arjunsharma/gobot/facts"
	"github.com/arjunsharma/gobot/images"
	"github.com/arjunsharma/gobot/kittens"
	"github.com/arjunsharma/gobot/pugs"
	"github.com/arjunsharma/gobot/soothe"
	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/pugs/bomb/{count}", pugs.BombHandler)
	router.HandleFunc("/pugs/random", pugs.RandomHandler)

	router.HandleFunc("/kittens/random", kittens.RandomHandler)
	router.HandleFunc("/kittens/bomb/{count}", kittens.BombHandler)

	router.HandleFunc("/facts/cat", facts.CatHandler)

	router.HandleFunc("/images/{query}", images.QueryHandler)

	router.HandleFunc("/soothe", soothe.SootheHandler)

	router.HandleFunc("/carlton", carlton.DanceHandler)

	router.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		writer.WriteHeader(200)
		fmt.Fprintf(writer, "NICE!")
	})
	http.Handle("/", router)
	http.ListenAndServe(":"+os.Getenv("PORT"), nil)
}
