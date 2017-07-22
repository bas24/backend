package main

import (
	"bytes"
	"github.com/nfnt/resize"
	"image"
	"image/jpeg"
	"net/http"
)

func encodeImage(image image.Image) (bytes.Buffer, error) {
	buff := new(bytes.Buffer)
	err := jpeg.Encode(buff, image, nil)
	if err != nil {
		return *buff, err
	}
	return *buff, nil
}

func getImagesInBytes(url string) (map[int][]byte, error) {
	mp := map[int][]byte{}
	res, err := http.Get(url)
	if err != nil || res.StatusCode != 200 {
		return mp, err
	}
	defer res.Body.Close()

	m, _, err := image.Decode(res.Body)
	if err != nil {
		return mp, err
	}
	// resize image.Image
	m600 := resize.Resize(600, 0, m, resize.Lanczos3)
	m200 := resize.Resize(200, 0, m, resize.Lanczos3)
	buff600, err := encodeImage(m600)
	if err != nil {
		return mp, err
	}
	buff200, err := encodeImage(m200)
	if err != nil {
		return mp, err
	}
	mp[600] = buff600.Bytes()
	mp[200] = buff200.Bytes()

	return mp, nil
}
