package model

import (
	"fmt"

	"github.com/open-falcon/common/utils"
)

// Transfer发送给Judge的数据结构
type JudgeItem struct {
	Endpoint  string            `json:"endpoint"`
	Metric    string            `json:"metric"`
	Value     float64           `json:"value"`
	Timestamp int64             `json:"timestamp"`
	JudgeType string            `json:"judgeType"` //GAUGE或COUNTER或DERIVE
	Tags      map[string]string `json:"tags"`
}

func (this *JudgeItem) String() string {
	return fmt.Sprintf("<Endpoint:%s, Metric:%s, Value:%f, Timestamp:%d, JudgeType:%s Tags:%v>",
		this.Endpoint,
		this.Metric,
		this.Value,
		this.Timestamp,
		this.JudgeType,
		this.Tags)
}

func (this *JudgeItem) PrimaryKey() string {
	return utils.Md5(utils.PK(this.Endpoint, this.Metric, this.Tags))
}

// judge保存历史数据的结构
type HistoryData struct {
	Timestamp int64   `json:"timestamp"`
	Value     float64 `json:"value"`
}
