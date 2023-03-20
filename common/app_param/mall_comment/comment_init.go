package mall_comment

const (
	CommentCanImageCount = 9 //总能够上传的图片数
)

type (
	CommentForEdit struct {
		SendLevel     string            `json:"send_level"`     //快递包装评分
		DeliveryLevel string            `json:"delivery_level"` //送货速度评分
		PackingLevel  string            `json:"packing_level"`  //配送员服务
		CreatedAt     string            `json:"created_at"`     //订单生成时间
		OrderId       string            `json:"order_id"`       //订单号
		SubOrderId    string            `json:"sub_order_id"`   //子订单号
		SkuList       []*CommentSkuItem `json:"sku_list"`
	}
	CommentSkuItem struct {
		SkuInfo       CommentSku          `json:"sku_info"`
		CommentLevel  string              `json:"comment_level"` //商品评分
		Mark          string              `json:"mark"`
		Videos        []*CommentVideoItem `json:"videos"`          //视频
		Images        []*CommentImageItem `json:"images"`          //图片
		ImageCount    int                 `json:"image_count"`     //图片数
		CanImageCount int                 `json:"can_image_count"` //总能上传数
	}

	CommentSku struct {
		SkuName      string        `json:"sku_name"`
		SkuProperty  string        `json:"sku_property"`
		ThumbnailURL string        `json:"thumbnail_url"`
		SkuId        string        `json:"sku_id"`
		SpuId        string        `json:"spu_id"`
		ShopId       int64         `json:"shop_id"`
		Price        string        `json:"price"`
		SaleType     uint8         `json:"sale_type"`
		SaleTypeName string        `json:"sale_type_name"`
		HaveGift     uint8         `json:"have_gift"`
		Gifts        []*CommentSku `json:"gifts,omitempty"` //赠品
		Href         string        `json:"href"`
	}
)

func NewCommentSkuItem() (res *CommentSkuItem) {
	res = &CommentSkuItem{
		Videos: make([]*CommentVideoItem, 0, CommentCanImageCount),
		Images: make([]*CommentImageItem, 0, CommentCanImageCount),
	}
	return
}

func (r *CommentSkuItem) DeferLogic() {
	r.ImageCount = len(r.Images) + len(r.Videos)
	r.CanImageCount = CommentCanImageCount
}
