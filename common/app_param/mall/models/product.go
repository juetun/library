package models

import (
	"encoding/json"
	"fmt"
	"github.com/juetun/library/common/app_param/upload_operate"
	"github.com/juetun/library/common/const_apply"
	"strconv"
	"strings"
	"time"

	"github.com/juetun/base-wrapper/lib/base"
	"gorm.io/gorm"
)

const (
	TablePrefix = "mall_" // 表明前缀
)

const DefaultImageShow = "//dev-file.iviewui.com/userinfoPDvn9gKWYihR24SpgC319vXY8qniCqj4/avatar"
const SpuImageDivide = "#$#"
const (
	ProductStatusTmp           int8 = iota - 1 //草稿中(ID初始化中)
	ProductStatusAll                           //全部数据
	ProductStatusOnline                        // 在售
	ProductStatusManuscript                    // 草稿中
	ProductStatusInit                          // 仓库中
	ProductStatusOffLine                       // 已下架
	ProductStatusOnlineAtTime                  // 定时上架
	ProductStatusDeprecated                    // 删除
	ProductStatusWaitingOnline                 // 待上架
	ProductStatusWaitingDeal                   // 待处理

)
const (
	ProductHasOnlineYes uint8 = iota + 1 //商品曾上过架
	ProductHasOnlineNo                   //商品未曾上过架
)
const (
	RelateTypeMall uint8 = 1 // 商品关联类型 其他电商商品
)

const (
	SaleTypeGeneral     uint8 = iota + 1 // 普通商品
	SaleTypePreSale                      // SaleTypePreSale 全款预售
	SaleTypeDown                         // 定金预售
	SaleTypeActivity                     // 活动信息
	SaleTypeIntentional                  // 意向金预售
)

const (
	SettleTypeCurrent uint8 = iota + 1 // 现结
	SettleTypeMonth                    // 月结
)

//商品在前端是否展示（非正常显示逻辑除外 如：草稿中数据,详情页展示通过特殊逻辑绕开）
const (
	ProductShopTypeNotShow = "not_show" //不展示
	ProductShopTypeShow    = "show"     //展示
)
const (
	FreightTypeExpressDelivery uint8 = iota + 1 //快递
	FreightTypeLogistics                        //物流
	FreightTypeElectronic                       //电子凭证
)

const (
	SpuHaveGiftYes uint8 = iota + 1
	SpuHaveGiftNo
)
const (
	SpuSupportCommentYes uint8 = iota + 1
	SpuSupportCommentNo
)
const (
	DeliveryTimeTypeTime uint8 = iota + 1 //固定时间
	DeliveryTimeTypeDay                   //固定天数 （付款时间开始延迟天数发货）
)
const (
	DeliveryTimeTypeDayDefault int64 = 7 //默认延迟天发货
)

//定金预售最多可延迟支付时间的范围
const DownPayDelayPayLimit = 5 * time.Minute

const (
	ProductFreightNeedYes uint8 = iota + 1
	ProductFreightNeedNo
	ProductFreightVirtual
)

const (
	SupportFreightYes uint8 = iota + 1 //支持
	SupportFreightNo
)

var (
	SliceProductFreightNeed = base.ModelItemOptions{
		{
			Value: ProductFreightNeedYes, //需要实物
			Label: "需发货",
		},
		{
			Value: ProductFreightVirtual, //虚拟发货
			Label: "虚拟发货",
		},
		{
			Value: ProductFreightNeedNo, //不需要实物
			Label: "勿需发货",
		},
	}
	SliceDeliveryTimeType = base.ModelItemOptions{
		{
			Value: DeliveryTimeTypeTime, //固定时间
			Label: "固定时间",
		},
		{
			Value: DeliveryTimeTypeDay, //固定天数
			Label: "固定天数",
		},
	}
	SliceProductHasOnlineYes = base.ModelItemOptions{
		{
			Value: ProductHasOnlineYes, //上过架
			Label: "是",
		},
		{
			Value: ProductHasOnlineNo, //未曾上过架
			Label: "否",
		},
	}
	SliceSpuSupportComment = base.ModelItemOptions{
		{
			Value: SpuSupportCommentYes,
			Label: "支持",
		},
		{
			Value: SpuSupportCommentNo,
			Label: "不支持",
		},
	}
	SliceSpuHaveGift = base.ModelItemOptions{
		{
			Value: SpuHaveGiftYes,
			Label: "有赠品",
		},
		{
			Value: SpuHaveGiftNo,
			Label: "无赠品",
		},
	}
	SliceFreightType = base.ModelItemOptions{
		{
			Value: FreightTypeExpressDelivery,
			Label: "快递",
		},
		{
			Value: FreightTypeLogistics,
			Label: "物流",
		},
		{
			Value: FreightTypeElectronic,
			Label: "电子凭证",
		},
	}
	SliceSettleType = base.ModelItemOptions{
		{
			Value: SettleTypeCurrent,
			Label: "现结",
		},
		{
			Value: SettleTypeMonth,
			Label: "月结",
		},
	}
	SliceProductStatus = base.ModelItemOptions{
		{
			Value: ProductStatusManuscript,
			Label: "编辑中...",
		},
		{
			Value: ProductStatusOnline,
			Label: "出售中",
		},
		{
			Value: ProductStatusTmp,
			Label: "ID初始化",
		},
		{
			Value: ProductStatusInit,
			Label: "仓库中",
		},
		{
			Value: ProductStatusWaitingOnline,
			Label: "待上架",
		},
		{
			Value: ProductStatusWaitingDeal,
			Label: "待处理",
		},
		{
			Value: ProductStatusOffLine,
			Label: "已下架",
		},
		{
			Value: ProductStatusOnlineAtTime,
			Label: "定时上架",
		},
	}
	SliceRelateType = base.ModelItemOptions{
		{
			Value: RelateTypeMall,
			Label: "电商",
		},
	}
	SliceSaleType = base.ModelItemOptions{
		{
			Value: SaleTypeGeneral,
			Label: "",
		},
		{
			Value: SaleTypePreSale,
			Label: "预售",
		},
		{
			Value: SaleTypeDown,
			Label: "定金预售",
		},
		{
			Value: SaleTypeIntentional,
			Label: "意向金",
		},
		{
			Value: SaleTypeActivity,
			Label: "商品活动",
		},
	}
)
var (
	ProductTagList = []ProductTag{
		{ID: 1, Label: "正品保证", DefaultChecked: true},
		{ID: 2, Label: "三年质保"},
		{ID: 3, Label: "七天无理由退换"},
		{ID: 4, Label: "极速退款"},
		{ID: 5, Label: "免举证退换货"},
		{ID: 6, Label: "赠保价险"},
		{ID: 7, Label: "延保服务"},
		{ID: 8, Label: "赠运费险"},
		{ID: 9, Label: "闪电到家"},
	}
	SpuStatusTabsSlice = SpuStatusTabs{
		{
			Label:     "仓库中",
			Type:      ProductStatusInit,
			Status:    []int8{ProductStatusInit},
			ShowCount: true,
		},
		{
			Label:     "出售中", //包括在售和定时上架
			Type:      ProductStatusOnline,
			Status:    []int8{ProductStatusOnline, ProductStatusOnlineAtTime},
			ShowCount: false,
		},
		{
			Label:     "已下架",
			Type:      ProductStatusOffLine,
			Status:    []int8{ProductStatusOffLine},
			ShowCount: false,
		},
		{
			Label:     "定时上架",
			Type:      ProductStatusOnlineAtTime,
			Status:    []int8{ProductStatusOnlineAtTime},
			ShowCount: true,
		},
		//{
		//	Label:     "编辑中",
		//	Type:      ProductStatusManuscript,
		//	Status:    []int8{ProductStatusManuscript},
		//	ShowCount: true,
		//},
		{
			Label: "全部商品",
			Type:  ProductStatusAll,
			Status: []int8{
				//ProductStatusManuscript,   //草稿中
				ProductStatusInit,         // 仓库中
				ProductStatusOnline,       // 在售
				ProductStatusOffLine,      // 已下架
				ProductStatusOnlineAtTime, // 定时上架},
			},
		},
	}
	SpuAdminStatusTabsSlice = SpuStatusTabs{
		{
			Label:     "出售中",
			Type:      ProductStatusOnline,
			Status:    []int8{ProductStatusOnline},
			ShowCount: false,
		},
		{
			Label:     "仓库中",
			Type:      ProductStatusInit,
			Status:    []int8{ProductStatusInit},
			ShowCount: true,
		},
		{
			Label:     "已下架",
			Type:      ProductStatusOffLine,
			Status:    []int8{ProductStatusOffLine},
			ShowCount: true,
		},
		{
			Label: "全部商品",
			Type:  ProductStatusAll,
			Status: []int8{
				//ProductStatusManuscript, //草稿中
				ProductStatusInit,         // 仓库中
				ProductStatusOnline,       // 在售
				ProductStatusOffLine,      // 已下架
				ProductStatusOnlineAtTime, // 定时上架},
			},
			ShowCount: false,
		},
	}
)

type (
	Product struct {
		ProductID              string                         `gorm:"column:id;primary_key;type:bigint(20);not null;default:0;comment:商品SPUID" json:"product_id,omitempty"`
		Title                  string                         `gorm:"column:title;type:varchar(200);not null;default:'';comment:标题" json:"title,omitempty"`
		UserHid                int64                          `gorm:"column:user_hid;default:0;index:idx_userHid,priority:1;type:bigint(20);not null;comment:发布人用户ID" json:"user_hid,omitempty"`
		Thumbnail              string                         `gorm:"column:thumbnail;type:varchar(255);not null;default:'';comment:封面图ID" json:"thumbnail,omitempty"`
		ThumbnailURL           string                         `gorm:"-" json:"thumbnail_url"`
		Image                  string                         `gorm:"column:image;type:varchar(800);not null;default:'';comment:图片json数组" json:"image,omitempty"`
		ImageURL               []*upload_operate.ProductImage `gorm:"-" json:"-"`
		BrandId                int64                          `gorm:"column:brand_id;type:bigint(20);not null;default:0;comment:品牌ID" json:"brand_id,omitempty"`
		Video                  string                         `gorm:"column:video;type:varchar(255);not null;default:'';comment:视频" json:"video,omitempty"`
		VideoURL               string                         `gorm:"-" json:"-" `
		ShopId                 int64                          `gorm:"column:shop_id;index:idx_shop_id,priority:1;default:0;type:bigint(20);not null;comment:店铺ID" json:"shop_id,omitempty"`
		Status                 int8                           `gorm:"column:status;index:idx_shop_id,priority:2;index:idx_time,priority:2;default:0;type:tinyint(2);not null;comment:状态" json:"status,omitempty"`
		SubTitle               string                         `gorm:"column:sub_title;type:varchar(800);not null;default:'';comment:副标题" json:"sub_title,omitempty"`
		MinPrice               string                         `gorm:"column:min_price;index:idx_price,priority:1;default:0;type:decimal(10,2);not null;comment:最低价" json:"min_price,omitempty"`
		MaxPrice               string                         `gorm:"column:max_price;index:idx_price,priority:2;default:0;type:decimal(10,2);not null;comment:最高价" json:"max_price,omitempty"`
		MinPriceCost           string                         `gorm:"column:min_price_cost;default:0;type:decimal(10,2);not null;comment:最低市场价" json:"min_price_cost,omitempty"`
		MaxPriceCost           string                         `gorm:"column:max_price_cost;default:0;type:decimal(10,2);not null;comment:最高市场价" json:"max_price_cost,omitempty"`
		MinDownPayment         string                         `gorm:"column:min_down_payment;default:0;type:decimal(10,2);not null;comment:定金最低价" json:"min_down_payment,omitempty"`
		MaxDownPayment         string                         `gorm:"column:max_down_payment;default:0;type:decimal(10,2);not null;comment:定金最高价" json:"max_down_payment,omitempty"`
		MinDownPaymentDecr     string                         `gorm:"column:min_down_payment_decr;default:0;type:decimal(10,2);not null;comment:定金最低抵扣" json:"min_down_payment_decr,omitempty"`
		MaxDownPaymentDecr     string                         `gorm:"column:max_down_payment_decr;default:0;type:decimal(10,2);not null;comment:定金最高价抵扣" json:"max_down_payment_decr,omitempty"`
		IntentionalMin         string                         `gorm:"column:intentional_min;default:0;type:decimal(10,2);not null;comment:意向金最低额" json:"intentional_min,omitempty"`
		IntentionalMax         string                         `gorm:"column:intentional_max;default:0;type:decimal(10,2);not null;comment:意向金最高额" json:"intentional_max,omitempty"`
		IntentionalDecrMin     string                         `gorm:"column:intentional_decr_min;default:0;type:decimal(10,2);not null;comment:意向金最低抵扣" json:"intentional_decr_min,omitempty"`
		IntentionalDecrMax     string                         `gorm:"column:intentional_decr_max;default:0;type:decimal(10,2);not null;comment:意向金最高抵扣" json:"intentional_decr_max,omitempty"`
		TagIds                 string                         `gorm:"column:tag_ids;type:varchar(300);not null;default:'';comment:标签数据" json:"tag_ids,omitempty"`
		TagIdsArray            []int64                        `json:"-" gorm:"-"`
		TagAttr                string                         `gorm:"column:tag_attr;type:varchar(600);not null;default:'';comment:属性标签数据" json:"tag_attr,omitempty"`
		ServiceIds             string                         `gorm:"column:service_ids;type:varchar(300);not null;default:'';comment:支持服务列表" json:"service_ids,omitempty"`
		Keywords               string                         `gorm:"column:keywords;type:varchar(300);not null;default:'';comment:关键词" json:"keywords,omitempty"`
		SaleNum                int                            `gorm:"column:sale_num;type:bigint(20);not null;default:0;comment:销量(数据可能不及时)" json:"sale_num,omitempty"`
		FreightNeed            uint8                          `gorm:"column:freight_need;type:tinyint(2);default:1;not null;comment:是否需要发实物1-需要 2-不需要" json:"freight_need,omitempty"`
		FreightType            uint8                          `gorm:"column:freight_type;type:tinyint(2);default:1;not null;comment:快递方式 1-快递 2-EMS" json:"freight_type,omitempty"`
		FreightTemplate        int64                          `gorm:"column:freight_template;type:bigint(20);default:0;not null;comment:运费模板ID" json:"freight_template,omitempty"`
		DeliveryTimeType       uint8                          `gorm:"column:delivery_time_type;type:tinyint(2);default:2;not null;comment:发货时间类型 1-固定时间 2-固定天数" json:"delivery_time_type,omitempty"`
		DelayDays              int64                          `gorm:"column:delay_days;not null;type: bigint(15);default:7;comment:付款时间后延迟发货天数"  json:"delay_days,omitempty"`
		DeliveryTime           *base.TimeNormal               `gorm:"column:delivery_time;comment:预售预计发货时间" json:"delivery_time,omitempty"` // 预计发货时间
		TotalStock             int64                          `gorm:"column:total_stock;type:bigint(20);not null;default:0;comment:总库存数" json:"total_stock,omitempty,omitempty"`
		CategoryId             int64                          `gorm:"column:category_id;type:bigint(20);not null;default:0;comment:所属类目（二级类目）" json:"category_id,omitempty"`
		FinalCategoryId        int64                          `gorm:"column:final_category_id;type:bigint(20);not null;default:0;comment:所属类目(最终类目)" json:"final_category_id,omitempty"`
		SaleType               uint8                          `gorm:"column:sale_type;not null;type: tinyint(2);index:idx_time,priority:3;index:idx_price,priority:3;default:1;comment:销售类型1-普通商品 2-全款预售 3-定金预售"  json:"sale_type,omitempty"`
		PullOnTime             *base.TimeNormal               `gorm:"column:pull_on_time;index:idx_time,priority:1;index:idx_shop_id,priority:3;comment:定时上架时间" json:"pull_on_time,omitempty"`
		PullOffTime            *base.TimeNormal               `gorm:"column:pull_off_time;index:idx_time,priority:2;comment:定时下架时间" json:"pull_off_time,omitempty"`
		PullOffReason          string                         `gorm:"column:pull_off_reason;not null;type: varchar(80);default:'';comment:下架原因"   json:"pull_off_reason,omitempty"`
		IntentionalStime       base.TimeNormal                `gorm:"column:intentional_stime;not null;default:'2000-01-01 00:00:00';comment:意向金开售时间" json:"intentional_stime,omitempty"`
		IntentionalOtime       base.TimeNormal                `gorm:"column:intentional_otime;default:'2000-01-01 00:00:00';comment:意向金截止时间" json:"intentional_otime,omitempty"`
		SaleOnlineTime         base.TimeNormal                `gorm:"column:sale_online_time;not null;default:CURRENT_TIMESTAMP;comment:开售时间(可购买时间)" json:"sale_online_time,omitempty"`
		SaleOverTime           *base.TimeNormal               `gorm:"column:sale_over_time;comment:可购买截止时间" json:"sale_over_time,omitempty"`
		FinalStartTime         base.TimeNormal                `gorm:"column:final_start_time;not null;default:CURRENT_TIMESTAMP;comment:尾款开始时间" json:"final_start_time,omitempty"`
		FinalOverTime          base.TimeNormal                `gorm:"column:final_over_time;not null;default:CURRENT_TIMESTAMP;comment:尾款结束时间" json:"final_over_time,omitempty"`
		ShowInList             uint8                          `gorm:"column:show_in_list;type:tinyint(2);default:1;not null;comment:是否在推荐列表展示 1-展示(默认) 2-不展示" json:"show_in_list,omitempty"`
		SaleCountShow          uint8                          `gorm:"column:sale_count_show;type:bigint(20);not null;default:0;comment:销量超过数时展示销量" json:"sale_count_show,omitempty"`
		RelateType             uint8                          `gorm:"column:relate_type;not null;type: tinyint(1);default:0;comment:关联类型 0-无关联 1-电商"  json:"relate_type,omitempty"`
		RelateItemId           string                         `gorm:"column:relate_item_id;not null;type: varchar(80);default:'';comment:关联数据ID"   json:"relate_item_id,omitempty"`
		RelateBuyCount         int64                          `gorm:"column:relate_buy_count;not null;type: bigint(15);default:0;comment:关联购买人数"  json:"relate_buy_count,omitempty"`
		RelateBuyAMount        string                         `gorm:"column:relate_buy_amount;not null;type: decimal(15,2);default:0;comment:关联购买金额"  json:"relate_buy_amount,omitempty"`
		SettleType             uint8                          `gorm:"column:settle_type;not null;type: tinyint(2);default:1;comment:结算方式 1-现结 2-月结" json:"settle_type,omitempty"` // 结算方式 1：现结 2：月结
		JoinActivityId         string                         `gorm:"column:join_activity_id;not null;type: varchar(80);default:'';comment:加入活动的活动ID"   json:"join_activity_id,omitempty"`
		HaveGift               uint8                          `gorm:"column:have_gift;not null;type: tinyint(2);default:2;comment:是否有赠品 1-有赠品 2-无赠品" json:"have_gift,omitempty"` // 结算方式 1：现结 2：月结
		FlagTester             uint8                          `gorm:"column:flag_tester;not null;type: tinyint(2);default:1;comment:是否为测试数据 1-不是 2-是"  json:"flag_tester,omitempty"`
		SupportComment         uint8                          `gorm:"column:support_comment;not null;type: tinyint(2);default:1;comment:是否支持平路n 1-支持 2-不支持"  json:"support_comment,omitempty"`
		ApplyUserHid           int64                          `gorm:"column:apply_user_hid;default:0;type:bigint(20);not null;comment:审核人用户ID" json:"apply_user_hid,omitempty"`
		ApplyAt                base.TimeNormal                `gorm:"column:apply_at;not null;default:'2000-01-01 00:00:00';comment:审核时间" json:"apply_at,omitempty"`
		ApplyMark              string                         `gorm:"column:apply_mark;type:varchar(500);not null;default:'';comment:审核备注" json:"apply_mark,omitempty"`
		RecommendWeight        int64                          `gorm:"column:rec_weight;not null;type:bigint(20);default:10000;comment:商品推荐权重因子" json:"rec_weight" `
		CurrentRecommendWeight float64                        `gorm:"column:current_rec_weight;not null;type:decimal(22,2);default:0;comment:商品推荐权重" json:"current_rec_weight" `
		HasOnline              uint8                          `gorm:"column:has_online;not null;type: tinyint(2);default:2;comment:是否上架过 1-上架过 2-未上架"  json:"has_online,omitempty"`
		DataOtherAttr          string                         `gorm:"column:data_other_attr;type:varchar(1000);not null;default:'';comment:其他属性" json:"data_other_attr,omitempty"`
		CreatedAt              base.TimeNormal                `gorm:"column:created_at;not null;default:CURRENT_TIMESTAMP" json:"created_at,omitempty"`
		UpdatedAt              base.TimeNormal                `gorm:"column:updated_at;not null;default:CURRENT_TIMESTAMP" json:"updated_at,omitempty"`
		DeletedAt              *base.TimeNormal               `gorm:"column:deleted_at;" json:"-"`
	}
	ProductDataOtherAttr struct {
		IntentionalFreightNeed uint8 `json:"ifn"` //意向金是否需发货
		DownFreightNeed        uint8 `json:"dfn"` //定金是否需发货
	}
	ProductTag struct {
		ID             int64  `json:"id"`
		Label          string `json:"label"`
		DefaultChecked bool   `json:"default_checked"` //默认是否选中
	}
	ServiceItem struct {
		Value    int64  `json:"value"`
		Label    string `json:"label"`
		Disabled bool   `json:"disabled"`
		Checked  bool   `json:"checked"`
	}

	SpuStatusCount struct {
		Status int8  `json:"status" gorm:"column:status"`
		Count  int64 `json:"count" gorm:"column:count"`
	}
	SpuStatusTab struct {
		Label     string `json:"label"`
		Type      int8   `json:"type"`
		Count     int64  `json:"-"`
		ShowCount bool   `json:"show_count"` //是否展示数量
		Status    []int8 `json:"status"`
	}
	SpuStatusTabs []SpuStatusTab

	PageTag struct {
		Label     string `json:"label"`
		Plain     bool   `json:"plain"`
		TextColor string `json:"text_color"`
		Closable  bool   `json:"closable"`  //标签是否可以关闭
		Checkable bool   `json:"checkable"` //标签是否可以选择
		Checked   bool   `json:"checked"`   //标签的选中状态
		Type      string `json:"type"`      //	标签的样式类型，可选值为 border、dot或不填	String	-
		Color     string `json:"color"`     //标签颜色，预设颜色值为default、primary、success、warning、error、blue、green、red、yellow、pink、magenta、volcano、orange、gold、lime、cyan、geekblue、purple，你也可以自定义颜色值。	String	default
		Name      string `json:"name"`      //	当前标签的名称，使用 v-for，并支持关闭时，会比较有用	String | Number	-
		Size      string `json:"size"`      // large、medium、default
	}
	TitleTags []*PageTag
)

func (r *ProductDataOtherAttr) Default() {
	if r.DownFreightNeed == 0 {
		r.DownFreightNeed = ProductFreightNeedNo
	}
	if r.IntentionalFreightNeed == 0 {
		r.IntentionalFreightNeed = ProductFreightNeedNo
	}
	return
}

func (r SpuStatusTabs) GetMap() (res map[int8]SpuStatusTab) {
	res = make(map[int8]SpuStatusTab, len(r))
	for _, item := range r {
		res[item.Type] = item
	}
	return
}

func (r *Product) ParseProductDataOtherAttr() (res *ProductDataOtherAttr, err error) {
	if r.DataOtherAttr == "" {
		return
	}
	res = &ProductDataOtherAttr{}
	if err = json.Unmarshal([]byte(r.DataOtherAttr), res); err != nil {
		return
	}
	res.Default()
	return
}

func (r *Product) SetProductDataOtherAttr(productDataOtherAttr *ProductDataOtherAttr) (err error) {
	if productDataOtherAttr == nil {
		return
	}
	var bt []byte
	if bt, err = json.Marshal(productDataOtherAttr); err != nil {
		return
	}
	r.DataOtherAttr = string(bt)
	return
}

func (r *Product) DefaultBeforeAdd() {
	if r.MinDownPayment == "" {
		r.MinDownPayment = "0"
	}
	if r.MinPrice == "" {
		r.MinPrice = "0"
	}
	if r.MaxPrice == "" {
		r.MaxPrice = "0"
	}
	if r.MinPriceCost == "" {
		r.MinPriceCost = "0"
	}
	if r.MaxPriceCost == "" {
		r.MaxPriceCost = "0"
	}
	if r.MinDownPayment == "" {
		r.MinDownPayment = "0"
	}
	if r.MaxDownPayment == "" {
		r.MaxDownPayment = "0"
	}
	if r.MaxDownPayment == "" {
		r.MaxDownPayment = "0"
	}
	if r.RelateBuyAMount == "" {
		r.RelateBuyAMount = "0"
	}
	if r.FreightType == 0 {
		r.FreightType = FreightTypeExpressDelivery
	}
	if r.FlagTester == 0 {
		r.FlagTester = const_apply.FlagTesterNo
	}
	if r.SaleType == 0 {
		r.SaleType = SaleTypeGeneral
	}
	if r.Status == 0 || r.Status == ProductStatusManuscript {
		r.Status = ProductStatusInit
	}
	if r.HaveGift == 0 {
		r.HaveGift = SpuHaveGiftNo
	}
}

func (r *Product) GetPageTags() (res []*PageTag) {
	res = make([]*PageTag, 0, 3)
	saleTypeTag := GetPageTagsWithSaleType(r.SaleType)
	if saleTypeTag != nil {
		res = append(res, saleTypeTag)
	}

	testTag := GetPageTagsTester(r.FlagTester)
	if testTag != nil {
		res = append(res, testTag)
	}
	return
}

func GetPageTagsTester(FlagTester uint8) (dt *PageTag) {
	switch FlagTester {
	case const_apply.FlagTesterYes:
		dt = NewPageTag()
		dt.Label = "测试"
		dt.Color = "error"
	case 0:
		dt = NewPageTag()
		dt.Label = "是否测试异常"
		dt.Color = "error"
	}
	return
}

func GetProductFreightNeed(ProductFreightNeed uint8) (res string) {
	var ok bool
	MapRelateType, _ := SliceProductFreightNeed.GetMapAsKeyUint8()
	if res, ok = MapRelateType[ProductFreightNeed]; ok {
		return
	}
	return
}

func (r *Product) ParseProductFreightNeed() (res string) {
	res = GetProductFreightNeed(r.FreightNeed)
	return
}

func GetPageTagsWithSaleType(saleType uint8) (dt *PageTag) {
	switch saleType {
	case SaleTypeIntentional: //意向金
		mapSaleType, _ := SliceSaleType.GetMapAsKeyUint8()
		dt = NewPageTag()
		dt.Color = "info"
		dt.Label = mapSaleType[saleType]
	case SaleTypePreSale:
		mapSaleType, _ := SliceSaleType.GetMapAsKeyUint8()
		dt = NewPageTag()
		dt.Color = "success"
		dt.Label = mapSaleType[saleType]
	case SaleTypeDown:
		mapSaleType, _ := SliceSaleType.GetMapAsKeyUint8()
		dt = NewPageTag()
		dt.Label = mapSaleType[saleType]
		dt.Color = "warning"
	case SaleTypeGeneral:
		//dt = NewPageTag()
		//dt.Label = "普通商品"
		//dt.Color = "green"
	case 0:
		dt = NewPageTag()
		dt.Label = "销售类型异常"
		dt.Color = "error"
	}
	return
}
func (r *Product) ParseServiceItem() (res []ServiceItem, serviceIds []int64, err error) {
	res = make([]ServiceItem, 0, len(ProductTagList))
	serviceIds = make([]int64, 0, len(ProductTagList))
	var mapService map[int64]bool
	if r.ServiceIds != "" {
		if err = json.Unmarshal([]byte(r.ServiceIds), &serviceIds); err != nil {
			return
		}
		mapService = make(map[int64]bool, len(serviceIds))
		for _, it := range serviceIds {
			if it <= 0 {
				return
			}
			mapService[it] = true
		}
	}
	nullServiceId := len(serviceIds) == 0 //判断当前是否是初始化没有选中服务项目
	var serviceItem ServiceItem
	for _, item := range ProductTagList {
		serviceItem = ServiceItem{
			Value:   item.ID,
			Label:   item.Label,
			Checked: item.DefaultChecked, //如果默认必选的话，则该选择空间禁用
		}
		if item.DefaultChecked { //如果默认必选的话，则该选择空间禁用
			serviceItem.Disabled = true
			if nullServiceId {
				serviceIds = append(serviceIds, item.ID)
			}
		}
		if _, ok := mapService[item.ID]; ok {
			serviceItem.Checked = true
		}
		res = append(res, serviceItem)
	}

	return
}
func (r *Product) TableName() string {
	return fmt.Sprintf("%sproduct", TablePrefix)
}

func (r *Product) GetTableComment() (res string) {
	res = "商品SPU表"
	return
}

func (r *Product) GetSpuId() (spuId int64, err error) {
	spuId, err = strconv.ParseInt(r.ProductID, 10, 64)
	return
}

//当前详情页是否正常展示
func (r *Product) CanShowPreheat(currentTimes ...time.Time) (showFlag string, msg string) {
	var current time.Time
	if len(currentTimes) > 0 {
		current = currentTimes[0]
	} else {
		current = time.Now()
	}
	showFlag = ProductShopTypeShow
	r.defaultSaleOverTime()
	//如果没有设置预热开始时间，则默认选择开售时间
	// 如果预热开始时间晚于开售时间，则以开售时间为准
	if r.PullOnTime == nil || r.PullOnTime.Time.IsZero() || r.PullOnTime.Time.After(r.SaleOnlineTime.Time) {
		r.PullOnTime = &r.SaleOnlineTime
	}

	// 如果预热开始时间晚于开售时间，则以开售时间为准
	if r.PullOffTime == nil || r.PullOffTime.Time.IsZero() || r.PullOffTime.Time.After(r.SaleOnlineTime.Time) { //如果没有设置预热结束时间，则默认选择开售时间
		r.PullOffTime = r.SaleOverTime
	}

	//如果当前时间还没到预热开始时间
	if current.Before(r.PullOnTime.Time) {
		showFlag = ProductShopTypeNotShow
		msg = "未到预热开始时间"
		return
	}

	return
}

// 默认销售截止时间
func (r *Product) defaultSaleOverTime() {
	if r.SaleOverTime == nil {
		r.SaleOverTime = &base.TimeNormal{Time: r.SaleOnlineTime.Time.Add(100 * 24 * time.Hour)}
	}
}

// 根据状态，判断商品是否能够被购买
// Return pageMessage 界面展示不能购买原因
// Return systemMark 系统记录不能购买原因 用于记录日志使用
func (r *Product) JudgeCanBuyStatus(currentTimes ...time.Time) (ok bool, pageMessage, systemMark string) {
	current := r.getCurrentTime(currentTimes...)
	r.defaultSaleOverTime()
	switch r.Status {
	case ProductStatusOnline:
		if current.Before(r.SaleOnlineTime.Time) {
			systemMark = fmt.Sprintf("商品(SPU)在售时间不正确,商品暂不支持购买(SaleOnlineTime:%+v)", r.SaleOnlineTime)
			pageMessage = "商品暂不支持购买"
			return
		}
		if current.After(r.SaleOverTime.Time) {
			systemMark = fmt.Sprintf("商品(SPU)在售时间不正确,已超过在售时间(SaleOverTime:%+v)", r.SaleOverTime)
			pageMessage = "已超过在售时间"
			return
		}
		ok = true
		return
	default:
		systemMark = fmt.Sprintf("商品(SPU)状态不为在售中(%d)", r.Status)
		pageMessage = "商品暂不支持购买"
	}
	return
}

func (r *Product) getCurrentTime(currentTimes ...time.Time) (current time.Time) {
	if len(currentTimes) > 0 {
		current = currentTimes[0]
	} else {
		current = time.Now()
	}
	return
}

//判断商品当前时间是否满足能够购买
func (r *Product) JudgeCanBuyWithTime(currentTimes ...time.Time) (ok bool, msg string) {
	current := r.getCurrentTime(currentTimes...)

	//如果当前时间在销售时间范围内，则商品可以被购买
	if current.After(r.SaleOnlineTime.Time) && current.Before(r.SaleOverTime.Time) {
		ok = true
		return
	}
	msg = "商品不在可购买时间范围内"
	return
}

// ParseRelateType 获取关联关系类型
func (r *Product) ParseRelateType() (res string) {
	var ok bool
	MapRelateType, _ := SliceRelateType.GetMapAsKeyUint8()
	if res, ok = MapRelateType[r.RelateType]; ok {
		return
	}
	return
}

// ParseRelateType 是否有赠品
func (r *Product) ParseSpuHaveGifts() (res string) {
	return ParseSpuHaveGifts(r.HaveGift)
}

func ParseSpuHaveGifts(haveGift uint8) (res string) {
	var ok bool
	MapRelateType, _ := SliceSpuHaveGift.GetMapAsKeyUint8()
	if res, ok = MapRelateType[haveGift]; ok {
		return
	}
	return
}

func (r *Product) ParseSettleType() (res string) {
	var ok bool
	MapSettleType, _ := SliceSettleType.GetMapAsKeyUint8()
	if res, ok = MapSettleType[r.SettleType]; ok {
		return
	}
	return
}
func ParseSaleType(SaleType uint8) (res string) {
	var ok bool
	MapSaleType, _ := SliceSaleType.GetMapAsKeyUint8()
	if res, ok = MapSaleType[SaleType]; ok {
		return
	}
	res = fmt.Sprintf("未知类型(%d)", SaleType)
	return
}

func (r *Product) ParseSaleType() (res string) {
	return ParseSaleType(r.SaleType)
}

//SPU状态转换
func ProductParseStatus(status int8) (res string) {
	var ok bool
	MapProductStatus, _ := SliceProductStatus.GetMapAsKeyInt8()
	if res, ok = MapProductStatus[status]; ok {
		return
	}
	res = fmt.Sprintf("未知状态(%d)", status)
	return
}
func (r *Product) ParseStatus() (res string) {
	res = ProductParseStatus(r.Status)
	return
}

func (r *Product) ImageMarshal(images []string) {
	r.Image = strings.Join(images, SpuImageDivide)
}

func (r *Product) ImageUnmarshal() (res []string) {
	res = strings.Split(r.Image, SpuImageDivide)
	return
}

func (r *Product) SetImage(images string) {
	r.Image = images
}

func (r *Product) SetVideo(video *upload_operate.UploadVideo) {
	if video == nil {
		return
	}
	r.Video = video.ToString()
	return
}

func (r *Product) BeforeCreate(tx *gorm.DB) (err error) {
	_ = tx

	return
}

func (r *Product) BeforeUpdate(tx *gorm.DB) (err error) {
	_ = tx

	return
}

func (r *Product) UnmarshalBinary(data []byte) (err error) {
	err = json.Unmarshal(data, r)
	return
}

//实现 序列化方法 encoding.BinaryMarshaler
func (r *Product) MarshalBinary() (data []byte, err error) {
	data, err = json.Marshal(r)
	return
}
func (r *Product) AfterCreate(tx *gorm.DB) (err error) {

	return
}
func (r *Product) GetID() string {
	return r.ProductID
}
func (r *Product) SaltForHID() string {
	return r.TableName()
}

func (r *Product) SetTagIds(tagIds []int64) {
	r.TagIds = "[]"
	if len(tagIds) > 0 {
		bt, _ := json.Marshal(tagIds)
		r.TagIds = string(bt)
	}
}

func (r *Product) ParseTagIds() (tagIds []int64, err error) {
	if r.TagIds == "" {
		return
	}
	err = json.Unmarshal([]byte(r.TagIds), &r.TagIdsArray)
	return
}

//判断商品是否已经过了尾款支付时间
func (r *Product) JudgeSaleTypeDownFinalExpire(currentTimes ...time.Time) (ok bool) {
	currentTime := r.getCurrentTime(currentTimes...)
	if currentTime.After(r.FinalOverTime.Time) {
		ok = true
	}
	return
}

//判断当前商品是否能够支付尾款
func (r *Product) JudgeSaleTypeDownFinal(currentTimes ...time.Time) (notCanPay bool, pageMessage, systemMessage string) {
	notCanPay = true
	currentTime := r.getCurrentTime(currentTimes...)
	if currentTime.After(r.FinalStartTime.Time) && currentTime.Before(r.FinalOverTime.Time) {
		notCanPay = false
		return
	}

	if currentTime.Before(r.FinalOverTime.Time) {
		pageMessage = "未到支付尾款时间"
		systemMessage = fmt.Sprintf("current:%+v,FinalStartTime:%+v,FinalOverTime:%+v", currentTime, r.FinalStartTime, r.FinalOverTime)
		return
	}

	if r.JudgeSaleTypeDownFinalExpire(currentTime) {
		pageMessage = "尾款支付已过期"
		systemMessage = fmt.Sprintf("current:%+v,FinalStartTime:%+v,FinalOverTime:%+v", currentTime, r.FinalStartTime, r.FinalOverTime)
		return
	}
	return
}

func (r *Product) ParseFlagTester() (res string) {
	sliceFlagTester, _ := const_apply.SliceFlagTester.GetMapAsKeyUint8()
	if dt, ok := sliceFlagTester[r.FlagTester]; ok {
		res = dt
		return
	}
	res = fmt.Sprintf("未知(%d)", r.FlagTester)
	return
}

func (r *Product) GetDefaultThumbnail() (res string) {
	if r.Thumbnail == "" {
		r.Thumbnail = DefaultImageShow
	}
	res = r.Thumbnail
	return
}

//获取尾款结束时间
func (r *Product) GetFinalOverTime(payingFlags ...bool) (res base.TimeNormal) {
	var (
		payingFlag bool
	)
	if len(payingFlags) > 0 {
		payingFlag = payingFlags[0]
	}
	if payingFlag {
		res = r.FinalOverTime.Add(DownPayDelayPayLimit)
		return
	}
	res = r.FinalOverTime
	return
}

//定金结束时间
func (r *Product) GetFirstOverTime(payingFlags ...bool) (res base.TimeNormal) {
	var (
		payingFlag bool
	)
	if len(payingFlags) > 0 {
		payingFlag = payingFlags[0]
	}
	if payingFlag {
		res = r.SaleOverTime.Add(DownPayDelayPayLimit)
		return
	}
	res = *r.SaleOverTime
	return
}

//意向金结束时间
func (r *Product) GetIntentionalOTime(payingFlags ...bool) (res base.TimeNormal) {

	var (
		payingFlag bool
	)
	if len(payingFlags) > 0 {
		payingFlag = payingFlags[0]
	}
	if payingFlag {
		res = r.IntentionalOtime.Add(DownPayDelayPayLimit)
		return
	}
	res = r.IntentionalOtime
	return
}

func (r *Product) FlagCanUpdateStatus(wiLLUpdateStatus int8) (err error) {

	var StatusCanUpdateMap = map[int8]map[int8]bool{
		ProductStatusTmp: {ProductStatusManuscript: true,}, //草稿中(ID初始化中)
		ProductStatusManuscript: {
			ProductStatusInit:         true,
			ProductStatusOnline:       true,
			ProductStatusOnlineAtTime: true,
		}, // 草稿中

		ProductStatusInit:         {ProductStatusOnline: true,},                                                      // 仓库中
		ProductStatusOnline:       {ProductStatusOffLine: true,},                                                     // 在售
		ProductStatusOffLine:      {ProductStatusOnline: true, ProductStatusInit: true,},                             // 已下架
		ProductStatusOnlineAtTime: {ProductStatusOnline: true, ProductStatusInit: true, ProductStatusOffLine: true,}, // 定时上架
	}
	var (
		ok        bool
		statusMap map[int8]bool
	)
	if statusMap, ok = StatusCanUpdateMap[r.Status]; !ok {
		err = fmt.Errorf("当前不支持你选择的状态修改(F:%d,T:%d)", r.Status, wiLLUpdateStatus)
		return
	}
	if _, ok := statusMap[wiLLUpdateStatus]; !ok {
		err = fmt.Errorf("该商品状态暂不支持你要修改的状态,请刷新页面重试")
		return
	}

	return
}

func (r *Product) GetHaveVideo() (res bool) {
	if r.Video != "" {
		res = true
	}
	return
}

//定金预售尾款可支付时间有一个缓冲时间 当前暂定为5分钟
func (r *Product) DownPayCanPayFinal(timeNow base.TimeNormal, isDelays ...bool) (res bool) {
	isDelay := false
	if len(isDelays) > 0 {
		isDelay = isDelays[0]
	}

	//如果支持延迟验证
	if isDelay {
		//如果定金预售的开始时间在当前时间之前 且定金预售的结束时间(定金预售)在当前时间之后，则可用支付
		if r.FinalStartTime.Before(timeNow.Time) && timeNow.Time.Before(r.GetFinalOverTime(isDelay).Time) {
			res = true
		}
	} else {
		if r.FinalStartTime.Before(timeNow.Time) && r.FinalOverTime.After(timeNow.Time) {
			res = true
		}
	}

	return
}

func NewPageTag() (res *PageTag) {
	res = &PageTag{}
	res.Default()
	return
}

func (r *PageTag) Default() {
	r.Checked = true
	r.Size = "default"
	r.Type = "border"
}
