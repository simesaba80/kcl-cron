package task

import (
	"fmt"
	"io"
	"net/http"
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
	fmt.Println(resp.Status)
	fmt.Println(string(body))
}
