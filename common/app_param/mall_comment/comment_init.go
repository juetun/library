package mall_comment

import (
	"github.com/juetun/base-wrapper/lib/base"
	"github.com/juetun/library/common/app_param/order_create/parameters"
	"github.com/juetun/library/common/app_param/upload_operate"
	"github.com/juetun/library/common/recommend"
)

const (
	CommentCanImageCount = 9 //总能够上传的图片数
)
const (
	CommentForEditAnonymousYes = iota + 1 //是匿名评论
	CommentForEditAnonymousNo             //不是匿名评论
)

var (
	SliceCommentForEditAnonymous = base.ModelItemOptions{
		{
			Label: "是",
			Value: CommentForEditAnonymousYes,
		},
		{
			Label: "否",
			Value: CommentForEditAnonymousNo,
		},
	}
)

type (
	CommentForEdit struct {
		SendLevel        float32           `json:"send_level"`        //快递包装评分
		DeliveryLevel    float32           `json:"delivery_level"`    //送货速度评分
		PackingLevel     float32           `json:"packing_level"`     //配送员服务
		CreatedAt        string            `json:"created_at"`        //订单生成时间
		OrderId          string            `json:"order_id"`          //订单号
		SubOrderId       string            `json:"sub_order_id"`      //子订单号
		Status           uint8             `json:"status"`            //订单状态
		SubStatus        uint8             `json:"sub_status"`        //子单状态
		HaveComment      bool              `json:"have_comment"`      //是否已评论
		ShopId           int64             `json:"shop_id"`           //店铺ID
		Anonymous        uint8             `json:"anonymous"`         //是否匿名评论
		ActComprehensive bool              `json:"act_comprehensive"` //是否提交综合评价
		HideCommonBtn    bool              `json:"hide_common_btn"`   //是否隐藏评论提交按钮
		UploadDataType   string            `json:"upload_data_type"`  //上传类型
		SkuList          []*CommentSkuItem `json:"sku_list"`          //商品信息
	}
	CommentSkuItem struct {
		SkuInfo        CommentSku          `json:"sku_info"`
		Mark           string              `json:"mark"`
		Videos         []*CommentVideoItem `json:"videos"`           //视频
		Images         []*CommentImageItem `json:"images"`           //图片
		CommentScore   float32             `json:"comment_score"`    //商品评论等级
		ImageCount     int                 `json:"image_count"`      //图片数
		CanImageCount  int                 `json:"can_image_count"`  //总能上传数
		CommentAt      string              `json:"comment_at"`       //评论时间
		AddComment     string              `json:"add_comment"`      //追平内容
		ShowSkuComment bool                `json:"show_sku_comment"` //是否显示评论
		CanComment     bool                `json:"can_comment"`      //是否能够评论
		HasComment     bool                `json:"has_comment"`      //是否已评论
		HasAddComment  bool                `json:"has_add_comment"`  //是否已追评
		AddImages      []*CommentImageItem `json:"add_images"`       //追平图片
	}

	CommentSku struct {
		SkuName       string        `json:"sku_name"`
		SkuProperty   string        `json:"sku_property"`
		ThumbnailURL  string        `json:"thumbnail_url"`
		SkuId         string        `json:"sku_id"`
		SkuNum        int64         `json:"sku_num"`
		SpuId         string        `json:"spu_id"`
		ShopId        int64         `json:"shop_id"`
		Status        uint8         `json:"status"` //订单状态
		StatusName    string        `json:"status_name"`
		SubStatus     uint8         `json:"sub_status"` //子单状态
		SubStatusName string        `json:"sub_status_name"`
		Price         string        `json:"price"`
		SaleType      uint8         `json:"sale_type"`
		SaleTypeName  string        `json:"sale_type_name"`
		HaveGift      uint8         `json:"have_gift"`
		Gifts         []*CommentSku `json:"gifts,omitempty"` //赠品
		Href          interface{}   `json:"href"`
	}
)

func NewCommentForEdit() (res *CommentForEdit) {
	res = &CommentForEdit{
		SendLevel:      5,
		DeliveryLevel:  5,
		PackingLevel:   5,
		SkuList:        make([]*CommentSkuItem, 0),
		Anonymous:      CommentForEditAnonymousYes,
		UploadDataType: recommend.AdDataDataTypeOrderComment,
	}
	return
}

//判断所有已收货的商品是否已评论
// haveNotReceipt 是否有未确认收货的商品 true-有 false-没有
func (r *CommentForEdit) flagAllHasComment() (allHasComment, haveNotReceipt bool) {

	allHasComment = true
	for _, item := range r.SkuList {
		switch item.SkuInfo.SubStatus {
		case parameters.OrderStatusGoodSendFinished: //已收货
			if !item.HasComment {
				allHasComment = false
			}
		case parameters.OrderStatusGoodSending: //已发货，未收货
			haveNotReceipt = true
		}

	}
	return
}

func (r *CommentForEdit) allReceiptAndComment() {
	r.ActComprehensive = true
	r.HaveComment = true
	for _, item := range r.SkuList {
		item.ShowSkuComment = true
	}
	r.HideCommonBtn = true //订单都已评论，直接展示 ，无法再操作
	return
}

//所有已收货，部分未评论
func (r *CommentForEdit) allReceiptAndNotAllComment() {
	r.ActComprehensive = true
	r.HaveComment = false
	for _, item := range r.SkuList {
		if item.HasComment { //已评论的隐藏
			continue
		}
		//未评论的显示
		item.ShowSkuComment = true
	}
	r.HideCommonBtn = false
	return
}

//如果有未收货的商品
func (r *CommentForEdit) haveNotReceipt() {
	r.ActComprehensive = false
	r.HaveComment = false
	r.HideCommonBtn = true
	for _, item := range r.SkuList {
		switch item.SkuInfo.SubStatus {
		case parameters.OrderStatusGoodSending: //已发货未收货
			continue
		case parameters.OrderStatusGoodSendFinished: //已收货
			item.ShowSkuComment = true
			r.HideCommonBtn = false
		case parameters.OrderStatusHasComment, parameters.OrderStatusHasCommentAuto: // 已评价 自动评价
			item.ShowSkuComment = false
		}
	}
	if r.HideCommonBtn {
		for _, item := range r.SkuList {
			switch item.SkuInfo.SubStatus {
			case parameters.OrderStatusHasComment, parameters.OrderStatusHasCommentAuto: // 已评价 自动评价
				item.ShowSkuComment = true
			}
		}
	}
	return
}

func (r *CommentForEdit) Default() (err error) {
	if r.SkuList == nil {
		r.SkuList = make([]*CommentSkuItem, 0)
	}
	if r.Anonymous == 0 {
		r.Anonymous = CommentForEditAnonymousYes
	}
	var (
		allHasComment  bool //所有已评论的商品
		haveNotReceipt bool
	)
	allHasComment, haveNotReceipt = r.flagAllHasComment()

	if !haveNotReceipt { //没有未收货的商品
		if allHasComment { //如果所有的商品已评论,且没有未发货的商品
			r.allReceiptAndComment()
			return
		}
		//已收货，部分未评论
		r.allReceiptAndNotAllComment()
		return
	}
	//如果有未收货的商品
	r.haveNotReceipt()

	return
}

func NewCommentSkuItem() (res *CommentSkuItem) {
	res = &CommentSkuItem{
		CommentScore: 5,
		Videos:       make([]*CommentVideoItem, 0, CommentCanImageCount),
		Images:       make([]*CommentImageItem, 0, CommentCanImageCount),
	}
	return
}

func (r *CommentSkuItem) DeferLogic() {
	r.ImageCount = len(r.Images) + len(r.Videos)

	r.CanImageCount = CommentCanImageCount
}

func (r *CommentVideoItem) ParseFromVideo(video *upload_operate.UploadVideo) {
	r.Src = video.Src
	r.ShowData = video.GetShowUrl()
}

func (r *CommentImageItem) ParseFromImg(img *upload_operate.UploadFile) {
	r.ImgUrl = img.Src
	r.BigImgUrl = img.Src
	r.SmallImgUrl = img.Src
}
