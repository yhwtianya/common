package utils

import (
	"fmt"
)

// 生成监测点key，由endpoint/metric/tags构成
func PK(endpoint, metric string, tags map[string]string) string {
	if tags == nil || len(tags) == 0 {
		return fmt.Sprintf("%s/%s", endpoint, metric)
	}
	return fmt.Sprintf("%s/%s/%s", endpoint, metric, SortedTags(tags))
}

// 生成监测点key，由endpoint/metric构成
func PK2(endpoint, counter string) string {
	return fmt.Sprintf("%s/%s", endpoint, counter)
}

// 生成监测点UUID,由endpoint/metric/tags/dstype/step构成
func UUID(endpoint, metric string, tags map[string]string, dstype string, step int) string {
	if tags == nil || len(tags) == 0 {
		return fmt.Sprintf("%s/%s/%s/%d", endpoint, metric, dstype, step)
	}
	return fmt.Sprintf("%s/%s/%s/%s/%d", endpoint, metric, SortedTags(tags), dstype, step)
}

// 返回md5(PK)
func Checksum(endpoint string, metric string, tags map[string]string) string {
	pk := PK(endpoint, metric, tags)
	return Md5(pk)
}

// 返回md5(UUID)
func ChecksumOfUUID(endpoint, metric string, tags map[string]string, dstype string, step int64) string {
	return Md5(UUID(endpoint, metric, tags, dstype, int(step)))
}
