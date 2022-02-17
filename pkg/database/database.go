package db

import (
	"database/sql"
	"fmt"
	"os"
	"strings"

	_ "github.com/denisenkom/go-mssqldb"
	"github.com/doug-martin/goqu/v9"
	_ "github.com/doug-martin/goqu/v9/dialect/mysql"
	_ "github.com/doug-martin/goqu/v9/dialect/sqlite3"
	_ "github.com/doug-martin/goqu/v9/dialect/sqlserver"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/lib/pq"
	_ "github.com/mattn/go-sqlite3"
	_ "gopkg.in/doug-martin/goqu.v5/adapters/postgres"
)

type Database struct {
	Type     string
	Username string
	Password string
	DbName   string
	Host     string
	Port     string
	SSLMode  string
	TimeZone string
	Charset  string
}

// creates connections and returns database query builders and its connections and error if anything wrong happened
func New(dbs map[string]Database) (map[string]*goqu.Database, map[string]*sql.DB, error) {
	outs := map[string]*goqu.Database{}
	cons := map[string]*sql.DB{}

	for k, v := range dbs {
		dialect := goqu.Dialect(strings.ToLower(v.Type))
		config := ""

		switch strings.ToLower(v.Type) {
		case "mysql":
			config = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", v.Username, v.Password, v.Host, v.Port, v.DbName)
		case "sqlite3":
			if _, err := os.Stat(v.DbName); err != nil {
				_, err := os.Create(v.DbName)
				if err != nil {
					return nil, nil, err
				}
			}
			config = fmt.Sprintf("file:%s?cache=shared&mode=rw", v.DbName)
		case "postgres":
			config = fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s", v.Host, v.Port, v.Username, v.Password, v.DbName, v.SSLMode)
		case "mssql":
			config = fmt.Sprintf("server=%s;user id=%s;password=%s;port=%s;database=%s;", v.Host, v.Username, v.Password, v.Port, v.DbName)
		default:
			return nil, nil, fmt.Errorf("%s database not supported", v.Type)
		}

		c, err := sql.Open(v.Type, config)
		if err != nil {
			return nil, nil, err
		}

		cons[fmt.Sprintf("%s,%s", k, v.Type)] = c
		outs[fmt.Sprintf("%s,%s", k, v.Type)] = dialect.DB(c)
	}

	return outs, cons, nil
}

func CloseDBs(cons map[string]*sql.DB) {
	for _, con := range cons {
		con.Close()
	}
}
