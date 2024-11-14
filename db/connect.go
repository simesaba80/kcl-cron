package db

import (
	"database/sql"
	"log"

	"github.com/simesaba80/go-cron/utils"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
	"github.com/uptrace/bun/driver/pgdriver"
	"github.com/uptrace/bun/extra/bundebug"
)

var DB *bun.DB
var SqlDB *sql.DB

var PostDB *bun.DB
var PostSqlDB *sql.DB

func Connectdb() {
	// Bunを使ってDB接続
	dsn := utils.DBURL1
	SqlDB = sql.OpenDB(pgdriver.NewConnector(pgdriver.WithDSN(dsn)))
	DB = bun.NewDB(SqlDB, pgdialect.New())
	if err := SqlDB.Ping(); err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
		return
	}

	dsnp := utils.DBURL2
	PostSqlDB = sql.OpenDB(pgdriver.NewConnector(pgdriver.WithDSN(dsnp)))
	PostDB = bun.NewDB(PostSqlDB, pgdialect.New())
	if err := PostSqlDB.Ping(); err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
		return
	}
	// クエリーフックを追加
	DB.AddQueryHook(bundebug.NewQueryHook(
		bundebug.WithVerbose(true),
	))
	PostDB.AddQueryHook(bundebug.NewQueryHook(
		bundebug.WithVerbose(true),
	))
}
