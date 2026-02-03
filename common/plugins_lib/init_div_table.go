package plugins_lib

import (
	"bytes"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/juetun/base-wrapper/lib/base"
	"github.com/juetun/base-wrapper/lib/base/sub_treasury_impl"
	"html/template"
	"net/http"
)

func GetAppDivTable(c *gin.Context) {
	var options = base.ModelItemOptions{}
	for _, item := range sub_treasury_impl.DiffDbAndTableConfig {
		options = append(options, base.ModelItemOption{
			Value: item.Struct.TableName(),
			Label: item.Struct.GetTableComment(),
		})
	}
	c.JSON(http.StatusOK, base.Result{Code: base.SuccessCode, Data: options, Msg: ""})
	return

}

func InitDivSql(c *gin.Context) {
	var (
		err     error
		options = ""
		arg     struct {
			App       string `json:"app" form:"app"`
			TableName string `json:"table_name" form:"table_name"`
			BaseSql   string `json:"base_sql" form:"base_sql"`
		}
		dbNameSpace       []string
		tableNum          int64
		sql, nowTableName string
	)
	if err = c.Bind(&arg); err != nil {
		return
	}
	for _, item := range sub_treasury_impl.DiffDbAndTableConfig {
		nowTableName = item.Struct.TableName()
		if nowTableName != arg.TableName {
			continue
		}

		dbNameSpace, tableNum = item.Struct.GetDBAndTableNumber()
		for _, dbName := range dbNameSpace {
			var i int64
			for i = 0; i < tableNum; i++ {
				if sql, err = renderDivTable(arg.BaseSql, dbName, fmt.Sprintf("%s%d", nowTableName, i)); err != nil {
					return
				}
				options += fmt.Sprintf("%v<br/>", sql)
			}
		}

	}
	c.JSON(http.StatusOK, base.Result{Code: base.SuccessCode, Data: options, Msg: ""})
	return
}

func renderDivTable(tpl, dbName, tableName string) (res string, err error) {
	res = tpl
	// 创建一个新的模板
	t := template.New("test")
	if t, err = t.Parse(tpl); err != nil { // 解析模板字符串
		return
	}
	// 准备数据
	data := struct {
		DbName    string
		TableName string
	}{
		DbName:    dbName,
		TableName: tableName,
	}
	var tplResult bytes.Buffer
	// 执行模板，输出结果到标准输出（或任何io.Writer）
	if err = t.Execute(&tplResult, data); err != nil {
		return
	}
	res = tplResult.String()
	return
}
