package configman

import (
        "time"
        "github.com/vlence/gossert"
)

// Represents a name-value pair.
type NamedValue interface {
        // Name returns the name of the setting.
        Name() string

        // Type returns the type of the setting's value.
        Type() Type

        // String returns the setting's value as a string.
        String() string
}

// A name-value pair. T must conform to one of the Type values exported,
// except Unsupported.
type Setting[T any] struct {
        hasName
        hasDescription
        canBeDeprecated
        canBeCreated
        canBeUpdated

        typ Type
        value T
}

func (setting *Setting[T]) Type() Type {
        gossert.Ok(nil != setting, "setting: cannot return type of nil setting")
        return setting.typ
}

func (setting *Setting[T]) String() string {
        gossert.Ok(nil != setting, "setting: cannot return string representation of nil setting")

        // todo
        return ""
}