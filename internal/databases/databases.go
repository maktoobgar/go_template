package databases

import (
	"database/sql"
	"fmt"
	"strings"

	"github.com/maktoobgar/go_template/internal/config"
	g "github.com/maktoobgar/go_template/internal/global"
	db "github.com/maktoobgar/go_template/pkg/database"
)

func New(cfg *config.Config) (map[string]*sql.DB, error) {
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

func SetConnections(cons map[string]*sql.DB) error {
	mainOrTest := "test"
	if !g.CFG.Debug {
		mainOrTest = "main"
	}
	for k, v := range cons {
		dbName := strings.Split(k, ",")[0]
		dbType := strings.Split(k, ",")[1]
		if dbName == mainOrTest {
			g.DB = v
		}
		switch dbType {
		case "postgres":
			g.PostgresCons[dbName] = v
		case "sqlite3":
			g.SqliteCons[dbName] = v
		case "mysql":
			g.MySQLCons[dbName] = v
		case "mssql":
			g.SqlServerCons[dbName] = v
		default:
			return fmt.Errorf("%s database not supported", strings.Split(k, ",")[1])
		}
		g.AllCons[dbName] = v
	}

	return nil
}

func Setup(cfg *config.Config) error {
	cons, err := New(cfg)
	if err != nil {
		return err
	}

	err = SetConnections(cons)
	if err != nil {
		return err
	}

	return nil
}
