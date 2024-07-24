package app_param

import "github.com/juetun/base-wrapper/lib/base"

const (
	ForceUpgradeDefault = iota //不升级 忽略
	ForceUpgradeDialog         //弹窗提示升级
	ForceUpgradeYes            //强制用户升级app
)

var (
	SliceForceUpgrade = base.ModelItemOptions{
		{
			Value: ForceUpgradeDefault,
			Label: "不升级",
		},
		{
			Value: ForceUpgradeDialog,
			Label: "弹窗提示升级",
		},
		{
			Value: ForceUpgradeYes,
			Label: "强制用户升级app",
		},
	}
)
