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

	"github.com/kolesa-team/go-webp/encoder"
	libWebp "github.com/kolesa-team/go-webp/webp"
	"golang.org/x/image/bmp"

	//	libWebp "github.com/kolesa-team/go-webp/webp"
	pb "image.oscto.icu/proto"
)

func (l *Logic) ToWebP(ctx context.Context, request *pb.ToWebPRequest, response *pb.ToWebPResponse) error {

	uri := strings.Split(request.Url, "/")
	path := fmt.Sprintf("/%s", time.Now().Format("20060102"))
	if err := Download(request.Url, path, uri[len(uri)-1]); err != nil {
		return err
	}

	filename := fmt.Sprintf("%s%s/%s", rootPath, path, uri[len(uri)-1])
	file, err := os.Open(filename)
	defer file.Close()
	if err != nil {
		return err
	}

	var img image.Image
	fileType, _ := FileTypeGet(path)
	switch fileType {
	case "image/gif":
		img, err = gif.Decode(file)
	case "image/png":
		img, err = png.Decode(file)
	case "image/webp":
		return nil
	case "image/bmp":
		img, err = bmp.Decode(file)
	default:
		img, err = jpeg.Decode(file)
	}
	newFile := fmt.Sprintf("%s%s/new-%s.webp", rootPath, path, uri[len(uri)-1])
	out, err := os.Create(newFile)
	defer out.Close()
	if err != nil {
		return err
	}
	options, err := encoder.NewLossyEncoderOptions(encoder.PresetDefault, 75)
	if err != nil {
		return err
	}

	if err := libWebp.Encode(out, img, options); err != nil {
		return err
	}

	response.Path = fmt.Sprintf("%s/new-%s.webp", path, uri[len(uri)-1])

	return nil
}
