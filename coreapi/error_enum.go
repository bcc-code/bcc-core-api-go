// Code generated by go-enum DO NOT EDIT.
// Version: 0.5.5
// Revision: b9e7d1ac24b2b7f6a5b451fa3d21706ffd8d79e2
// Build Date: 2023-01-30T01:49:43Z
// Built By: goreleaser

package coreapi

import (
	"errors"
	"fmt"
)

const (
	// ErrorCodeNotFound is a ErrorCode of type not-found.
	ErrorCodeNotFound ErrorCode = "not-found"
	// ErrorCodeInvalidQuery is a ErrorCode of type invalid-query.
	ErrorCodeInvalidQuery ErrorCode = "invalid-query"
	// ErrorCodeInvalidBody is a ErrorCode of type invalid-body.
	ErrorCodeInvalidBody ErrorCode = "invalid-body"
	// ErrorCodeInvalidKeyReference is a ErrorCode of type invalid-key-reference.
	ErrorCodeInvalidKeyReference ErrorCode = "invalid-key-reference"
	// ErrorCodeDuplicateUniqueKey is a ErrorCode of type duplicate-unique-key.
	ErrorCodeDuplicateUniqueKey ErrorCode = "duplicate-unique-key"
	// ErrorCodeCannotParseToken is a ErrorCode of type cannot-parse-token.
	ErrorCodeCannotParseToken ErrorCode = "cannot-parse-token"
	// ErrorCodeInvalidToken is a ErrorCode of type invalid-token.
	ErrorCodeInvalidToken ErrorCode = "invalid-token"
	// ErrorCodeMissingScopes is a ErrorCode of type missing-scopes.
	ErrorCodeMissingScopes ErrorCode = "missing-scopes"
	// ErrorCodeInternalError is a ErrorCode of type internal-error.
	ErrorCodeInternalError ErrorCode = "internal-error"
	// ErrorCodeUnauthorized is a ErrorCode of type unauthorized.
	ErrorCodeUnauthorized ErrorCode = "unauthorized"
	// ErrorCodeInvalidAppUid is a ErrorCode of type invalid-app-uid.
	ErrorCodeInvalidAppUid ErrorCode = "invalid-app-uid"
	// ErrorCodeUnknownErrorResponse is a ErrorCode of type unknown-error-response.
	ErrorCodeUnknownErrorResponse ErrorCode = "unknown-error-response"
)

var ErrInvalidErrorCode = errors.New("not a valid ErrorCode")

// String implements the Stringer interface.
func (x ErrorCode) String() string {
	return string(x)
}

// String implements the Stringer interface.
func (x ErrorCode) IsValid() bool {
	_, err := ParseErrorCode(string(x))
	return err == nil
}

var _ErrorCodeValue = map[string]ErrorCode{
	"not-found":              ErrorCodeNotFound,
	"invalid-query":          ErrorCodeInvalidQuery,
	"invalid-body":           ErrorCodeInvalidBody,
	"invalid-key-reference":  ErrorCodeInvalidKeyReference,
	"duplicate-unique-key":   ErrorCodeDuplicateUniqueKey,
	"cannot-parse-token":     ErrorCodeCannotParseToken,
	"invalid-token":          ErrorCodeInvalidToken,
	"missing-scopes":         ErrorCodeMissingScopes,
	"internal-error":         ErrorCodeInternalError,
	"unauthorized":           ErrorCodeUnauthorized,
	"invalid-app-uid":        ErrorCodeInvalidAppUid,
	"unknown-error-response": ErrorCodeUnknownErrorResponse,
}

// ParseErrorCode attempts to convert a string to a ErrorCode.
func ParseErrorCode(name string) (ErrorCode, error) {
	if x, ok := _ErrorCodeValue[name]; ok {
		return x, nil
	}
	return ErrorCode(""), fmt.Errorf("%s is %w", name, ErrInvalidErrorCode)
}

// MarshalText implements the text marshaller method.
func (x ErrorCode) MarshalText() ([]byte, error) {
	return []byte(string(x)), nil
}

// UnmarshalText implements the text unmarshaller method.
func (x *ErrorCode) UnmarshalText(text []byte) error {
	tmp, err := ParseErrorCode(string(text))
	if err != nil {
		return err
	}
	*x = tmp
	return nil
}
