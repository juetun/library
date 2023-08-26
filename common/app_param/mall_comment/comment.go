package mall_comment

import (
	"github.com/juetun/base-wrapper/lib/base"
	"github.com/juetun/library/common/app_param/upload_operate/ext_up"
)

type (
	UInfo struct {
		Content  string `json:"content"`
		NickName string `json:"nickname"`
		Score    int    `json:"score"`
		Avatar   string `json:"avatar"`
		Time     string `json:"time"`
		Size     string `json:"size"`
		Replay   int64  `json:"replay"`
		Like     int64  `json:"like"`
	}
	CommentVideoItem struct {
		Src             string `json:"src"`
		ext_up.ShowData        //视频地址
	}
	CommentImageItem struct {
		Src         string `json:"src"`
		SmallImgUrl string `json:"smallImgUrl"`
		BigImgUrl   string `json:"bigImgUrl"`
		ImgUrl      string `json:"imgUrl"`
	}
	CommentFollow struct {
		Days    int      `json:"days"`
		Content string   `json:"content"`
		Images  []string `json:"images"`
	}
	CommentItem struct {
		Videos   []*CommentVideoItem `json:"videos"` //视频
		Images   []*CommentImageItem `json:"images"` //图片
		Follow   CommentFollow       `json:"follow"`
		UserInfo UInfo               `json:"info"` //用户信息
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

func (r *ArgAddComment) Default(ctx *base.Context) (err error) {
	if r.TimeNow.IsZero() {
		r.TimeNow = base.GetNowTimeNormal()
	}
	return
}
