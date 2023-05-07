package const_apply

const (
	ApplyStatusOk          int8 = iota + 1 //入驻状态审核通过
	ApplyStatusInvalid                     //已失效
	ApplyStatusFailure                     //入驻状态审核失败
	ApplyStatusInit                        //入驻状态初始化或编辑中
	ApplyStatusAuditing                    //审核中
	ApplyStatusTimeEditing                 //编辑中
	ApplyStatusTimeInvalid                 //超时(失效)
)

