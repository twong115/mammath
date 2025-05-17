package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"os/signal"

	"github.com/gorilla/websocket"
)

func main() {
	c, _, err := websocket.DefaultDialer.Dial("ws://localhost:8080/", nil)
	if err != nil {
		log.Fatal("Dial:", err)
	}
	defer c.Close()

	done := make(chan struct{})

	go func() {
		defer close(done)
		for {
			_, message, err := c.ReadMessage()
			if err != nil {
				fmt.Println("Disconnected from server.")
				return
			}
			fmt.Println("\n[SERVER] " + string(message))
			fmt.Print("> ")
		}
	}()

	input := bufio.NewScanner(os.Stdin)
    fmt.Println("Enter your username: ")
	fmt.Print("> ")
	for input.Scan() {
		text := input.Text()
		if err := c.WriteMessage(websocket.TextMessage, []byte(text)); err != nil {
			log.Println("WriteMessage:", err)
			return
		}
	}

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt)

	<-interrupt
	log.Println("Interrupted")
}
