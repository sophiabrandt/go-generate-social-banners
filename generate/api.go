package generate

import (
	"image"

	"github.com/fogleman/gg"
	"github.com/pkg/errors"
)

const (
	// InputImage holds the path to the background image
	InputImage = "./background.jpg"
	// Output Path for the generated image
	OutputImage = "./social-banner.png"
)

// Load image file from disk
func (app *AppEnv) LoadImage(inputImage string) (image.Image, error) {
	img, err := gg.LoadImage(inputImage)
	if err != nil {
		return image.Rect(0, 0, 0, 0), errors.Wrap(err, "load background image")
	}
	return img, nil
}

// Render image
func (app *AppEnv) RenderImage(img image.Image) {
	app.dc = gg.NewContext(1000, 420)
	app.dc.DrawImage(img, 0, 0)
}

// Save the image
func (app *AppEnv) SaveImage(outputFileName string) error {
	err := app.dc.SavePNG(outputFileName)
	if err != nil {
		return errors.Wrap(err, "save png")
	}
	return nil
}
