package recommend

type (
	DataItem struct {
		Title         string            `json:"title,omitempty"`       //标题
		PreTags       []*DataItemTag    `json:"pre_tags,omitempty"`    //前缀标签
		DataType      int8              `json:"data_type"`             //数据类型
		DataId        string            `json:"data_id"`               //数据ID
		HaveVideo     bool              `json:"have_video,omitempty"`  //是否有视频
		Img           string            `json:"img,omitempty"`         //头图
		Price         string            `json:"price"`                 //价格
		DataValue     []*DataItemDetail `json:"data_value,omitempty"`  //详情
		SuffixTags    []*DataItemTag    `json:"suffix_tags,omitempty"` //后缀标签
		ShowShop      bool              `json:"show_shop"`             //是否显示店铺名
		ShopName      string            `json:"shop_name,omitempty"`   //店铺名称
		CurrentAccUId int64             `json:"current_acc_uid"`       //获取数据的用户ID
		ShopId        int64             `json:"shop_id,omitempty"`     //店铺ID
		ExtraMsg      string            `json:"extra_msg"`             //携带的其他信息
	}
	DataItemTag struct {
		Label string `json:"label"` //类型名称
		Color string `json:"color"` //显示颜色
		Value string `json:"value"` //类型值
	}
	DataItemDetail struct {
		Label string `json:"label"` //类型名称
		Key   string `json:"key"`   //类型的KEY
		Color string `json:"color"` //显示颜色
		Value string `json:"value"` //类型值
	}
)
