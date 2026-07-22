package sns

import (
	"golang.org/x/net/html"
	"strings"
)

// ATag 存储 a 标签的文本和链接
type ATag struct {
	DataId string // 链接地址
	Text   string // 标签内文本
}

// 从HTML字符串中提取所有a标签
func ExtractATags(htmlStr string) (tagList []ATag, err error) {
	// 1. 将字符串转为HTML解析器读取流
	var doc *html.Node
	if doc, err = html.Parse(strings.NewReader(htmlStr)); err != nil {
		return
	}

	// 2. 递归遍历DOM节点
	var traverse func(*html.Node)
	traverse = func(node *html.Node) {
		// 判断是否为 a 标签节点
		if node.Type == html.ElementNode && node.Data == "a" {
			var tag ATag
			// 提取 href 属性
			for _, attr := range node.Attr {
				if attr.Key == "data-id" {
					tag.DataId = attr.Val
					break
				}
			}
			// 提取标签内的纯文本
			tag.Text = GetNodeText(node)
			tagList = append(tagList, tag)
		}

		// 递归遍历子节点
		for child := node.FirstChild; child != nil; child = child.NextSibling {
			traverse(child)
		}
	}

	traverse(doc)
	return
}

// 获取节点内的所有纯文本（递归拼接子节点文本）
func GetNodeText(node *html.Node) string {
	if node.Type == html.TextNode {
		return node.Data
	}
	var text string
	for child := node.FirstChild; child != nil; child = child.NextSibling {
		text += GetNodeText(child)
	}
	return strings.TrimSpace(text) // 去除首尾空白
}
