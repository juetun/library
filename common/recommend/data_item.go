package recommend

type (
	DataItem struct {
		Title        string            `json:"title"`          //标题
		PreTags      []*DataItemTag    `json:"pre_tags"`       //前缀标签
		DataType     int8              `json:"data_type"`      //数据类型
		DataTypeName string            `json:"data_type_name"` //数据类型名称
		DataId       string            `json:"data_id"`        //数据ID
		HaveVideo    bool              `json:"have_video"`     //是否有视频
		Img          string            `json:"img"`            //头图
		Price        string            `json:"price"`          //价格
		DataValue    []*DataItemDetail `json:"data_value"`     //详情
		SuffixTags   []*DataItemTag    `json:"suffix_tags"`    //后缀标签
		ExtraMsg     string            `json:"extra_msg"`      // 携带的其他信息
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
