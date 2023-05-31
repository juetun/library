package comment

import (
	"fmt"
	"github.com/juetun/base-wrapper/lib/base"
	"github.com/juetun/base-wrapper/lib/common"
)

const (
	ActTypeAdd    = "add"
	ActTypeCancel = "cancel"
)

const (
	ActTypeLove     = "love"      //点赞
	ActTypeSee      = "see"       //喜欢
	ActTypeCollect  = "collect"   //收藏
	ActTypeMyAttend = "my_attend" //关注
	ActTypeAttendMe = "attend_me" //粉丝
)

//数据库存储的数据
const (
	ActTypeLoveValue     uint8 = iota + 1 //点赞
	ActTypeSeeValue                       //喜欢
	ActTypeCollectValue                   //收藏
	ActTypeMyAttendValue                  //关注
	ActTypeAttendMeValue                  //粉丝
)

var (
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
	}
)

type (
	ArgActData struct {
		common.HeaderInfo
		UserHid  int64           `json:"user_hid" form:"user_hid"`
		ActType  string          `json:"act_type" form:"act_type"`
		DataType int32           `json:"data_type" form:"data_type"`
		DataIds  []string        `json:"data_ids" form:"data_ids"`
		DoType   string          `json:"do_type" form:"do_type"`
		TimeNow  base.TimeNormal `json:"time_now" form:"time_now"`
	}
	ResultActDat struct {
		Result bool `json:"result"`
	}
)

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
