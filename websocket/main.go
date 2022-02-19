package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
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
	http.Handle("/", http.FileServer(http.Dir("static")))
	http.HandleFunc("/ws", socketHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))

}

func socketHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("hi!")
	upgrader.CheckOrigin = func(r *http.Request) bool {
		return true
	}
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}
	defer conn.Close()
	for {
		messageType, message, err := conn.ReadMessage()
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
		fmt.Println(test)
		err = conn.WriteJSON(&test)
		fmt.Println(err)

	}
	log.Println(conn)
}
