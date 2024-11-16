package crud

import (
	"context"
	"fmt"

	"github.com/simesaba80/go-cron/db"
)

func GetFitbitUserID() []db.ToGetFitbitData {
	ctx := context.Background()
	var fitbitUserData []db.ToGetFitbitData
	err := db.DB.NewSelect().Model(&fitbitUserData).Column("fitbit_user_id", "access_token").Order("created_at DESC").Scan(ctx)
	if err != nil {
		fmt.Println("err")
		panic(err)
	}
	return fitbitUserData
}
