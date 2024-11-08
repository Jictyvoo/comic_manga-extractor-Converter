package cbxr

import (
	"errors"
	"io"
	"iter"

	"github.com/Jictyvoo/comic_manga-extractor-Converter/internal/utils"
)

type (
	FileResult        utils.ResultErr[[]byte]
	FileName          string
	FileContentStream interface {
		io.ReaderAt
		io.ReadSeeker
	}
)

type Extractor interface {
	FileSeq() iter.Seq2[FileName, FileResult]
}

var ErrUnsupportedFormat = errors.New("unsupported file format")
