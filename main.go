package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/arjunsharma/gobot/carlton"
	"github.com/arjunsharma/gobot/facts"
	"github.com/arjunsharma/gobot/images"
	"github.com/arjunsharma/gobot/kittens"
	"github.com/arjunsharma/gobot/names"
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

	router.HandleFunc("/names/random", names.CsvNameHandler).Headers("Accept", "text/csv")
	router.HandleFunc("/names/random", names.JsonNameHandler)
	router.HandleFunc("/names/bomb/{count}", names.CsvNameHandler).Headers("Accept", "text/csv")
	router.HandleFunc("/names/bomb/{count}", names.JsonNameHandler)

	router.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		writer.WriteHeader(200)
		routes := make(map[string]interface{})
		routes["routes"] = map[string]interface{}{
			"pug_bomb":      "/pugs/bomb/:count",
			"pug_random":    "/pugs/random",
			"kitten_bomb":   "/kittens/bomb/:count",
			"kitten_random": "/kittens/random",
			"cat_facts":     "/facts/cat",
			"image_query":   "/images/:query",
			"names_random":  "/names/random",
			"name_bomb":     "/names/bomb/:count",
			"soothe":        "/soothe",
			"carlton":       "/carlton",
		}

		jsonResponse, err := json.Marshal(routes)
		if err != nil {
			fmt.Println(err)
			writer.WriteHeader(http.StatusInternalServerError)
			return
		}

		writer.Write(jsonResponse)
		return
	})
	http.Handle("/", router)
	http.ListenAndServe(":"+os.Getenv("PORT"), nil)
}
