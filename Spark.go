package main

import (
	"os"
	"fmt"
	"net/http"
	"encoding/json"
	"bytes"
)

func main() {

	argone := os.Args[1]
	fmt.Println("Sending to Spark Room: " + argone)

	token := "<Your_Spark_Token>"
	roomid := "<Spark_Room_ID>"
	url := "https://api.ciscospark.com/v1/messages"
	msg := SparkRoomMessage{RoomId: roomid, Text: argone}
	jsonString, err := json.Marshal(msg)
	if err != nil {
		panic(err)
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonString))
	req.Header.Set("Authorization", "Bearer "+token)
	req.Header.Set("Content-type", "application/json")
	client := &http.Client{}
	resp, err := client.Do(req)
	resp.Body.Close()

}

type SparkRoomMessage struct {
	RoomId string `json:"roomId"`
	Text   string `json:"text"`
}
