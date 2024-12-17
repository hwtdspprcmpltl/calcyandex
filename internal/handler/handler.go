package handler

import (
	"encoding/json"
	"net/http"

	"github.com/hwtdspprcmpltl/calcyandex/internal/calc"
)

type PostRequest struct {
	Expression string `json:"expression"`
}

type HTTPAnswer struct {
	Result float64 `json:"result,omitempty"`
	Error  string  `json:"error,omitempty"`
}

func HandleCalculator(w http.ResponseWriter, r *http.Request) {
	var request PostRequest

	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		w.WriteHeader(422) //не так с входными(декод)
		answer := HTTPAnswer{Error: "неправильный ввод (ошибка с декодером)"}
		json.NewEncoder(w).Encode(answer)
		return
	}

	result, err := calc.Calculate(string(request.Expression))
	if err != nil {
		w.WriteHeader(500) //здесь надо все ошибки задебагить
		return
	}

	w.WriteHeader(200)
	answer := HTTPAnswer{Result: result}
	json.NewEncoder(w).Encode(answer)

}
