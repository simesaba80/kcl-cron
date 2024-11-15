package task

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/simesaba80/go-cron/crud"
	"github.com/simesaba80/go-cron/db"
	"github.com/simesaba80/go-cron/utils"
)

func GetSleep() {
	// ここに定期実行したい処理を書く
	//リクエスト先のURLの指定とGETリクエストの送信
	client := &http.Client{}
	resp, err := client.Get(utils.SLEEPURL)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	//bodyデータの読み取り
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	//json形式の[]byteから構造体へパース(構造体はdb/model.goに定義)
	//パース先は構造体単体でもスライスでもOK
	users := []db.User{}
	err = json.Unmarshal(body, &users)
	if err != nil {
		panic(err)
	}
	//crudにSQLを定義して関数呼び出しで保存
	crud.SaveUsersData(users)
	//ログ出力
	fmt.Println(resp.Status)
	fmt.Println(users)
}
