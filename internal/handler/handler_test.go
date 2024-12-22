package handler_test

import (
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/hwtdspprcmpltl/calcyandex/internal/handler"
)

const (
	valid_endpoint = "/api/v1/calculate"
)

func setupRouter() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/api/v1/calculate", handler.HandleCalculator)
	return mux
}

func testHandler(t *testing.T, method string, requestBody string, expectedCode int, expectedBody string, endpoint string) {
	t.Helper()

	req := httptest.NewRequest(method, endpoint, strings.NewReader(requestBody))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	router := setupRouter()
	router.ServeHTTP(w, req)
	res := w.Result()
	defer res.Body.Close()

	if res.StatusCode != expectedCode {
		t.Errorf("неправильный статус код: получил %d, хотел %d", res.StatusCode, expectedCode)
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		t.Fatalf("не удалось прочитать тело ответа: %v", err)
	}

	if strings.TrimSpace(string(body)) != expectedBody {
		t.Errorf("неправильное тело ответа получил %s, хотел %s", strings.TrimSpace(string(body)), expectedBody)
	}
}

func TestHandleCalculator_SuccessCases(t *testing.T) {
	testHandler(t, http.MethodPost, `{"expression": "2+2*2"}`, http.StatusOK, `{"result":6}`, valid_endpoint)
	testHandler(t, http.MethodPost, `{"expression": "(3+7)/2"}`, http.StatusOK, `{"result":5}`, valid_endpoint)
}

func TestHandleCalculator_InvalidMethod(t *testing.T) {
	testHandler(t, http.MethodGet, ``, http.StatusMethodNotAllowed, `{"error":"не пост запрос"}`, valid_endpoint)
	testHandler(t, http.MethodPut, ``, http.StatusMethodNotAllowed, `{"error":"не пост запрос"}`, valid_endpoint)
}

func TestHandleCalculator_CalculationErrors(t *testing.T) {
	testHandler(t, http.MethodPost, `{"expression": "1/0"}`, http.StatusUnprocessableEntity, `{"error":"деление на ноль"}`, valid_endpoint)
	testHandler(t, http.MethodPost, `{"expression": "(2+2))"}`, http.StatusUnprocessableEntity, `{"error":"ошибка в расставлении скобок"}`, valid_endpoint)
}

func TestHandleCalculator_CalcilationErrors(t *testing.T) {
	testHandler(t, http.MethodPost, `{"expression": "1+20"}`, http.StatusNotFound, "404 page not found", "/wrong/endpoint")
	testHandler(t, http.MethodPost, `{"expression": "1+0"}`, http.StatusNotFound, "404 page not found", "/api/v2/sdfgdgdfge")
}
