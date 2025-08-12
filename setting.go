package configman

// A Setting is a named value. Settings do not exist on their own. You
// need to create a Config first and then add settings to the config.
type Setting interface {
        // Name returns the name of this setting.
        Name() string

        // Desc returns the description of this setting.
        Desc() string

        // Value returns the value of this setting. Use Type
        // to find the type of the value.
        Value() any

        // SetValue sets the value of this setting. The value
        // given must be the same type as the setting. If a
        // value of a different type is given then ErrTypeMismatch
        // is returned. Returns true if the value of this setting
        // has been changed successfully.
        SetValue(any) (bool, error)

        // IsValidValue returns true if the given value can be set
        // as this setting's value.
        IsValidValue(any) bool

        // Type returns the type of this setting's value.
        Type() Type

        // SetDesc sets the description of this setting. Returns true if
        // the description was updated.
        SetDesc(desc string) (bool, error)

        // Depr deprecates this setting. Returns true if the
        // setting was deprecated. Deprecating an already
        // deprecated setting does nothing and returns true.
        Depr(reason string) (bool, error)

        // SetDeprReason sets the deprecation reason. Returns
        // true if the deprecation reason was updated. If the
        // setting is not deprecation it does nothing and
        // returns false.
        SetDeprReason(reason string) (bool, error)

        // Config returns the config that this setting belongs to.
        Config() Config
}
