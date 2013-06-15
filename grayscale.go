package grayscale

import (
  "image"
  "os"
  "log"
  "image/color"
  _ "image/jpeg"
)

func handleError(err error) {
  if err != nil {
    log.Fatal(err)
  }
}

// Given a path to a jpeg, convert it to grayscale and return as an Image
func GrayScale(fname string) image.Image {
  file, err := os.Open(fname)
  handleError(err)
  defer file.Close()

  originalImg, _, err := image.Decode(file)
  handleError(err)

  bounds := originalImg.Bounds()
  w, h := bounds.Max.X, bounds.Max.Y

  grayImg := image.NewGray(bounds)

  for x := 0; x < w; x++ {
    for y := 0; y < h; y++ {
      originalColor := originalImg.At(x, y)
      grayColor := color.GrayModel.Convert(originalColor)
      grayImg.Set(x, y, grayColor)
    }
  }

  return grayImg
}
