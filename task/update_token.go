package task

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/simesaba80/go-cron/crud"
	"github.com/simesaba80/go-cron/db"
)

func UpdateAccessToken() {
	// ここに定期実行したい処理を書く
	client := &http.Client{}
	resp, err := client.Get("http://app:8080/")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	//bodyデータの読み取り
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	//json形式の[]byteから構造体へパース
	//構造体のフィールドに`json:"name"`のようにタグをつけることで、jsonのキーと構造体のフィールドを紐付ける(タグがなければフィールド名と同じキーを探す)
	//`json:"-"のフィールドは無視される`
	users := []db.User{}
	err = json.Unmarshal(body, &users)
	if err != nil {
		panic(err)
	}
	crud.SaveUsersData(users)
	fmt.Println(resp.Status)
	fmt.Println(users)
}
