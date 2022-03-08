package tests

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"

	"github.com/doug-martin/goqu/v9"
	_ "github.com/doug-martin/goqu/v9/dialect/sqlite3"
	_ "github.com/mattn/go-sqlite3"
)

var testDBName = "test.db"
var testDBType = "sqlite3"

func init() {
	// Recreate test.db file
	if _, err := os.Stat(testDBName); err == nil {
		if err = os.Remove(testDBName); err != nil {
			log.Fatalln(err)
		}
	}
	_, err := os.Create(testDBName)
	if err != nil {
		log.Fatalln(err)
	}

	// Up migrations
	err = exec.Command("sql-migrate", "up").Run()
	if err != nil {
		log.Fatalln(err)
	}
}

// Create sqlite3 connection
func New() *goqu.Database {
	dialect := goqu.Dialect(strings.ToLower(testDBType))

	config := fmt.Sprintf("file:%s?cache=shared&mode=rw", testDBName)

	c, err := sql.Open(testDBType, config)
	if err != nil {
		log.Fatalln(err)
	}

	return dialect.DB(c)
}
