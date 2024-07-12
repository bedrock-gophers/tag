package tag

import "go.mongodb.org/mongo-driver/bson"

// bsonMarshaler is a Marshaler that uses the go-bson package to marshal and unmarshal data.
type bsonMarshaler struct{}

// Marshal ...
func (bsonMarshaler) Marshal(v interface{}) ([]byte, error) {
	return bson.Marshal(v)
}

// Unmarshal ...
func (bsonMarshaler) Unmarshal(data []byte, v interface{}) error {
	return bson.Unmarshal(data, v)
}

// MarshalBSON ...
func (t *Tags) MarshalBSON() ([]byte, error) {
	return marshalTags(t, bsonMarshaler{})
}

// UnmarshalBSON ...
func (t *Tags) UnmarshalBSON(b []byte) error {
	return unmarshalTags(t, b, bsonMarshaler{})
}
