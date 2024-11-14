package crud

import (
	"context"
	"time"

	"github.com/simesaba80/go-cron/db"
)

func SaveUsersData(users []db.User) {
	for i := range users {
		users[i].ID = 0
		users[i].CreatedAt = time.Now()
	}
	ctx := context.Background()
	db.PostDB.NewInsert().Model(&users).Exec(ctx)
}
