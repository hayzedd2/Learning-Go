package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"example.com/investmentCalculator/structproject/blog"
)

func main() {
	title, content := getBlogData()
	blogData, err := blog.New(title, content)
	if err != nil {
		fmt.Println(err)
		return
	}
	blogData.BlogDisplay()
	err = blogData.SaveToFile()
	if err != nil {
		fmt.Println("Saving blog failed")
		return
	}
	fmt.Println("Saved blog successfully")


}

func getBlogData() (string, string) {
	blogTitle := getUserInput("Blog Title:")
	blogContent := getUserInput("Blog Content:")
	return blogTitle, blogContent
}
func getUserInput(prompt string) (string) {
	fmt.Printf("%v ", prompt)
	reader := bufio.NewReader(os.Stdin)
	text, err:= reader.ReadString('\n')
	if err!= nil{
		return ""
	}
	text = strings.TrimSuffix(text, "\n")
	text = strings.TrimSuffix(text, "\r")
	return text
}
