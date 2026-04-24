package sns

import (
	"bytes"
	"fmt"
	"github.com/juetun/base-wrapper/lib/base"
	"github.com/juetun/base-wrapper/lib/common"
	"github.com/juetun/library/common/app_param"
	"github.com/juetun/library/common/recommend"
	"golang.org/x/net/html"
	"strings"
)

const (
	EditOptionTypeRing = recommend.AdDataDataTypeRing
	EditOptionTypeUser = recommend.AdDataDataTypeUser
)

type (
	ArgRingsId struct {
		common.HeaderInfo
		app_param.RequestUser
		base.ArgGetByNumberIds
	}
	ResultRingMap  map[string]ResultRingItem
	EditOptionItem struct {
		ID    int64  `gorm:"column:id" json:"id"`
		Title string `gorm:"column:title" json:"title,omitempty"`
	}
	LoadEditorItem struct {
		Value  string `json:"value"`
		Label  string `json:"label"`
		Prefix string `json:"prefix"`
		Type   string `json:"type"`
	}
	ResultRingItem struct {
		ID             int64  `json:"id"`
		Title          string `json:"title,omitempty"`
		Thumbnail      string `json:"thumbnail,omitempty"`
		ThumbnailURL   string `json:"thumbnail_url"`
		ShowLevel      uint8  `json:"show_level"`
		IsRecommend    uint8  `json:"is_recommend"` //是否推荐圈子
		Key            string `json:"key,omitempty"`
		PingYin        string `json:"ping_yin,omitempty"`
		PingYinFirst   string `json:"ping_yin_first,omitempty"`
		Description    string `json:"description"`
		SubDescription string `json:"sub_description"`
		Status         int8   `json:"status,omitempty"`
		UserHid        int64  `json:"user_hid"`
	}
)

func (r *LoadEditorItem) ParseFromEditOptionItem(option *EditOptionItem, optionType string) {
	r.Value = option.GetPk(optionType)
	r.Label = option.Title
	r.Type = optionType
	switch optionType {
	case EditOptionTypeRing:
		r.Prefix = "#"
	case EditOptionTypeUser:
		r.Prefix = "@"
	}
	return
}

func (r *EditOptionItem) GetPk(optionType string) (res string) {
	res = fmt.Sprintf("%v|%v", optionType, r.ID)
	return
}

func (r *ArgRingsId) Default(ctx *base.Context) (err error) {
	return
}

// 遍历 DOM 树，移除 a 标签的 contenteditable="false"
func removeATagEditable(n *html.Node) {
	// 只处理 a 元素
	if n.Type == html.ElementNode {
		switch n.Data {
		case "a":
			newAttr := make([]html.Attribute, 0, len(n.Attr))
			for _, attr := range n.Attr {
				// 移除 contenteditable="false"
				if !(attr.Key == "contenteditable" && strings.EqualFold(attr.Val, "false")) {
					newAttr = append(newAttr, attr)
				}
			}
			n.Attr = newAttr
		case "img":
			for i := range n.Attr {
				if n.Attr[i].Key == "src" {
					n.Attr[i].Val = "" // 直接设为空字符串
				}
			}
		}
	}

	// 递归子节点
	for child := n.FirstChild; child != nil; child = child.NextSibling {
		removeATagEditable(child)
	}
}

// 找到 <body> 节点
func findBody(n *html.Node) *html.Node {
	if n.Type == html.ElementNode && n.Data == "body" {
		return n
	}
	for child := n.FirstChild; child != nil; child = child.NextSibling {
		if b := findBody(child); b != nil {
			return b
		}
	}
	return nil
}

// 处理 HTML：移除 a 标签的 contenteditable="false"，并返回 body 内部内容（不带 body 标签）
func ProcessHTML(htmlStr string) (res string, err error) {
	doc, err := html.Parse(strings.NewReader(htmlStr))
	if err != nil {
		return
	}

	// 1. 先修改所有 a 标签
	removeATagEditable(doc)

	var body *html.Node
	// 2. 找到 body 节点
	if body = findBody(doc); body == nil {
		body = doc // 没有 body 就用整个文档
	}

	// 3. 只渲染 body 内部内容，不渲染 <body> 标签
	var buf bytes.Buffer
	var child *html.Node
	for child = body.FirstChild; child != nil; child = child.NextSibling {
		if err = html.Render(&buf, child); err != nil {
			return
		}
	}
	res = buf.String()
	return
}

// 给所有 a 标签添加 contenteditable="false"
func addContentEditableToA(n *html.Node) {
	// 只处理 a 标签
	if n.Type == html.ElementNode && n.Data == "a" {
		hasEditable := false
		// 遍历现有属性，检查是否已经有 contenteditable
		for i, attr := range n.Attr {
			if attr.Key == "contenteditable" {
				// 如果已经存在，直接强制设为 false
				n.Attr[i].Val = "false"
				hasEditable = true
				break
			}
		}
		// 如果不存在，新增属性
		if !hasEditable {
			n.Attr = append(n.Attr, html.Attribute{
				Key: "contenteditable",
				Val: "false",
			})
		}
	}

	// 递归遍历子节点
	for child := n.FirstChild; child != nil; child = child.NextSibling {
		addContentEditableToA(child)
	}
}

// 处理 HTML：给所有 a 标签添加 contenteditable="false"，并返回 body 内部内容（不带外层标签）
func AddProcessHTML(htmlStr string) (string, error) {
	// 解析 HTML
	doc, err := html.Parse(strings.NewReader(htmlStr))
	if err != nil {
		return "", err
	}

	// 给所有 a 标签添加 contenteditable="false"
	addContentEditableToA(doc)

	// 找到 body，只输出内部内容
	body := findBody(doc)
	if body == nil {
		body = doc
	}

	// 渲染结果（不带 <html> <body>）
	var buf bytes.Buffer
	for child := body.FirstChild; child != nil; child = child.NextSibling {
		if err := html.Render(&buf, child); err != nil {
			return "", err
		}
	}

	return buf.String(), nil
}
