package app_param

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/juetun/base-wrapper/lib/base"
)

type (
	ControllerExcelImport interface {
		ExcelImportHeaderRelate(c *gin.Context)

		//excel导入的参数校验
		ExcelImportValidate(c *gin.Context)

		//数据同步
		ExcelImportSyncData(c *gin.Context)
	}

	//Excel导入服务需要定义的接口，对应服务上需要实现这些方法和调用接口
	ServiceExcelImport interface {
		//excel导入的header关系
		ExcelImportHeaderRelate(args *ArgExcelImportHeaderRelate) (res *ResultExcelImportHeaderRelate, err error)

		//excel导入的参数校验
		ExcelImportValidate(args *ArgExcelImportValidateAndSync) (res []*ExcelImportDataItem, err error)

		//数据同步
		ExcelImportSyncData(args *ArgExcelImportValidateAndSync) (res []*ExcelImportDataItem, err error)
	}

	ArgExcelImportHeaderRelate struct {
		Scene string `json:"scene" form:"scene"`
	}
	ResultExcelImportHeaderRelate struct {
		Desc             string         `json:"desc"`    //文案说明
		SheetHeadersInfo []*SheetHeader `json:"headers"` //每个sheet的表头信息
	}
	SheetHeader struct {
		SheetIndex int64                          `json:"sheet_index"` //sheet的序号 从0 开始
		SheetName  string                         `json:"sheet_name"`  //sheetName 名称 默认Sheet1
		Headers    []*ExcelImportHeaderRelateItem `json:"headers"`     //表头信息
	}
	ArgExcelImportValidateAndSync struct {
		Scene             string                 `json:"scene" form:"scene"`
		Data              []*ExcelImportDataItem `json:"data" form:"data"`
		CurrentUserShopId int64                  `json:"current_user_shop_id" form:"current_user_shop_id"`
		CurrentUid        int64                  `json:"current_uid" form:"current_uid"`
		TimeNow           base.TimeNormal        `json:"time_now" form:"time_now"`
	}
	ExcelImportHeaderRelateItem struct {
		Type            string          `json:"type,omitempty"`       //列类型，可选值为  index、selection、expand、html
		Title           string          `json:"title,omitempty"`      //列中文标题
		ColumnKey       string          `json:"column_key,omitempty"` //字段的 key
		Index           int64           `json:"index,omitempty"`      //列序号 如:第一列：0, 第二列：1
		ClassName       string          `json:"className,omitempty"`  //列的样式名称
		Width           int             `json:"width,omitempty"`      //列宽
		MinWidth        int             `json:"minWidth,omitempty"`   //最小列宽
		MaxWidth        int             `json:"maxWidth,omitempty"`   //最大列宽
		Align           string          `json:"align,omitempty"`      //显示位置
		Tooltip         bool            `json:"tooltip,omitempty"`    //开启后，文本将不换行，超出部分显示为省略号，并用 Tooltip 组件显示完整内容
		ValidateHandler ValidateHandler `json:"-"`
	}
	ValidateHandler     func(header *ExcelImportHeaderRelateItem, data *ExcelImportDataItem) (errorItem *ImportErrMsgInfo)
	ExcelImportDataItem struct {
		Id             int64             `gorm:"column:id" json:"id,omitempty"`
		Line           int64             `gorm:"column:line" json:"line,omitempty"`
		Data           string            `gorm:"column:data" json:"data,omitempty"`
		SheetName      string            `gorm:"column:sheet_name" json:"sheet_name,omitempty"`
		ValidateStatus uint8             `gorm:"-" json:"validate_status,omitempty"` //验证状态是否通过
		ErrMsg         string            `gorm:"-" json:"err_msg,omitempty"`         //错误信息提示
		DataMap        map[string]string `json:"-" gorm:"-"`
	}
	ImportErrMsgInfos []*ImportErrMsgInfo
	ImportErrMsgInfo  struct {
		ColumnKey      string `json:"k,omitempty"`
		ValidateStatus uint8  `json:"vs,omitempty"`
		Msg            string `json:"msg,omitempty"`
	}
)

const (
	ExcelImportDataValidateStatusInit       uint8 = iota //导入数据初始化
	ExcelImportDataValidateStatusOk                      //校验成功
	ExcelImportDataValidateStatusFailure                 //校验失败
	ExcelImportDataValidateStatusWarning                 //可忽略
	ExcelImportDataValidateStatusImportOk                //导入完成
	ExcelImportDataValidateStatusValidating              //校验中
	ExcelImportDataValidateStatusSyncing                 //同步中
)

var (
	SliceExcelImportDataValidateStatus = base.ModelItemOptions{
		{
			Value: ExcelImportDataValidateStatusInit,
			Label: "初始化",
		},
		{
			Value: ExcelImportDataValidateStatusOk,
			Label: "校验成功",
		},
		{
			Value: ExcelImportDataValidateStatusFailure,
			Label: "校验失败",
		},
		{
			Value: ExcelImportDataValidateStatusWarning,
			Label: "可忽略",
		},
		{
			Value: ExcelImportDataValidateStatusImportOk,
			Label: "导入完成",
		},
		{
			Value: ExcelImportDataValidateStatusValidating,
			Label: "校验中",
		},
		{
			Value: ExcelImportDataValidateStatusSyncing,
			Label: "同步中",
		},
	}
)

func (r *ExcelImportDataItem) SetErrMsg(importErrMsgInfos ImportErrMsgInfos) (err error) {
	if len(importErrMsgInfos) == 0 {
		r.ErrMsg = ""
		return
	}
	r.ErrMsg, err = importErrMsgInfos.ToString()
	return
}

func (r *ExcelImportDataItem) GetDataMap() (dataValue map[string]string, err error) {
	if r.Data != "" {
		if err = json.Unmarshal([]byte(r.Data), &dataValue); err != nil {
			return
		}
	}
	return
}

func (r *ImportErrMsgInfos) ToString() (res string, err error) {
	if r == nil {
		res = ""
	}
	var bt []byte
	if bt, err = json.Marshal(r); err != nil {
		return
	}
	res = string(bt)
	return
}

func (r *SheetHeader) GetSheetHeaderMap() (res map[string]*ExcelImportHeaderRelateItem) {
	res = make(map[string]*ExcelImportHeaderRelateItem, len(r.Headers))
	for _, item := range r.Headers {
		res[item.Title] = item
	}
	return
}

func (r *ArgExcelImportValidateAndSync) Default(ctx *base.Context) (err error) {
	return
}

func (r *ArgExcelImportValidateAndSync) ToJson() (res []byte, err error) {
	if r == nil {
		r = &ArgExcelImportValidateAndSync{}
	}

	res, err = json.Marshal(r)
	return
}

func (r *ExcelImportDataItem) GetId() (res int64) {
	return r.Id
}

func ExcelImportHeaderRelate(c *gin.Context, srv ServiceExcelImport) (data *ResultExcelImportHeaderRelate, err error) {
	var (
		arg ArgExcelImportHeaderRelate
	)
	data = &ResultExcelImportHeaderRelate{}
	if err = c.Bind(&arg); err != nil {
		return
	}
	if data, err = srv.ExcelImportHeaderRelate(&arg); err != nil {
		return
	}
	return
}

func ExcelImportValidate(c *gin.Context, srv ServiceExcelImport) (data []*ExcelImportDataItem, err error) {

	var (
		arg ArgExcelImportValidateAndSync
	)
	data = make([]*ExcelImportDataItem, 0)
	if err = c.Bind(&arg); err != nil {
		return
	}
	if arg.TimeNow.IsZero() {
		arg.TimeNow = base.GetNowTimeNormal()
	}
	if data, err = srv.ExcelImportValidate(&arg); err != nil {
		return
	}
	return
}

func ExcelImportSyncData(c *gin.Context, srv ServiceExcelImport) (data []*ExcelImportDataItem, err error) {
	var (
		arg ArgExcelImportValidateAndSync
	)
	data = []*ExcelImportDataItem{}
	if err = c.Bind(&arg); err != nil {
		return
	}
	if arg.TimeNow.IsZero() {
		arg.TimeNow = base.GetNowTimeNormal()
	}
	if data, err = srv.ExcelImportSyncData(&arg); err != nil {
		return
	}
	return
}
