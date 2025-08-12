package configman

// A Store implements how configs and settings are stored in disk and
// later retrieved.
type Store interface {
        // CreateConfig creates a new config with the given name and description.
        CreateConfig(name, desc string) (Config, error)

        // GetConfig returns the config with the given name if it exists otherwise
        // nil.
        GetConfig(name string) (Config, error)

        // GetConfigs returns all configs.
        GetConfigs() (configs []Config, err error)
}
