package io

import (
	"io"

	"github.com/v2ray/v2ray-core/common"
	"github.com/v2ray/v2ray-core/common/alloc"
)

// Writer extends io.Writer with alloc.Buffer.
type Writer interface {
	common.Releasable
	// Write writes an alloc.Buffer into underlying writer.
	Write(*alloc.Buffer) error
}

// AdaptiveWriter is a Writer that writes alloc.Buffer into underlying writer.
type AdaptiveWriter struct {
	writer io.Writer
}

// NewAdaptiveWriter creates a new AdaptiveWriter.
func NewAdaptiveWriter(writer io.Writer) *AdaptiveWriter {
	return &AdaptiveWriter{
		writer: writer,
	}
}

// Write implements Writer.Write(). Write() takes ownership of the given buffer.
func (this *AdaptiveWriter) Write(buffer *alloc.Buffer) error {
	defer buffer.Release()
	for !buffer.IsEmpty() {
		nBytes, err := this.writer.Write(buffer.Value)
		if err != nil {
			return err
		}
		buffer.SliceFrom(nBytes)
	}
	return nil
}

func (this *AdaptiveWriter) Release() {
	this.writer = nil
}
