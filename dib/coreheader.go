package dib

import (
	"encoding/binary"
)

type BitmapCoreHeader struct {
	Size         uint32
	BitmapWidth  uint16
	BitmapHeight uint16
	ColorPlanes  uint16
	BitsPerPixel uint16
}

func (h *BitmapCoreHeader) MarshalBinary() (data []byte, err error) {
	panic("not implemented")
}

func (h *BitmapCoreHeader) UnmarshalBinary(data []byte) error {
	if len(data) < int(TypeBitmapCoreHeader) {
		return ErrDIBIsTooSmall
	}

	h.Size = binary.LittleEndian.Uint32(data[:4])
	h.BitmapWidth = binary.LittleEndian.Uint16(data[4:6])
	h.BitmapHeight = binary.LittleEndian.Uint16(data[6:8])
	h.ColorPlanes = binary.LittleEndian.Uint16(data[8:10])
	h.BitsPerPixel = binary.LittleEndian.Uint16(data[10:12])

	return nil
}

func (h *BitmapCoreHeader) Type() Type {
	return Type(h.Size)
}
