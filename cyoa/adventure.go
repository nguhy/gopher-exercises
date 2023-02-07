package cyoa

import (
	"encoding/json"
	"io"
)

type Story map[string]Chapter

func JSONStory(r io.Reader) (Story, error) {

	dec := json.NewDecoder(r)

	var story Story

	if err := dec.Decode(&story); err != nil {
		return nil, err
	}

	return story, nil
}

type Chapter struct {
	Title      string   `json:"title"`
	Paragraphs []string `json:"paragraphs"`
	Options    []Option `json:"options"`
}

type Option struct {
	Text    string `json:"text"`
	Chapter string `json:"chapter"`
}
