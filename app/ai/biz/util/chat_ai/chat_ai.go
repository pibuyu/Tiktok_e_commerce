package chat_ai

import (
	"context"
	"fmt"
	"github.com/cloudwego/eino-ext/components/model/ark"
	"github.com/cloudwego/eino/components/model"
	"github.com/cloudwego/eino/components/prompt"
	"github.com/cloudwego/eino/schema"
	"io"
	"log"
	"time"
)

const (
	API_KEY = "fe60c743-1bbe-45a2-a745-4212b87a5938"
	MODEL   = "ep-20250216115158-fwqw7"
)

func main() {

	// 使用模版创建messages
	log.Printf("===create messages===\n")
	messages := CreateMessagesFromTemplate("你好")
	log.Printf("messages: %+v\n\n", messages)

	// 创建llm
	//log.Printf("===create llm===\n")
	//cm := CreateOpenAIChatModel(ctx)
	//// cm := createOllamaChatModel(ctx)
	//log.Printf("create llm success\n\n")

	log.Printf("===llm generate===\n")
	result := Generate("你好")

	fmt.Println("--------------------------")
	log.Printf("result: %v", result.Content)
	fmt.Println("--------------------------")
}

var (
	CM  model.ChatModel
	ctx context.Context
)

func Init() {
	ctx = context.Background()
	CM = CreateOpenAIChatModel(ctx)
}

func CreateMessagesFromTemplate(question string) []*schema.Message {
	template := prompt.FromMessages(schema.FString,

		// 用户消息模板
		schema.UserMessage("{question}"),
	)

	// 使用模板生成消息
	messages, err := template.Format(context.Background(), map[string]any{
		"question": question,
	})
	if err != nil {
		log.Fatalf("format template failed: %v\n", err)
	}
	return messages
}

func CreateOpenAIChatModel(ctx context.Context) model.ChatModel {
	//key := os.Getenv("sk-tY1on5iExVvL8wTpiYKyMr5f3SXqvqd9XEjhIgpA9IQOG6o4")
	//chatModel, err := openai.NewChatModel(ctx, &openai.ChatModelConfig{
	//	Model:   "gpt-4o-mini", // 使用的模型版本
	//	APIKey:  key,           // OpenAI API 密钥
	//	BaseURL: "https://api.chatanywhere.tech/v1",
	//})
	// 创建一个时间持续时间变量
	timeout := 30 * time.Second
	chatModel, err := ark.NewChatModel(ctx, &ark.ChatModelConfig{
		APIKey:  API_KEY,
		Region:  "cn-beijing",
		Model:   MODEL,
		Timeout: &timeout,
	})
	if err != nil {
		log.Fatalf("create openai chat model failed, err=%v", err)
	}
	return chatModel
}

func Generate(text string) *schema.Message {
	in := CreateMessagesFromTemplate(text)
	result, err := CM.Generate(ctx, in)
	if err != nil {
		log.Fatalf("llm generate failed: %v", err)
	}
	return result
}

func stream(text string) *schema.StreamReader[*schema.Message] {

	in := CreateMessagesFromTemplate(text)
	result, err := CM.Stream(ctx, in)
	if err != nil {
		log.Fatalf("llm generate failed: %v", err)
	}
	return result
}

func reportStream(sr *schema.StreamReader[*schema.Message]) {
	defer sr.Close()

	i := 0
	for {
		message, err := sr.Recv()
		if err == io.EOF {
			return
		}
		if err != nil {
			log.Fatalf("recv failed: %v", err)
		}
		log.Printf("message[%d]: %+v\n", i, message)
		i++
	}
}
