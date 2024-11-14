package main

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	"github.com/go-co-op/gocron"
	"github.com/simesaba80/go-cron/utils"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
	"github.com/uptrace/bun/driver/pgdriver"
	"github.com/uptrace/bun/extra/bundebug"
)

func funcC() {
	if err := SqlDB.Ping(); err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
		return
	}
	fmt.Println("Ping OK")
}

var DB *bun.DB
var SqlDB *sql.DB

func Connectdb() {
	// Bunを使ってDB接続
	dsn := utils.DBURL
	SqlDB = sql.OpenDB(pgdriver.NewConnector(pgdriver.WithDSN(dsn)))
	DB = bun.NewDB(SqlDB, pgdialect.New())
	if err := SqlDB.Ping(); err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
		return
	}
	// クエリーフックを追加
	DB.AddQueryHook(bundebug.NewQueryHook(
		bundebug.WithVerbose(true),
	))
}

func main() {
	utils.LoadConfig()
	Connectdb()
	defer DB.Close()
	scheduler := gocron.NewScheduler(time.Local)
	scheduler.Every(5).Second().Do(funcC) // 30分に1回
	scheduler.StartBlocking()
}
