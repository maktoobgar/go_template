package databases

import (
	"database/sql"
	"fmt"
	"strings"

	"github.com/doug-martin/goqu/v9"
	"github.com/maktoobgar/go_template/internal/config"
	g "github.com/maktoobgar/go_template/internal/global"
	db "github.com/maktoobgar/go_template/pkg/database"
)

func New(cfg *config.Config) (map[string]*goqu.Database, map[string]*sql.DB, error) {
	dbs := cfg.Databases

	in := map[string]db.Database{}
	for _, v := range dbs {
		in[v.Name] = db.Database{
			Type:     v.Type,
			Username: v.Username,
			Password: v.Password,
			DbName:   v.DBName,
			Host:     v.Host,
			Port:     v.Port,
			SSLMode:  v.SSLMode,
			TimeZone: v.TimeZone,
			Charset:  v.Charset,
		}
	}

	return db.New(in)
}

func SetDBs(dbs map[string]*goqu.Database) error {
	for k, v := range dbs {
		dbName := strings.Split(k, ",")[0]
		dbType := strings.Split(k, ",")[1]
		switch dbType {
		case "postgres":
			if dbName == "main" && g.DB == nil {
				g.DB = v
			}
			g.Postgres[dbName] = v
		case "sqlite3":
			if dbName == "main" && g.DB == nil {
				g.DB = v
			}
			g.Sqlite[dbName] = v
		case "mysql":
			if dbName == "main" && g.DB == nil {
				g.DB = v
			}
			g.MySQL[dbName] = v
		case "mssql":
			if dbName == "main" && g.DB == nil {
				g.DB = v
			}
			g.SqlServer[dbName] = v
		default:
			return fmt.Errorf("%s database not supported", strings.Split(k, ",")[1])
		}
	}

	return nil
}

func SetConnections(cons map[string]*sql.DB) error {
	for k, v := range cons {
		switch strings.Split(k, ",")[1] {
		case "postgres":
			g.PostgresCons[strings.Split(k, ",")[0]] = v
		case "sqlite3":
			g.SqliteCons[strings.Split(k, ",")[0]] = v
		case "mysql":
			g.MySQLCons[strings.Split(k, ",")[0]] = v
		case "mssql":
			g.SqlServerCons[strings.Split(k, ",")[0]] = v
		default:
			return fmt.Errorf("%s database not supported", strings.Split(k, ",")[1])
		}
	}

	return nil
}

func Setup(cfg *config.Config) error {
	dbs, cons, err := New(cfg)
	if err != nil {
		return err
	}

	err = SetDBs(dbs)
	if err != nil {
		return err
	}

	err = SetConnections(cons)
	if err != nil {
		return err
	}

	return nil
}
