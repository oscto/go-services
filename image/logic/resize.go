package logic

import (
	"context"
	"fmt"
	"image"
	"image/gif"
	"image/jpeg"
	"image/png"
	"os"
	"strings"
	"time"

	"errors"

	libWebp "github.com/kolesa-team/go-webp/webp"
	"github.com/nfnt/resize"
	"golang.org/x/image/bmp"
	"golang.org/x/image/webp"
	pb "image.oscto.icu/proto"
)

func (l *Logic) Resize(ctx context.Context, request *pb.CallRequest, response *pb.CallResponse) error {

	uri := strings.Split(request.Url, "/")
	path := fmt.Sprintf("/%s", time.Now().Format("20060102"))
	if err := Download(request.Url, path, uri[len(uri)-1]); err != nil {
		return err
	}

	oldFile := fmt.Sprintf("%s%s%s/%s", rootPath, filePath, path, uri[len(uri)-1])
	file, err := os.Open(oldFile)
	defer file.Close()
	if err != nil {
		return err
	}
	var img image.Image
	fileType, err := FileTypeGet(oldFile)
	if err != nil {
		return err
	}
	switch fileType {
	case "image/gif":
		img, err = gif.Decode(file)
	case "image/png":
		img, err = png.Decode(file)
	case "image/webp":
		img, err = webp.Decode(file)
	case "image/bmp":
		img, err = bmp.Decode(file)
	case "image/jpeg":
		img, err = jpeg.Decode(file)
	default:
		return errors.New("image type not support")
	}

	newImg := resize.Resize(uint(request.Width), uint(request.Height), img, resize.Lanczos3)

	newFile, err := os.Create(fmt.Sprintf("%s%s%s/new_%s", rootPath, filePath, path, uri[len(uri)-1]))
	defer newFile.Close()
	if err != nil {
		return err
	}

	switch fileType {
	case "image/gif":
		err = gif.Encode(newFile, newImg, nil)
	case "image/png":
		err = png.Encode(newFile, newImg)
	case "image/webp":
		err = libWebp.Encode(newFile, newImg, nil)
	case "image/bmp":
		err = bmp.Encode(newFile, newImg)
	case "image/jpeg":
		err = jpeg.Encode(newFile, newImg, nil)
	default:
		return errors.New("image type not support")
	}
	if err != nil {
		return err
	}

	response.Path = fmt.Sprintf("%s%s/new_%s", filePath, path, uri[len(uri)-1])

	return nil
}
