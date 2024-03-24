package notifierbot

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type JsonResponse struct {
	Welcome string
}

func Home(
	w http.ResponseWriter,
	r *http.Request,
) {
	_ = json.NewEncoder(w).Encode(JsonResponse{
		Welcome: "Привет! Это главная страница.",
	})
}

func Test(
	w http.ResponseWriter,
	r *http.Request,
) {
	fmt.Print(r.URL.Query().Get("q"))
	_ = json.NewEncoder(w).Encode(JsonResponse{
		Welcome: "Hello World!",
	})
}
