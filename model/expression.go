package model

import (
	"fmt"

	"github.com/open-falcon/common/utils"
)

// Expression是基于Metric的
type Expression struct {
	Id         int               `json:"id"`
	Metric     string            `json:"metric"`     //endpoint放到了tag中保存
	Tags       map[string]string `json:"tags"`       //保存了tags,endpoint也作为tag进行保存
	Func       string            `json:"func"`       // e.g. max(#3) all(#3)
	Operator   string            `json:"operator"`   // e.g. < !=
	RightValue float64           `json:"rightValue"` // critical value
	MaxStep    int               `json:"maxStep"`
	Priority   int               `json:"priority"`
	Note       string            `json:"note"`
	ActionId   int               `json:"actionId"`
}

func (this *Expression) String() string {
	return fmt.Sprintf(
		"<Id:%d, Metric:%s, Tags:%v, %s%s%s MaxStep:%d, P%d %s ActionId:%d>",
		this.Id,
		this.Metric,
		this.Tags,
		this.Func,
		this.Operator,
		utils.ReadableFloat(this.RightValue),
		this.MaxStep,
		this.Priority,
		this.Note,
		this.ActionId,
	)
}

// Hbs响应Judge同步Expression的数据结构
type ExpressionResponse struct {
	Expressions []*Expression `json:"expressions"`
}
