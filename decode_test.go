package bmp

import (
	"encoding/json"
	"os"
	"os/user"
	"path/filepath"
)

func ExampleDecode() {
	userInfo, err := user.Current()
	if err != nil {
		panic(err)
	}

	var f *os.File

	f, err = os.Open(filepath.Join(userInfo.HomeDir, "Downloads", "courtyard.bmp"))
	if err != nil {
		panic(err)
	}

	defer f.Close()

	var bitmap *Bitmap

	bitmap, err = Decode(f)
	if err != nil {
		panic(err)
	}

	// Print the resulting header and DIB as JSON.
	enc := json.NewEncoder(os.Stdout)
	enc.SetIndent("", "    ")
	if err = enc.Encode(bitmap); err != nil {
		panic(err)
	}

	// Output:
	// {
	//    "Header": {
	//        "Identifier": "BM",
	//        "Size": 720056,
	//        "ReservedA": 0,
	//        "ReservedB": 0,
	//        "Offset": 54
	//    },
	//    "DIB": {
	//        "Size": 40,
	//        "BitmapWidth": 600,
	//        "BitmapHeight": -400,
	//        "ColorPlanes": 1,
	//        "BitsPerPixel": 24,
	//        "CompressionMethod": 0,
	//        "ImageSize": 720002,
	//        "HorizontalResolution": 11811,
	//        "VerticalResolution": 11811,
	//        "ColorsInPalette": 0,
	//        "ImportantColors": 0
	//    }
	// }
}
