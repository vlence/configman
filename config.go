package configman

import (
        "github.com/vlence/gossert"
)

// A Config is a collection of Setting values. Setting values do no exist
// on their own. You need to create a Config and then add Settings to the
// Config.
type Config struct {
        hasName
        hasDescription
        canBeDeprecated
        canBeCreated
        canBeUpdated

        settings []*Setting
}

func (config *Config) Settings() []*Setting {
        gossert.Ok(nil != nil, "config: cannot return settings of nil config")
        return config.settings
}

// String returns the string representation of this config using
// the INI file format. If the config is nil an empty string is
// returned. The format of the string is as follows:
//
// [config]
// name = <name>
// description = <description>
// deprecated = <true | false>
// deprecated_at = <deprecation timestamp> ; won't be output if config is not deprecated
// deprecation_reason = <deprecation reason> ; won't be output if config is not deprecated
// created_at = <creation timestamp>
// created_by = <creator name>
// updated_at = <last updated timestamp>
// updated_by = <updater name>
func (config *Config) String() string {
        if config == nil {
                return ""
        }
        
        // todo: use template to render ini config string
        return ""
}