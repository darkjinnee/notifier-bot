package notifierbot

import (
	"github.com/darkjinnee/notifierbot/pkg/adapter/httpx"
)

func Home(ctx httpx.Context) {
	r := httpx.DataResponse{
		Data: []string{"apple", "banana", "cherry"},
	}

	ctx.JsonResponse(r, 200)
}

func Test(ctx httpx.Context) {
	r := httpx.SuccessResponse{
		Message: "Запрос успешно обработан",
		Data:    nil,
	}

	ctx.JsonResponse(r, 202)
}
