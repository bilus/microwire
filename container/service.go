package container

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
)

type Service struct {
	Name string
	Path string
	Host string
}

func (s Service) proxy(w http.ResponseWriter, r *http.Request) {
	appURL := s.appURL(r)
	pr, err := http.NewRequest(r.Method, appURL.String(), r.Body)
	if err != nil {
		s.internalServerError(w, err)
		return
	}
	pr.Header = r.Header
	resp, err := http.DefaultClient.Do(pr)
	if err != nil {
		s.internalServerError(w, err)
		return
	}
	for k, vs := range resp.Header {
		w.Header().Del(k)
		for _, v := range vs {
			w.Header().Add(k, v)
		}
	}
	w.WriteHeader(resp.StatusCode)
	_, err = io.Copy(w, resp.Body)
	if err != nil {
		// Too late to change status code. But this error usually means client disconnected.
		return
	}
}

func (s Service) appURL(r *http.Request) *url.URL {
	url := *r.URL
	url.Host = s.Host
	url.Scheme = "http"
	return &url
}

func (s Service) internalServerError(w http.ResponseWriter, err error) {
	_, _ = w.Write([]byte(fmt.Sprintf("Error while proxying content from service %q: %v", s.Name, err)))
	w.WriteHeader(http.StatusInternalServerError)
}
