package main

import (
	"io"

	"github.com/Jictyvoo/comic_manga-extractor-Converter/internal/services/outdirwriter"
)

type FileWriterWrapper struct {
	outdirwriter.WriterHandle
}

func NewFileWriterWrapper(extractDir string) *FileWriterWrapper {
	return &FileWriterWrapper{
		outdirwriter.NewWriterHandle(extractDir),
	}
}

func (f FileWriterWrapper) Close() error {
	return nil
}

func (f FileWriterWrapper) Shutdown() error {
	return f.WriterHandle.OnFinish()
}

func (f FileWriterWrapper) Process(filename string, data []byte) {
	_ = f.WriterHandle.Handler(
		filename, func(writer io.Writer) error {
			_, err := writer.Write(data)
			return err
		},
	)
}
