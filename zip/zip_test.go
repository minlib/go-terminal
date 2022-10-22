package zip

import (
	"fmt"
	"testing"
)

func TestZip(t *testing.T) {
	err := Zip("D:\\temp\\bean.zip", "D:\\temp\\bean", "D:\\temp\\json.txt")
	if err == nil {
		fmt.Println("Zip successful.")
	} else {
		fmt.Println("Zip failure,", err)
	}
}

func TestUnzip(t *testing.T) {
	err := Unzip("D:\\temp\\bean.zip", "D:\\temp\\bean_zip")
	if err == nil {
		fmt.Println("Unzip successful.")
	} else {
		fmt.Println("Unzip failure,", err)
	}
}
