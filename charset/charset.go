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

// BytesToString 二进制数组转指定字符编码的字符串
// @byte 字节流
// @charset 字符编码
func BytesToString(byte []byte, charset Charset) string {
	var str string
	switch charset {
	case GB18030:
		var decodeBytes, _ = simplifiedchinese.GB18030.NewDecoder().Bytes(byte)
		str = string(decodeBytes)
	case GBK:
		var decodeBytes, _ = simplifiedchinese.GBK.NewDecoder().Bytes(byte)
		str = string(decodeBytes)
	case UTF8:
		fallthrough
	default:
		str = string(byte)
	}
	return str
}
