package model

import (
	"fmt"

	"github.com/open-falcon/common/utils"
)

type Strategy struct {
	Id         int               `json:"id"`
	Metric     string            `json:"metric"`
	Tags       map[string]string `json:"tags"`
	Func       string            `json:"func"`       // e.g. max(#3) all(#3)
	Operator   string            `json:"operator"`   // e.g. < !=
	RightValue float64           `json:"rightValue"` // critical value
	MaxStep    int               `json:"maxStep"`
	Priority   int               `json:"priority"` // 告警级别
	Note       string            `json:"note"`
	Tpl        *Template         `json:"tpl"`
}

func (this *Strategy) String() string {
	return fmt.Sprintf(
		"<Id:%d, Metric:%s, Tags:%v, %s%s%s MaxStep:%d, P%d, %s, %v>",
		this.Id,
		this.Metric,
		this.Tags,
		this.Func,
		this.Operator,
		utils.ReadableFloat(this.RightValue),
		this.MaxStep,
		this.Priority,
		this.Note,
		this.Tpl,
	)
}

// Strategy是基于主机的
type HostStrategy struct {
	Hostname   string     `json:"hostname"`
	Strategies []Strategy `json:"strategies"`
}

// Hbs响应Judge同步Strategies的数据结构
type StrategiesResponse struct {
	HostStrategies []*HostStrategy `json:"hostStrategies"`
}
