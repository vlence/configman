package sqlstore

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/vlence/configman"
	"github.com/vlence/gossert"
)

var errGetConfig = fmt.Errorf("sqlstore: failed to get config")
var errPrepStmts = fmt.Errorf("sqlstore: failed to prepare sql statements")
var errGetConfigs = fmt.Errorf("sqlstore: failed to get configs")
var errScanConfig = fmt.Errorf("sqlstore: failed to scan config")
var errConfigsTable = fmt.Errorf("sqlstore: failed to create configs table")
var errCreateConfig = fmt.Errorf("sqlstore: failed to create config")
var errSettingsTable = fmt.Errorf("sqlstore: failed to create settings table")

type RowScanner interface {
        Scan(dest ...any) error
}

// SqlStore is a configman.Store that uses a SQL database as its storage
// engine.
type SqlStore struct {
        // The underlying SQL database
        db *sql.DB

        // Prepared statement. Execute it to get a config by name.
        getConfigStmt    *sql.Stmt
        getConfigsStmt   *sql.Stmt
        setDescStmt      *sql.Stmt
        createConfigStmt *sql.Stmt
}

// NewSqlStore creates a new SqlStore using the given *sql.DB.
func NewSqlStore(db *sql.DB) (*SqlStore, error) {
        var err error

        gossert.Ok(db != nil, "sqlstore: received nil instead of *sql.DB")

        store := new(SqlStore)
        store.db = db

        if err = store.init(); err != nil {
                return nil, err
        }

        if err = store.prepStmts(); err != nil {
                return nil, err
        }

        return store, nil
}

// init creates all the tables and indices required.
func (store *SqlStore) init() error {
        var err error

        if err = store.initConfigsTable(); err != nil {
                return err
        }

        if err = store.initSettingsTable(); err != nil {
                return err
        }

        return nil
}

// initConfigsTable creates the configs table and its indices.
func (store *SqlStore) initConfigsTable() error {
        var tx *sql.Tx
        var txErr, commitErr, rollbackErr, execErr error

        tx, txErr = store.db.Begin()

        if txErr != nil {
                return errors.Join(errConfigsTable, txErr)
        }

        _, execErr = tx.Exec(`
        CREATE TABLE IF NOT EXISTS configs (
                id INTEGER PRIMARY KEY,
                name TEXT NOT NULL,
                desc TEXT,
                created_at INTEGER NOT NULL,
                updated_at INTEGER NOT NULL
        )
        `)

        if execErr != nil {
                if rollbackErr = tx.Rollback(); rollbackErr != nil {
                        panic(errors.Join(errConfigsTable, rollbackErr, execErr))
                }

                return errors.Join(errConfigsTable, execErr)
        }

        _, execErr = tx.Exec(`
        CREATE INDEX IF NOT EXISTS config_name_index ON configs (name)
        `)

        if execErr != nil {
                if rollbackErr = tx.Rollback(); rollbackErr != nil {
                        panic(errors.Join(errConfigsTable, rollbackErr, execErr))
                }

                return errors.Join(errConfigsTable, execErr)
        }

        commitErr = tx.Commit()

        if commitErr != nil {
                return errors.Join(errConfigsTable, commitErr)
        }

        return nil
}

func (store *SqlStore) initSettingsTable() error {
        var tx *sql.Tx
        var txErr, commitErr, rollbackErr, execErr error

        if tx, txErr = store.db.Begin(); txErr != nil {
                return errors.Join(errSettingsTable, txErr)
        }

        _, execErr = tx.Exec(`
                CREATE TABLE IF NOT EXISTS settings (
                        id INTEGER PRIMARY KEY,
                        name TEXT NOT NULL,
                        desc TEXT NOT NULL,
                        created_at INTEGER NOT NULL,
                        updated_at INTEGER NOT NULL,
                        deprecated BOOLEAN NOT NULL DEFAULT FALSE,
                        deprecation_reason TEXT NOT NULL DEFAULT '',
                        deprecated_at INTEGER NOT NULL,
                        config_id INTEGER NOT NULL,
                        config_name TEXT NOT NULL,
                        value_type INTEGER NOT NULL,
                        uint32_value INTEGER,
                        uint64_value INTEGER,
                        int32_value INTEGER,
                        int64_value INTEGER,
                        float32_value INTEGER,
                        float64_value INTEGER,
                        bool_value BOOLEAN,
                        string_value TEXT
                )
        `)

        if execErr != nil {
                if rollbackErr = tx.Rollback(); rollbackErr != nil {
                        panic(errors.Join(errSettingsTable, rollbackErr, execErr))
                }

                return errors.Join(errSettingsTable, execErr)
        }

        _, execErr = tx.Exec(`
                CREATE INDEX IF NOT EXISTS settings_configname_name_index ON settings (
                        config_name,
                        name
                )
        `)

        if execErr != nil {
                if rollbackErr = tx.Rollback(); rollbackErr != nil {
                        panic(errors.Join(errSettingsTable, rollbackErr, execErr))
                }

                return errors.Join(errSettingsTable, execErr)
        }

        if commitErr = tx.Commit(); commitErr != nil {
                return errors.Join(errSettingsTable, commitErr)
        }

        return nil
}

// prepStmts prepares all SQL statements that will be used
// to manage configs in the SQL database.
func (store *SqlStore) prepStmts() error {
        var err error

        store.getConfigStmt, err = store.db.Prepare("SELECT * FROM configs WHERE name = ?")

        if err != nil {
                return errors.Join(errPrepStmts, err)
        }

        store.getConfigsStmt, err = store.db.Prepare("SELECT * FROM configs")

        if err != nil {
                return errors.Join(errPrepStmts, err)
        }

        store.setDescStmt, err = store.db.Prepare(`
                UPDATE configs
                SET desc = ?,
                    updated_at = ?
                WHERE id = ?
        `)

        if err != nil {
                return errors.Join(errPrepStmts, err)
        }

        store.createConfigStmt, err = store.db.Prepare(`
                INSERT INTO configs (
                        name,
                        desc,
                        created_at,
                        updated_at
                ) VALUES (?, ?, ?, ?)
        `)

        if err != nil {
                return errors.Join(errPrepStmts, err)
        }

        return nil
}

// GetConfig finds the config with the given name and returns it.
// If a config with the given name does not exist then nil is returned.
func (store *SqlStore) GetConfig(name string) (configman.Config, error) {
        config, err := store.scanConfig(store.getConfigStmt.QueryRow(name))

        if err != nil {
                return nil, err
        }

        if config == nil {
                return nil, nil
        }

        return config, nil
}

func (store *SqlStore) GetConfigs() ([]configman.Config, error) {
        var rows *sql.Rows
        var err error
        var config configman.Config

        configs := make([]configman.Config, 0)

        if rows, err = store.getConfigsStmt.Query(); err != nil {
                return configs, errors.Join(errGetConfigs, err)
        }

        defer rows.Close()

        for rows.Next() {
                if config, err = store.scanConfig(rows); err != nil {
                        return configs, errors.Join(errGetConfigs, err)
                }

                if config == nil {
                        continue
                }

                configs = append(configs, config)
        }

        if err = rows.Err(); err != nil {
                return configs, errors.Join(errGetConfigs, err)
        }

        return configs, nil
}

// CreateConfig creates a new config using the given name and description
// and returns it.
func (store *SqlStore) CreateConfig(name, desc string) (configman.Config, error) {
        var err error
        var rows int64
        var result sql.Result

        now := time.Now()
        result, err = store.createConfigStmt.Exec(name, desc, now.Unix(), now.Unix())

        if err != nil {
                return nil, errors.Join(errCreateConfig, err)
        }

        if rows, err = result.RowsAffected(); err != nil {
                return nil, errors.Join(errCreateConfig, err)
        }

        if rows == 0 {
                log.Println("warn: no rows affected when creating config")
        }

        config := newSqlConfig(store)
        config.name = name
        config.desc = desc
        config.createdAt = now
        config.updatedAt = now

        if config.id, err = result.LastInsertId(); err != nil {
                return config, errors.Join(errCreateConfig, err)
        }

        gossert.Ok(config.id != configIdUnknown, "sqlstore: created config with unknown id")

        return config, nil
}

func (store *SqlStore) CreateSetting(name string, typ configman.Type, value any) (configman.Setting, error) {
        switch typ {
        case configman.Int32:
                v, ok := value.(int32)
                gossert.Ok(ok, fmt.Sprintf("sqlstore: expected int32 value but got value of type %T", value))

                return store.createInt32Setting(name, v)
        }
}

func (store *SqlStore) createInt32Setting(name string, value int32) (configman.Setting, error) {
}

// scanConfig scans the given row and returns a *SqlConfig. If no
// rows were returned then nil is returned.
func (store *SqlStore) scanConfig(row RowScanner) (*SqlConfig, error) {
        var id int64
        var name, desc string
        var createdAt, updatedAt int64

        err := row.Scan(
                &id,
                &name,
                &desc,
                &createdAt,
                &updatedAt,
        )

        if err == sql.ErrNoRows {
                return nil, nil
        }

        if err != nil {
                return nil, errors.Join(errScanConfig, err)
        }

        config := newSqlConfig(store)
        config.id = id
        config.name = name
        config.desc = desc
        config.createdAt = time.Unix(createdAt, 0)
        config.updatedAt = time.Unix(updatedAt, 0)

        return config, nil
}
