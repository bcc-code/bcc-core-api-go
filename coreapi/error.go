//go:generate go-enum --marshal

package coreapi

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Error struct {
	Code    ErrorCode `json:"code"`
	Message string    `json:"message"`
}

/*
ENUM(

	not-found
	invalid-query,
	invalid-body,
	invalid-key-reference,
	duplicate-unique-key,
	cannot-parse-token,
	invalid-token,
	missing-scopes,
	internal-error,
	unauthorized,
	invalid-app-uid,
	unknown-error-response,

)
*/
type ErrorCode string //@name ErrorCode

type ErrorResponse struct {
	Error Error `json:"error"`
}

func newError(response *http.Response) error {
	apiError := &ErrorResponse{}

	if err := json.NewDecoder(response.Body).Decode(apiError); err != nil {
		return &Error{
			Code:    ErrorCodeUnknownErrorResponse,
			Message: fmt.Errorf("failed to decode json error response payload: %w", err).Error(),
		}
	}

	return &apiError.Error
}

// Error formats the error into a string representation.
func (m *Error) Error() string {
	return fmt.Sprintf("%s: %s", m.Code, m.Message)
}
