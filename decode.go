package bmp

import (
	"errors"
	"io"

	"github.com/somebadcode/bmp/dib"
	"github.com/somebadcode/bmp/header"
)

type Bitmap struct {
	Header header.Header
	DIB    dib.DIB
}

var (
	ErrUnsupportedFormat = errors.New("unsupported format")
)

func Decode(r io.Reader) (*Bitmap, error) {
	h, err := header.Decode(r)
	if err != nil {
		return nil, err
	}

	var d dib.DIB

	d, err = dib.Decode(r)
	if err != nil {
		return nil, err
	}

	return &Bitmap{
		Header: h,
		DIB:    d,
	}, nil
}
