package dib

import (
	"bufio"
	"encoding"
	"encoding/binary"
	"errors"
	"fmt"
	"io"
)

type Type uint32

type DIB interface {
	Type() Type
	encoding.BinaryUnmarshaler
	encoding.BinaryMarshaler
}

const dibOffset = 14

const (
	TypeBitmapCoreHeader Type = 12
	TypeBitmapInfoHeader Type = 40
	/*DIBOS21BitmapHeader               = 12
	DIBOS22xBitmapHeader              = 64
	DIBOS22xBitmapHeaderSmall         = 16
	DIBBitmapV2InfoHeader             = 52
	DIBBitmapV3InfoHeader             = 56
	DIBBitmapV4Header                 = 108
	DIBBitmapV5Header                 = 124*/
)

var (
	ErrDIBIsTooSmall    = errors.New("dib header is too small")
	ErrInvalidDIBHeader = errors.New("invalid dib header")
)

func Decode(rd io.Reader) (DIB, error) {
	r := bufio.NewReader(rd)

	sizeBuf, err := r.Peek(4)
	if err != nil {
		return nil, fmt.Errorf("invalid DIB header: %w", err)
	}

	size := binary.LittleEndian.Uint32(sizeBuf)

	var d DIB

	switch Type(size) {
	case TypeBitmapCoreHeader:
		d = &BitmapCoreHeader{}
	case TypeBitmapInfoHeader:
		d = &BitmapInfoHeader{}
	default:
		return nil, fmt.Errorf("unsupported DIB format (%d): %w", size, ErrInvalidDIBHeader)
	}

	b := make([]byte, size)
	if _, err = io.ReadFull(r, b); err != nil {
		return nil, fmt.Errorf("failed to read DIB header: %w", err)
	}

	if err = d.UnmarshalBinary(b); err != nil {
		return nil, fmt.Errorf("failed to parse DIB header: %w", err)
	}

	return d, nil
}
