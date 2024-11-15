package task

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func UpdateAccessToken() {
	// ここに定期実行したい処理を書く
	type reqBody struct {
		UserID         string `json:"user_id"`
		MealName       string `json:"meal_name"`
		CaloriePer100g int    `json:"calorie_per_100g"`
		Date           string `json:"date"`
	}
	reqBody1 := reqBody{
		UserID:         "uids893njkdf89",
		MealName:       "カレー",
		CaloriePer100g: 100,
		Date:           "2024-11-15",
	}
	fmt.Println(reqBody1)
	jsonData, err := json.Marshal(reqBody1)
	client := &http.Client{}
	resp, err := client.Post("http://app:8080/meal/create", "application/json", bytes.NewBuffer(jsonData))
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
