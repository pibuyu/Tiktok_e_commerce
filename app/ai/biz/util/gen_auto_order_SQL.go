package util

import (
	"fmt"
	"github.com/Blue-Berrys/Tiktok_e_commerce/app/ai/biz/dal/model"
	"github.com/Blue-Berrys/Tiktok_e_commerce/app/ai/biz/util/chat_ai"
	"github.com/google/uuid"
	"regexp"
	"strconv"
	"strings"
)

const (
	AutoOrderPrompt = `
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
新增订单要求：{要求}
其他要求：
1、只需添加order和order_item表数据，product中的id需要根据给出的要求进行查询得到。
2、给出对应的一段sql语句即可，不用给出多余非sql语句内容
3、已知值：order_id = {order_id}, user_id = {user_id}, email = "{email}"
4、只包含insert语句即可。
你现在是一名生成sql语句的ai大模型，以上信息为order（订单表），order_item（订单中包含的产品表），product（产品表），请严格上述的表数据，以及要求，进行生成对应的sql语句，并按照一下格式给出：
{sql}
若完全没有给出收货地址，则不生成sql语句，提示用户请给出收货地址再为你自动下单，若给出一部分收货地址则自行分析补充即可。
`
)

// GenAutoOrderSQL 生成自动下单sql语句，
// state（0：失败，1：生成所需信息不全（还需收获地址），2：成功）
func GenAutoOrderSQL(question string, user *model.User) (s []string, state int) {

	text := strings.ReplaceAll(AutoOrderPrompt, "[]", "`")
	text = strings.ReplaceAll(text, "{要求}", question)
	orderId, _ := uuid.NewRandom()
	text = strings.ReplaceAll(text, "{order_id}", orderId.String())
	text = strings.ReplaceAll(text, "{user_id}", strconv.FormatInt(int64(user.Id), 10))
	text = strings.ReplaceAll(text, "{email}", user.Email)
	text = strings.ReplaceAll(text, "{sql}", "```sql\n XXXXXX \n```")

	fmt.Println("role:\n", text)

	result := chat_ai.Generate(text)
	sqlText := result.Content

	fmt.Println("订单查询sql语句：\n", sqlText)

	return SplitSQLStatements(sqlText)
}

// SplitSQLStatements 将多条 SQL 语句分割成字符串数组
func SplitSQLStatements(sql string) (s []string, state int) {
	// 去除注释和空行
	cleaned := extractSQLCode(sql)

	var result []string

	if cleaned == "" {
		result = append(result, sql)
		return result, 1
	}

	// 按分号分割
	statements := strings.Split(cleaned, ";")

	// 去除每个语句前后的空白字符，并过滤掉空字符串
	for _, stmt := range statements {
		trimmedStmt := strings.TrimSpace(stmt)
		if trimmedStmt != "" {
			result = append(result, trimmedStmt)
		}
	}

	return result, 2
}

// extractSQLCode 匹配以 "```sql" 开头和 "```" 结尾的内容
func extractSQLCode(input string) string {
	// 定义正则表达式
	re := regexp.MustCompile("(?s)```sql(.*?)```")
	matches := re.FindStringSubmatch(input)
	if len(matches) > 1 {
		sqlContent := matches[1]
		fmt.Println("提取的 SQL 内容：")
		fmt.Println(sqlContent)

		return sqlContent
	} else {
		fmt.Println("未找到匹配的 SQL 内容")
		return ""
	}
}
