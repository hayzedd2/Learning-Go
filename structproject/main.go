package main

import (
	"bufio"
	"example.com/investmentCalculator/structproject/blog"
	"example.com/investmentCalculator/structproject/todo"
	"fmt"
	"os"
	"strings"
)

type saver interface {
	SaveToFile() error
}

type saveAndDisplay interface {
	saver
	Display()
}

func main() {
	title, content := getBlogData()
	todoText := getUserInput("Create a new todo: ")
	todo, err := todo.New(todoText)
	if err != nil {
		fmt.Println(err)
		return
	}
	blogData, err := blog.New(title, content)
	if err != nil {
		return
	}
	err = displayAndSaveData(blogData)
	if err != nil {
		return
	}
	displayAndSaveData(todo)
}

func displayAndSaveData(data saveAndDisplay) error {
	data.Display()
	return data.SaveToFile()
}
func getBlogData() (string, string) {
	blogTitle := getUserInput("Blog Title:")
	blogContent := getUserInput("Blog Content:")
	return blogTitle, blogContent
}
// func saveData(data saver) error {
// 	err := data.SaveToFile()
// 	if err != nil {
// 		fmt.Println("Saving blog failed")
// 		return err
// 	}
// 	fmt.Println("Saved blog successfully")
// 	return nil
// }

func getUserInput(prompt string) string {
	fmt.Printf("%v ", prompt)
	reader := bufio.NewReader(os.Stdin)
	text, err := reader.ReadString('\n')
	if err != nil {
		return ""
	}
	text = strings.TrimSuffix(text, "\n")
	text = strings.TrimSuffix(text, "\r")
	return text
}

// func add[T int | float64 | string](a, b T) T {
// 	return a + b
// }
