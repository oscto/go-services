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
	path string
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

func Download(url, path, filename string) error {

	if err := os.MkdirAll(path, os.ModePerm); err != nil {
		return err
	}

	var res *http.Response
	var err error
	if strings.Contains(url, "https") {
		tlr := &http.Transport{
			TLSClientConfig: &tls.Config{
				InsecureSkipVerify: true,
			},
		}
		client := &http.Client{Transport: tlr}
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
	if err = ioutil.WriteFile(fmt.Sprintf(`%v/%v`, path, filename), data, 0777); err != nil {
		return err
	}

	return nil
}
