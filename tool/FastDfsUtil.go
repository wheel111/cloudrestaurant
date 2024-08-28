package tool

import (
	"bufio"
	"github.com/tedcy/fdfs_client"
	"log"
	"os"
	"strings"
)

func UploadFile(fileName string) string {
	client, err := fdfs_client.NewClientWithConfig("./config/fastdfs.conf")
	defer client.Destory()
	if err != nil {
		log.Fatal(err.Error())
		return ""
	}
	fileId, err := client.UploadByFilename(fileName)
	if err != nil {
		log.Fatal(err.Error())
		return ""
	}
	return fileId
}

func FileServerAddr() string {
	file, err := os.Open("./config/ffastdfs.conf")
	if err != nil {
		return ""
	}
	reader := bufio.NewReader(file)
	for {
		line, err := reader.ReadString('\n')
		line = strings.TrimSuffix(line, "\n")
		str := strings.SplitN(line, "=", 2)
		switch str[0] {
		case "http_port":
			return str[1]
		}
		if err != nil {
			return ""
		}
	}
}
