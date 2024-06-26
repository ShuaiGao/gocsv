package gocsv

import (
	"context"
	"errors"
	"testing"
)

func TestUnmarshalToCallback_ReaderError(t *testing.T) {
	type Dummy struct{}
	var reader = &errorReader{}
	ctx := context.Background()

	err := UnmarshalToCallback(ctx, reader, func(Dummy) {})
	if !errors.Is(err, readerErr) {
		t.Error("UnmarshalToCallback should return first reader error")
	}

	err = UnmarshalDecoderToCallback(ctx, newSimpleDecoderFromReader(reader), func(Dummy) {})
	if !errors.Is(err, readerErr) {
		t.Error("UnmarshalDecoderToCallback should return first reader error")
	}

	err = UnmarshalToCallbackWithError(ctx, reader, func(Dummy) error { return nil })
	if !errors.Is(err, readerErr) {
		t.Error("UnmarshalToCallbackWithError should return first reader error")
	}
}

type errorReader struct{}

func (e *errorReader) Read([]byte) (n int, err error) {
	return 0, readerErr
}

var readerErr = errors.New("reader error")
