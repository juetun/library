package mall_comment

import (
	"encoding/json"
	"github.com/juetun/base-wrapper/lib/base"
	"github.com/juetun/library/common/app_param/upload_operate/ext_up"
)

type (
	UInfo struct {
		Id             int64       `json:"id"`
		Content        string      `json:"content"`          //评论内容
		NickName       string      `json:"nickname"`         //用户昵称
		Score          int         `json:"score"`            //评分
		Avatar         string      `json:"avatar"`           //用户头图
		Time           string      `json:"time"`             //评论时间
		Size           string      `json:"size"`             //购买商品属性详情
		Replay         int64       `json:"replay"`           //评论回复数
		Like           int64       `json:"like"`             //商品评论点赞赞数
		SpuId          string      `json:"spu_id"`           //商品SPU_ID
		SkuId          string      `json:"sku_id"`           //商品Sku_ID
		HaveExtComment bool        `json:"have_ext_comment"` //是否有追评
		OrderId        string      `json:"order_id"`         //订单号
		SubOrderId     string      `json:"sub_order_id"`     //子单号
		Num            int64       `json:"num"`              //商品数量
		UserLink       interface{} `json:"user_link"`        //用户跳转连接
	}
	CommentVideoItem struct {
		Src             string `json:"src"` //视频源地址
		ext_up.ShowData                     //视频地址
	}
	CommentImageItem struct {
		Src         string `json:"src"`         //原图
		SmallImgUrl string `json:"smallImgUrl"` //小图
		BigImgUrl   string `json:"bigImgUrl"`   //大图
		ImgUrl      string `json:"imgUrl"`      //正常图片
	}
	CommentFollow struct {
		Days     int      `json:"days"`      //追平和评论时间间隔
		Content  string   `json:"content"`   //评论内容
		Images   []string `json:"images"`    //评论图片
		ImageNum int64    `json:"image_num"` //追评图片数
	}
	CommentItem struct {
		Videos   []*CommentVideoItem `json:"videos"` //视频
		Images   []*CommentImageItem `json:"images"` //图片
		Follow   *CommentFollow      `json:"follow"` //回复数据信息
		UserInfo UInfo               `json:"info"`   //用户信息
	}
	OrderComment struct {
		ShopGoodBit bool           `json:"shop_good_bit"` //是否展示好评度
		GoodBit     string         `json:"good_bit"`      //好评度
		Comment     []*CommentItem `json:"Comment"`       //评论列表
		Number      int64          `json:"number"`        //评论数量
	}

	ArgAddComment struct {
		TimeNow          base.TimeNormal   `json:"time_now" form:"time_now"`
		ActComprehensive bool              `json:"act_comprehensive"` //是否提交综合评价
		List             []*AddCommentItem `json:"list" form:"list"`
	}
	AddCommentItem struct {
		OrderId      string `json:"order_id" form:"order_id"`
		SubOrderId   string `json:"sub_order_id" form:"sub_order_id"`
		ShopId       int64  `json:"shop_id" form:"shop_id"`
		SpuId        string `json:"spu_id" form:"spu_id"`
		SkuId        string `json:"sku_id" form:"sku_id"`
		IsNewComment bool   `json:"is_new_comment" form:"is_new_comment"` //是否是本次操作提交评论
	}
)

func NewOrderComment() (res *OrderComment) {
	res = &OrderComment{ShopGoodBit: true, GoodBit: "100%", Comment: []*CommentItem{}}
	return
}

func (r *OrderComment) UnmarshalBinary(data []byte) (err error) {
	err = json.Unmarshal(data, r)
	return
}

//实现 序列化方法 encoding.BinaryMarshaler
func (r *OrderComment) MarshalBinary() (data []byte, err error) {
	if r == nil {
		data = []byte("{}")
		return
	}
	data, err = json.Marshal(r)
	return
}

func (r *ArgAddComment) Default(ctx *base.Context) (err error) {
	if r.TimeNow.IsZero() {
		r.TimeNow = base.GetNowTimeNormal()
	}
	return
}
