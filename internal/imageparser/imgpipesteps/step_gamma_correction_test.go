package imgpipesteps

import (
	"image"
	"testing"

	"github.com/Jictyvoo/comic_manga-extractor-Converter/internal/imageparser"
	"github.com/Jictyvoo/comic_manga-extractor-Converter/pkg/imgutils"
)

func TestStepGammaCorrectionImage_PerformExec(t *testing.T) {
	testCases := []struct {
		name        string
		gammaValue  float64
		inputImg    image.Image
		expectedImg image.Image
	}{
		{
			name:       "Colored image",
			gammaValue: 0.19,
			inputImg: &image.RGBA{
				Pix: []uint8{
					0x78, 0x78, 0x78, 0xff, 0xb7, 0xb7, 0xb7, 0xff, 0x42, 0x42, 0x42, 0xff, 0xec, 0xec, 0xec,
					0xff, 0xcc, 0xcc, 0xcc, 0xff, 0x93, 0x93, 0x93, 0xff, 0xa6, 0xa6, 0xa6, 0xff, 0x0, 0x0,
					0x0, 0xff, 0xff, 0xff, 0xff, 0xff,
				},
				Stride: 12,
				Rect:   image.Rect(0, 0, 3, 3),
			},
			expectedImg: &image.RGBA{
				Pix: []uint8{
					0xdc, 0xdc, 0xdc, 0xff, 0xef, 0xef, 0xef, 0xff, 0xc5, 0xc5, 0xc5, 0xff, 0xfb, 0xfb, 0xfb, 0xff,
					0xf4, 0xf4, 0xf4, 0xff, 0xe5, 0xe5, 0xe5, 0xff, 0xeb, 0xeb, 0xeb, 0xff, 0x0, 0x0, 0x0, 0xff,
					0xff, 0xff, 0xff, 0xff,
				},
				Stride: 12, Rect: image.Rectangle{Min: image.Point{X: 0, Y: 0}, Max: image.Point{X: 3, Y: 3}},
			},
		},
		{
			name:       "Gray Image",
			gammaValue: 0.62,
			inputImg: &image.Gray{
				Pix:    []uint8{0x4c, 0x96, 0x1d, 0xe2, 0xb3, 0x69, 0x80, 0x0, 0xff},
				Stride: 3,
				Rect:   image.Rect(0, 0, 3, 3),
			},
			expectedImg: &image.Gray{
				Pix:    []uint8{0x78, 0xb7, 0x42, 0xec, 0xcc, 0x93, 0xa6, 0x0, 0xff},
				Stride: 3,
				Rect:   image.Rectangle{Min: image.Point{X: 0, Y: 0}, Max: image.Point{X: 3, Y: 3}},
			},
		},
	}

	for _, tCase := range testCases {
		t.Run(tCase.name, func(t *testing.T) {
			step := NewStepGammaCorrection()
			var (
				state = imageparser.PipeState{Img: tCase.inputImg}
				opts  = imageparser.ProcessOptions{Gamma: tCase.gammaValue}
			)

			// Perform grayscale conversion
			if err := step.PerformExec(&state, opts); err != nil {
				t.Fatalf("PerformExec: %v", err.Error())
			}

			// Validate that the output matches the expected grayscale image
			result := state.Img
			if !imgutils.IsImageEqual(result, tCase.expectedImg) {
				t.Errorf(
					"expected: %#v, actual: %#v", tCase.expectedImg, result,
				)
			}
		})
	}
}
