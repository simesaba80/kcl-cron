package main

import (
	"fmt"
	"time"

	"github.com/go-co-op/gocron"
)

func funcC() {
	fmt.Println("C")
}

func main() {

	scheduler := gocron.NewScheduler(time.Local)
	scheduler.Every(5).Second().Do(funcC) // 30分に1回
	scheduler.StartBlocking()
}
