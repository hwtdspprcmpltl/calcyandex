package handler_test

import (
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/hwtdspprcmpltl/calcyandex/internal/handler"
)

func testHandler(t *testing.T, method string, requestBody string, expectedCode int, expectedBody string) {
	t.Helper()

	req := httptest.NewRequest(method, "/api/v1/calculate", strings.NewReader(requestBody))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	handler.HandleCalculator(w, req)
	res := w.Result()
	defer res.Body.Close()

	if res.StatusCode != expectedCode {
		t.Errorf("unexpected status code: got %d, want %d", res.StatusCode, expectedCode)
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		t.Fatalf("failed to read response body: %v", err)
	}

	if strings.TrimSpace(string(body)) != expectedBody {
		t.Errorf("unexpected response body: got %s, want %s", strings.TrimSpace(string(body)), expectedBody)
	}
}

func TestHandleCalculator_SuccessCases(t *testing.T) {
	testHandler(t, http.MethodPost, `{"expression": "2+2*2"}`, http.StatusOK, `{"result":6}`)
	testHandler(t, http.MethodPost, `{"expression": "(3+7)/2"}`, http.StatusOK, `{"result":5}`)
}

func TestHandleCalculator_InvalidMethod(t *testing.T) {
	testHandler(t, http.MethodGet, ``, http.StatusMethodNotAllowed, `{"error":"не пост запрос"}`)
	testHandler(t, http.MethodPut, ``, http.StatusMethodNotAllowed, `{"error":"не пост запрос"}`)
}

func TestHandleCalculator_CalculationErrors(t *testing.T) {
	testHandler(t, http.MethodPost, `{"expression": "1/0"}`, http.StatusUnprocessableEntity, `{"error":"деление на ноль"}`)
	testHandler(t, http.MethodPost, `{"expression": "(2+2))"}`, http.StatusUnprocessableEntity, `{"error":"ошибка в расставлении скобок"}`)
}
