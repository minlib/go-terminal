package tar

import (
	"fmt"
	"testing"
)

func TestTar(t *testing.T) {
	err := Tar("C:\\Users\\Admin\\Desktop\\bean.tar.gz", "C:\\Users\\Admin\\Desktop\\bean", "C:\\Users\\Admin\\Desktop\\json.txt")
	//err := Tar("C:\\Users\\Admin\\Desktop\\bean3.tar.gz", "C:\\Users\\Admin\\Desktop\\json.txt")
	//err := Tar("C:\\Users\\Admin\\Desktop\\bean.tar.gz", "C:\\Users\\Admin\\Desktop\\bean")
	fmt.Println(err)
}

func TestUnTar(t *testing.T) {
	err := UnTar("C:\\Users\\Admin\\Desktop\\bean.tar.gz", "C:\\Users\\Admin\\Desktop\\bean5")
	fmt.Println(err)
}
