package task

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/simesaba80/go-cron/crud"
	"github.com/simesaba80/go-cron/db"
)

func SaveActivity() {
	// ここに定期実行したい処理を書く
	fitbitUserData := crud.GetFitbitUserID()
	fmt.Println(fitbitUserData[0])
	client := &http.Client{}
	req, err := http.NewRequest("GET", "https://api.fitbit.com/1/user/"+fitbitUserData[0].FitbitUserID+"/activities/date/2024-11-17.json", nil)
	if err != nil {
		panic(err)
	}
	req.Header.Set("authorization", "Bearer "+fitbitUserData[0].AccessToken)
	req.Header.Set("Content-Type", "application/json")
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	//bodyデータの読み取
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	type summary struct {
		CaloriesOut int `json:"caloriesOut"`
		Steps       int `json:"steps"`
	}
	type activities struct {
		Summary summary `json:"summary"`
	}
	//json形式の[]byteから構造体へパース
	//構造体のフィールドに`json:"name"`のようにタグをつけることで、jsonのキーと構造体のフィールドを紐付ける(タグがなければフィールド名と同じキーを探す)
	//`json:"-"のフィールドは無視される`
	responsData := activities{}
	if err := json.Unmarshal(body, &responsData); err != nil {
		panic(err)
	}
	activitiesData := db.Activities{
		UserID:   fitbitUserData[0].FitbitUserID,
		Calories: responsData.Summary.CaloriesOut,
		Steps:    responsData.Summary.Steps,
	}
	result := crud.AddActivities(activitiesData)
	fmt.Println(result)

	fmt.Println(resp.Status)
}
