package model

import (
	"encoding/json"
	"io"
)

type Story map[string]Chapter

type Chapter struct {
	Title      string   `json:"title"`
	Paragraphs []string `json:"story"`
	Options    []Option `json:"options"`
}

type Option struct {
	ChapterReference string `json:"arc"`
	Text             string `json:"text"`
}

func JsonStory(r io.Reader) (story Story, err error) {
	decoder := json.NewDecoder(r)
	err = decoder.Decode(&story)
	return
}
