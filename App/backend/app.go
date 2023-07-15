package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"

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
	} else if strings.Contains(code, "delete_") {
		code = strings.TrimPrefix(code, "delete_")
		index, _ := strconv.Atoi(code)
		return processDelete(index)
	} else if strings.Contains(code, "save_content_to_") {
		code = strings.TrimPrefix(code, "save_content_to_")
		simicolonIndex := strings.Index(code, ":")
		index, _ := strconv.Atoi(code[:simicolonIndex])
		if index >= len(postfiles) || index < 0 {
			return []byte("Error: Index out of range")
		}

		postfiles[index].revised = true
		postfiles[index].content = code[simicolonIndex+1:]
		fmt.Println(postfiles[index].content)
		return []byte("Saved")
	} else {
		switch code {
		case "Hello From The Client!":
			return []byte("Hello From The Server!")
		case "new_post":
			return processNew()
		case "initilize_all_posts":
			fmt.Println(len(postfiles))
			return processOpenPosts()
		default:
			return []byte("Error")
		}
	}
}

func processDelete(index int) []byte {
	fmt.Println("Delete")
	if index >= len(postfiles) {
		return []byte("Error: Index out of range")
	}
	if len(postfiles) == 1 {
		postfiles = postfiles[:0]
	} else {
		postfiles = append(postfiles[:index], postfiles[index+1:]...)
	}

	return processOpenPosts()

}
func processNew() []byte {
	fmt.Println("New")
	addNewPost(time.Now().Format("2006-01-02 15:04:05"))
	return processOpenPosts()
}

func addNewPost(date string) {
	newPost := NewPost(true, "New Post", date, "")
	// newPost.revised = true
	// newPost.title = "title: New Post"
	// newPost.date = date
	// newPost.content = ""
	postfiles = append([]Post{*newPost}, postfiles...) //这里还可以做的更好，但是需要调整前端保存文件地址的方式

}

func NewPost(revised bool, title string, date string, content string) *Post {
	return &Post{
		revised: revised,
		title:   title,
		date:    date,
		content: content,
	}
}
func processOpen(index int) []byte {
	fmt.Println("Open")
	fmt.Println(len(postfiles))
	fmt.Println(index)
	if index >= len(postfiles) || index < 0 {
		return []byte("Error: Index out of range")
	}
	return []byte("post_content_" + postfiles[index].content)
}
func processOpenPosts() []byte {
	fmt.Println("Open Posts")
	messages := ""
	for _, post := range postfiles {
		messages += "post_title_"
		messages += post.title
		messages += "\n"
	}
	// for i := len(postfiles) - 1; i >= 0; i-- {
	// 	messages += postfiles[i].title
	// 	messages += "\n"
	// }
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
