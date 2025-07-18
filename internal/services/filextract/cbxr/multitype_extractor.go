package cbxr

import (
	"context"
	"io"
	"iter"
	"time"

	"github.com/mholt/archives"
)

func checkFileFormat(filename string, file io.Reader) (io.Reader, archives.Extractor, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	format, fileReader, err := archives.Identify(ctx, filename, file)
	if err != nil {
		return nil, nil, err
	}

	// It must be an filextract
	if ex, ok := format.(archives.Extractor); ok {
		return fileReader, ex, nil
	}

	return nil, nil, ErrUnsupportedFormat
}

type (
	MultiZipRarExtractor struct {
		format     archives.Extractor
		fileReader io.Reader
		timeout    time.Duration
	}
	archivesExtractInteract struct {
		yield          func(FileName, FileResult) bool
		stopExtracting bool
	}
)

func NewMultiZipRarExtractor(
	filename string,
	fileReader FileContentStream,
) (*MultiZipRarExtractor, error) {
	reader, format, err := checkFileFormat(filename, fileReader)
	if err != nil {
		return nil, err
	}

	return &MultiZipRarExtractor{
		fileReader: reader,
		format:     format,
		timeout:    9000 * time.Second,
	}, nil
}

func (aei *archivesExtractInteract) handleFile(_ context.Context, f archives.FileInfo) error {
	// Skip all remaining files
	if aei.stopExtracting {
		return nil
	}

	reader, err := f.Open()
	if err != nil {
		return err
	}
	defer reader.Close()

	filename := f.NameInArchive
	if f.IsDir() {
		return nil
	}

	result := FileResult{}
	result.Data, result.Error = io.ReadAll(reader)
	if !aei.yield(FileName(filename), result) {
		aei.stopExtracting = true
		return context.Canceled
	}

	return nil
}

func (ext MultiZipRarExtractor) FileSeq() iter.Seq2[FileName, FileResult] {
	return func(yield func(FileName, FileResult) bool) {
		aei := archivesExtractInteract{yield: yield}
		ctx, cancel := context.WithTimeout(context.Background(), ext.timeout)
		defer cancel()

		// Use nil to extract all files
		err := ext.format.Extract(ctx, ext.fileReader, aei.handleFile)
		if err != nil {
			if !yield("", FileResult{Error: err}) {
				return
			}
		}

		return
	}
}
