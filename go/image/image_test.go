package image

import (
	"bytes"
	"github.com/chai2010/webp"
	"github.com/disintegration/imaging"
	"github.com/stretchr/testify/require"
	"image"
	"io"
	"os"
	"testing"
)

func TestCompressJPEGFile(t *testing.T) {
	f, err := os.Open("img.jpg")
	require.Nil(t, err)
	imageBytes, err := io.ReadAll(f)
	require.Nil(t, err)
	t.Logf("file size is %d KB", len(imageBytes)/1024)
	resultImage := new(bytes.Buffer)
	imageObject, _, err := image.Decode(bytes.NewReader(imageBytes))
	require.Nil(t, err)
	err = imaging.Encode(resultImage, imageObject, imaging.JPEG, imaging.JPEGQuality(95))
	require.Nil(t, err)
	resultByte := resultImage.Bytes()
	t.Logf("file size of result image %d KB", len(resultByte)/1024)
	t.Logf("compression efficiency is %f", 1-float64(len(resultByte))/float64(len(imageBytes)))
	f, err = os.OpenFile("imgCompress.jpeg", os.O_RDWR|os.O_CREATE, os.ModePerm)
	require.Nil(t, err)
	_, err = f.Write(resultByte)
	require.Nil(t, err)
}

func TestCompressionEfficiencyForWebP(t *testing.T) {
	f, err := os.Open("imgCompress.jpeg")
	require.Nil(t, err)
	imageBytes, err := io.ReadAll(f)
	require.Nil(t, err)
	imageObject, _, err := image.Decode(bytes.NewReader(imageBytes))
	require.Nil(t, err)
	var result bytes.Buffer

	err = webp.Encode(&result, imageObject, &webp.Options{Quality: 80})
	resultByte := result.Bytes()
	t.Logf("file size of result webp image %d KB", len(resultByte)/1024)
	t.Logf("compression efficiency is %f", 1-float64(len(resultByte))/float64(len(imageBytes)))
	f, err = os.OpenFile("imgCompressWebP.webp", os.O_RDWR|os.O_CREATE, os.ModePerm)
	require.Nil(t, err)
	_, err = f.Write(resultByte)
	require.Nil(t, err)
}

func TestCompressWebPDirectly(t *testing.T) {
	f, err := os.Open("img.jpg")
	require.Nil(t, err)
	imageBytes, err := io.ReadAll(f)
	require.Nil(t, err)
	imageObject, _, err := image.Decode(bytes.NewReader(imageBytes))
	require.Nil(t, err)
	var result bytes.Buffer

	err = webp.Encode(&result, imageObject, &webp.Options{Quality: 80})
	resultByte := result.Bytes()
	t.Logf("file size of result webp image %d KB", len(resultByte)/1024)
	t.Logf("compression efficiency is %f", 1-float64(len(resultByte))/float64(len(imageBytes)))
	f, err = os.OpenFile("imgCompressWebPDirectly.webp", os.O_RDWR|os.O_CREATE, os.ModePerm)
	require.Nil(t, err)
	_, err = f.Write(resultByte)
	require.Nil(t, err)
}
