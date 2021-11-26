package logic

import (
	"crypto/tls"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

type Logic struct {
}

var (
	rootPath = "/"
	filePath = "tmp"
)

func Register() *Logic {
	return &Logic{}
}

func FileTypeGet(filePath string) (string, error) {

	file, err := os.Open(filePath)
	defer file.Close()
	buffer := make([]byte, 512)
	_, err = file.Read(buffer)
	if err != nil {
		return "", err
	}
	fileType := http.DetectContentType(buffer)

	return fileType, nil
}

func Download(url, filepath, filename string) error {

	if err := os.MkdirAll(fmt.Sprintf("%s%s%s", rootPath, filePath, filepath), os.ModePerm); err != nil {
		return err
	}

	var res *http.Response
	var err error
	if strings.Contains(url, "https") {
		tls := &http.Transport{
			TLSClientConfig: &tls.Config{
				InsecureSkipVerify: true,
			},
		}
		client := &http.Client{Transport: tls}
		res, err = client.Get(url)
	} else {
		res, err = http.Get(url)
	}

	if err != nil {
		return err
	}
	data, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return err
	}
	defer res.Body.Close()
	if err = ioutil.WriteFile(fmt.Sprintf("%s%s%s%s", rootPath, filePath, filepath, filename), data, 0777); err != nil {
		return err
	}

	return nil
}
