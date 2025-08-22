package configman

import (
        "time"
        "github.com/vlence/gossert"
)

// A name-value pair. T must conform to one of the Type values exported,
// except Unsupported.
type Setting struct {
        hasName
        hasDescription
        canBeDeprecated
        canBeCreated
        canBeUpdated

        typ Type
        value any
}

func (setting *Setting) Type() Type {
        gossert.Ok(nil != setting, "setting: cannot return type of nil setting")
        return setting.typ
}

func (setting *Setting) String() string {
        gossert.Ok(nil != setting, "setting: cannot return value of nil setting as string")

        // todo
        return ""
}

func (setting *Setting) Value() any {
        gossert.Ok(nil != setting, "setting: cannot return value of nil setting")
        return setting.value
}