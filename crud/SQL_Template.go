package crud

import (
	"context"
	"time"

	"github.com/simesaba80/go-cron/db"
)

// 引数は保存したいデータの構造体または構造体スライス
func SQLName(users []db.User) {
	//構造体のCreatedAtを現在時間に設定
	for i := range users {
		users[i].ID = 0
		users[i].CreatedAt = time.Now()
	}
	//PostDBに対してInsertを実行
	//SQLの実行にはcontext.Contextが必要なのでcontext.Background()を渡す
	ctx := context.Background()
	db.PostDB.NewInsert().Model(&users).Exec(ctx)
	//bunのインサート周りのドキュメント→https://bun.uptrace.dev/guide/query-insert.html
}
