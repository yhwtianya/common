package model

import (
	"fmt"
	"math"

	MUtils "github.com/open-falcon/common/utils"
)

// Transfer发送给Graph的数据结构
// DsType 即RRD中的Datasource的类型：GAUGE|COUNTER|DERIVE
type GraphItem struct {
	Endpoint  string            `json:"endpoint"`
	Metric    string            `json:"metric"`
	Tags      map[string]string `json:"tags"`
	Value     float64           `json:"value"`
	Timestamp int64             `json:"timestamp"`
	DsType    string            `json:"dstype"`
	Step      int               `json:"step"`
	Heartbeat int               `json:"heartbeat"`
	Min       string            `json:"min"`
	Max       string            `json:"max"`
}

func (this *GraphItem) String() string {
	return fmt.Sprintf(
		"<Endpoint:%s, Metric:%s, Tags:%v, Value:%v, TS:%d %v DsType:%s, Step:%d, Heartbeat:%d, Min:%s, Max:%s>",
		this.Endpoint,
		this.Metric,
		this.Tags,
		this.Value,
		this.Timestamp,
		MUtils.UnixTsFormat(this.Timestamp),
		this.DsType,
		this.Step,
		this.Heartbeat,
		this.Min,
		this.Max,
	)
}

// GraphItem的Endpoint、Metric、Tags构成Key
func (this *GraphItem) PrimaryKey() string {
	return MUtils.PK(this.Endpoint, this.Metric, this.Tags)
}

// GraphItem的Endpoint、Metric、Tags构成md5sum
func (t *GraphItem) Checksum() string {
	return MUtils.Checksum(t.Endpoint, t.Metric, t.Tags)
}

// GraphItem的Endpoint、Metric、Tags、DsType、Step构成UUID
func (this *GraphItem) UUID() string {
	return MUtils.UUID(this.Endpoint, this.Metric, this.Tags, this.DsType, this.Step)
}

// Rpc GraphQuery的入参
// ConsolFun 是RRD中的概念，比如：MIN|MAX|AVERAGE
type GraphQueryParam struct {
	Start     int64  `json:"start"`
	End       int64  `json:"end"`
	ConsolFun string `json:"consolFuc"`
	Endpoint  string `json:"endpoint"`
	Counter   string `json:"counter"`
}

// Rpc GraphQuery的出参
type GraphQueryResponse struct {
	Endpoint string     `json:"endpoint"`
	Counter  string     `json:"counter"`
	DsType   string     `json:"dstype"`
	Step     int        `json:"step"`
	Values   []*RRDData `json:"Values"` //大写为了兼容已经再用这个api的用户
}

// 页面上已经可以看到DsType和Step了，直接带进查询条件，Graph更易处理
type GraphAccurateQueryParam struct {
	Checksum  string `json:"checksum"`
	Start     int64  `json:"start"`
	End       int64  `json:"end"`
	ConsolFun string `json:"consolFuc"`
	DsType    string `json:"dsType"`
	Step      int    `json:"step"`
}

type GraphAccurateQueryResponse struct {
	Values []*RRDData `json:"values"`
}

type JsonFloat float64

func (v JsonFloat) MarshalJSON() ([]byte, error) {
	f := float64(v)
	if math.IsNaN(f) || math.IsInf(f, 0) {
		return []byte("null"), nil
	} else {
		return []byte(fmt.Sprintf("%f", f)), nil
	}
}

// rrd值结构
type RRDData struct {
	Timestamp int64     `json:"timestamp"`
	Value     JsonFloat `json:"value"`
}

func NewRRDData(ts int64, val float64) *RRDData {
	return &RRDData{Timestamp: ts, Value: JsonFloat(val)}
}

func (this *RRDData) String() string {
	return fmt.Sprintf(
		"<RRDData:Value:%v TS:%d %v>",
		this.Value,
		this.Timestamp,
		MUtils.UnixTsFormat(this.Timestamp),
	)
}

// 请求参数，用于查询指标对应的rrrd属性
type GraphInfoParam struct {
	Endpoint string `json:"endpoint"`
	Counter  string `json:"counter"`
}

// 响应参数，代表指标对应的rrd属性
type GraphInfoResp struct {
	ConsolFun string `json:"consolFun"`
	Step      int    `json:"step"`
	Filename  string `json:"filename"`
}

// 响应参数，代表指标及其rrd属性
type GraphFullyInfo struct {
	Endpoint  string `json:"endpoint"`
	Counter   string `json:"counter"`
	ConsolFun string `json:"consolFun"`
	Step      int    `json:"step"`
	Filename  string `json:"filename"`
	Addr      string `json:"addr"`
}

type GraphLastParam struct {
	Endpoint string `json:"endpoint"`
	Counter  string `json:"counter"`
}

type GraphLastResp struct {
	Endpoint string   `json:"endpoint"`
	Counter  string   `json:"counter"`
	Value    *RRDData `json:"value"`
}
