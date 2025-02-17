package util

import (
	"fmt"
	"github.com/Blue-Berrys/Tiktok_e_commerce/app/ai/biz/util/chat_ai"
	"strconv"
	"strings"
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
其他要求：给定user_id = {user_id}
你现在是一名生成sql语句的ai大模型，以上信息为order（订单表），order_item（订单中包含的产品表），product（产品表），请严格上述的表数据，以及查询订单要求，进行生成对应的sql语句。并且只需给出一条sql语句即可，不用给出多余的内容。
`
)

func GenQueryOrderSQL(question string, userId int32) (sqlText string) {
	text := QueryOrderPrompt
	text = strings.ReplaceAll(text, "[]", "`")
	text = strings.ReplaceAll(text, "{要求}", question)
	text = strings.ReplaceAll(text, "{user_id}", strconv.FormatInt(int64(userId), 10))
	fmt.Println("role:\n", text)

	result := chat_ai.Generate(text)
	sqlText = result.Content

	fmt.Println("订单查询sql语句：\n", sqlText)
	return sqlText
}
