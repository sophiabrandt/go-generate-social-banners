package generate

import (
	"flag"
	"fmt"
	"os"

	"github.com/fogleman/gg"
)

// AppEnv holds the local context for the application.
type AppEnv struct {
	inputImg    string      // background image
	outputImg   string      // file name for the generated image
	defaultText string      // default text (domain name)
	title       string      // title
	dc          *gg.Context // drawContext for images
}

// CLI runs the generate command line app and returns its exit status.
func CLI(args []string) int {
	var app AppEnv
	err := app.fromArgs(args)
	if err != nil {
		return 2
	}
	if err = app.run(); err != nil {
		fmt.Fprintf(os.Stderr, "Runtime error: %v\n", err)
		return 1
	}
	return 0
}

func (app *AppEnv) fromArgs(args []string) error {
	fl := flag.NewFlagSet("go-generate-social-banners", flag.ContinueOnError)
	fl.StringVar(
		&app.inputImg, "i", InputImage, "Path to background image to generate banner from",
	)
	fl.StringVar(
		&app.outputImg, "o", OutputImage, "Full path of the image to generate",
	)
	fl.StringVar(
		&app.defaultText, "d", DefaultText, "Default Text (domain name)",
	)
	fl.StringVar(
		&app.title, "t", Title, "Title",
	)
	if err := fl.Parse(args); err != nil {
		return err
	}
	return nil
}

func (app *AppEnv) run() error {
	// load image
	imgLoaded, err := app.LoadImage(app.inputImg)
	if err != nil {
		return err
	}

	// render image with semi-transparent overlay
	app.RenderImage(imgLoaded)

	// add default text
	if err := app.AddDefaultText(app.defaultText); err != nil {
		return err
	}

	// add title
	if err := app.AddTitle(app.title); err != nil {
		return err
	}

	// save image
	if err := app.SaveImage(app.outputImg); err != nil {
		return err
	}
	return nil
}
