package configman

import "errors"

var ErrTypeMismatch = errors.New("the data type of the value does not match the setting's type")

var ErrUnsupportedType = errors.New("unsupported value data type")

// A SettingGetter can get settings from some kind of storage location.
type SettingGetter interface {
        // Get gets the setting(s) with the given name(s). The number of settings
        // returned may be lesser than the names given if, for example, the
        // settings don't exist.
        Get(name ...string) ([]*Setting, error)
}

// A SettingSetter can store new settings and update the value of existing settings.
type SettingSetter interface {
        // Set sets the value of the setting with the given name and value.
        // If the setting does not exist then it is created. The type of setting
        // is automatically detected. See Type for list of valid types. If the
        // data type of value is unsupported then ErrSupportedType is returned.
        Set(name string, value any) error
}

// A SettingDeprecator can deprecate settings.
type SettingDeprecator interface {
        // Deprecate deprecates the given setting. If a setting with the given
        // name does not exist then nothing happens. Optionally a deprecation
        // reason can also be provided. This function can also be used to
        // update the deprecation reason of a setting.
        Deprecate(name string, reason ...string) (*Setting, error)
}
