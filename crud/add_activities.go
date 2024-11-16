package crud

import (
	"context"
	"time"

	"github.com/simesaba80/go-cron/db"
)

func AddActivities(activity db.Activities) db.Activities {
	// ここに定期実行したい処理を書く
	ctx := context.Background()
	activity.CreatedAt = time.Now()
	activity.Date = time.Now().Format("2006-01-02")
	db.DB.NewInsert().Model(&activity).Exec(ctx)
	return activity
}
