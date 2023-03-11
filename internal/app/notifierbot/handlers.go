package notifierbot

import (
	"encoding/json"
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
