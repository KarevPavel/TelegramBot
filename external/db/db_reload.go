package db

import (
	"bitbucket.org/y4cxp543/telegram-bot/constants"
	"database/sql"
	"fmt"
	"log"
)

type DbReload struct {
	connectionString string
	db *sql.DB
}

func NewDbReload() *DbReload {
	return &DbReload{
		connectionString: connectionUrl(),
	}
}

func (dbReload DbReload) Connect(connectionString string) *sql.DB {
	db, err := sql.Open("postgres", connectionString)
	if err != nil {
		log.Panicln(err)
	}
	defer db.Close()
	err = db.Ping()
	if err != nil {
		log.Panicln(err)
	}
	return db
}

func (dbReload DbReload) execQuery(query string, args...string) sql.Result {
	exec, err := dbReload.db.Exec(query, args)
	if err != nil {
		log.Panicln(err)
	}
	return exec
}

func connectionUrl() string {
	var connCfg = constants.Config.DbConnection;
	return fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		connCfg.Host, connCfg.Port, connCfg.User, connCfg.Password, connCfg.DbName)
}