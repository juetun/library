package recommend

import (
	"fmt"
	"github.com/juetun/base-wrapper/lib/base"
)

const (
	DataItemShowTypeCard       = "card"        //普通card类型
	DataItemShowTypeImgList    = "img_list"    //图片列表类型
	DataItemShowTypeCardRight  = "card_right"  //card图片展示在右边
	DataItemShowTypeCardDouble = "card_double" //双列
)

//是否在推荐列表中展示
const (
	ShowInListYes uint8 = iota + 1 //展示
	ShowInListNo                   //不展示
)

var (
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
		DataId        string                     `json:"data_id"`              //数据ID
		Link          interface{}                `json:"link"`                 //链接地址 小程序对象DataItemLinkMina
		HaveVideo     bool                       `json:"have_video,omitempty"` //是否有视频
		ImgData       string                     `json:"-"`
		Img           string                     `json:"img,omitempty"`         //头图
		DataValue     map[string]*DataItemDetail `json:"data_value,omitempty"`  //详情
		CanBuy        bool                       `json:"can_buy"`               //是否能够购买
		ShowError     bool                       `json:"show_error"`            //是否展示错误提示，不显示商品其他内容 true-商品不在列表展示（详情页提示错误信息）
		ShopManager   bool                       `json:"shop_manager"`          //当前用户是否为店铺管理员
		Mark          string                     `json:"mark"`                  //备注
		SuffixTags    []*DataItemTag             `json:"suffix_tags,omitempty"` //后缀标签
		ShowShop      bool                       `json:"show_shop"`             //是否显示店铺名
		ShopName      string                     `json:"shop_name,omitempty"`   //店铺名称
		CurrentAccUId int64                      `json:"current_acc_uid"`       //获取数据的用户ID
		ShopId        int64                      `json:"shop_id,omitempty"`     //店铺ID
		ShopLink      interface{}                `json:"shop_link,omitempty"`   //链接地址 小程序对象DataItemLinkMina
		ShopIcon      string                     `json:"shop_icon,omitempty"`   //店铺Icon
		ExtraMsg      string                     `json:"extra_msg"`             //携带的其他信息
		ShowType      string                     `json:"show_type"`             //展示样式 默认card
		Children      []*DataItem                `json:"children,omitempty"`    //子列表
		ShowTime      string                     `json:"show_time,omitempty"`
		OtherData     interface{}                `json:"other_data,omitempty"` //其他数据
		Pk            string                     `json:"pk"`                   //数据的唯一KEy

		PageName     string `json:"-"` //页面名称 内部使用参数不对前端展示
		PageConfigId string `json:"-"`
	}
	DataItemLinkMina struct {
		PageName string                 `json:"page_name"`
		Query    map[string]interface{} `json:"query"`
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

//参数默认值
func (r *DataItem) Default() {
	if r.ShowType == "" {
		r.ShowType = DataItemShowTypeCard
	}
	if len(r.Children) > 0 {
		r.ShowType = DataItemShowTypeImgList
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
