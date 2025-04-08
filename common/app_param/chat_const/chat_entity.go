package chat_const

import (
	"fmt"
	"github.com/juetun/library/common/app_param"
)

type (
	FromEntityInfo struct {
		FromType       uint8  `json:"from_type"`
		FromId         string `json:"from_id"`
		CurrentUserStr string `json:"current_user_str"`
	}
)

func GetFromEntityInfo(currentUserRole string, requestUser *app_param.RequestUser) (res *FromEntityInfo, err error) {
	res = &FromEntityInfo{}
	switch currentUserRole {
	case CurrentUserRoleCustomer:
		res.FromType = ConstChatTokenToShop
		res.FromId = fmt.Sprintf("%v", requestUser.UShopId)
	case CurrentUserRoleUser:
		res.FromType = ConstChatTokenToUser
		res.FromId = fmt.Sprintf("%v", requestUser.UUserHid)
	case CurrentUserRoleConsole: //官方客服
	default:
		err = fmt.Errorf("请选择当前用户角色信息(code:4")
		return
	}
	res.CurrentUserStr = fmt.Sprintf("%v_%v", res.FromType, res.FromId)
	return
}
