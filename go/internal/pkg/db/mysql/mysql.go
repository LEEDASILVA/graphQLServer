package mysql

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/golang-migrate/migrate"
	"github.com/golang-migrate/migrate/database/mysql"
	_ "github.com/golang-migrate/migrate/source/file"
)

var DB *sql.DB

func checkError(err error) {
	if err != nil {
		log.Panic(err)
	}
}

// connection to the data base
func InitDB() {
	db, err := sql.Open("mysql", "root:mysqlpass@(172.17.0.2:3306)/hackernews")
	checkError(err)
	if err = db.Ping(); err != nil {
		log.Panic(err)
	}
	DB = db
}

// this will aplly yhe migrations on the path internal/pkg/db/migrations/mysql
// just like the command `migrate -database mysql://root:*****@(172.17.0.2:3306)/hackernews -path internal/pkg/db/migrations/mysql up`
func Migrate() {
	if err := DB.Ping(); err != nil {
		log.Panic(err)
	}
	driver, err := mysql.WithInstance(DB, &mysql.Config{})
	checkError(err)
	m, err := migrate.NewWithDatabaseInstance(
		"file://internal/pkg/db/migrations/mysql",
		"mysql",
		driver,
	)
	checkError(err)
	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		log.Fatal(err)
	}
}
