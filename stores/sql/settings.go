package sqlstore

import "github.com/vlence/configman"

type SqlSetting struct {
        typ configman.Type
        name string
        value any
        deprecated bool
        description string
        deprecationReason string

        store *SqlStore
        config *SqlConfig
}

func NewSqlSetting(name string, typ configman.Type, value any) *SqlSetting {
        setting := new(SqlSetting)

        setting.typ = typ
        setting.name = name
        setting.value = value

        return setting
}
