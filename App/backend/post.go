package main

import (
	"fmt"
	"strings"
)

type Post struct {
	revised    bool
	title      string
	date       string
	tags       []string
	categories []string
	toc        bool
	cover      string
	thumbnail  string
	content    string
}

func (p *Post) setTitle(title string) { // set title
	p.title = title
}

func (p Post) getTitle() string {
	return p.title
}

func (p Post) String() string {
	return fmt.Sprintf("Title: %s", p.title)
}

func (p *Post) setFileTitle() {
	fileTitle := strings.Replace(p.title, " ", "-", -1)
	fileTitle = strings.Replace(fileTitle, ".", "", -1)
	p.title = fileTitle
}

type ByDate []Post

func (p ByDate) Len() int {
	return len(p)
}
func (p ByDate) Less(i, j int) bool {
	return p[i].date < p[j].date
}
func (p ByDate) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}
