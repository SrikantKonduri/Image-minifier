// utils_test.go
package utils_test

import (
	"bytes"
	"image"
	"image/color"
	"image/jpeg"
	"image_server/utils"
	"os"
	"testing"
)

func TestDownloadImage(t *testing.T) {
	// Provide a valid image URL for testing
	imageURL := "https://cdn.mos.cms.futurecdn.net/qxRqURamUXuNeQr46zfyo6.jpeg"

	img, err, msg := utils.DownloadImage(imageURL)

	// Check if there is no error
	if err != nil {
		t.Errorf("DownloadImage failed with error: %v, message: %s", err, msg)
	}

	// Check if the returned image is of type image.Image
	if _, ok := img.(image.Image); !ok {
		t.Errorf("DownloadImage did not return an image")
	}
}

func TestResizeImage(t *testing.T) {
	// Create a sample image for testing
	img := image.NewRGBA(image.Rect(0, 0, 800, 600))

	resizedImg, err := utils.ResizeImage(img)

	// Check if there is no error
	if err != nil {
		t.Errorf("ResizeImage failed with error: %v", err)
	}

	// Check if the resized image has the correct dimensions
	if resizedImg.Bounds().Dx() != 1024 {
		t.Errorf("ResizeImage did not resize the image to the expected width")
	}
}

// Add similar tests for CompressImage and SaveImage

func TestCompressImage(t *testing.T) {

	img := image.NewRGBA(image.Rect(0, 0, 640, 480))
	for y := 0; y < 480; y++ {
		for x := 0; x < 640; x++ {
			img.SetRGBA(x, y, color.RGBA{
				R: uint8(x % 256),
				G: uint8(y % 256),
				B: uint8((x + y) % 256),
				A: 255,
			})
		}
	}

	// Compress the image with quality 80
	compressedImg, err := utils.CompressImage(img, 80)
	if err != nil {
		t.Fatalf("Failed to compress image: %v", err)
	}

	// Decode the compressed image
	decodedImg, err := jpeg.Decode(bytes.NewReader(compressedImg))
	if err != nil {
		t.Fatalf("Failed to decode compressed image: %v", err)
	}

	// Check the dimensions of the decoded image
	width := decodedImg.Bounds().Dx()
	if width != 640 {
		t.Errorf("Decoded image has incorrect width: got %d, want 640", width)
	}

	height := decodedImg.Bounds().Dy()
	if height != 480 {
		t.Errorf("Decoded image has incorrect height: got %d, want 480", height)
	}
}

func TestSaveImage(t *testing.T) {
	img := []byte{0x89, 0x50, 0x4E, 0x47, 0x0D, 0x0A, 0x1A, 0x0A}
	// Save the image
	filePath, err := utils.SaveImage(img, "./testImages/", "https://tinyjpg.com/images/social/website.jpg")

	// Check if there is no error
	if err != nil {
		t.Errorf("SaveImage failed with error: %v", err)
	}

	// Check if the file exists
	_, err = os.Stat(filePath)
	if os.IsNotExist(err) {
		t.Errorf("SaveImage did not create the file")
	}

	// Optionally, check additional properties of the saved file if needed
}
