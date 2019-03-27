package utils

import (
	"sort"
)

// TODO 以下的部分, 考虑放到公共组件库
// 返回排序后的key
func KeysOfMap(m map[string]string) []string {
	keys := make(sort.StringSlice, len(m))
	i := 0
	for key, _ := range m {
		keys[i] = key
		i++
	}

	keys.Sort()
	return []string(keys)
}
