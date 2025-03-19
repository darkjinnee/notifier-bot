package httpx

import (
	"encoding/json"
	goerr "github.com/darkjinnee/go-err"
	"net/http"
	"strings"
)

const (
	ErrFailedToEncodeResponse = "[Error] httpx.Abort: Failed to encode response"
	ErrListenFailed           = "[Error] httpx.Listen: Failed to listen to address"
)

var JSONHeaders = []Header{
	{
		Key:   "Accept",
		Value: "application/json",
	},
	{
		Key:   "Content-Type",
		Value: "application/json",
	},
}

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

func (ctx Context) JsonResponse(r interface{}, c int) {
	ctx.ResponseWriter.WriteHeader(c)
	err := json.NewEncoder(ctx.ResponseWriter).Encode(r)
	if err != nil {
		goerr.Log(err, ErrFailedToEncodeResponse)
	}
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
	err := json.NewEncoder(w).Encode(ErrorResponse{
		Message: http.StatusText(status),
		Errors:  []string{http.StatusText(status)},
	})
	if err != nil {
		goerr.Log(err, ErrFailedToEncodeResponse)
	}
}

func MappedHandler(routes map[string]Route, m *http.ServeMux) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if isJSONHeader(r) {
			SetHeaders(w, JSONHeaders)
		}

		route, exists := routes[r.URL.Path]
		if !exists {
			Abort(w, r, http.StatusNotFound)
			return
		}

		if !isMethod(r, route.Method) {
			Abort(w, r, http.StatusMethodNotAllowed)
			return
		}

		m.ServeHTTP(w, r)
	}
}

func MuxWrapper(r []Route) http.Handler {
	m := http.NewServeMux()
	routes := make(map[string]Route)

	for _, i := range r {
		m.HandleFunc(i.Pattern, i.Boot)
		routes[i.Pattern] = i
	}

	return MappedHandler(routes, m)
}

func (i Route) Boot(
	w http.ResponseWriter,
	r *http.Request,
) {

	SetHeaders(w, i.Headers)
	i.Handler(Context{
		ResponseWriter: w,
		Request:        r,
	})
}

func Listen(r []Route, addr string) {
	m := MuxWrapper(r)
	err := http.ListenAndServe(
		addr,
		m,
	)
	goerr.Fatal(
		err,
		ErrListenFailed,
	)
}
