package imageparser

import (
	"bytes"
	"image"
	"image/color"
	"image/draw"
	"image/jpeg"
	"io"

	"github.com/Jictyvoo/comic_manga-extractor-Converter/internal/deviceprof"
)

type Options struct {
	ForceColor       bool
	TargetResolution deviceprof.Resolution // Assuming this is the dimensions for thumbnail
}

type ImageParser struct {
	options Options
	image   image.Image
	name    string
}

func NewParser(imageBytes []byte, opt Options) (parser ImageParser, err error) {
	parser = ImageParser{options: opt}
	if parser.image, parser.name, err = image.Decode(bytes.NewReader(imageBytes)); err != nil {
		return
	}

	return
}

func (p *ImageParser) AdaptManga() (err error) {
	pipe := NewImagePipeline(
		[]PipeStep{
			NewStepAutoContrast(0, 0),
			NewStepRescale(p.options.TargetResolution),
			NewStepGrayScale(),
		},
	)

	if p.image, err = pipe.Process(p.image); err != nil {
		return
	}

	return
}

func (p *ImageParser) Save(outDst io.Writer) error {
	options := jpeg.Options{Quality: 85}
	if err := jpeg.Encode(outDst, p.image, &options); err != nil {
		return err
	}
	return nil
}

func createDrawImage(img image.Image, bounds image.Rectangle) draw.Image {
	switch img.ColorModel() {
	case color.GrayModel, color.Gray16Model:
		return image.NewGray(bounds)
	case color.RGBAModel, color.RGBA64Model, color.NRGBAModel, color.NRGBA64Model:
		return image.NewRGBA(bounds)
	}

	return image.NewRGBA(bounds)
}
