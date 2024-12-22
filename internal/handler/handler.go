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
	if r.URL.Path != "/api/v1/calculate" {
		w.WriteHeader(http.StatusNotFound)
		answer := HTTPAnswer{Error: "неправильный путь"}
		json.NewEncoder(w).Encode(answer)
		return
	}
	w.Header().Set("Content-Type", "application/json")

	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		answer := HTTPAnswer{Error: "не пост запрос"}
		json.NewEncoder(w).Encode(answer)
		return
	}

	var request PostRequest

	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		w.WriteHeader(http.StatusUnprocessableEntity)
		json.NewEncoder(w).Encode(HTTPAnswer{Error: "неправильный формат ввода(ошибка при декодировании json)"}) //поправь
		return
	}

	result, err := calc.Calculate(string(request.Expression))
	if err != nil {
		if err == calc.Errors.DivisionByZero || err == calc.Errors.MismatchedParentheses {
			w.WriteHeader(http.StatusUnprocessableEntity)
			answer := HTTPAnswer{Error: err.Error()}
			json.NewEncoder(w).Encode(answer)
			return
		} else {
			answer := HTTPAnswer{Error: "ошибка сервера"}
			json.NewEncoder(w).Encode(answer)
			return
		}
	}
	w.WriteHeader(http.StatusOK)
	answer := HTTPAnswer{Result: result}
	json.NewEncoder(w).Encode(answer)

}
