package carlton

import (
	"math/rand"
	"net/http"

	"github.com/arjunsharma/gobot/presenters"
)

func DanceHandler(writer http.ResponseWriter, request *http.Request) {
	imageSet := []string{
		"http://imgur.com/t4EZhsD.gif",
		"http://imgur.com/7YIfBiV.gif",
		"http://imgur.com/4K7eo3s.gif",
		"http://imgur.com/kKXW55q.gif",
		"http://imgur.com/OE5KuZH.gif",
	}
	danceIndex := int(5 * rand.Float32())

	presenters.PresentJson("carlton", imageSet[danceIndex], writer)
	return
}
