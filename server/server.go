package main

import (
	"fmt"
	"log"
	"net/http"
	"sync"
	"github.com/twong115/mammath/questions"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true // allow all connections
	},
}

var clients = make(map[*websocket.Conn]bool)
var broadcast = make(chan string)
var mu sync.Mutex

var currQuestion = questions.GenerateSimplePolynomial(3)

func main() {
	http.HandleFunc("/", handleConnections)
	go handleMessages()

	fmt.Println("Server started on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func handleConnections(w http.ResponseWriter, r *http.Request) {
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("Upgrade:", err)
		return
	}
	defer ws.Close()

	fmt.Println("A client connected!")

	mu.Lock()
	clients[ws] = true
	questionStr := "Question: " + currQuestion.GetQuestionString()
	ws.WriteMessage(websocket.TextMessage, []byte(questionStr))

	mu.Unlock()

	for {
		_, msg, err := ws.ReadMessage()
		if err != nil {
			mu.Lock()
			delete(clients, ws)
			mu.Unlock()
			break
		}
		broadcast <- string(msg)
	}
}

func broadcastMessage(msg string) {
	for client := range clients {
		err := client.WriteMessage(websocket.TextMessage, []byte(msg))
		if err != nil {
			client.Close()
			delete(clients, client)
		}
	}
}


func handleMessages() {
	for {
		userAns := <-broadcast
		mu.Lock()
		if currQuestion.GetSolutionString() == userAns {
			broadcastMessage("A user has gotten the correct answer: " + currQuestion.GetSolutionString())
			currQuestion = questions.GenerateSimplePolynomial(3)
			broadcastMessage("New question: " + currQuestion.GetQuestionString())
		} else {
			broadcastMessage("A user has guessed the wrong answer")
		}
		mu.Unlock()
	}
}
