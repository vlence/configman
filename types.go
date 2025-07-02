package configman

// The type of a setting's value.
type Type uint8

const (
        Unsupported Type = 0  // Unsupported data type
        Uint8       Type = 1  // 1 byte unsigned integer
        Uint16      Type = 2  // 2 byte unsigned integer
        Uint32      Type = 3  // 4 byte unsigned integer
        Uint64      Type = 4  // 8 byte unsigned integer
        Int8        Type = 5  // 1 byte signed integer
        Int16       Type = 6  // 2 byte signed integer
        Int32       Type = 7  // 4 byte signed integer
        Int64       Type = 8  // 8 byte signed integer
        Float32     Type = 9  // 4 byte IEEE 754 single precision floating point number
        Float64     Type = 10 // 8 byte IEEE 754 double precision floating point number
        Bool        Type = 11 // boolean
        String      Type = 12 // string
)

// TypeOf returns a Type value that represents the type of v.
func TypeOf(v any) Type {
        switch v.(type) {
        case uint8:
                return Uint8
        case uint16:
                return Uint16
        case uint32:
                return Uint32
        case uint64:
                return Uint64
        case int8:
                return Int8
        case int16:
                return Int16
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
