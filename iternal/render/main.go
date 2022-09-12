package render

import (
	"fmt"
	"image"
	"mime/multipart"
	"os"

	"github.com/ybrbnf2080/genRand/iternal/app"
	"github.com/ybrbnf2080/genRand/iternal/render"
	"github.com/ybrbnf2080/genRand/iternal/transform"
)

func RenderPict(file multipart.File, height int, width int) string {
	var Transforms = transform.NewTransform(height, width, -1)

	image, _, err := image.Decode(file)
	file.Close()
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s: %v\n", err)
		os.Exit(1)
	}
	if Transforms.CompressCoof == -1 && Transforms.Height > 0 {
		Transforms.CompressCoof = (image.Bounds().Dx() / Transforms.Height) * 2
	}
	return app.Convert(image, Transforms, render.Drawer(render.SybmolSel(app.ColorMap, 0)))
}
