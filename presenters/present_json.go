package presenters

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func PresentJson(rootKey string, responseObj interface{}, writer http.ResponseWriter) {
	presentedResponse := make(map[string]interface{})
	presentedResponse[rootKey] = responseObj

	jsonResponse, err := json.Marshal(presentedResponse)
	if err != nil {
		fmt.Println(err)
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}

	writer.Write(jsonResponse)

}
