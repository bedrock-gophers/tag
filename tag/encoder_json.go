package tag

import "encoding/json"

// jsonMarshaler is a Marshaler that uses the encoding/json package to marshal and unmarshal data.
type jsonMarshaler struct{}

// Marshal ...
func (jsonMarshaler) Marshal(v interface{}) ([]byte, error) {
	return json.Marshal(v)
}

// Unmarshal ...
func (jsonMarshaler) Unmarshal(data []byte, v interface{}) error {
	return json.Unmarshal(data, v)
}

// MarshalJSON ...
func (t *Tags) MarshalJSON() ([]byte, error) {
	return marshalTags(t, jsonMarshaler{})
}

// UnmarshalJSON ...
func (t *Tags) UnmarshalJSON(b []byte) error {
	return unmarshalTags(t, b, jsonMarshaler{})
}
