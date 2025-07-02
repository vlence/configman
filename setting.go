package configman

// The type of a setting's value.
type Type uint8

const (
        // 1 byte unsigned integer
        Uint8   Type = 1
        // 2 byte unsigned integer
        Uint16  Type = 2
        // 4 byte unsigned integer
        Uint32  Type = 3
        // 8 byte unsigned integer
        Uint64  Type = 4
        // 1 byte signed integer
        Int8    Type = 5
        // 2 byte signed integer
        Int16   Type = 6
        // 4 byte signed integer
        Int32   Type = 7
        // 8 byte signed integer
        Int64   Type = 8
        // 4 byte IEEE 754 single precision floating point number
        Float32 Type = 9
        // 8 byte IEEE 754 double precision floating point number
        Float64 Type = 10
        // boolean
        Bool    Type = 11
        // string
        String  Type = 12
)

// A setting is a value with a name and optional description. The type
// of the value, once set, cannot be changed. To use a value of a different
// type create a new setting.
//
// Settings can have descriptions. It is recommended to provide descriptions
// for settings. Use descriptions to communicate what the setting's value is
// meant to be used for and what kind of values should be stored.
//
// When a setting is no longer needed you can remove them. It is recommended
// that settings are deprecated first with a deprecation reason. This gives
// your users some time to update their apps before the setting is removed.
// Once a setting is deprecated it cannot be reversed.
type Setting struct {
        typ               Type
        name              string
        desc              string
        value             any
        deprecated        bool
        deprecationReason string
}

// Deprecated returns true if this setting is deprecated
func (s *Setting) Deprecated() bool {
        return s.deprecated
}

// Name returns the name of this setting.
func (s *Setting) Name() string {
        return s.name
}

// Description returns the current description of this setting.
// If given a new description the current description will be
// replaced and the new description will be returned.
func (s *Setting) Description(newDesc ...string) (string, error) {
        if len(newDesc) > 0 {
                s.desc = newDesc[0]
        }

        return s.desc, nil
}

// Deprecate marks this setting as deprecated. Once a setting has
// been deprecated it cannot be reversed. Returns true if the
// setting was deprecated successfully otherwise returns false
// with an error. A deprecation reason can be provided optionally.
// This function can also be used to update the deprecation reason.
func (s *Setting) Deprecate(reason ...string) (bool, error) {
        if len(reason) > 0 {
                s.deprecationReason = reason[0]
        }

        s.deprecated = true

        return true, nil
}
