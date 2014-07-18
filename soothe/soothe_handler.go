package soothe

import (
	"fmt"
	"math/rand"
	"net/http"

	"github.com/arjunsharma/gobot/presenters"
)

func SootheHandler(writer http.ResponseWriter, request *http.Request) {
	calmIndex := int(33*rand.Float32()) + 1
	calmUrl := fmt.Sprintf("http://calmingmanatee.com/img/manatee%d.jpg", calmIndex)

	presenters.PresentJson("soothe", calmUrl, writer)
	return
}
