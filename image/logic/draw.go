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

	"golang.org/x/image/bmp"
	pb "image.oscto.icu/proto"
)

func (l *Logic) Draw(ctx context.Context, request *pb.DrawRequest, response *pb.DrawResponse) error {

	uri := strings.Split(request.Url, "/")
	path := fmt.Sprintf("%s/", time.Now().Format("20060102"))
	if err := Download(request.Url, path, uri[len(uri)-1]); err != nil {
		return err
	}
	in, err := os.Open(request.Url)
	defer in.Close()
	if err != nil {
		return err
	}
	newFile := fmt.Sprintf("%s%s/new-%s", rootPath, path, uri[len(uri)-1])
	out, err := os.Create(newFile)
	defer out.Close()
	if err != nil {
		return err
	}
	x0 := int(request.X0)
	x1 := int(request.X1)
	y0 := int(request.Y0)
	y1 := int(request.Y1)
	canvas, fm, err := image.Decode(in)
	bounds := canvas.Bounds()

	if x0 < 0 || x0 > bounds.Dx() {
		x0 = 0
	}
	if y0 < 0 || y0 > bounds.Dy() {
		y0 = 0
	}
	if x1 > bounds.Dx() {
		x1 = bounds.Dx()
	}
	if y1 > bounds.Dy() {
		y1 = bounds.Dy()
	}

	if err != nil {
		return err
	}
	switch fm {
	case "png":
		switch canvas.(type) {
		case *image.NRGBA:
			img := canvas.(*image.NRGBA)
			subImg := img.SubImage(image.Rect(x0, y0, x1, y1)).(*image.NRGBA)
			return png.Encode(out, subImg)
		case *image.RGBA:
			img := canvas.(*image.RGBA)
			subImg := img.SubImage(image.Rect(x0, y0, x1, y1)).(*image.RGBA)
			return png.Encode(out, subImg)
		}
	case "gif":
		img := canvas.(*image.Paletted)
		subImg := img.SubImage(image.Rect(x0, y0, x1, y1)).(*image.Paletted)
		return gif.Encode(out, subImg, &gif.Options{})
	case "bmp":
		img := canvas.(*image.RGBA)
		subImg := img.SubImage(image.Rect(x0, y0, x1, y1)).(*image.RGBA)
		return bmp.Encode(out, subImg)
	default:
		img := canvas.(*image.YCbCr)
		subImg := img.SubImage(image.Rect(x0, y0, x1, y1)).(*image.YCbCr)
		return jpeg.Encode(out, subImg, nil)
	}
	response.Path = fmt.Sprintf("/%snew_%s", path, uri[len(uri)-1])
	return nil
}
