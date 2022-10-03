package zip

import (
	"fmt"
	"testing"
)

func TestZip(t *testing.T) {
	err := Zip("C:\\Users\\Administrator\\Desktop\\bean.zip", "C:\\Users\\Administrator\\Desktop\\json.txt")
	if err == nil {
		fmt.Println("压缩成功")
	} else {
		fmt.Println("压缩失败，", err)
	}
}

func TestUnzip(t *testing.T) {
	err := Unzip("C:\\Users\\Administrator\\Desktop\\bean.zip", "C:\\Users\\Administrator\\Desktop\\bean")
	if err == nil {
		fmt.Println("解压成功")
	} else {
		fmt.Println("解压失败，", err)
	}
}
