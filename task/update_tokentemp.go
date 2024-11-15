package task

import (
	"fmt"
	"io"
	"net/http"
)

func Test() {
	// ここに定期実行したい処理を書く
	client := &http.Client{}
	resp, err := client.Get("https://www.fitbit.com/oauth2/authorize?response_type=code&client_id=23PRT4&redirect_uri=https://lifecampusu.jp/callback&scope=activity" + "%20" + "heartrate" + "%20" + "sleep")
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
	// meal := []db.Meal{}
	// err = json.Unmarshal(body, &meal)
	// if err != nil {
	// 	panic(err)
	// }
	// // crud.SaveUsersData(meal)
	fmt.Println(string(body))
	fmt.Println(resp.Status)
	// fmt.Println(meal)
}
