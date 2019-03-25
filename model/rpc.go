package model

import (
	"fmt"
)

// rpc ping响应
// code == 0 => success
// code == 1 => bad request
type SimpleRpcResponse struct {
	Code int `json:"code"`
}

func (this *SimpleRpcResponse) String() string {
	return fmt.Sprintf("<Code: %d>", this.Code)
}

// 空请求，用于rpc ping
type NullRpcRequest struct {
}
