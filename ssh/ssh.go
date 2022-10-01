package ssh

import (
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"path"
	"time"

	"github.com/pkg/sftp"
	"golang.org/x/crypto/ssh"
)

// Client SSH客户端对象
type Client struct {
	host     string
	port     string
	username string
	password string
	*ssh.Client
}

// NewClient 创建SshClient
func NewClient(host, port, username, password string) *Client {
	if port == "" {
		port = "22"
	}
	c := new(Client)
	c.host = host
	c.port = port
	c.username = username
	c.password = password
	auth := make([]ssh.AuthMethod, 0)
	auth = append(auth, ssh.Password(c.password))

	config := &ssh.ClientConfig{
		User:    c.username,
		Auth:    auth,
		Timeout: 30 * time.Second,
		HostKeyCallback: func(hostname string, remote net.Addr, key ssh.PublicKey) error {
			return nil
		},
	}

	addr := c.host + ":" + c.port
	client, err := ssh.Dial("tcp", addr, config) //连接ssh
	if err != nil {
		log.Fatal("连接ssh失败", err)
	}

	c.Client = client
	return c
}

// RunCmd 运行Shell命令
func (c *Client) RunCmd(cmd string) (string, error) {
	session, err := c.Client.NewSession()
	if err != nil {
		fmt.Println("err:", err)
		return "", err
	}
	defer session.Close()
	runResult, err := session.CombinedOutput(cmd)
	if err != nil {
		fmt.Println("err:", err)
		return "", err
	}
	return string(runResult), nil
}

// NewFtpClient 创建一个FtpClient
func (c *Client) NewFtpClient() *sftp.Client {
	ftpClient, err := sftp.NewClient(c.Client)
	if err != nil {
		fmt.Println("创建ftp客户端失败", err)
		panic(err)
	}
	return ftpClient
}

// MkdirAll 创建文件夹
func (c *Client) MkdirAll(path string) error {
	ftpClient := c.NewFtpClient()
	defer ftpClient.Close()
	return ftpClient.MkdirAll(path)
}

// UploadFile 上传文件
// @localPath 本地路径
// @remotePath 远程目录路径
// @remoteFileName 远程文件名
func (c *Client) UploadFile(localPath string, remotePath string, remoteFileName string) {
	ftpClient := c.NewFtpClient()
	defer ftpClient.Close()

	fmt.Println(localPath, remoteFileName)
	srcFile, err := os.Open(localPath)
	if err != nil {
		fmt.Println("打开文件失败", err)
		panic(err)
	}
	defer srcFile.Close()

	fn := path.Join(remotePath, remoteFileName)

	err = ftpClient.MkdirAll(remotePath)
	if err != nil {
		fmt.Println("创建文件夹失败", err)
		panic(err)
	}

	destFile, e := ftpClient.Create(fn)
	if e != nil {
		fmt.Println("创建文件失败", e)
		panic(e)
	}
	defer destFile.Close()

	buffer := make([]byte, 1024000)
	for {
		n, err := srcFile.Read(buffer)
		destFile.Write(buffer[:n])
		//注意，由于文件大小不定，不可直接使用buffer，否则会在文件末尾重复写入，以填充1024的整数倍
		if err != nil {
			if err == io.EOF {
				fmt.Println("upload finish")
				break
			} else {
				fmt.Println("读取文件出错", err)
				panic(err)
			}
		}
	}
}

// DownloadFile 下载文件
// @remotePath 远程路径
// @localPath 本地目录路径
// @localFilename 本地文件名
func (c *Client) DownloadFile(remotePath string, localPath string, localFilename string) {
	ftpClient := c.NewFtpClient()
	defer ftpClient.Close()

	srcFile, err := ftpClient.Open(remotePath)
	if err != nil {
		fmt.Println("文件读取失败", err)
		panic(err)
	}
	defer srcFile.Close()

	destFile, e := os.Create(path.Join(localPath, localFilename))
	if e != nil {
		fmt.Println("文件创建失败", e)
		panic(e)
	}
	defer destFile.Close()
	if _, err = srcFile.WriteTo(destFile); err != nil {
		fmt.Println("文件写入失败", err)
		panic(err)
	}
	fmt.Println("文件下载成功")
}
