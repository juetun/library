package const_apply

import "github.com/juetun/base-wrapper/lib/base"

const (
	ApplyStatusOk          int8 = iota + 1 //入驻状态审核通过
	ApplyStatusInvalid                     //已失效
	ApplyStatusFailure                     //入驻状态审核失败
	ApplyStatusInit                        //入驻状态初始化或编辑中
	ApplyStatusAuditing                    //审核中
	ApplyStatusTimeEditing                 //编辑中
	ApplyStatusTimeInvalid                 //超时(失效)
)

const (
	InitYes = iota + 1 //初始化
	InitNo             // 非初始化
)

var (
	SliceInit = base.ModelItemOptions{
		{
			Label: "初始化",
			Value: InitYes,
		},
		{
			Label: "非初始化",
			Value: InitNo,
		},
	}
)
