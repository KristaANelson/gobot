package pugs

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
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

	response, err := http.Get(fmt.Sprintf("http://pugme.herokuapp.com/bomb?count=%d", count))
	defer response.Body.Close()

	if err != nil {
		fmt.Println(err)
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}
	body, err := ioutil.ReadAll(response.Body)

	resp := make(map[string]interface{})
	err = json.Unmarshal(body, &resp)
	if err != nil {
		fmt.Println(err)
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}

	presenters.PresentJson("pugs", resp["pugs"], writer)
	return
}
