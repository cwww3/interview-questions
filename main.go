package main

import (
	"strings"
)

func main() {
	isUniq("111")
}

func isUniq(str string) bool {
	// 返回 rune类型长度加1
	if strings.Count(str, "") > 3000 {
		return false
	}
	// 标准ASCII码0-127  127-255为扩展ASCII码
	for _, s := range str {
		if s > 127 {
			return false
		}
		// 通过strings.Count判断是否重复
		if strings.Count(str, string(s)) > 1 {
			return false
		}
		// 通过 index判断是否重复
		if strings.Index(str, string(s)) != strings.LastIndex(str, string(s)) {
			return false
		}
	}
	return true
}
