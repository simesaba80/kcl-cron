package task

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/simesaba80/go-cron/db"
)

func Test() {
	// ここに定期実行したい処理を書く
	client := &http.Client{}
	req, err := http.NewRequest("GET", "https://api.fitbit.com/1.2/user/C5N3BG/sleep/date/2024-11-16.json", nil)
	if err != nil {
		panic(err)
	}
	req.Header.Set("authorization", "Bearer eyJhbGciOiJIUzI1NiJ9.eyJhdWQiOiIyM1BSVDQiLCJzdWIiOiJDNU4zQkciLCJpc3MiOiJGaXRiaXQiLCJ0eXAiOiJhY2Nlc3NfdG9rZW4iLCJzY29wZXMiOiJyYWN0IHJociByc2xlIiwiZXhwIjoxNzMxNzk5ODA5LCJpYXQiOjE3MzE3NzEwMDl9.vkGsFDrCNYvklvZ-Nqh7MyOtec1tV8RBzSOJexF5VWk")
	req.Header.Set("Content-Type", "application/json")
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	//bodyデータの読み取り
	type sleepStage struct {
		Deep  int `json:"deep"`
		Light int `json:"light"`
		Rem   int `json:"rem"`
		Wake  int `json:"wake"`
	}
	type sleepSummary struct {
		Stages             sleepStage `json:"stages"`
		TotalMinutesAsleep int        `json:"totalMinutesAsleep"`
		TotalSleepRecords  int        `json:"totalSleepRecords"`
		TotalTimeInBed     int        `json:"totalTimeInBed"`
	}
	type bodySleep struct {
		Summary sleepSummary `json:"summary"`
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	//json形式の[]byteから構造体へパース
	//構造体のフィールドに`json:"name"`のようにタグをつけることで、jsonのキーと構造体のフィールドを紐付ける(タグがなければフィールド名と同じキーを探す)
	//`json:"-"のフィールドは無視される`
	// jsonBytes := []byte(body)
	// err = json.Unmarshal(body, &jsonBytes)
	// if err != nil {
	// 	panic(err)
	// }
	// crud.SaveUsersData(meal)
	responsData := bodySleep{}
	if err := json.Unmarshal(body, &responsData); err != nil {
		panic(err)
	}
	fmt.Println("body: " + string(body))
	fmt.Println("responsData: ", responsData)
	sleep := db.Sleep{
		UserID:     "C5N3BG",
		Hours:      responsData.Summary.TotalMinutesAsleep,
		DeepSleep:  responsData.Summary.Stages.Deep,
		LightSleep: responsData.Summary.Stages.Light,
		RemSleep:   responsData.Summary.Stages.Rem,
		Wake:       responsData.Summary.Stages.Wake,
		Date:       "2024-11-16",
	}
	fmt.Println(sleep)
	fmt.Println(resp.Status)
	// fmt.Println(meal)
}
