package tar

import (
	"fmt"
	"testing"
)

func TestTar(t *testing.T) {
	err := Tar("D:\\temp\\bean.tar.gz", "D:\\temp\\bean", "D:\\temp\\json.txt")
	if err == nil {
		fmt.Println("Tar successful.")
	} else {
		fmt.Println("Tar failure,", err)
	}
}

func TestUntar(t *testing.T) {
	err := Untar("D:\\temp\\bean.tar.gz", "D:\\temp\\bean_tar")
	if err == nil {
		fmt.Println("Untar successful.")
	} else {
		fmt.Println("Untar failure,", err)
	}
}
