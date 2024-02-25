package recommend

import (
	"fmt"
	"github.com/juetun/base-wrapper/lib/base"
	"github.com/juetun/library/common/app_param/comment"
	"github.com/juetun/library/common/app_param/upload_operate"
	"net/url"
	"strings"
)

const (
	DataItemShowTypeCard       = "card"        //普通card类型
	DataItemShowTypeCardShop   = "card_shop"   //普通店铺类型
	DataItemShowTypeSns        = "sns"         //社交类型(钓点和圈子动态使用)
	DataItemShowTypeImgList    = "img_list"    //图片列表类型
	DataItemShowTypeCardRight  = "card_right"  //card图片展示在右边
	DataItemShowTypeCardDouble = "card_double" //双列
	DataItemShowTypeUser       = "card_user"   //用户信息
)

//是否在推荐列表中展示
const (
	ShowInListYes uint8 = iota + 1 //展示
	ShowInListNo                   //不展示
)

const (
	BadgeTypeNull uint8 = iota //空 默认
	BadgeTypeNum               //数字
	BadgeTypeDot               //点
)

var (
	SliceBadgeType = base.ModelItemOptions{
		{
			Label: "数字",
			Value: BadgeTypeNum,
		},
		{
			Label: "点",
			Value: BadgeTypeDot,
		},
		{
			Label: "默认(空)",
			Value: BadgeTypeNull,
		},
	}

	SliceShowInList = base.ModelItemOptions{
		{
			Label: "展示",
			Value: ShowInListYes,
		},
		{

			Label: "不展示",
			Value: ShowInListNo,
		},
	}
)

type (
	DataItem struct {
		Title         string                     `json:"title,omitempty"`      //标题
		PreTags       []*DataItemTag             `json:"pre_tags,omitempty"`   //前缀标签
		DataType      string                     `json:"data_type"`            //数据类型
		DataTypeStr   string                     `json:"data_type_str"`        //数据类型字符串 点赞 收藏使用的类型
		DataId        string                     `json:"data_id"`              //数据ID
		Link          interface{}                `json:"link,omitempty"`       //链接地址 小程序对象DataItemLinkMina
		HaveVideo     bool                       `json:"have_video,omitempty"` //是否有视频
		VideoInfo     *upload_operate.VideoInfo  `json:"video_info,omitempty"` //视频信息
		ImgData       string                     `json:"-"`
		Img           string                     `json:"img,omitempty"`          //头图
		Imgs          []string                   `json:"imgs,omitempty"`         //多条图片
		DataValue     map[string]*DataItemDetail `json:"data_value,omitempty"`   //详情
		CanBuy        bool                       `json:"can_buy,omitempty"`      //是否能够购买
		ShowError     bool                       `json:"show_error,omitempty"`   //是否展示错误提示，不显示商品其他内容 true-商品不在列表展示（详情页提示错误信息）
		ShopManager   bool                       `json:"shop_manager,omitempty"` //当前用户是否为店铺管理员
		Mark          string                     `json:"mark,omitempty"`         //备注
		SuffixTags    []*DataItemTag             `json:"suffix_tags,omitempty"`  //后缀标签
		ShowShop      bool                       `json:"show_shop,omitempty"`    //是否显示店铺名
		ShopName      string                     `json:"shop_name,omitempty"`    //店铺名称
		CurrentAccUId int64                      `json:"current_acc_uid"`        //获取数据的用户ID
		CreateUid     int64                      `json:"create_uid,omitempty"`   //信息发布者
		ShopId        int64                      `json:"shop_id,omitempty"`      //店铺ID
		ShopLink      interface{}                `json:"shop_link,omitempty"`    //链接地址 小程序对象DataItemLinkMina
		ShopIcon      string                     `json:"shop_icon,omitempty"`    //店铺Icon
		ExtraMsg      string                     `json:"extra_msg,omitempty"`    //携带的其他信息
		ShowType      string                     `json:"show_type"`              //展示样式 默认card
		Children      []*DataItem                `json:"children,omitempty"`     //子列表
		ShowTime      string                     `json:"show_time,omitempty"`    //展示时间
		OtherData     interface{}                `json:"other_data,omitempty"`   //其他数据
		Pk            string                     `json:"pk"`                     //数据的唯一KEy
		BadgeType     string                     `json:"badge_type,omitempty"`   //徽标类型 num-数字 dot-点 空不填
		BadgeString   string                     `json:"badge_string,omitempty"` //徽标值    "100" "10"
		PageName      string                     `json:"-"`                      //页面名称 内部使用参数不对前端展示
		PageConfigId  int64                      `json:"-"`
	}
	DataItemLinkMina struct {
		PageName string                 `json:"page_name"`
		Query    map[string]interface{} `json:"query,omitempty"`
	}
	DataItemTag struct {
		Type      string `json:"type"`                //标签类型，可选值为primary success danger warning	默认	default
		Label     string `json:"label"`               //类型名称
		Color     string `json:"color,omitempty"`     //标签颜色
		TextColor string `json:"textColor,omitempty"` //文本颜色，优先级高于color属性	String	white
		Plain     bool   `json:"plain"`               //是否为空心样式	Boolean	false
		Round     bool   `json:"round"`               //是否为圆角样式	Boolean	false
		Mark      bool   `json:"mark"`                //是否为标记样式
	}
	DataItemDetail struct {
		Type      string `json:"type,omitempty"`      //标签类型，可选值为primary success danger warning	默认	default
		Value     string `json:"value"`               //值
		Label     string `json:"label"`               //类型名称
		Color     string `json:"color,omitempty"`     //标签颜色
		TextColor string `json:"textColor,omitempty"` //文本颜色，优先级高于color属性	String	white
		Plain     bool   `json:"plain"`               //是否为空心样式	Boolean	false
		Round     bool   `json:"round"`               //是否为圆角样式	Boolean	false
		Mark      bool   `json:"mark"`                //是否为标记样式
	}
)

//获取广告唯一Id字符串
func (r *DataItem) GetUniqueKey() (res string) {
	return GetUniqueKey(r.DataType, r.DataId)
}

//解析广告唯一Id字符串
func (r *DataItem) ParseUniqueKey() (DataType, DataId string) {
	return ParseUniqueKey(r.Pk)
}

//参数默认值
func (r *DataItem) Default() {
	if r.ShowType == "" {
		r.ShowType = DataItemShowTypeCard
	}
	if len(r.Children) > 0 {
		r.ShowType = DataItemShowTypeImgList
	}
	if r.BadgeType == fmt.Sprintf("%v", BadgeTypeNum) {
		if r.BadgeString == "" {
			r.BadgeString = "0"
		}
	}
	return
}

func (r *DataItem) GetUrlValue() (res url.Values) {
	if r.DataId != "" {
		res = ParseHttp(r.DataId)
	}
	return
}

func (r *DataItem) GetShareKey() (res string) {
	res = fmt.Sprintf("%v_%v_%v", comment.ActTypeShare, r.DataType, r.DataId)
	return

}

func (r *DataItem) GetCollectKey() (res string) {
	res = fmt.Sprintf("%v_%v_%v", comment.ActTypeCollect, r.DataType, r.DataId)

	return
}

func (r *DataItem) GetLoveKey() (res string) {
	res = fmt.Sprintf("%v_%v_%v", comment.ActTypeLove, r.DataType, r.DataId)

	return
}

func (r *DataItem) GetCommentKey() (res string) {
	res = fmt.Sprintf("%v_%v_%v", comment.ActTypeComment, r.DataType, r.DataId)
	return
}

func ParseHttp(clientUrl string) (values url.Values) {
	u, _ := url.Parse(clientUrl) //将string解析成*URL格式
	if u.RawQuery == "" && u.Path != "" {
		values, _ = url.ParseQuery(u.Path) //返回Values类型的字典
	} else if u.RawQuery != "" {
		values, _ = url.ParseQuery(u.RawQuery) //返回Values类型的字典
	}
	return
}

func (r *DataItemTag) Default() {
	if r.TextColor == "" {
		r.TextColor = "white"
	}
	if r.Type == "" {
		r.Type = "default"
	}
	return
}

//获取广告唯一Id字符串
func GetUniqueKey(DataType string, DataId string) (res string) {
	return fmt.Sprintf("%s-%s", DataType, DataId)
}

//获取广告唯一Id字符串
func ParseUniqueKey(pk string) (DataType, DataId string) {
	if pk == "" {
		return
	}
	dataSlice := strings.Split(pk, "-")
	l := len(dataSlice)
	switch l {
	case 0:
		return
	case 1:
		DataType = dataSlice[0]
	default:
		DataType = dataSlice[0]
		DataId = dataSlice[1]
	}
	return
}
