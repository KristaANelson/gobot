package numberformatter

import (
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func NumberFormatterHandler(writer http.ResponseWriter, request *http.Request) {
	number := mux.Vars(request)["number"]
	if number == "" {
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}

	_, err := strconv.Atoi(number)
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}

	writer.Write([]byte(Format(number)))
}

func Format(number string) string {
	return number + stNumber(number) + ndNumber(number) + rdNumber(number) + thNumber(number)
}

func stNumber(number string) string {
	if string(number[len(number)-1]) == "1" && !isANumberInTheTeens(number) {
		return "st"
	} else {
		return ""
	}
}

func ndNumber(number string) string {
	if string(number[len(number)-1]) == "2" && !isANumberInTheTeens(number) {
		return "nd"
	} else {
		return ""
	}
}

func rdNumber(number string) string {
	if string(number[len(number)-1]) == "3" && !isANumberInTheTeens(number) {
		return "rd"
	} else {
		return ""
	}
}

func thNumber(number string) string {
	lastDigit := string(number[len(number)-1])
	if isANumberInTheTeens(number) || (lastDigit != "1" && lastDigit != "2" && lastDigit != "3") {
		return "th"
	}

	return ""
}

func isANumberInTheTeens(number string) bool {
	if len(number) > 1 {
		return string(number[len(number)-2]) == "1"
	} else {
		return false
	}
}
