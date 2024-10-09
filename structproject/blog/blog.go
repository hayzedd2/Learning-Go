package blog

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"strings"
	"time"
)

type Blog struct {
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"createdAt"`
}

func (b Blog) Display() {
	fmt.Printf("Your blog is titled %v and the content is:\n\n%v\n\n", b.Title, b.Content)

}

func (b Blog) SaveToFile() error {
	fileName := strings.ReplaceAll(b.Title, " ", "_")
	fileName = strings.ToLower(fileName) + ".json"
	json, err := json.Marshal(b)
	if err != nil {
		return err
	}
	return os.WriteFile(fileName, json, 0644)
}
func New(title, content string) (Blog, error) {
	if title == "" || content == "" {
		return Blog{}, errors.New("must be greater than one character")
	}
	return Blog{
		Title:   title,
		Content: content,
	}, nil

}
