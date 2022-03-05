package helpers

import (
	"campmart/models"
	"io/ioutil"
	"log"
)

func ProcessImageAndReturnNames(productImages []models.ProductImage, productID string) ([]string, error) {
	// range over the project images and store in path "website-pub/images/products"
	for _, image := range productImages {
		// read bytes from image file
		bs, err := ioutil.ReadAll(image.File)
		if err != nil {
			log.Printf("Error processing image: %v, errormsg{%v}", image.Name, err)
		}

		// pattern is a th name of each image file which is
		// a string generated by concatenating
		// databaseId, random string generated by ioutil.Tempfile and the image extension
		pattern := productID + "*" + image.Extension
		tempfile, err := ioutil.TempFile("website-pub/images/products", pattern)
		if err != nil {
			log.Printf("Erorr creating tempfile for %v, errormsg{%v}", image.Name, err)
		}
		defer tempfile.Close()

		tempfile.Write(bs)
	}

	filesInProductDir, err := ioutil.ReadDir("website-pub/images/products")
	if err != nil {
		log.Printf("Error reading product images drectory: %v", err)
		return []string{}, err
	}

	var img_names []string

	for _, file := range filesInProductDir {
		img_names = append(img_names, file.Name())
	}

	return img_names, nil
}