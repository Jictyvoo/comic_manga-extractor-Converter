package cbxr

import (
	"errors"
	"iter"

	"github.com/Jictyvoo/comic_manga-extractor-Converter/internal/utils"
)

type (
	FileResult utils.ResultErr[[]byte]
	FileName   string
)

type Extractor interface {
	FileSeq() iter.Seq2[FileName, FileResult]
}

var ErrUnsupportedFormat = errors.New("unsupported file format")
