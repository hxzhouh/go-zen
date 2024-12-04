package webdev

import (
	"fmt"
	"io"
	"os"
	"time"

	"github.com/studio-b12/gowebdav"
)

type WebDav struct {
	ServerUrl string
	UserName  string
	PassWord  string
	FilePath  string
}

func InitWebDav(url, userName, passWord, FilePath string) *WebDav {
	return &WebDav{
		ServerUrl: url,
		UserName:  userName,
		PassWord:  passWord,
		FilePath:  FilePath,
	}
}
func (*WebDav) SyncFile() error {
	// WebDAV 服务器地址和凭据
	serverURL := "https://example.com/webdav/"
	username := "your_username"
	password := "your_password"

	// 初始化 WebDAV 客户端
	client := gowebdav.NewClient(serverURL, username, password)

	// 验证连接
	err := client.Connect()
	if err != nil {
		fmt.Println("Failed to connect to WebDAV server:", err)
		return err
	}
	fmt.Println("Connected to WebDAV server")

	// 定时任务，每 10 秒拉取一次
	ticker := time.NewTicker(10 * time.Second)
	defer ticker.Stop()

	for range ticker.C {
		err := pullFile(client, "/remote/path/to/file.txt", "./local/path/file.txt")
		if err != nil {
			fmt.Println("Failed to pull file:", err)
		} else {
			fmt.Println("File pulled successfully")
		}
	}
	return err
}

// pullFile 从 WebDAV 服务器下载文件到本地
func pullFile(client *gowebdav.Client, remotePath, localPath string) error {
	// 从远程服务器读取文件
	fileReader, err := client.ReadStream(remotePath)
	if err != nil {
		return fmt.Errorf("failed to read remote file: %w", err)
	}
	defer fileReader.Close()

	// 打开本地文件准备写入
	localFile, err := os.Create(localPath)
	if err != nil {
		return fmt.Errorf("failed to create local file: %w", err)
	}
	defer localFile.Close()

	// 将远程文件写入到本地
	_, err = io.Copy(localFile, fileReader)
	if err != nil {
		return fmt.Errorf("failed to copy file content: %w", err)
	}
	return nil
}
