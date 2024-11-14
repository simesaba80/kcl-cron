package task

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/simesaba80/go-cron/db"
)

func GetAllUser() {
	// ここに定期実行したい処理を書く
	client := &http.Client{}
	resp, err := client.Get("http://app:8080/user")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	users := []db.User{}
	err = json.Unmarshal(body, &users)
	if err != nil {
		panic(err)
	}
	fmt.Println(resp.Status)
	fmt.Println(users)
}
