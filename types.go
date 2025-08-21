package configman

import "errors"

// Represents a valid data type of value that can be stored in a setting.
type Type uint8

const (
        Unsupported Type = 0 // Unsupported data type
        Int32       Type = 1 // 4 byte signed integer
        Int64       Type = 2 // 8 byte signed integer
        Float32     Type = 3 // 4 byte IEEE 754 single precision floating point number
        Float64     Type = 4 // 8 byte IEEE 754 double precision floating point number
        Bool        Type = 5 // boolean
        String      Type = 6 // string
)

var ErrTypeMismatch = errors.New("configman: type of value does not match expected type")
var ErrUnsupportedType error = errors.New("configman: unknown or unsupported type of value")

// TypeOf returns a Type value that represents the type of v.
func TypeOf(v any) Type {
        switch v.(type) {
        case int32:
                return Int32
        case int64:
                return Int64
        case float32:
                return Float32
        case float64:
                return Float64
        case bool:
                return Bool
        case string:
                return String
        default:
                return Unsupported
        }
}

// String returns the name of the type as a string.
func (t Type) String() string {
        switch (t) {
        case Int32:
                return "int32"
        case Int64:
                return "int64"
        case Float32:
                return "float32"
        case Float64:
                return "float64"
        case Bool:
                return "bool"
        case String:
                return "string"
        default:
                return "unsupported"
        }
}
