package util

import (
	"github.com/Blue-Berrys/Tiktok_e_commerce/app/ai/biz/util/chat_ai"
	"testing"
)

func TestChatAI(t *testing.T) {
	chat_ai.Init()
	GenQueryOrderSQL("查询 2024 年 1 月 1 日之后创建的订单的订单 ID 和对应的产品名称，以及价格")
}
