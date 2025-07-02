package configman

import "regexp"

var nameRegexp, _ = regexp.Compile(`^[a-zA-Z][_a-zA-Z0-9]*$`)

// IsValidName returns true if the given name can be used as the name.
// Names must start with an alphabet followed by any number of alphabets
// and numbers.
func IsValidName(name string) bool {
        return nameRegexp.MatchString(name)
}
