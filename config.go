package configman

import (
	"errors"
	"fmt"
)

// A config is a collection of settings. To create and use settings you need to create
// a config first.
type Config struct {
        // The name of the config.
        name string

        // The description of the config.
        desc string
}

// This error is thrown when a config is created using an invalid name.
// See IsValidConfigName for an explanation of a valid name.
var ErrInvalidConfigName = errors.New("invalid config name")

func NewConfig(name, desc string) (*Config, error) {
        if !IsValidName(name) {
                return nil, ErrInvalidConfigName
        }

        return &Config{name, desc}, nil
}

func (c *Config) Name() string {
        return c.name
}

func (c *Config) Desc() string {
        return c.desc
}

func (c *Config) UpdateDesc(newDesc string) {
        c.desc = newDesc
}

// Settings returns a list of settings in this config that
// match the given names. To get all settings call this
// function without any arguments. The list of settings
// returned may be smaller than the list of names given.
func (c *Config) Settings(names ...string) []*Setting {
        settings := make([]*Setting, 0)

        return settings
}

// Setting returns the setting with the given name. If no
// such setting is found then nil is returned.
func (c *Config) Setting(name string) *Setting {
        return nil
}

// Sets the value of the setting with the given name. If a
// setting with that name does not exist then a new setting
// with that name is created. If a setting does exist and
// type value does not match the type of the setting then
// an error is returned.
func (c *Config) Set(name string, value any) (*Setting, error) {
        // check if setting exists

        // if setting does not exist

        setting := new(Setting)
        setting.name = name
        setting.value = value

        switch value.(type) {
        case uint8:
                setting.typ = Uint8
        case uint16:
                setting.typ = Uint16
        case uint32:
                setting.typ = Uint32
        case uint64:
                setting.typ = Uint64
        case int8:
                setting.typ = Int8
        case int16:
                setting.typ = Int16
        case int32:
                setting.typ = Int32
        case int64:
                setting.typ = Int64
        case float32:
                setting.typ = Float32
        case float64:
                setting.typ = Float64
        case bool:
                setting.typ = Bool
        case string:
                setting.typ = String
        default:
                return nil, fmt.Errorf("value of type %T not supported", value)
        }

        return setting, nil
}
