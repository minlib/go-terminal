package ssh

import (
	"fmt"
	"testing"
)

func TestNewClient(t *testing.T) {
	// 虚拟机
	c := NewClient("192.168.116.137", "22", "root", "yuda")
	// 上传程序文件
	c.UploadFile("C:\\Users\\Admin\\Desktop\\main.tar.gz", "/home/app", "main.tar.gz")
	// 解压文件并删除压缩包
	tarDist, _ := c.RunCmd(`cd /home/app;
	tar -zxvf main.tar.gz;
	rm -rf main.tar.gz;
	`)
	fmt.Println(tarDist)
}
