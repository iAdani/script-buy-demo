package scrape

import (
	"context"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/PullRequestInc/go-gpt3"
)

func Classify(query string) string {
	apiKey := os.Getenv("OPENAI_API_KEY")

	classes := []string{"Clothing, Shoes & Accessories", "Home & Patio", "Baby", "Electronics", "School & Office",
		"Toys", "Sports, Fitness & Outdoors", "Entertainment", "Beauty & Personal Care", "Health",
		"Household Essentials", "Pets", "Grocery"}
	editedClasses := "{ '" + strings.Join(classes, "', '") + "' }"

	ctx := context.Background()
	client := gpt3.NewClient(apiKey)

	resp, err := client.ChatCompletion(ctx, gpt3.ChatCompletionRequest{
		Model: gpt3.GPT3Dot5Turbo,
		Messages: []gpt3.ChatCompletionRequestMessage{
			{
				Role:    "system",
				Content: "categories : " + editedClasses,
			},
			{
				Role:    "user",
				Content: "respond only with the category for " + query + " or 'none' if there is no category",
			},
		},
		N: 1,
	})
	if err != nil {
		log.Fatalln(err)
	}
	result := resp.Choices[0].Message.Content
	for _, str := range classes {
		if str == result {
			fmt.Println("ChatGPT Classifier" + str)
			return result
		}
	}
	return "none"
}
