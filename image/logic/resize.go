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

	"errors"

	libWebp "github.com/kolesa-team/go-webp/webp"
	"github.com/nfnt/resize"
	"golang.org/x/image/bmp"
	"golang.org/x/image/webp"
	pb "image.oscto.icu/proto"
)

func (l *Logic) Resize(ctx context.Context, request *pb.CallRequest, response *pb.CallResponse) error {

	uri := strings.Split(request.Url, "/")
	if err := Download(request.Url, path, uri[len(uri)-1]); err != nil {
		return err
	}

	filePath := fmt.Sprintf("%s/%s", path, uri[len(uri)-1])
	file, err := os.Open(filePath)
	defer file.Close()
	if err != nil {
		return err
	}
	var img image.Image
	fileType, err := FileTypeGet(filePath)
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

	newFile, err := os.Create(fmt.Sprintf("%s/new_%s", path, uri[len(uri)-1]))
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

	response.Path = fmt.Sprintf("/new_%s", uri[len(uri)-1])

	return nil
}
