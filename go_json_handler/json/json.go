// Package json implements a wrapper for encoding/json package.
// We can create encoding/json mock with below command.
//   mockgen -source=json/json.go -destination json/json_mock.go -package=json
package json

import (
	"encoding/json"
)

// JsonInterface defines encoding/json methods.
type JsonInterface interface {
	MarshalIndent(v interface{}, prefix string, indent string) ([]byte, error)
	Unmarshal(data []byte, v interface{}) error
}

// Json implements JsonInterface.
type Json struct {
}

// NewJson returns Json instance.
func NewJson() *Json {
	return &Json{}
}

// MarshalIndent wraps encoding/json.MarshalIndent.
func (j Json) MarshalIndent(v interface{}, prefix string, indent string) ([]byte, error) {
	return json.MarshalIndent(v, prefix, indent)
}

// Unmarshal wraps encoding/json.Unmarshal.
func (j Json) Unmarshal(data []byte, v interface{}) error {
	return json.Unmarshal(data, v)
}
