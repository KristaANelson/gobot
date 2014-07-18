package pugs

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/arjunsharma/gobot/presenters"
)

func RandomHandler(writer http.ResponseWriter, request *http.Request) {
	response, err := http.Get("http://pugme.herokuapp.com/random")
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

	presenters.PresentJson("pug", resp["pug"], writer)
	return
}
