package http

import (
	"encoding/json"
	goerr "github.com/darkjinnee/go-err"
	"net/http"
)

type JsonResponseErr struct {
	Message string   `json:"message"`
	Errors  []string `json:"errors"`
}

type Header struct {
	Key   string
	Value string
}

type Route struct {
	Headers []Header
	Method  string
	Pattern string
	Handler func(
		w http.ResponseWriter,
		r *http.Request,
	)
}

func (i Route) Boot(
	w http.ResponseWriter,
	r *http.Request,
) {
	for _, h := range i.Headers {
		w.Header().Set(h.Key, h.Value)
	}
	if r.URL.Path != i.Pattern {
		Abort(w, r, http.StatusNotFound)
		return
	}
	if r.Method != i.Method {
		Abort(w, r, http.StatusMethodNotAllowed)
		return
	}

	i.Handler(w, r)
}

func Abort(
	w http.ResponseWriter,
	r *http.Request,
	status int,
) {
	w.WriteHeader(status)
	switch {
	case status == http.StatusMethodNotAllowed:
	case status == http.StatusNotFound:
		err := json.NewEncoder(w).Encode(JsonResponseErr{
			Message: http.StatusText(status),
			Errors: []string{
				http.StatusText(status),
			},
		})
		goerr.Log(
			err,
			"[Error] http.Abort: Failed to encode response",
		)
	default:
		return
	}
}

func Listen(r []Route, addr string) {
	m := http.NewServeMux()
	for _, i := range r {
		m.HandleFunc(i.Pattern, i.Boot)
	}

	err := http.ListenAndServe(
		addr,
		m,
	)
	goerr.Fatal(
		err,
		"[Error] http.Listen: Failed to listen to address",
	)
}
