package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/gorilla/websocket"
)

func main() {
	// var hexoDir string
	// fmt.Print("Input your directory: ")
	// fmt.Scan(&hexoDir)
	hexoDir := "/Users/xander/Documents/Hexo"
	ParseHexoDirectory(hexoDir)
	PrintAllPostsContent()
	fmt.Println("Waiting for connection...")
	setupRoutes()
	http.ListenAndServe(":8080", nil)
	fmt.Println("System Works")
}

// define Upgrader an set Read and Write buffer size
var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	// check coonectino origin
	CheckOrigin: func(r *http.Request) bool { return true },
}

// define reader which will listen for new messages being sent to our WebSocket endpoint
func reader(conn *websocket.Conn) {
	for {
		messageType, p, err := conn.ReadMessage()
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println("Reading...")
		fmt.Println("Recived: " + string(p))
		retMsg := processCode(string(p))
		fmt.Println("Sending: " + string(retMsg))
		if err := conn.WriteMessage(messageType, retMsg); err != nil {
			log.Println(err)
			return
		}
	}
}

func processCode(code string) []byte {
	if strings.Contains(code, "open_") {
		code = strings.TrimPrefix(code, "open_")
		index, _ := strconv.Atoi(code)
		return processOpen(index)
	} else {
		switch code {
		case "Hello From The Client!":
			return []byte("Hello From The Server!")
		case "New":
			return processNew()
		case "save":
			processSave()
			return []byte("Saved")
		case "openPosts":
			return processOpenPosts()
		default:
			return []byte("Error")
		}
	}

}
func processNew() []byte {
	fmt.Println("New")
	return []byte("New")
}
func processOpen(index int) []byte {
	fmt.Println("Open")
	return []byte("post_content: " + postfiles[index].content)
}
func processOpenPosts() []byte {
	fmt.Println("Open Posts")
	messages := ""
	for _, post := range postfiles {
		messages += post.title
		messages += "\n"
	}
	messages = strings.TrimRight(messages, "\n")
	return []byte(messages)
}

func processSave() {
	fmt.Println("Save")
}

// define our WebSocket endpoint
func serveWs(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.Host)

	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Client Connected")
	reader(ws)
}

func setupRoutes() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Simple Server!")
		http.HandleFunc("/ws", serveWs)
	})
}

func savePost(p Post) {
	//	TODO: Save post to file
}

func openPost(file string) {
	// TODO: Open post from file
}

func Deploy() bool {
	//	TODO: Deploy to github
	return true
}
