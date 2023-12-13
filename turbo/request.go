package turbo

import (
	"net/http"
	"strings"
)

func IsTurbo(r *http.Request) bool {
	return strings.Contains(r.Header.Get("accept"), "text/vnd.turbo-stream.html") ||
		r.Header.Get("Turbo-Frame") != ""
}
