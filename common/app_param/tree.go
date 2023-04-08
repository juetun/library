package app_param

type (
	TreeOption struct {
		Title           string        `json:"title"`           //标题
		Value           string        `json:"value"`           //值
		Expand          bool          `json:"expand"`          //是否展开直子节点
		Disabled        bool          `json:"disabled"`        //禁掉响应
		DisableCheckbox bool          `json:"disableCheckbox"` //禁掉 checkbox
		Selected        bool          `json:"selected"`        //是否选中子节点
		Checked         bool          `json:"checked"`         //是否勾选(如果勾选，子节点也会全部勾选)
		Contextmenu     bool          `json:"contextmenu"`     //是否支持右键菜单
		Indeterminate   bool          `json:"indchk"`          //设置 indeterminate 状态，只负责样式控制 (自定义选择时使用)
		Children        []*TreeOption `json:"children"`        //子节点属性数组
	}
)
