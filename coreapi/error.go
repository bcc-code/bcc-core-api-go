package coreapi

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Error struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

type ErrorResponse struct {
	Error Error `json:"error"`
}

func newError(response *http.Response) error {
	apiError := &ErrorResponse{}

	if err := json.NewDecoder(response.Body).Decode(apiError); err != nil {
		return &Error{
			Code:    http.StatusText(response.StatusCode),
			Message: fmt.Errorf("failed to decode json error response payload: %w", err).Error(),
		}
	}

	return &apiError.Error
}

// Error formats the error into a string representation.
func (m *Error) Error() string {
	return fmt.Sprintf("%s: %s", m.Code, m.Message)
}
