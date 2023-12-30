package utils

import (
	"bytes"
	"fmt"
	"image"
	"image/jpeg"
	"io"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"strings"

	"github.com/google/uuid"
	"github.com/nfnt/resize"
)

func DownloadImage(imageURL string) (image.Image, error, string) {
	resp, err := http.Get(imageURL)
	if err != nil {
		return nil, err, "Getting error"
	}
	defer resp.Body.Close()
	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err, "Reading error"
	}
	img, _, err := image.Decode(bytes.NewReader(data))
	if err != nil {
		return nil, err, "decoding Error"
	}
	return img, nil, ""
}

func ResizeImage(img image.Image) (image.Image, error) {
	imgResized := resize.Resize(1024, 0, img, resize.Lanczos3)
	return imgResized, nil
}

func CompressImage(img image.Image, qFactor int) ([]byte, error) {
	buf := new(strings.Builder)
	err := jpeg.Encode(buf, img, &jpeg.Options{Quality: qFactor})
	if err != nil {
		return nil, err
	}
	return []byte(buf.String()), nil
}

func SaveImage(img []byte, dir string, comp_url string) (string, error) {

	parsedURL, err := url.Parse(comp_url)
	if err != nil {
		fmt.Println("Error parsing URL:", err)
		return "", err
	}

	imageName := filepath.Base(parsedURL.Path)
	imageUUID := uuid.New()
	imageName = fmt.Sprintf("%s_%s", imageUUID.String(), imageName)

	outputFileName := filepath.Join(dir, imageName)
	err = os.MkdirAll(dir, os.ModePerm)
	if err != nil {
		return "", err
	}
	f, err := os.Create(outputFileName)
	if err != nil {
		return "", err
	}
	defer f.Close()

	_, err = io.Copy(f, strings.NewReader(string(img)))
	if err != nil {
		return "", err
	}

	return outputFileName, err
}
