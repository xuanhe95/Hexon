package main

import (
	"fmt"
	"strings"
)

type Post struct {
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
