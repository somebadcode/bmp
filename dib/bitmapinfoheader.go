package dib

import (
	"encoding/binary"
)

type BitmapInfoHeader struct {
	Size                 uint32
	BitmapWidth          int32
	BitmapHeight         int32
	ColorPlanes          uint16
	BitsPerPixel         uint16
	CompressionMethod    uint32
	ImageSize            uint32
	HorizontalResolution int32
	VerticalResolution   int32
	ColorsInPalette      uint32
	ImportantColors      uint32
}

func (h *BitmapInfoHeader) MarshalBinary() (data []byte, err error) {
	data = make([]byte, TypeBitmapInfoHeader)

	binary.LittleEndian.PutUint32(data[14-dibOffset:18-dibOffset], h.Size)
	binary.LittleEndian.PutUint32(data[18-dibOffset:22-dibOffset], uint32(h.BitmapWidth))
	binary.LittleEndian.PutUint32(data[22-dibOffset:26-dibOffset], uint32(h.BitmapHeight))
	binary.LittleEndian.PutUint16(data[26-dibOffset:28-dibOffset], h.ColorPlanes)
	binary.LittleEndian.PutUint16(data[28-dibOffset:30-dibOffset], h.BitsPerPixel)
	binary.LittleEndian.PutUint32(data[30-dibOffset:34-dibOffset], h.CompressionMethod)
	binary.LittleEndian.PutUint32(data[34-dibOffset:38-dibOffset], h.ImageSize)
	binary.LittleEndian.PutUint32(data[38-dibOffset:42-dibOffset], uint32(h.VerticalResolution))
	binary.LittleEndian.PutUint32(data[42-dibOffset:46-dibOffset], uint32(h.HorizontalResolution))
	binary.LittleEndian.PutUint32(data[46-dibOffset:50-dibOffset], h.ColorsInPalette)
	binary.LittleEndian.PutUint32(data[50-dibOffset:54-dibOffset], h.ImportantColors)

	return data, nil
}

func (h *BitmapInfoHeader) UnmarshalBinary(data []byte) error {
	if len(data) < int(TypeBitmapInfoHeader) {
		return ErrDIBIsTooSmall
	}

	h.Size = binary.LittleEndian.Uint32(data[14-dibOffset : 18-dibOffset])
	h.BitmapWidth = int32(binary.LittleEndian.Uint32(data[18-dibOffset : 22-dibOffset]))
	h.BitmapHeight = int32(binary.LittleEndian.Uint32(data[22-dibOffset : 26-dibOffset]))
	h.ColorPlanes = binary.LittleEndian.Uint16(data[26-dibOffset : 28-dibOffset])
	h.BitsPerPixel = binary.LittleEndian.Uint16(data[28-dibOffset : 30-dibOffset])
	h.CompressionMethod = binary.LittleEndian.Uint32(data[30-dibOffset : 34-dibOffset])
	h.ImageSize = binary.LittleEndian.Uint32(data[34-dibOffset : 38-dibOffset])
	h.VerticalResolution = int32(binary.LittleEndian.Uint32(data[38-dibOffset : 42-dibOffset]))
	h.HorizontalResolution = int32(binary.LittleEndian.Uint32(data[42-dibOffset : 46-dibOffset]))
	h.ColorsInPalette = binary.LittleEndian.Uint32(data[46-dibOffset : 50-dibOffset])
	h.ImportantColors = binary.LittleEndian.Uint32(data[50-dibOffset : 54-dibOffset])

	return nil
}

func (h *BitmapInfoHeader) Type() Type {
	return Type(h.Size)
}
