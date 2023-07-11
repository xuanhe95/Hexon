package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func ParseHexoDirectory(dir string) {
	if dir[len(dir)-1] == '/' {
		dir += "source/_posts/"
	} else {
		dir += "/source/_posts/"
	}
	LoadPosts(dir)
}
func LoadPost(filename string) Post { // load one post
	// TODO: Load posts from file
	var content []string
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("Open File Failed: ", err)
		//return nil
	}
	defer file.Close()

	var scanner bufio.Scanner = *bufio.NewScanner(file)

	var meta bool
	var metaCount int
	var p Post
	for scanner.Scan() {
		var line string = scanner.Text()
		//fmt.Println(line)
		if line == "---" && metaCount < 2 {
			meta = !meta
			metaCount++
			continue
		}
		if meta {
			if strings.HasPrefix(line, "title:") {
				strings.Replace(line, "title:", "", 1)
				strings.TrimSpace(line)
				p.title = line
			} else if strings.HasPrefix(line, "date:") {
				strings.Replace(line, "date:", "", 1)
				strings.TrimSpace(line)
				p.date = line
			} else if strings.HasPrefix(line, "tags:") {
				strings.Replace(line, "tags:", "", 1)
				strings.TrimSpace(line)
				tags := strings.Split(line, ",")
				for _, tag := range tags {
					tag = strings.TrimSpace(tag)
					p.tags = append(p.tags, tag)
				}
			} else if strings.HasPrefix(line, "categories:") {
				strings.Replace(line, "categories:", "", 1)
				strings.TrimSpace(line)
				categories := strings.Split(line, ",")
				for _, category := range categories {
					category = strings.TrimSpace(category)
					p.categories = append(p.categories, category)
				}
			} else if strings.HasPrefix(line, "toc:") {
				if strings.Contains(line, "true") {
					p.toc = true
				} else {
					p.toc = false
				}
			} else if strings.HasPrefix(line, "cover:") {
				strings.Replace(line, "cover:", "", 1)
				strings.Replace(line, ">-", "", 1)
				strings.TrimSpace(line)
				p.cover = line
			} else if strings.HasPrefix(line, "thumbnail:") {
				strings.Replace(line, "thumbnail:", "", 1)
				strings.Replace(line, ">-", "", 1)
				strings.TrimSpace(line)
				p.thumbnail = line
			}

		} else {
			content = append(content, line)
		}
	}
	p.content = strings.Join(content, "\n")
	if err := scanner.Err(); err != nil {
		fmt.Println("File Load Error: ", err)
	} else {
		fmt.Println("File Load Success!")
	}
	return p
}

var postfiles []Post

func LoadPosts(dir string) ([]string, error) {
	var posts []string
	files, err := os.ReadDir(dir)
	if err != nil {
		return posts, err
	}

	for _, file := range files {
		filePath := dir + file.Name()
		if !file.IsDir() && strings.HasSuffix(file.Name(), ".md") {
			fmt.Println("Opening " + file.Name() + " ...")
			posts = append(posts, filePath)
			post := LoadPost(filePath)
			postfiles = append(postfiles, post)
			fmt.Println("Open Success!")
		}
	}
	return posts, nil
}
