package rest

import (
	"encoding/json"
	"errors"
	"fmt"
	e "github.com/statistico/statistico-grpc-gateway/internal/app/errors"
	"net/http"
)

func parseJsonError(e error) (int, error) {
	if err, ok := e.(*json.SyntaxError); ok {
		err := fmt.Errorf("request body contains badly-formed JSON (at position %d)", err.Offset)
		return http.StatusBadRequest, err
	}

	if err, ok := e.(*json.UnmarshalTypeError); ok {
		return http.StatusUnprocessableEntity, err
	}

	return http.StatusBadRequest, e
}

func handleError(w http.ResponseWriter, err error) {
	if err == e.ErrorInternalServerError {
		errorResponse(w, http.StatusInternalServerError, errors.New("internal server error"))
		return
	}

	if err == e.ErrorBadGateway {
		errorResponse(w, http.StatusBadGateway, errors.New("bad gateway"))
		return
	}

	failResponse(w, http.StatusUnprocessableEntity, err)
}
