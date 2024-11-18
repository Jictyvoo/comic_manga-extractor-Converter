package imageparser

import (
	"image"
	"image/color"
	"image/draw"

	"github.com/Jictyvoo/comic_manga-extractor-Converter/pkg/imgutils"
)

type (
	paletteFactoryStep interface {
		UpdateDrawFactory(fac imgutils.DrawImageFactory)
	}
	UnitStep interface {
		PixelStep(imgColor color.Color) color.Color
	}

	PipeStep interface {
		PerformExec(state *pipeState, opts processOptions) (err error)
		paletteFactoryStep
	}
)

type baseImageStep struct {
	fac imgutils.DrawImageFactory
}

func (s *baseImageStep) drawImage(img image.Image, bounds image.Rectangle) draw.Image {
	if s.fac != nil {
		return s.fac.CreateDrawImage(img, bounds)
	}
	return imgutils.NewDrawFromImgColorModel(img, bounds)
}

func (s *baseImageStep) UpdateDrawFactory(fac imgutils.DrawImageFactory) {
	s.fac = fac
}
