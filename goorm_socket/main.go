package main

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/websocket"
	"github.com/joho/godotenv"

	"goorm_socket/config"
	"goorm_socket/utils"
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
	config.ConnectDB()
	if !config.KafkaSetting() {
		panic("kafka setting error")
	}
	utils.ErrorCheck(err)
	producer := config.KafkaProduce() //produce는 서버당 하나 생성
	config.ConnectBroker()
	defer producer.ChatProducer.Close()

	http.Handle("/", http.FileServer(http.Dir("static")))
	http.HandleFunc("/ws", socketHandler)
	//웹소켓 포트 연결
	log.Fatal(http.ListenAndServe(":"+os.Getenv("WEBSOCKET_PORT"), nil))

	//웹 사이트(백엔드) 포트 연결

}

//url parameter -> r.URL.Query()
//각 소켓은 하나의 유저 -> 하나의 유저는 참여하고 있는 채팅방의 메시지를 받을 수 있어야함
//-> 채팅 ID로 topic 생성, all topic 생성 -> 유저는 참여한 채팅방 ID의 토픽을 구독
func socketHandler(w http.ResponseWriter, r *http.Request) {
	upgrader.CheckOrigin = func(r *http.Request) bool {
		return true
	}
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		utils.ErrorCheck(err)
		return
	}
	defer conn.Close()
	params := r.URL.Query()
	userID := params.Get("userid")
	if userID == "" {
		utils.ErrorCheck(errors.New("websocket url error"))
		conn.WriteJSON(&config.ResultJson{
			Command: "error",
			Value:   "userid Error",
		})
		return
	}
	//토픽 생성하기
	fmt.Println(userID)
	//소켓 연결 됐을 시 topic 확인 후 생성 코드 필요
	config.MakeTopic(userID)
	// config.Producer.Send(userID, "messageTest", 1)

	go config.KafkaConsumer() //Consumer는 사용자당 하나 생성
	for {
		messageType, message, err := conn.ReadMessage() //사용자에게만 보낼 때 사용
		if err != nil {
			log.Println("err", err)
		}
		fmt.Println(string(message), "    ", messageType)
		// config.Producer.Send("2", string(message))//topic send 테스트
		config.ConverToFunc(message)
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
