package sqlstore

import (
	"database/sql"
	"time"

	"github.com/vlence/configman"
	"github.com/vlence/gossert"
)

const configIdUnknown int64 = -1

type SqlConfig struct {
        id int64
        name string
        desc string
        createdAt time.Time
        updatedAt time.Time
        store *SqlStore
}

func newSqlConfig(store *SqlStore) *SqlConfig {
        config := &SqlConfig{
                id: configIdUnknown,
                store: store,
        }

        return config
}

func (config *SqlConfig) Name() string {
        return config.name
}

func (config *SqlConfig) Desc() string {
        return config.desc
}

func (config *SqlConfig) SetDesc(desc string) (bool, error) {
        gossert.Ok(config.id != configIdUnknown, "sqlstore: attempting to update description of config with unknown id")

        var affected int64
        var result sql.Result
        var err error

        now := time.Now()
        result, err = config.store.setDescStmt.Exec(desc, now.Unix(), config.id)

        if err != nil {
                return false, nil
        }

        affected, err = result.RowsAffected()
        descUpdated := affected > 0

        if descUpdated {
                config.desc = desc
                config.updatedAt = now
        }

        return descUpdated, err
}

func (config *SqlConfig) NewSetting(name string, typ configman.Type, value any) (configman.Setting, error) {
        config.store.db.Exec(``)
}
