package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/websocket"
	"github.com/joho/godotenv"

	"goorm_socket/config"
)

const socketBufferSize = 1024

var (
	//http request와 http response를 websocket에 맞게 커스터마이징 시켜준다.
	upgrader = &websocket.Upgrader{}
)

type testJson struct {
	A string
	B int
	C map[string]string
}

//javascript console에서 w = new WebSocket("ws://15.164.220.65:8080/ws") 시 실행
//w.send("message")로 테스트
func main() {
	err := godotenv.Load("../.env")
	if !config.KafkaSetting() {
		panic("kafka setting error")
	}
	log.Println(err)
	http.Handle("/", http.FileServer(http.Dir("static")))
	http.HandleFunc("/ws", socketHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))

}

//url parameter -> r.URL.Query()
func socketHandler(w http.ResponseWriter, r *http.Request) {
	upgrader.CheckOrigin = func(r *http.Request) bool {
		return true
	}
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}
	defer conn.Close()
	params := r.URL.Query()
	userID := params.Get("userid")
	if userID == "" {
		fmt.Println("websocket url error")
		return
	}
	//토픽 생성하기
	fmt.Println(userID)
	config.MakeTopic(userID)
	//소켓 연결 됐을 시 topic 확인 후 생성 코드 필요
	go config.KafkaProduce()
	time.Sleep(1 * time.Second)
	go config.KafkaConsumer()
	for {
		messageType, message, err := conn.ReadMessage() //사용자에게만 보낼 때 사용
		if err != nil {
			log.Println("err", err)
		}
		fmt.Println(string(message), "    ", messageType)
		test := testJson{
			A: "a1",
			B: 1,
			C: map[string]string{
				"test": "1",
			},
		}
		err = conn.WriteJSON(&test)
		if err != nil {
			log.Println(err)
		}

	}
}
