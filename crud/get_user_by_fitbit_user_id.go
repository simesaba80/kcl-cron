package crud

import (
	"context"
	"fmt"

	"github.com/simesaba80/go-cron/db"
)

func GetUserByFitbitUserID(fitbitUserID string) (db.User, error) {
	user := db.User{}
	ctx := context.Background()
	fmt.Println("fitbitUserID:" + fitbitUserID)
	if err := db.DB.NewSelect().Model(&user).Where("fitbit_user_id = ?", fitbitUserID).Scan(ctx); err != nil {
		return user, err
	}
	return user, nil
}
