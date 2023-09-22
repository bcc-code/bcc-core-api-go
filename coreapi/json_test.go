package coreapi

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

type TestType struct {
	Prop struct{ Str } `json:"prop"`
}

type Str struct {
	Val string
}

// MarshalText implements the text marshaller method.
func (x Str) MarshalText() ([]byte, error) {
	return []byte(x.Val), nil
}

// UnmarshalText implements the text unmarshaller method.
func (x *Str) UnmarshalText(text []byte) error {
	x.Val = string(text)
	return nil
}

func TestMarshal(t *testing.T) {
	marshalVal := TestType{}
	marshalVal.Prop.Str.Val = "Hello"

	data, err := json.Marshal(marshalVal)
	assert.NoError(t, err)
	assert.Equal(t, []byte(`{"prop":"Hello"}`), data) // This works

	var unmarshalVal TestType

	err = json.Unmarshal([]byte(`{"prop":"Hello"}`), &unmarshalVal)
	assert.NoError(t, err) // Fails with json: cannot unmarshal string into Go struct field TestType.prop of type struct { coreapi.Str }
	assert.Equal(t, Str{"Hello"}, unmarshalVal.Prop.Str)
}
