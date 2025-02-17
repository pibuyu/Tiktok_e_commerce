package service

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/Blue-Berrys/Tiktok_e_commerce/app/ai/biz/dal/mysql"
	"github.com/Blue-Berrys/Tiktok_e_commerce/app/ai/biz/util"
	ai "github.com/Blue-Berrys/Tiktok_e_commerce/rpc_gen/kitex_gen/ai"
	"log"
)

type QueryOrderService struct {
	ctx context.Context
} // NewQueryOrderService new QueryOrderService
func NewQueryOrderService(ctx context.Context) *QueryOrderService {
	return &QueryOrderService{ctx: ctx}
}

// Run create note info
func (s *QueryOrderService) Run(req *ai.OrderQueryRequest) (resp *ai.OrderQueryResponse, err error) {
	// Finish your business logic.
	fmt.Println("请求的内容为：", req)

	queryOrderSQL := util.GenQueryOrderSQL(req.Message, req.UserId)

	rows, err := mysql.DB.Raw(queryOrderSQL).Rows()
	defer rows.Close()

	data := get(rows)

	// 使用 json.Marshal 将数据转换为 JSON 字符串
	jsonData, err := json.Marshal(data)
	if err != nil {
		log.Fatalf("Error marshaling to JSON: %v", err)
	}

	formatString := util.FormatData(string(jsonData))
	resp = &ai.OrderQueryResponse{Data: formatString}

	return resp, nil
}

func get(rows *sql.Rows) []map[string]interface{} {
	// 创建用于存储每行数据的切片
	var results []map[string]interface{}

	// 获取列名
	columns, err := rows.Columns()
	if err != nil {
		fmt.Println("获取列名失败:", err)
		return nil
	}

	// 创建用于扫描的切片
	values := make([]interface{}, len(columns))
	valuePtrs := make([]interface{}, len(columns))
	for i := range values {
		valuePtrs[i] = &values[i]
	}

	// 遍历查询结果
	for rows.Next() {
		// 扫描行数据
		err := rows.Scan(valuePtrs...)
		if err != nil {
			fmt.Println("扫描行数据失败:", err)
			return nil
		}

		// 创建一个map来存储当前行的数据
		row := make(map[string]interface{})
		for i, col := range columns {
			var v interface{}
			val := values[i]
			b, ok := val.([]byte)
			if ok {
				v = string(b)
			} else {
				v = val
			}
			row[col] = v
		}

		// 将当前行的数据添加到结果切片中
		results = append(results, row)
	}

	// 检查是否有错误
	if err = rows.Err(); err != nil {
		fmt.Println("遍历结果集时出错:", err)
		return nil
	}

	// 打印查询结果
	for _, row := range results {
		fmt.Println(row)
	}
	return results
}
