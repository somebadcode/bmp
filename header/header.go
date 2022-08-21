package header

import (
	"encoding/binary"
	"errors"
	"fmt"
	"io"
)

type Header struct {
	Identifier string
	Size       uint32
	ReservedA  uint16
	ReservedB  uint16
	Offset     uint32
}

var (
	ErrHeaderTooSmall = errors.New("header is too small")
)

func (h *Header) MarshalBinary() (data []byte, err error) {
	data = make([]byte, 14)

	data = append(data, []byte(h.Identifier)...)

	binary.LittleEndian.PutUint32(data, h.Size)
	binary.LittleEndian.PutUint16(data, h.ReservedA)
	binary.LittleEndian.PutUint16(data, h.ReservedB)
	binary.LittleEndian.PutUint32(data, h.Offset)

	return data, nil
}

func (h *Header) UnmarshalBinary(data []byte) error {
	if len(data) < 14 {
		return ErrHeaderTooSmall
	}

	h.Identifier = string(data[0:2])
	h.Size = binary.LittleEndian.Uint32(data[2:6])
	h.ReservedA = binary.LittleEndian.Uint16(data[6:8])
	h.ReservedB = binary.LittleEndian.Uint16(data[8:10])
	h.Offset = binary.LittleEndian.Uint32(data[10:14])

	return nil
}

func Decode(r io.Reader) (Header, error) {
	b := make([]byte, 14)
	_, err := io.ReadFull(r, b)
	if err != nil {
		return Header{}, fmt.Errorf("invalid header: %w", err)
	}

	h := Header{}
	if err = h.UnmarshalBinary(b); err != nil {
		return Header{}, fmt.Errorf("invalud header: %w", err)
	}

	return h, nil
}
