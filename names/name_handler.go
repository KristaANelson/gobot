package names

import (
	"encoding/csv"
	"fmt"
	"net/http"
	"strconv"

	"github.com/arjunsharma/gobot/presenters"
	"github.com/gorilla/mux"
)

func JsonNameHandler(writer http.ResponseWriter, request *http.Request) {
	var err error
	countStr := mux.Vars(request)["count"]
	count := 1
	if countStr != "" {
		count, err = strconv.Atoi(countStr)
		if err != nil {
			fmt.Println(err)
			writer.WriteHeader(http.StatusInternalServerError)
			return
		}
	}

	nameList := []string{}
	for i := 0; i < count; i++ {
		name := randomName()
		nameList = append(nameList, name)
	}

	presenters.PresentJson("names", nameList, writer)
	return
}

func CsvNameHandler(writer http.ResponseWriter, request *http.Request) {
	var err error
	countStr := mux.Vars(request)["count"]
	count := 1
	if countStr != "" {
		count, err = strconv.Atoi(countStr)
		if err != nil {
			fmt.Println(err)
			writer.WriteHeader(http.StatusInternalServerError)
			return
		}
	}

	csvWriter := csv.NewWriter(writer)

	for i := 0; i < count; i++ {
		err = csvWriter.Write([]string{randomName()})
		if err != nil {
			fmt.Println(err)
			writer.WriteHeader(http.StatusInternalServerError)
			return
		}
	}

	csvWriter.Flush()

	return
}
