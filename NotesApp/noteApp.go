package main

import (
	"bufio"
	"example.com/note-app/note"
	"example.com/note-app/todo"
	"fmt"
	"os"
	"strings"
)

type Saver interface {
	Save() error
}
type Displayer interface {
	Display()
}
type Outputable interface {
	Saver
	Displayer
}

func add[T int | float64 | string](a, b T) T {
	return a + b

}
func printSomething(value any) {
	switch value.(type) {
	case int:
		fmt.Print("int: %v", value)
	case float64:
		fmt.Print("float %v", value)
	default:
		fmt.Print(value)
	}

	typedVal, ok := value.(int)
	if ok {
		fmt.Print(add(typedVal, 1))
	}
}
func saveData(data Saver) error {
	err := data.Save()
	if err != nil {
		fmt.Println(err)
		fmt.Println("saving note failed")
		return err
	}
	fmt.Println("saving success")
	return nil
}
func outputData(data Outputable) error {
	data.Display()
	return saveData(data)
}

func main() {
	note, err := getNoteInput()
	if err != nil {
		fmt.Println(err)
		return
	}

	err = outputData(note)
	if err != nil {
		return
	}

	newTodo, err := getTodoData()
	if err != nil {
		return
	}
	err = outputData(newTodo)

	if err != nil {
		return
	}
	printSomething("test")
}
func getTodoData() (todo.Todo, error) {
	return todo.New(getUserInput("todoText: "))
}
func getNoteInput() (note.Note, error) {
	title := getUserInput("note title:")

	content := getUserInput("note content")

	return note.New(title, content)
}
func getUserInput(prompt string) string {

	fmt.Print(prompt)
	reader := bufio.NewReader(os.Stdin)
	text, err := reader.ReadString('\n')
	if err != nil {
		return ""
	}
	text = strings.TrimSuffix(text, "\n")
	text = strings.TrimSuffix(text, "\r")

	return text
}
