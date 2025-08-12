package configman

// A Config is a collection of Setting values. Setting values do no exist
// on their own. You need to create a Config and then add Settings to the
// Config.
type Config interface {
        // Name returns the name of this config.
        Name() string

        // Desc returns the description of this config.
        Desc() string

        // SetDesc sets this config's description.
        SetDesc(desc string) (bool, error)

        // NewSetting creates a new setting with the given name, type and value.
        NewSetting(string, Type, any) (Setting, error)
}
