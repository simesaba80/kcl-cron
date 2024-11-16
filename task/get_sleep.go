package task

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/simesaba80/go-cron/crud"
	"github.com/simesaba80/go-cron/db"
)

func SaveSleepData() {
	// ここに定期実行したい処理を書く
	fitbitUserData := crud.GetFitbitUserID()
	fmt.Println(fitbitUserData[0])
	client := &http.Client{}
	req, err := http.NewRequest("GET", "https://api.fitbit.com/1.2/user/"+fitbitUserData[0].FitbitUserID+"/sleep/date/2024-11-16.json", nil)
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
	responsData := bodySleep{}
	if err := json.Unmarshal(body, &responsData); err != nil {
		panic(err)
	}
	sleep := db.Sleep{
		UserID:     "C5N3BG",
		Minutes:    responsData.Summary.TotalMinutesAsleep,
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
