package todo

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"strings"
)

type Todo struct {
	Text string `json:"text"`
}

func New(text string) (Todo, error) {
	if text == "" {
		return Todo{}, errors.New("invalid input")
	}
	return Todo{
		Text: text,
	}, nil
}
func (todo Todo) Display() {
	fmt.Printf("your note todo: %v", todo.Text)
}

func (todo Todo) Save() error {
	filename := strings.ReplaceAll("todoFile", " ", "_")
	filename = strings.ToLower(filename) + ".json"
	json, err := json.Marshal(todo)
	if err != nil {
		return err
	}
	err2 := os.WriteFile(filename, json, 0644)
	if err2 != nil {
		return err2
	}
	return nil
}
