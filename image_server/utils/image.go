package utils

import (
	"bytes"
	"image"
	"io"
	"net/http"
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
