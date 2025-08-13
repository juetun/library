package chat_const

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/juetun/base-wrapper/lib/app/app_obj"
	"github.com/juetun/base-wrapper/lib/base"
	"github.com/juetun/base-wrapper/lib/plugins/rpc"
	"github.com/juetun/base-wrapper/lib/utils"
	"github.com/juetun/library/common/app_param"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"
)

type (
	SendMsgChat struct {
		MsgId           string `json:"msg_id,omitempty" form:"msg_id"`       // 消息ID
		ChatType        uint8  `json:"chat_type,omitempty" form:"chat_type"` // 聊天类型 单聊 、群聊
		FromId          string `json:"from_id,omitempty" form:"from_id"`     // 发送消息方
		FromType        uint8  `json:"from_type,omitempty" form:"from_type"` //
		ToId            string `json:"to_id,omitempty" form:"to_id"`         // 发送给的目标用户或群
		ToType          uint8  `json:"to_type,omitempty" form:"to_type"`     //
		ToPathType      string `json:"to_path_type",form:"to_path_type"`
		ManagerUserHid  string `json:"manager_user_hid,omitempty" form:"manager_user_hid"`   // 管理员ID
		Content         string `json:"content,omitempty" form:"content"`                     // 消息内容
		IsCustomer      bool   `json:"is_customer,omitempty" form:"is_customer"`             // 是否是客服聊天
		CustomerType    uint8  `json:"customer_type,omitempty" form:"customer_type"`         // 客服主体(CustomerManagerDataTypeShop-店铺实体  CustomerManagerDataTypePlatform-系统客服)
		ToCustomerUser  string `json:"to_customer_user,omitempty" form:"to_customer_user"`   // 聊天信息产生时的管理员用户ID
		CurrentUserRole string `json:"current_user_role,omitempty" form:"current_user_role"` //当前用户角色
		SendType        uint8  `json:"send_type,omitempty" form:"send_type"`                 // 发送方式 1-全双工 2-发送方可见 3-接收方可见
		ReplayType      uint8  `json:"replay_type,omitempty" form:"replay_type"`             //客户端响应数据类型 ，提醒客户端刷新消息内容
	}
	ArgGetChatToken struct {
		ArgGetChatTokenOther
		CurrentUserRole string `json:"current_user_role,omitempty" form:"current_user_role"` //当前用户聊天中的角色 customer-客服;user-普通用户
		IsCustomer      bool   `json:"is_customer,omitempty" form:"is_customer"`             //当前数据是否是与客服聊天
	}
	ArgChatSend struct {
		app_param.RequestUser
		Data []*struct {
			ArgGetChatToken
			SendMsgChat    SendMsgChat `json:"send_msg_chat" form:"send_msg_chat"`
			MsgType        int         `json:"msg_type" form:"msg_type"`
			CurrentUserStr string      `json:"cus" form:"cus"`
		}
	}
	SendMsgChatParam struct {
		ConnectToken  string `json:"connect_token,omitempty"` // 聊天的token信息
		ToUserNotRead int64  `json:"not_read"`                // 消息接收方未读消息数
		MsgType       int    `json:"msg_type"`                // 消息类型（）
		TimeStamp     int64  `json:"-"`                       // 消息发送的时间（单位：纳秒）
	}
	SendMsgCommon struct {
		Code            int    `json:"code,omitempty"`              // 当前请求是否异常状态码
		App             string `json:"app,omitempty"`               // 应用名称
		XTimeStamp      int64  `json:"x_time_stamp,omitempty"`      // 消息发送的时间（单位：毫秒）
		ErrMsg          string `json:"err_msg,omitempty"`           // 如果错误，错误提示
		XSign           string `json:"x_sign,omitempty"`            // 签名
		DataType        uint8  `json:"data_type,omitempty"`         // 数据类型
		NeedApply       bool   `json:"need_apply,omitempty"`        // 是否需要调用反垃圾审核
		SecWebsocketKey string `json:"sec_websocket_key,omitempty"` // 连接的唯一KEY
	}
	SendMsg struct {
		SendMsgCommon
		SendMsgChatParam
		TimeNow base.TimeNormal `json:"-"`
		Content string          `json:"content,omitempty"`  //普通内容json字符串
		MsgInfo *SendMsgChat    `json:"msg_info,omitempty"` //聊天内容信息（只限聊天场景使用）
	}
	ResultSendChatMessage struct {
		Result bool `json:"result"`
	}
	SendMsgHandler func(sendMsg *SendMsg)
)

func (r *ArgGetChatToken) needAttendArea() (res bool) {
	area := []uint8{ConstChatTokenNeedAttendNo, ConstChatTokenNeedAttendUser, ConstChatTokenNeedAttendEachOther}
	for _, i2 := range area {
		if i2 == r.NeedAttend {
			res = true
			return
		}
	}
	return
}

func (r *ArgChatSend) Default(ctx *base.Context) (err error) {
	for _, item := range r.Data {
		if err = item.ArgGetChatToken.Default(ctx); err != nil {
			return
		}
		if item.MsgType == 0 {
			item.MsgType = ChatMsgTypeText
		}
		if item.SendMsgChat.ChatType == 0 {
			item.SendMsgChat.ChatType = ChatMsgChatTypeSingle
		}
		if err = item.SendMsgChat.Default(); err != nil {
			return
		}
		item.SendMsgChat.InitMsgId(item.MsgType)
		var (
			fromEntityInfo *FromEntityInfo
		)
		if fromEntityInfo, err = GetFromEntityInfo(item.CurrentUserRole, &r.RequestUser); err != nil {
			return
		}
		item.FromType = fromEntityInfo.FromType
		item.FromId = fromEntityInfo.FromId
		item.CurrentUserStr = fromEntityInfo.CurrentUserStr
	}

	return
}

func (r *SendMsg) ToJson() (res []byte, err error) {
	res, err = json.Marshal(r)
	return
}

func (r *SendMsg) Default() (err error) {
	if r.MsgInfo == nil {
		r.MsgInfo = NewSendMsgChatText("")
	}
	r.MsgInfo.InitMsgId(r.MsgType)
	return
}

func ParseSendType(sendType uint8) (res string, err error) {
	var mapV map[uint8]string
	if mapV, err = SliceChatMsgSendType.GetMapAsKeyUint8(); err != nil {
		return
	}
	if dt, ok := mapV[sendType]; ok {
		res = dt
		return
	}
	err = fmt.Errorf("当前不支持数据类型(%d)", sendType)
	return
}

func (r *SendMsgChat) ParseSendType() (res string, err error) {
	return ParseSendType(r.SendType)
}

func NewSendMsgChatText(text string) (res *SendMsgChat) {
	res = &SendMsgChat{
		ChatType:       ChatMsgChatTypeSingle,
		ToId:           "",
		Content:        text,
		IsCustomer:     false,
		CustomerType:   0,
		ToCustomerUser: "",
		MsgId:          "",
		ToType:         0,
		SendType:       0,
	}
	_ = res.Default()
	return
}

func NewSendMsg(options ...SendMsgHandler) (res *SendMsg) {
	res = &SendMsg{MsgInfo: &SendMsgChat{}}
	for _, handler := range options {
		handler(res)
	}
	if res.XTimeStamp == 0 {
		res.TimeStamp = time.Now().UnixNano()
	}
	return
}

func GetClientLoadMoreChatListMsg(chatMsg *SendMsg) (replyByte []byte, err error) {
	chatMsg.Code = WebsocketChatList
	if replyByte, err = chatMsg.ToJson(); err != nil {
		return
	}
	return
}

func (r *SendMsg) GetSendMsgChat() (res *SendMsgChat, err error) {
	res = r.MsgInfo
	return
}

func SendMsgConnectToken(connectToken string) (res SendMsgHandler) {
	return func(replyMsg *SendMsg) {
		replyMsg.ConnectToken = connectToken
	}
}

func SendMsgDataType(dataType uint8) (res SendMsgHandler) {
	return func(replyMsg *SendMsg) {
		replyMsg.DataType = dataType
	}
}

func SendMsgType(msgType int) (res SendMsgHandler) {
	return func(replyMsg *SendMsg) {
		replyMsg.MsgType = msgType
	}
}

func SendMsgCode(code int) (res SendMsgHandler) {
	return func(replyMsg *SendMsg) {
		replyMsg.Code = code
	}
}

func SendMsgCodeSecWebsocketKey(secWebsocketKey string) (res SendMsgHandler) {
	return func(replyMsg *SendMsg) {
		replyMsg.SecWebsocketKey = secWebsocketKey
	}
}

func SendMsgChatMsgInfo(content *SendMsgChat) (res SendMsgHandler) {
	return func(replyMsg *SendMsg) {
		replyMsg.MsgInfo = content
	}
}

func SendMsgTimeNow(timeStamp int64) (res SendMsgHandler) {
	return func(replyMsg *SendMsg) {
		replyMsg.TimeStamp = timeStamp
	}
}

func SendMsgTimeNormal(timeStamp base.TimeNormal) (res SendMsgHandler) {
	return func(replyMsg *SendMsg) {
		replyMsg.TimeNow = timeStamp
	}
}

func (r *SendMsgChat) ToJson() (res string, err error) {
	var bt []byte
	if bt, err = json.Marshal(r); err != nil {
		return
	}
	res = string(bt)
	return
}

func (r *SendMsgChat) Default() (err error) {
	if r.ChatType == 0 {
		r.ChatType = ChatMsgChatTypeSingle
	}
	if r.SendType == 0 {
		r.SendType = ChatMsgSendTypeEachOther
	}
	return
}

func (r *SendMsgChat) InitMsgId(msgType int) {
	if r.MsgId == "" {
		base64Code := fmt.Sprintf("%v#%v#%v#%v#%v#%v#%v", r.ChatType, msgType, r.FromType, r.FromId, r.ToType, r.ToId, strings.Join(strings.Split(utils.Guid(""), "-"), ""))
		r.MsgId = base64.URLEncoding.EncodeToString([]byte(base64Code))
	}
	return
}

func (r *SendMsgChat) ParseMsgId(msgId string) (msgType int, err error) {
	r.MsgId = msgId
	if msgId == "" {
		return
	}
	var btValue []byte
	if btValue, err = base64.URLEncoding.DecodeString(r.MsgId); err != nil {
		return
	}

	list := strings.Split(string(btValue), "#")

	listTmp := make([]string, 7)
	for key, value := range list {
		listTmp[key] = value
	}
	r.ChatType, msgType, r.FromType, r.FromId, r.ToType, r.ToId = r.parseStringToUint8(listTmp[0]), r.parseStringToInt(listTmp[1]), r.parseStringToUint8(listTmp[2]), listTmp[3], r.parseStringToUint8(listTmp[4]), listTmp[5]
	return
}

func (r *SendMsgChat) parseStringToInt(value string) (res int) {
	if value == "" {
		return
	}
	res, _ = strconv.Atoi(value)
	return
}

func (r *SendMsgChat) parseStringToUint8(value string) (res uint8) {
	if value == "" {
		return
	}
	resTmp, _ := strconv.ParseUint(value, 10, 64)
	if resTmp > 256 {
		resTmp = 0
	} else {
		res = uint8(resTmp)
	}

	return
}

func (r *SendMsgChat) GetIsCustomer() (isCustomer bool) {
	isCustomer = (&ArgGetChatTokenOther{
		FromId:   r.FromId,
		FromType: r.FromType,
		ToId:     r.ToId,
		ToType:   r.ToType,
	}).GetIsCustomer()
	return
}

//批量发送聊天记录
func SendChatMessageBatch(ctx *base.Context, arg *ArgChatSend) (res *ResultSendChatMessage, err error) {
	res = &ResultSendChatMessage{}
	var value = url.Values{}

	ro := rpc.RequestOptions{
		Method:      http.MethodPost,
		AppName:     app_param.AppNameChat,
		URI:         "/chat/send_msgs",
		Header:      http.Header{},
		Value:       value,
		Context:     ctx,
		PathVersion: app_obj.App.AppRouterPrefix.Intranet,
	}
	if ro.BodyJson, err = json.Marshal(arg); err != nil {
		return
	}
	var data = struct {
		Code int                    `json:"code"`
		Data *ResultSendChatMessage `json:"data"`
		Msg  string                 `json:"message"`
	}{}
	err = rpc.NewHttpRpc(&ro).
		Send().
		GetBody().
		Bind(&data).Error
	if err != nil {
		return
	}
	res = data.Data
	return
}
