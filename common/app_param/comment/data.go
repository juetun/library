package comment

import (
	"fmt"
	"github.com/juetun/base-wrapper/lib/base"
	"github.com/juetun/base-wrapper/lib/common"
	"strings"
)

const (
	ActTypeAdd    = "add"
	ActTypeCancel = "cancel"
)
const (
	DivIdString = "^" // 唯一数据key的描述分隔符
)
const (
	AttendDataTypeUser         = "usr"           // 用户
	AttendDataTypeSpu          = "spu"           // 商品
	AttendDataTypeShop         = "shop"          // 店铺
	AttendDataTypeFishingSpots = "fishing_spots" // 钓点
	AttendDataTypeSns          = "article"       // 社交帖子
)
const (
	ActTypeLove     = "love"      //点赞
	ActTypeSee      = "see"       //喜欢
	ActTypeCollect  = "collect"   //收藏
	ActTypeShare    = "share"     //转发
	ActTypeMyAttend = "my_attend" //关注
	ActTypeAttendMe = "attend_me" //粉丝
	ActTypeComment  = "comment"   //评论
	ActTypeReplay   = "replay"    //回复
)

//数据库存储的数据
const (
	ActTypeLoveValue     uint8 = iota + 1 //点赞
	ActTypeSeeValue                       //喜欢
	ActTypeCollectValue                   //收藏
	ActTypeMyAttendValue                  //关注
	ActTypeAttendMeValue                  //粉丝
	ActTypeShareValue                     //转发数
	ActTypeCommentValue                   //评论
	ActTypeReplyValue                     //回复
)

var (
	SliceAttendDataType = base.ModelItemOptions{
		{
			Value: AttendDataTypeUser,
			Label: "用户",
		},
		{
			Value: AttendDataTypeSpu,
			Label: "商品",
		},
		{
			Value: AttendDataTypeShop,
			Label: "店铺",
		},
		{
			Value: AttendDataTypeFishingSpots,
			Label: "钓点",
		},
		{
			Value: AttendDataTypeSns,
			Label: "帖子",
		},
	}

	SliceCommonActType = base.ModelItemOptions{
		{
			Label: "添加",
			Value: ActTypeAdd,
		},
		{
			Label: "取消",
			Value: ActTypeCancel,
		},
	}

	MapCommonActType = map[string]uint8{
		ActTypeLove:     ActTypeLoveValue,
		ActTypeSee:      ActTypeSeeValue,
		ActTypeCollect:  ActTypeCollectValue,
		ActTypeMyAttend: ActTypeMyAttendValue,
		ActTypeAttendMe: ActTypeAttendMeValue,
		ActTypeShare:    ActTypeShareValue,
		ActTypeComment:  ActTypeCommentValue,
		ActTypeReplay:   ActTypeReplyValue,
	}
	SliceActType = base.ModelItemOptions{
		{
			Value: ActTypeLove,
			Label: "点赞",
		},
		{
			Value: ActTypeSee,
			Label: "喜欢",
		},
		{
			Value: ActTypeCollect,
			Label: "收藏",
		},
		{
			Value: ActTypeMyAttend,
			Label: "关注",
		},
		{
			Value: ActTypeAttendMe,
			Label: "粉丝",
		},
		{
			Value: ActTypeShare,
			Label: "转发",
		},
		{
			Value: ActTypeComment,
			Label: "评论",
		},
		{
			Value: ActTypeReplay,
			Label: "回复",
		},
	}
)

type (
	ArgGetNumberByKeys struct {
		CurrentUserHid    int64                  `json:"current_user_hid" form:"current_user_hid"`
		Keys              []string               `json:"keys" form:"keys"`
		TimeNow           base.TimeNormal        `json:"-" form:"-"`
		GetDataTypeCommon base.GetDataTypeCommon `json:"get_data_type_common" form:"get_data_type_common"`
	}
	ResultGetNumberByKeys map[string]*ResultGetNumberItem
	ResultGetNumberItem   struct {
		ActType    string `json:"act_type" form:"act_type"`
		DataType   string `json:"data_type" form:"data_type"`
		DataId     string `json:"data_id" form:"data_id"`
		CommentId  string `json:"comment_id" form:"comment_id"`
		Key        string `json:"key" form:"key"`
		Num        int64  `json:"num" form:"num"`                 //数量
		HasOperate bool   `json:"has_operate" form:"has_operate"` //当前用户是否已经操作过
	}
	ArgActData struct {
		common.HeaderInfo
		ShopId   int64           `json:"shop_id" form:"shop_id"`
		UserHid  int64           `json:"user_hid" form:"user_hid"`
		ActType  string          `json:"act_type" form:"act_type"`
		DataType string          `json:"data_type" form:"data_type"`
		DataIds  []*ActDataItem  `json:"data_ids" form:"data_ids"`
		DoType   string          `json:"do_type" form:"do_type"`
		TimeNow  base.TimeNormal `json:"time_now" form:"time_now"`
	}
	ActDataItem struct {
		DataId string `json:"data_id"`
		ShopId int64  `json:"shop_id"`
	}
	ResultActDat struct {
		Result bool `json:"result"`
	}
)

func (r *ArgGetNumberByKeys) Default(c *base.Context) (err error) {

	if r.TimeNow.IsZero() {
		r.TimeNow = base.GetNowTimeNormal()
	}
	return
}

func (r *ArgGetNumberByKeys) GetKey(ActType string, DataType string, DataId string) (res string) {
	return strings.Join([]string{ActType, DataType, DataId}, DivIdString)
}

func (r *ResultGetNumberItem) GetDataPk() (res string) {

	var sliceData []string
	if r.CommentId != "" {
		sliceData = []string{r.DataType, r.DataId, r.CommentId}
	} else {
		sliceData = []string{r.DataType, r.DataId}
	}
	res = strings.Join(sliceData, DivIdString)
	return
}

func (r *ArgGetNumberByKeys) ParseByKey(key string) (ActType, DataType, DataId, commentId string) {
	list := strings.Split(key, DivIdString)
	switch len(list) {
	case 1:
		ActType = list[0]
	case 2:
		ActType = list[0]
		DataType = list[1]
	case 3:
		ActType = list[0]
		DataType = list[1]
		DataId = list[2]
	case 4:
		ActType = list[0]
		DataType = list[1]
		DataId = list[2]
		DataId = list[4]
	}

	return
}

func (r *ArgActData) Default(c *base.Context) (err error) {
	if err = r.InitHeaderInfo(c.GinContext); err != nil {
		return
	}
	if err = r.validateDoType(); err != nil {
		return
	}
	if err = r.validateActType(); err != nil {
		return
	}
	if r.TimeNow.IsZero() {
		r.TimeNow = base.GetNowTimeNormal()
	}
	return
}

func (r *ArgActData) validateActType() (err error) {
	var mapActTypes map[string]string
	if mapActTypes, err = SliceActType.GetMapAsKeyString(); err != nil {
		return
	}
	if r.ActType == "" {
		err = fmt.Errorf("请选择您要操作的数据类型act_type")
		return
	}
	if _, ok := mapActTypes[r.ActType]; !ok {
		err = fmt.Errorf("未知操作的数据类型act_type")
		return
	}
	return
}

func (r *ArgActData) validateDoType() (err error) {
	var mapDoType map[string]string
	if mapDoType, err = SliceCommonActType.GetMapAsKeyString(); err != nil {
		return
	}
	if _, ok := mapDoType[r.DoType]; !ok {
		err = fmt.Errorf("当前不支持你的操作类型(do_type:%s)", r.DoType)
		return
	}

	return
}
