package main

import (
	"fmt"
	"log"
	"net/http"
	"sync"
	"github.com/twong115/mammath/questions"
	"github.com/twong115/mammath/Server/user"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true // allow all connections
	},
}

var (
	clients = make(map[*websocket.Conn]*user.User)
	broadcast = make(chan string)
	user_conn = make(chan *websocket.Conn)
	mu sync.Mutex
)

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
    _, msg, err := ws.ReadMessage()
    if err != nil {
			mu.Lock()
			delete(clients, ws)
			mu.Unlock()
    }

	currUser := user.New(string(msg), 0)
	clients[ws] = currUser
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
        user_conn <- ws
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
        ws := <- user_conn
        currUser, ok := clients[ws]
		if !ok {
			continue
		}
		userAns := <-broadcast
		mu.Lock()
		if currQuestion.GetSolutionString() == userAns {
			currUser.SetPoints(currUser.GetPoints()+1)
			broadcastMessage(fmt.Sprintf("%s has gotten the correct answer: %s", currUser.GetName(), currQuestion.GetQuestionString()))
            broadcastMessage(fmt.Sprintf("%s has gotten %d question(s) correct", currUser.GetName(), currUser.GetPoints()))
			currQuestion = questions.GenerateSimplePolynomial(3)
			broadcastMessage("New question: " + currQuestion.GetQuestionString())
		} else {
			broadcastMessage(currUser.GetName() + " has guessed the wrong answer")
		}
		mu.Unlock()
	}
}
