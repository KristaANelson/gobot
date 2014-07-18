package facts

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/arjunsharma/gobot/presenters"
)

func CatHandler(writer http.ResponseWriter, request *http.Request) {
	response, err := http.Get("http://catfacts-api.appspot.com/api/facts?number=1")
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

	presenters.PresentJson("fact", resp["facts"].([]interface{})[0].(string), writer)
	return
}
