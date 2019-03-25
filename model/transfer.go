package model

import (
	"fmt"
)

// 接收上报数据后的响应结构
type TransferResponse struct {
	Message string
	Total   int
	Invalid int   //无效计数
	Latency int64 //耗时
}

func (this *TransferResponse) String() string {
	return fmt.Sprintf(
		"<Total=%v, Invalid:%v, Latency=%vms, Message:%s>",
		this.Total,
		this.Invalid,
		this.Latency,
		this.Message,
	)
}
