package handler

import (
	"encoding/json"
	"net/http"
	"strings"

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
	w.Header().Set("Content-Type", "application/json")

	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		answer := HTTPAnswer{Error: "разрешен только пост запрос :<"}
		json.NewEncoder(w).Encode(answer)
	}

	var request PostRequest

	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		w.WriteHeader(http.StatusUnprocessableEntity)
		json.NewEncoder(w).Encode(HTTPAnswer{Error: "неправильный тип ввода (не декодается)"}) //поправь
		return
	}

	result, err := calc.Calculate(string(request.Expression))
	if err != nil {
		if strings.Contains(err.Error(), "invalid") {
			w.WriteHeader(http.StatusUnprocessableEntity)
			answer := HTTPAnswer{Error: "неправильное выражение"}
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
