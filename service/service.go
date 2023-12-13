package service

import "net/http"

func AddTurboStreamHeaders(w http.ResponseWriter) {
	w.Header().Add("Content-Type", "text/vnd.turbo-stream.html")
}

func AddTurboFrameHeaders(w http.ResponseWriter) {
	w.Header().Add("Content-Type", "text/html; turbo-stream")
}
