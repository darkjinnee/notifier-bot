package notifierbot

import (
	"encoding/json"
	goerr "github.com/darkjinnee/go-err"
	"github.com/darkjinnee/notifierbot/pkg/adapter/httpx"
)

func Home(ctx httpx.Context) {
	r := httpx.DataResponse{
		Data: []string{"apple", "banana", "cherry"},
	}

	ctx.ResponseWriter.WriteHeader(200)
	err := json.NewEncoder(ctx.ResponseWriter).Encode(r)
	if err != nil {
		goerr.Log(err, httpx.ErrFailedToEncodeResponse)
	}
}

func Test(ctx httpx.Context) {
	r := httpx.SuccessResponse{
		Message: "Запрос успешно обработан",
		Data:    nil,
	}

	ctx.ResponseWriter.WriteHeader(202)
	err := json.NewEncoder(ctx.ResponseWriter).Encode(r)
	if err != nil {
		goerr.Log(err, httpx.ErrFailedToEncodeResponse)
	}
}
