package util

import (
	"context"
	"fmt"
	"github.com/cloudwego/eino-ext/components/model/ark"
	"github.com/cloudwego/eino/components/model"
	"github.com/cloudwego/eino/components/prompt"
	"github.com/cloudwego/eino/schema"
	"io"
	"log"
	"strings"
	"time"
)

func main() {
	ctx := context.Background()

	// 使用模版创建messages
	log.Printf("===create messages===\n")
	messages := createMessagesFromTemplate("你好")
	log.Printf("messages: %+v\n\n", messages)

	// 创建llm
	log.Printf("===create llm===\n")
	cm := createOpenAIChatModel(ctx)
	// cm := createOllamaChatModel(ctx)
	log.Printf("create llm success\n\n")

	log.Printf("===llm generate===\n")
	result := generate(ctx, cm, messages)

	fmt.Println("--------------------------")
	log.Printf("result: %v", result.Content)
	fmt.Println("--------------------------")
}

var (
	CM model.ChatModel
)

const (
	QueryOrderPrompt = `
create table []order[]
(
    id             bigint auto_increment
        primary key,
    created_at     datetime(3)  null,
    updated_at     datetime(3)  null,
    order_id       varchar(256) null,
    user_id        int unsigned null,
    user_currency  longtext     null,
    email          longtext     null,
    street_address longtext     null,
    city           longtext     null,
    state          longtext     null,
    country        longtext     null,
    zip_code       longtext     null,
    order_state    longtext     null,
    constraint idx_order_order_id
        unique (order_id)
)
    engine = InnoDB;
create table order_item
(
    id             bigint auto_increment
        primary key,
    created_at     datetime(3)  null,
    updated_at     datetime(3)  null,
    product_id     int unsigned null,
    order_id_refer varchar(256) null,
    quantity       int          null,
    cost           float        null,
    constraint fk_order_order_items
        foreign key (order_id_refer) references []order[] (order_id)
)
    engine = InnoDB;

create index idx_order_item_order_id_refer
    on order_item (order_id_refer);

create table product
(
    id          bigint auto_increment
        primary key,
    created_at  datetime(3) null,
    updated_at  datetime(3) null,
    name        longtext    null,
    description longtext    null,
    picture     longtext    null,
    price       float       null
)
    engine = InnoDB;
查询订单要求：{要求}
你现在是一名生成sql语句的ai大模型，以上信息为order（订单表），order_item（订单中包含的产品表），product（产品表），请严格上述的表数据，以及查询订单要求，进行生成对应的sql语句。并且只需给出一条sql语句即可，不用给出多余的内容。
`
)

func Init() {
	ctx := context.Background()
	CM = createOpenAIChatModel(ctx)
}

func GetQueryOrderSQL(question string) (sqlText string) {
	ctx := context.Background()

	text := QueryOrderPrompt
	text = strings.ReplaceAll(text, "[]", "`")
	text = strings.ReplaceAll(text, "{要求}", question)
	fmt.Println("role:\n", text)

	messages := createMessagesFromTemplate(text)

	result := generate(ctx, CM, messages)
	sqlText = result.Content

	fmt.Println("订单查询sql语句：", sqlText)
	return ""
}

func GetAutoOrderSQL() {

}

func createMessagesFromTemplate(question string) []*schema.Message {
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

func createOpenAIChatModel(ctx context.Context) model.ChatModel {
	//key := os.Getenv("sk-tY1on5iExVvL8wTpiYKyMr5f3SXqvqd9XEjhIgpA9IQOG6o4")
	//chatModel, err := openai.NewChatModel(ctx, &openai.ChatModelConfig{
	//	Model:   "gpt-4o-mini", // 使用的模型版本
	//	APIKey:  key,           // OpenAI API 密钥
	//	BaseURL: "https://api.chatanywhere.tech/v1",
	//})
	// 创建一个时间持续时间变量
	timeout := 30 * time.Second
	chatModel, err := ark.NewChatModel(ctx, &ark.ChatModelConfig{
		APIKey:  "fe60c743-1bbe-45a2-a745-4212b87a5938",
		Region:  "cn-beijing",
		Model:   "ep-20250215215746-mq6nh",
		Timeout: &timeout,
	})
	if err != nil {
		log.Fatalf("create openai chat model failed, err=%v", err)
	}
	return chatModel
}

func generate(ctx context.Context, llm model.ChatModel, in []*schema.Message) *schema.Message {
	result, err := llm.Generate(ctx, in)
	if err != nil {
		log.Fatalf("llm generate failed: %v", err)
	}
	return result
}

func stream(ctx context.Context, llm model.ChatModel, in []*schema.Message) *schema.StreamReader[*schema.Message] {
	result, err := llm.Stream(ctx, in)
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
