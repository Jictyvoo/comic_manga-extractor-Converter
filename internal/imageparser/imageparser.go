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
		privateInternalStep()
	}
	stepIdentifier interface {
		StepID() string
	}
	UnitStep interface {
		PixelStep(imgColor color.Color) color.Color
		stepIdentifier
	}

	PipeStep interface {
		PerformExec(state *PipeState, opts ProcessOptions) (err error)
		paletteFactoryStep
		stepIdentifier
	}
)

type BaseImageStep struct {
	fac imgutils.DrawImageFactory
}

func NewBaseImageStep(palette color.Palette) BaseImageStep {
	return BaseImageStep{fac: imgutils.NewImageFactory(palette)}
}

func (s *BaseImageStep) privateInternalStep() {
	// Do nothing, this function only exists to make sure that all types compose this struct
}

func (s *BaseImageStep) DrawImage(colorModel color.Model, bounds image.Rectangle) draw.Image {
	if s.fac != nil {
		return s.fac.CreateDrawImage(colorModel, bounds)
	}
	return imgutils.NewDrawFromImgColorModel(colorModel, bounds)
}

func (s *BaseImageStep) UpdateDrawFactory(fac imgutils.DrawImageFactory) {
	s.fac = fac
}
