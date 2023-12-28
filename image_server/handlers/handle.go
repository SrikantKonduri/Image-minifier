package handlers

import (
	"database/sql"
	"fmt"
	"image_server/databases"
	"image_server/utils"

	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"
)

func HandleProduct(db *sql.DB, productId int64) {
	fmt.Println("Handling product with porduct ID: ", productId)
	imageURLs, err := databases.GetProductURLS(db, productId)
	if err != nil {
		fmt.Println("[x] Cannot find URLs: ", err)
		return
	}
	compressedURL, err := databases.GetCompressedURL(db, productId)
	if err != nil {
		fmt.Println("[x] Cannot find compreesed URL: ", err)
		return
	}
	newCompressedURL := compressedURL

	for _, url := range imageURLs {
		// Download the image
		// fmt.Printf("Downloadig file: %s\n", url)
		imageReader, err, msg := utils.DownloadImage(url)
		if err != nil {
			fmt.Printf("[x] Error downloading image from : %v %s\n", err, msg)
			continue
		}

		resImg, err := utils.ResizeImage(imageReader)
		if err != nil {
			fmt.Printf("[x] Error Resizing image from : %v\n", err)
			continue
		}

		img, err := utils.CompressImage(resImg, 50)

		if err != nil {
			fmt.Println("[x] Cannot compress image: ", err)
			continue
		}

		filePath, err := utils.SaveImage(img, "products/", url)

		if err != nil {
			fmt.Println("[x] Cannot save image: ", err)
			continue
		}
		newCompressedURL = newCompressedURL + "," + filePath
		fmt.Printf("[+] Image downloaded, compressed, and saved at: %s\n", filePath)
	}
	if compressedURL != newCompressedURL {
		err = databases.UpdateCompressedURL(db, productId, newCompressedURL)
		fmt.Println("Checking: ", err)
	}
}
