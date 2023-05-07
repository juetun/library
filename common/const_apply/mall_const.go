package const_apply

import "github.com/juetun/base-wrapper/lib/base"

const (
	FlagTesterNo  uint8 = iota + 1 // 不为测试数据
	FlagTesterYes                  // 为测试数据
)

var (
	SliceFlagTester = base.ModelItemOptions{
		{
			Value: FlagTesterNo,
			Label: "否",
		},
		{
			Value: FlagTesterYes,
			Label: "是",
		},
	}
)
