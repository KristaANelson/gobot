package kittens

import (
	"fmt"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/arjunsharma/gobot/presenters"
	"github.com/gorilla/mux"
)

func BombHandler(writer http.ResponseWriter, request *http.Request) {
	countStr := mux.Vars(request)["count"]
	count, err := strconv.Atoi(countStr)
	if err != nil {
		fmt.Println(err)
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}

	kittenUrls := []string{}
	for i := 0; i < count; i++ {
		height := (rand.Float32() * 400) + 100
		width := (rand.Float32() * 400) + 100
		kittenUrl := fmt.Sprintf("http://placekitten.com/%d/%d", int(height), int(width))
		kittenUrls = append(kittenUrls, kittenUrl)
	}

	presenters.PresentJson("kittens", kittenUrls, writer)
	return
}
