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
		UserID        string  `json:"user_id"`
		MealName      string  `json:"meal_name"`
		Calories      int     `json:"calories"`
		Protein       float64 `json:"protein"`
		Fat           float64 `json:"fat"`
		Carbohydrates float64 `json:"carbohydrates"`
		Salt          float64 `json:"salt"`
		Date          string  `json:"date"`
	}
	reqBody1 := reqBody{
		UserID:        "uids893njkdf89",
		MealName:      "カレー",
		Calories:      1000,
		Protein:       20.0,
		Fat:           30.0,
		Carbohydrates: 40.0,
		Salt:          5.0,
		Date:          "2024-11-15",
	}
	fmt.Println(reqBody1)
	jsonData, _ := json.Marshal(reqBody1)
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
