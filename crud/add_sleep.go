package crud

import (
	"context"
	"time"

	"github.com/simesaba80/go-cron/db"
)

func AddSleep(sleep db.Sleep) db.Sleep {
	ctx := context.Background()
	sleep.CreatedAt = time.Now()
	sleep.Date = time.Now().Format("2006-01-02")
	db.DB.NewInsert().Model(&sleep).Exec(ctx)
	return sleep
}
