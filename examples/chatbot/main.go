package main

import (
	"bufio"
	"context"
	"fmt"
	"github.com/yshujie/openai"
	"os"
)

func main() {

	client := openai.NewClient(os.Getenv("OPENAI_API_KEY"))

	fmt.Println("Conversation")
	fmt.Println("---------------------")
	fmt.Print("> ")
	s := bufio.NewScanner(os.Stdin)
	for s.Scan() {

		respContent, err := client.CreateChatCompletion(context.Background(), s.Text())
		if err != nil {
			fmt.Printf("ChatCompletion error: %v\n", err)
			continue
		}
		fmt.Printf("%s\n\n", respContent)
		fmt.Print("> ")
	}
}
