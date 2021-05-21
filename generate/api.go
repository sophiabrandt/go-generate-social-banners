package generate

import (
	"image"
	"image/color"
	"path/filepath"

	"github.com/fogleman/gg"
	"github.com/pkg/errors"
)

const (
	// InputImage holds the path to the background image
	InputImage = "./background.jpg"
	// Output Path for the generated image
	OutputImage = "./social-media-banner.png"
	// default text (domain name)
	DefaultText = "https://rockyourcode.com"
	// Title
	Title = "Programmatically generate social media images in Go"
)

// Load image file from disk
func (app *AppEnv) LoadImage(inputImage string) (image.Image, error) {
	img, err := gg.LoadImage(inputImage)
	if err != nil {
		return image.Rect(0, 0, 0, 0), errors.Wrap(err, "load background image")
	}
	return img, nil
}

// Render image with semi-transparent overlay
func (app *AppEnv) RenderImage(img image.Image) {
	// render basic image in 1000 x 420
	app.dc = gg.NewContext(1000, 420)
	app.dc.DrawImage(img, 0, 0)

	// add overlay
	margin := 20.0
	x := margin
	y := margin
	w := float64(app.dc.Width()) - (2.0 * margin)
	h := float64(app.dc.Height()) - (2.0 * margin)
	// black background with 80 % opacity (80 % x 255 = 204)
	app.dc.SetColor(color.RGBA{0, 0, 0, 204})
	app.dc.DrawRectangle(x, y, w, h)
	app.dc.Fill()
}

// Add default text (domain name)
func (app *AppEnv) AddDefaultText(text string) error {
	textColor := color.White
	fontPath := filepath.Join("fonts", "SpaceMono-Regular.ttf")
	if err := app.dc.LoadFontFace(fontPath, 30); err != nil {
		return errors.Wrap(err, "load font for default text")
	}
	r, g, b, _ := textColor.RGBA()
	mutedColor := color.RGBA{
		R: uint8(r),
		G: uint8(g),
		B: uint8(b),
		A: uint8(200),
	}
	app.dc.SetColor(mutedColor)
	marginX := 50.0
	marginY := 20.0
	textWidth, textHeight := app.dc.MeasureString(text)
	x := float64(app.dc.Width()) - textWidth - marginX
	y := float64(app.dc.Height()) - textHeight - marginY
	app.dc.DrawString(text, x, y)
	return nil
}

// Add title
func (app *AppEnv) AddTitle(title string) error {
	textColor := color.White
	textShadowColor := color.Black
	fontPath := filepath.Join("fonts", "BioRhyme-Bold.ttf")
	if err := app.dc.LoadFontFace(fontPath, 65); err != nil {
		return errors.Wrap(err, "load font for title")
	}
	textRightMargin := 60.0
	textTopMargin := 70.0
	x := textRightMargin
	y := textTopMargin
	maxWidth := float64(app.dc.Width()) - textRightMargin - textRightMargin
	app.dc.SetColor(textShadowColor)
	app.dc.DrawStringWrapped(title, x+1, y+1, 0, 0, maxWidth, 1.5, gg.AlignLeft)
	app.dc.SetColor(textColor)
	app.dc.DrawStringWrapped(title, x, y, 0, 0, maxWidth, 1.5, gg.AlignLeft)
	return nil
}

// Save the image
func (app *AppEnv) SaveImage(outputFileName string) error {
	err := app.dc.SavePNG(outputFileName)
	if err != nil {
		return errors.Wrap(err, "save png")
	}
	return nil
}
