package notifierbot

import (
	"encoding/json"
	goerr "github.com/darkjinnee/go-err"
	"net/http"
)

type ResponseJsonErr struct {
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
	if status == http.StatusMethodNotAllowed {
		err := json.NewEncoder(w).Encode(ResponseJsonErr{
			Message: http.StatusText(status),
			Errors: []string{
				http.StatusText(status),
			},
		})
		goerr.Log(
			err,
			"[Error] notifierbot.Abort: Failed to encode response",
		)
	}
}

func Listen(r []Route) {
	m := http.NewServeMux()
	for _, i := range r {
		m.HandleFunc(i.Pattern, i.Boot)
	}

	err := http.ListenAndServe(
		Conf.Listener.Address,
		m,
	)
	goerr.Fatal(
		err,
		"[Error] notifierbot.Listener: Failed to listen to address",
	)
}
