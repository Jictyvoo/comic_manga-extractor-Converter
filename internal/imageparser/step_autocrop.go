package imageparser

import (
	"image"

	"golang.org/x/image/draw"

	"github.com/Jictyvoo/comic_manga-extractor-Converter/internal/deviceprof"
)

type StepAutoCropImage struct {
	resolution deviceprof.Resolution
	isPixelArt bool
}

func NewStepAutoCrop(resolution deviceprof.Resolution) StepAutoCropImage {
	return StepAutoCropImage{resolution: resolution}
}

func (sgsi StepAutoCropImage) PerformExec(state *pipeState, _ processOptions) (err error) {
	resized := image.NewRGBA(image.Rect(0, 0, int(sgsi.resolution.Width), int(sgsi.resolution.Height)))

	drawInterpolator := draw.ApproxBiLinear
	if sgsi.isPixelArt {
		drawInterpolator = draw.NearestNeighbor
	}
	drawInterpolator.Scale(resized, resized.Bounds(), state.img, state.img.Bounds(), draw.Over, nil)

	state.img = resized
	return
}
