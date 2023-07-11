package main

import "fmt"

func PrintAllPostsTitle() {
	for _, post := range postfiles {
		fmt.Println(post.title)
	}
}

func PrintAllPostsDate() {
	for _, post := range postfiles {
		fmt.Println(post.date)
	}
}

func PrintAllPostsContent() {
	for _, post := range postfiles {
		fmt.Println(post.title)
		fmt.Println(post.date)
		fmt.Println(post.content)
		fmt.Println("=====================================")
	}
}
