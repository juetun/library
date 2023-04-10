package app_param

type (
	TreeOption struct {
		Title           string         `json:"title"`           //标题
		Value           string         `json:"value"`           //值
		Expand          bool           `json:"expand"`          //是否展开直子节点
		Disabled        bool           `json:"disabled"`        //禁掉响应
		DisableCheckbox bool           `json:"disableCheckbox"` //禁掉 checkbox
		Selected        bool           `json:"selected"`        //是否选中子节点
		Checked         bool           `json:"checked"`         //是否勾选(如果勾选，子节点也会全部勾选)
		Level           int            `json:"level"`           //级别
		Contextmenu     bool           `json:"contextmenu"`     //是否支持右键菜单
		Indeterminate   bool           `json:"indchk"`          //设置 indeterminate 状态，只负责样式控制 (自定义选择时使用)
		Tags            []*DataItemTag `json:"tags,omitempty"`
		Children        []*TreeOption  `json:"children"` //子节点属性数组
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
)
