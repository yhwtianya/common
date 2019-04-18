package model

import (
	"fmt"
)

// 保存主机id和hostname
type Host struct {
	Id   int
	Name string
}

func (this *Host) String() string {
	return fmt.Sprintf(
		"<id:%s,name:%s>",
		this.Id,
		this.Name,
	)
}
