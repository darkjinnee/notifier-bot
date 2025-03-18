package httpx

import (
	"encoding/json"
	"fmt"
	goerr "github.com/darkjinnee/go-err"
	"log"
	"net/http"
	"strings"
)

const (
	ErrFailedToEncodeResponse = "[Error] httpx.Abort: Failed to encode response"
	ErrUnexpectedStatusCode   = "[Error] httpx.Abort: Unexpected status code %d"
	ErrPathMismatch           = "[Error] httpx.Boot: Path mismatch, expected %s but got %s"
	ErrListenFailed           = "[Error] httpx.Listen: Failed to listen to address"
)

type ErrorResponse struct {
	Message string   `json:"message"`
	Errors  []string `json:"errors"`
}

type SuccessResponse struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

type DataResponse struct {
	Data interface{} `json:"data"`
}

type Header struct {
	Key   string
	Value string
}

type Context struct {
	http.ResponseWriter
	*http.Request
}

type Route struct {
	Headers []Header
	Method  string
	Pattern string
	Handler func(ctx Context)
}

func isJSONHeader(r *http.Request) bool {
	a := strings.Contains(r.Header.Get("Accept"), "application/json")
	c := strings.Contains(r.Header.Get("Content-Type"), "application/json")

	return a && c
}

func isMethod(r *http.Request, allowedMethod string) bool {
	return r.Method == allowedMethod
}

func SetHeaders(w http.ResponseWriter, headers []Header) {
	for _, h := range headers {
		if existing := w.Header().Get(h.Key); existing == "" {
			w.Header().Set(h.Key, h.Value)
		}
	}
}

func Abort(
	w http.ResponseWriter,
	r *http.Request,
	status int,
) {
	w.WriteHeader(status)
	switch status {
	case http.StatusMethodNotAllowed, http.StatusNotFound:
		err := json.NewEncoder(w).Encode(ErrorResponse{
			Message: http.StatusText(status),
			Errors:  []string{http.StatusText(status)},
		})
		if err != nil {
			goerr.Log(err, ErrFailedToEncodeResponse)
		}
	default:
		err := json.NewEncoder(w).Encode(ErrorResponse{
			Message: http.StatusText(status),
			Errors:  []string{http.StatusText(status)},
		})
		if err != nil {
			goerr.Log(err, ErrUnexpectedStatusCode)
		}
	}
}

func (i Route) Boot(
	w http.ResponseWriter,
	r *http.Request,
) {
	if isJSONHeader(r) {
		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("Accept", "application/json")
	}

	for key, values := range w.Header() {
		for _, value := range values {
			fmt.Printf("Header: %s = %s\n", key, value)
		}
	}

	if !isMethod(r, i.Method) {
		Abort(w, r, http.StatusMethodNotAllowed)
		return
	}

	if r.URL.Path != i.Pattern {
		log.Printf(ErrPathMismatch, i.Pattern, r.URL.Path)
		Abort(w, r, http.StatusNotFound)
		return
	}

	SetHeaders(w, i.Headers)
	i.Handler(Context{
		ResponseWriter: w,
		Request:        r,
	})
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
		ErrListenFailed,
	)
}
