package util

import "github.com/Blue-Berrys/Tiktok_e_commerce/app/ai/biz/util/chat_ai"

func FormatData(data string) string {

	text := `
整理一下数据为表格形式展示：
`
	text += data

	result := chat_ai.Generate(text)

	return result.Content
}
