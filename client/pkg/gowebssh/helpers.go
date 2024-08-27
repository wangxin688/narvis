package gowebssh

import (
	"bytes"
	"net/url"
	"strings"
)

func ByteContains(x, y []byte) (n []byte, contain bool) {
	index := bytes.Index(x, y)
	if index == -1 {
		return
	}
	lastIndex := index + len(y)
	n = append(x[:index], x[lastIndex:]...)
	return n, true
}

func UrlQueryUnescape(old string) (string, error) {
	// 客户端发送过来的数据是 url 编码过的，这里需要解码
	// url.QueryUnescape 会将'+'加号转换为' '空格。
	// 必须先替换 % ，再替换 +
	return url.QueryUnescape(strings.ReplaceAll(strings.ReplaceAll(old, "%", "%25"), "+", "%2b"))
}
