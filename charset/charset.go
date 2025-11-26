package charset

import (
	"golang.org/x/text/encoding/simplifiedchinese"
)

type Charset string

const (
	GB18030 = Charset("GB18030")
	GBK     = Charset("GBK")
	UTF8    = Charset("UTF-8")
)

// BytesToString 二进制数组转成指定字符编码的字符串
func BytesToString(b []byte, charset Charset) string {
	var bytes []byte
	switch charset {
	case GB18030:
		bytes, _ = simplifiedchinese.GB18030.NewDecoder().Bytes(b)
	case GBK:
		bytes, _ = simplifiedchinese.GBK.NewDecoder().Bytes(b)
	case UTF8:
		fallthrough
	default:
		bytes = b
	}
	return string(bytes)
}
