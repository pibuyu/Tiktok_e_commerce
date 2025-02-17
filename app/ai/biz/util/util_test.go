package util

import (
	"fmt"
	"github.com/Blue-Berrys/Tiktok_e_commerce/app/ai/biz/dal/model"
	"github.com/Blue-Berrys/Tiktok_e_commerce/app/ai/biz/util/chat_ai"
	"strconv"
	"strings"
	"testing"
)

func TestChatAI(t *testing.T) {
	chat_ai.Init()
	GenQueryOrderSQL("查询 2024 年 1 月 1 日之后创建的订单的订单 ID 和对应的产品名称，以及价格", 1)
}

func TestGenAutoOrderSQL(t *testing.T) {
	chat_ai.Init()
	sqlStrings, _ := GenAutoOrderSQL("我现在需要购买vivo x200 手机和Notebook，请为我自动下单, 邮政编码为'200000'，收货地址为中国上海市上海和平路123号", &model.User{Id: 1, Email: "3531095171@qq.com"})

	t.Logf("-------------------------------------------------")
	for _, sqlString := range sqlStrings {
		fmt.Println(sqlString)
	}
}

func TestA(t *testing.T) {
	text := `

[]sql
	-- 假设已经知道vivo x200手机的product_id为1001，Notebook的product_id为1002

	-- 插入订单记录
	INSERT INTO [1]order[1] (created_at, updated_at, order_id, user_id, user_currency, email, street_address, city, state, country, zip_code, order_state)
	VALUES (NOW(3), NOW(3), 'order_202311051015', 12345, 'CNY', 'example@example.com', '123 Main St', 'Anytown', 'State', 'China', '12345', 'processing');

	-- 获取插入订单的order_id
	SET @new_order_id = 'order_202311051015';

	-- 插入订单中vivo x200手机的产品记录
	INSERT INTO order_item (created_at, updated_at, product_id, order_id_refer, quantity, cost)
	VALUES (NOW(3), NOW(3), 1001, @new_order_id, 1, 3999.99);

	-- 插入订单中Notebook的产品记录
	INSERT INTO order_item (created_at, updated_at, product_id, order_id_refer, quantity, cost)
	VALUES (NOW(3), NOW(3), 1002, @new_order_id, 1, 8999.99);
	[]
`
	text = strings.ReplaceAll(text, "[1]", "`")
	text = strings.ReplaceAll(text, "[]", "```")

	fmt.Println(text)
	//
	//statements := SplitSQLStatements(text)
	//
	//for i, stmt := range statements {
	//	fmt.Printf("Statement %d:\n%s\n\n", i+1, stmt)
	//}
	//fmt.Println("已完成")
}

func TestB(t *testing.T) {
	var x int32
	x = 1
	fmt.Println(x)
	fmt.Println(strconv.FormatInt(int64(x), 10))
}
