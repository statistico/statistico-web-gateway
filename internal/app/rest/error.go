package rest

import (
	"encoding/json"
	"fmt"
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
