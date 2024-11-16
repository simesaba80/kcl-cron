package main

import (
	"fmt"
	"log"
	"time"

	"github.com/go-co-op/gocron"
	"github.com/simesaba80/go-cron/db"
	"github.com/simesaba80/go-cron/task"
	"github.com/simesaba80/go-cron/utils"
)

func funcC() {
	if err := db.SqlDB.Ping(); err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
		return
	}
	fmt.Println("Ping OK")
}

func main() {
	utils.LoadConfig()
	db.Connectdb()
	defer db.DB.Close()
	scheduler := gocron.NewScheduler(time.Local)
	scheduler.Every(15).Minute().Do(task.SaveActivity)  // 30分に1回
	scheduler.Every(15).Minute().Do(task.SaveSleepData) // 30秒に1回
	// scheduler.StartBlocking()
}
