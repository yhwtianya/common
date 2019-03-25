package model

import (
	"fmt"

	MUtils "github.com/open-falcon/common/utils"
)

// Metric值数据结构，rrd数据结构
type MetricValue struct {
	Endpoint  string      `json:"endpoint"`
	Metric    string      `json:"metric"`
	Value     interface{} `json:"value"`
	Step      int64       `json:"step"`
	Type      string      `json:"counterType"` //GAUGE或COUNTER或DERIVE
	Tags      string      `json:"tags"`
	Timestamp int64       `json:"timestamp"`
}

func (this *MetricValue) String() string {
	return fmt.Sprintf(
		"<Endpoint:%s, Metric:%s, Type:%s, Tags:%s, Step:%d, Time:%d, Value:%v>",
		this.Endpoint,
		this.Metric,
		this.Type,
		this.Tags,
		this.Step,
		this.Timestamp,
		this.Value,
	)
}

// Metric值数据结构，rrd数据结构
// Same As `MetricValue`
type JsonMetaData struct {
	Metric      string      `json:"metric"`
	Endpoint    string      `json:"endpoint"`
	Timestamp   int64       `json:"timestamp"`
	Step        int64       `json:"step"`
	Value       interface{} `json:"value"`
	CounterType string      `json:"counterType"` //GAUGE或COUNTER或DERIVE
	Tags        string      `json:"tags"`
}

func (t *JsonMetaData) String() string {
	return fmt.Sprintf("<JsonMetaData Endpoint:%s, Metric:%s, Tags:%s, DsType:%s, Step:%d, Value:%v, Timestamp:%d>",
		t.Endpoint, t.Metric, t.Tags, t.CounterType, t.Step, t.Value, t.Timestamp)
}

// Metric值数据结构，rrd数据结构
type MetaData struct {
	Metric      string            `json:"metric"`
	Endpoint    string            `json:"endpoint"`
	Timestamp   int64             `json:"timestamp"`
	Step        int64             `json:"step"`
	Value       float64           `json:"value"`
	CounterType string            `json:"counterType"` //GAUGE或COUNTER或DERIVE
	Tags        map[string]string `json:"tags"`
}

func (t *MetaData) String() string {
	return fmt.Sprintf("<MetaData Endpoint:%s, Metric:%s, Timestamp:%d, Step:%d, Value:%f, Tags:%v>",
		t.Endpoint, t.Metric, t.Timestamp, t.Step, t.Value, t.Tags)
}

// 返回Metric的Key
func (t *MetaData) PK() string {
	return MUtils.PK(t.Endpoint, t.Metric, t.Tags)
}

// GAUGE
// GAUGE 意为计量，其值简单地原样存储。用于可增减的值，如：温度或费用支出。

// DERIVE
// DERIVE 意为导数，关注值的变动，即导数。 这样的数据源通常为可计数的事件，如：邮件客户端启动后收到的邮件数。 相对于收件箱里的邮件总数，上次查看邮箱后新收到的邮件数量更值得关注。 该值可以根据以下公式转化为速率：
// rate = (value_new - value_old) / (time_new - time_old)
// 如果 value_new 小于 value_old，得出的 rate 为负。 如果设置最小值（min）为 0 ，这个数据点将被丢弃。 对于很少溢出（即达到最大值后回绕）的计数器推荐使用 DERIVE 数据源并设置最小值（min）为 0。 DERIVE 类型从 4.8 版本开始提供。

// COUNTER
// COUNTER 意为计数器，”正常“情况下与 DERIVE 一样。 细微的差异在于当 value_new 小于 value_old 时，COUNTER 数据源类型假设计数器已经“回绕”，计算速率的公式：
