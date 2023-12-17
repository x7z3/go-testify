package main

import (
	"fmt"
	"github.com/stretchr/testify/require"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestMainHandlerWhenCountMoreThanTotal(t *testing.T) {
	totalCount := 4
	req := httptest.NewRequest(http.MethodGet, fmt.Sprintf("/cafe?count=%d&city=moscow", 10), nil)

	res := getResponse(req)
	defer res.Body.Close()

	require.Equal(t, http.StatusOK, res.StatusCode)

	body, err := io.ReadAll(res.Body)
	require.NoError(t, err)
	items := strings.Split(string(body), ",")
	require.Len(t, items, totalCount, "Returned Coffee items count should be 4")
}

func TestMainHandlerWhenCityIsNotMoscow(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, fmt.Sprintf("/cafe?count=%d&city=asdf", 10), nil)

	res := getResponse(req)
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)

	require.NoError(t, err)
	require.Equal(t, "wrong count value", string(body))
	require.Equal(t, http.StatusBadRequest, res.StatusCode)
}

func TestMainHandlerBodyShouldBeNotEmpty(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, fmt.Sprintf("/cafe?count=%d&city=moscow", 10), nil)

	res := getResponse(req)
	defer res.Body.Close()

	require.Equal(t, http.StatusOK, res.StatusCode)

	body, err := io.ReadAll(res.Body)
	require.NoError(t, err)
	require.NotEmpty(t, body, "Response body should be not empty")
}

func getResponse(req *http.Request) *http.Response {
	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)

	res := responseRecorder.Result()
	return res
}
