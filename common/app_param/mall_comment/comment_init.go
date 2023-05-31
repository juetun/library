package mall_comment

import "github.com/juetun/library/common/app_param/upload_operate"

const (
	CommentCanImageCount = 9 //总能够上传的图片数
)

type (
	CommentForEdit struct {
		SendLevel     float32           `json:"send_level"`     //快递包装评分
		DeliveryLevel float32           `json:"delivery_level"` //送货速度评分
		PackingLevel  float32           `json:"packing_level"`  //配送员服务
		CreatedAt     string            `json:"created_at"`     //订单生成时间
		OrderId       string            `json:"order_id"`       //订单号
		SubOrderId    string            `json:"sub_order_id"`   //子订单号
		HaveComment   bool              `json:"have_comment"`   //是否已评论
		ShopId        int64             `json:"shop_id"`        //店铺ID
		SkuList       []*CommentSkuItem `json:"sku_list"`       //商品信息
	}
	CommentSkuItem struct {
		SkuInfo       CommentSku          `json:"sku_info"`
		Mark          string              `json:"mark"`
		Videos        []*CommentVideoItem `json:"videos"`          //视频
		Images        []*CommentImageItem `json:"images"`          //图片
		CommentScore  float32             `json:"comment_score"`   // 商品评论等级
		ImageCount    int                 `json:"image_count"`     //图片数
		CanImageCount int                 `json:"can_image_count"` //总能上传数
		CommentAt     string              `json:"comment_at"`      //评论时间
		AddComment    string              `json:"add_comment"`     //追平内容
		HasAddComment string              `json:"has_add_comment"` //是否已追评
		AddImages     []*CommentImageItem `json:"add_images"`      //追平图片
	}

	CommentSku struct {
		SkuName      string        `json:"sku_name"`
		SkuProperty  string        `json:"sku_property"`
		ThumbnailURL string        `json:"thumbnail_url"`
		SkuId        string        `json:"sku_id"`
		SkuNum       int64         `json:"sku_num"`
		SpuId        string        `json:"spu_id"`
		ShopId       int64         `json:"shop_id"`
		Price        string        `json:"price"`
		SaleType     uint8         `json:"sale_type"`
		SaleTypeName string        `json:"sale_type_name"`
		HaveGift     uint8         `json:"have_gift"`
		Gifts        []*CommentSku `json:"gifts,omitempty"` //赠品
		Href         interface{}   `json:"href"`
	}
)

func NewCommentForEdit() (res *CommentForEdit) {
	res = &CommentForEdit{
		SendLevel:     5,
		DeliveryLevel: 5,
		PackingLevel:  5,
	}
	return
}

func (r *CommentForEdit) Default() {

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
