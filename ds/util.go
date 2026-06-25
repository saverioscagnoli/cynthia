package ds

import (
	"encoding/json"
	"strconv"
)

// NumericNewtype is any newtype over an integer kind.
type NumericNewtype interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 |
		~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64
}

// UnmarshalNumeric unmarshals a JSON number OR quoted-number string into a numeric newtype.
// Usage: func (t *MyType) UnmarshalJSON(b []byte) error { return unmarshalNumeric(b, (*int)(t)) }
func unmarshalNumeric[T NumericNewtype](b []byte, out *T) error {
	// Try direct number first
	var n json.Number
	if err := json.Unmarshal(b, &n); err != nil {
		return err
	}

	i, err := strconv.ParseInt(n.String(), 10, 64)
	if err != nil {
		return err
	}

	*out = T(i)
	return nil
}

func unmarshalString[T ~string](b []byte, out *T) error {
	var s string
	if err := json.Unmarshal(b, &s); err != nil {
		return err
	}

	*out = T(s)
	return nil
}
