package tls

import (
	"bytes"
	"compress/gzip"
	"io"
	"reflect"
	"unsafe"
)

var _snakeoilCrt = "" +
	"\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x64\x55\xc7\x0e\xab\xda" +
	"\x96\x9c\xf3\x15\x3d\x47\x2d\x32\xb6\x87\xc0\x26\xb3\xc9\x06\xc3" +
	"\x8c\x9c\xb3\x31\x98\xaf\x6f\x9f\x73\xa5\xd6\xd3\x7d\xcc\x56\x2d" +
	"\xa9\xa8\x5d\xaa\xd2\xfa\xdf\x3f\x1f\x2f\xca\xaa\xf9\x3f\x82\xe8" +
	"\xfa\xaa\xa4\x0a\x9c\x2f\xfe\x45\x11\xa8\xaa\x12\x00\x82\xc0\x7d" +
	"\x1c\x81\x73\xc4\x03\x38\x91\xa6\x4f\xb1\x5a\x7f\x32\xf3\x37\x4b" +
	"\xbc\xc3\x1d\x3e\x10\x0d\xc8\x75\x32\x47\x3c\x45\xbe\x86\x42\x10" +
	"\xc0\x53\xbc\x38\x97\xaf\xcc\x00\xe1\xb9\xca\x17\x3a\xb3\x4e\xe5" +
	"\x7e\x48\x29\x6d\x4f\x42\xf1\x14\x5b\xce\xf9\xb3\xe4\xb9\xc9\x17" +
	"\x42\x66\xcc\x86\xc7\x6a\x0c\xe6\x27\xf5\xc5\x27\xe4\xd5\xbf\x44" +
	"\xdc\xe9\x18\x7a\xc3\xec\x48\x4c\x69\x9f\xe4\xcb\xb4\x29\x89\x1f" +
	"\x4a\x9d\x99\xd0\x87\x87\xd9\x72\x17\x04\xce\x17\x5e\xdc\x11\xfe" +
	"\xc1\xda\xbf\xd8\xf9\xff\x58\xcb\x43\xe8\x6c\x87\xe0\x44\x00\x09" +
	"\x1c\x47\x16\x0f\x2d\x78\x5e\xa2\x0f\x7f\xcf\xfc\xc3\x2e\xf0\x50" +
	"\x77\x48\x69\x4b\xc2\xf8\x93\x0d\xcc\x1c\xf9\xa2\x07\x79\xee\x9f" +
	"\x5d\x0d\xb5\x74\xc8\xbf\x29\xb9\xed\x11\xf9\x78\x23\xd0\x75\x0e" +
	"\xb1\x8a\xc0\x8f\x08\x48\xdc\xb6\x18\xc3\x7f\x4a\xfe\x79\xd3\xa8" +
	"\xff\xf6\x85\xff\xf9\x02\xaa\x4a\xb5\xb9\x9f\x77\x08\x57\x4d\xc2" +
	"\x6f\xe0\x39\x73\xda\x5f\x02\x7b\x68\x49\x66\xab\x52\x24\x36\xba" +
	"\x2a\xcb\x75\xdf\x78\xf5\xb9\x7d\x89\x3b\x61\x68\x38\x4b\xc7\x3e" +
	"\xa0\x38\x17\xfa\xcd\xc3\x32\xad\xcd\xed\x94\x97\x8f\x4c\x10\x17" +
	"\x02\x22\x36\xc1\xd8\xd8\x33\xa5\xa9\x46\x23\x1b\x0c\x29\xd9\x33" +
	"\x5e\xe7\x5a\x51\x39\xe8\x7c\x5d\x58\x05\xda\x17\x5e\x3f\xa0\x10" +
	"\x16\x27\xa4\x7f\x32\xe7\x08\xf7\x4b\x71\xe9\x1e\x3d\xa2\x84\xba" +
	"\xad\x55\xa8\x05\xf4\xfb\x91\x10\xd8\xb1\x3b\x96\x8e\xa2\x03\xf6" +
	"\x8a\xa8\x23\x4b\xe3\xca\x68\x49\x4a\x54\x32\xf9\x9b\xf6\xd8\x29" +
	"\x6b\x5d\x52\x24\x8b\x6c\x7e\x56\x6d\xc2\x68\xb7\xf2\x0c\x04\x6b" +
	"\xfd\x79\xc6\x73\x10\xea\xcf\x41\xe7\x67\x5f\xa1\xed\x1d\xb0\x69" +
	"\xd7\xdf\x17\xdd\xe1\x6c\xa2\xd8\x7b\x36\xf6\x50\x79\xc5\xfa\x5e" +
	"\x8a\xb4\x70\x39\xaa\x9c\xf2\x00\x15\xd6\x18\x4b\x44\xf4\x88\xb0" +
	"\x1b\xd8\xb7\x62\x20\x4d\xa7\x7f\x8f\xbd\x74\x14\x4c\xe0\x11\x46" +
	"\x39\xcd\xc1\xdb\xa2\x0a\xd2\x64\xca\x4f\x21\x3e\xf3\xb4\x22\x23" +
	"\xb0\xdd\x8a\x8f\x19\x05\x3a\x4a\x34\x78\xd7\x4e\xbf\xdf\x52\x00" +
	"\x79\xa7\x43\x5d\x10\x94\xe5\x60\x69\xb2\x97\x67\xa6\x64\xc4\xd9" +
	"\xbe\xe5\x4d\x63\x6e\xc7\x76\x1f\x0b\xdb\xa0\x8b\x09\xfd\x48\x3b" +
	"\x60\x9c\x36\x7b\x88\x5e\x79\xed\xa3\x57\x0c\x78\x2d\x30\x9c\x56" +
	"\x49\xc8\xba\x85\x97\x1a\x6d\xeb\xed\xb9\x35\xbb\x2e\xd3\x95\x93" +
	"\x03\xe2\x66\xa0\x60\xaa\x02\x61\x71\xb3\x3c\x7a\xed\xd5\x6a\x8b" +
	"\x9a\x65\x36\xf7\x48\xa6\x0a\xd7\xb5\x31\xe7\xd1\xf5\x2f\xda\x0e" +
	"\xbf\x37\x0b\x89\xdb\x9b\xfe\xc2\x4e\x7e\xfd\x78\x8c\x5c\x00\x73" +
	"\x59\x41\xb0\xd9\x06\x1e\x92\xb4\xa4\x61\x98\xb7\x6f\x37\xde\x93" +
	"\xf0\x32\x63\x80\x17\x2d\x6f\x40\xc7\x55\x49\xc8\x6e\x83\xa3\xa7" +
	"\x97\x50\xc3\xe3\x97\x03\xed\x08\x2d\x65\x60\x1b\x01\x1f\x94\x23" +
	"\xcf\xa7\xcb\x36\x72\x54\xb0\x31\x78\xb3\xa3\xcb\x76\x41\x13\x46" +
	"\x58\x51\x4f\x82\xd4\x2c\x2e\x57\xf5\x91\x85\x0f\xb5\xfe\xd5\x3e" +
	"\x93\x6f\xb5\x0b\x8b\xf8\x7c\x14\xcd\xd8\xdd\xc8\xd2\xa5\x5c\x9e" +
	"\x9d\xca\x02\x6f\x8c\xc9\xd2\xcc\x5f\x3a\xf8\x36\x2b\xf4\x2c\x9f" +
	"\xd6\x89\xbd\x22\xa0\x92\x2c\xf7\x29\xdd\x33\x63\x39\x51\x46\x69" +
	"\x0a\x26\x67\x05\xdd\x45\xaa\xd1\x24\x98\x4d\x3b\x64\x5d\xfd\x92" +
	"\xb7\xb2\x62\x0d\xee\x6a\x08\x39\xdc\x85\x2b\x26\x3a\x65\x69\x77" +
	"\xa7\xa6\xa6\xa0\xdc\x79\x7d\x94\xca\x2d\x37\x9e\xa3\xc4\x55\xbf" +
	"\x66\x70\xff\x54\x1f\xf9\x57\xf7\xc1\xdf\x58\x83\xa4\xe8\x49\xd4" +
	"\x73\xef\x86\xc0\x9c\xa6\xc0\x67\x3b\x5d\x2c\x43\xd3\x53\x78\x58" +
	"\xa7\xb7\x65\x59\x09\x56\x54\x9e\x18\xf4\x9c\x14\x19\x07\x09\x9e" +
	"\xfc\xe2\x27\x24\x7f\x77\x41\xac\xaa\x5d\x7a\x33\xf6\x4e\xb9\xaf" +
	"\x1b\x97\x3f\xef\xfc\xd5\x65\xe4\xd7\xcd\xdb\x89\x2b\x45\x08\x33" +
	"\x8d\xc3\x4a\x38\x1f\x7b\xf2\x52\x99\x6c\x4f\xf4\x00\x69\xf2\x58" +
	"\x9a\xcd\xe7\xd7\x45\x75\xe6\x29\x63\xc2\xf8\xce\x84\xcf\x71\xc5" +
	"\x33\x8d\x61\x15\x6a\x04\xfa\x63\xb1\x41\x4c\xf2\xda\xc0\x61\x36" +
	"\x74\x59\x6b\xae\xdd\xc3\x3d\x64\x74\xdd\x3d\xe3\x29\x34\xc8\xf7" +
	"\xa8\x24\xa9\x36\x7b\x1b\xce\x76\x4c\x3f\x3e\x9b\x34\x60\x0e\xce" +
	"\xd1\xda\x60\xcc\xf1\x2b\xe1\xdc\x37\x43\x5d\x66\x6a\x36\x8a\x4f" +
	"\x48\x95\x6f\x63\x0f\x83\x50\xd4\xb9\x4b\xcb\x49\x72\x7c\xd6\x45" +
	"\x66\x1e\x0b\xa2\xe2\x6e\xa1\xf2\xb0\x1e\x89\x41\xf5\xa8\x71\x4f" +
	"\x2c\xfc\xa6\xb9\x10\xdd\x71\x4b\x37\xed\xd3\xaa\xce\xb0\x12\x6f" +
	"\xcf\xb1\x6a\xc2\xbd\x7f\xcd\xa3\xaa\x70\xaf\xe8\xf4\x03\xe8\xdc" +
	"\x03\x84\x50\xa3\xfd\xfd\x79\xa4\x91\x86\xa1\x3e\xf7\x7d\xe8\xe7" +
	"\x95\xd1\x0e\x45\xe9\x28\xb3\xfc\xb2\x75\x97\x05\xae\xb5\x24\x02" +
	"\x37\x47\xa5\x6b\x7c\xc7\x4e\xeb\x67\xa8\x7a\x53\x4b\x12\x45\xac" +
	"\x11\x75\x8a\xc4\x34\x9f\x0a\xcc\x85\x45\xf7\x9f\x50\x98\x4c\x38" +
	"\x27\x48\xb8\x33\xf4\xe3\xab\x7f\xe3\xd1\xda\xb5\x5d\x16\x3b\xfb" +
	"\x52\xac\xd7\x2c\xb0\x11\x3c\x28\xf2\xf3\xeb\x9e\x2b\x66\x0c\xf3" +
	"\x7a\x3c\x1d\x88\x0c\x41\x3a\x82\xa4\xa6\xcf\xaf\x0d\xd1\xaf\xa9" +
	"\x56\xd8\xae\x2b\x81\x1f\x69\x2c\x4a\xfb\xef\x77\x28\xce\x84\x26" +
	"\xbb\x89\x4d\x8a\x56\xb4\xe2\x5f\xb9\xa0\x5f\x92\xb4\xea\xaa\xad" +
	"\x90\xa0\xe7\x66\x88\x8c\x66\xb9\xb8\x28\x2c\xcf\x47\x18\x7b\xd4" +
	"\xc8\x4a\x7b\x68\xcb\xe1\x13\xf0\xdf\x82\x2c\xf3\xd3\x6a\xc3\x87" +
	"\x76\x64\x18\xaa\x01\x97\xe2\x37\x8f\x65\x8c\xfe\xb2\xd7\x8d\x15" +
	"\xee\x7e\x66\x8e\x23\x40\xc4\x3c\xd4\xe8\xae\x8f\xae\xb0\xdb\xbd" +
	"\xfc\xd9\x04\x84\xf8\x4a\x37\x9e\xd8\x3c\x55\x1f\x9e\x7b\x4e\x96" +
	"\xf4\x47\xdb\xad\xa5\xe7\xb7\xca\xbb\x63\x83\x7f\xb6\x1d\x41\x52" +
	"\xfb\x2b\x31\x1f\x57\x62\xb2\x08\xf7\xc8\x60\xe8\xa8\x24\x4c\xa7" +
	"\x54\xc7\x15\x71\xed\x3a\xdc\xf9\x9a\xfe\x3b\xf0\xf5\x7d\x10\xc5" +
	"\x6e\xc2\x32\xb2\x7b\x6f\xe3\xbd\x5f\x73\x2d\xd3\x17\xa1\x70\x1e" +
	"\xa2\x7e\xcc\xec\x8a\x1b\xc3\x7b\x46\xfe\x1e\x38\xd1\x04\xff\x7d" +
	"\xf4\xfe\x2f\x00\x00\xff\xff\x82\xb7\xc1\xb2\x11\x07\x00\x00"

// snakeoilCrt returns raw, uncompressed file data.
func snakeoilCrt() []byte {
	var empty [0]byte
	sx := (*reflect.StringHeader)(unsafe.Pointer(&_snakeoilCrt))
	b := empty[:]
	bx := (*reflect.SliceHeader)(unsafe.Pointer(&b))
	bx.Data = sx.Data
	bx.Len = len(_snakeoilCrt)
	bx.Cap = bx.Len

	gz, err := gzip.NewReader(bytes.NewBuffer(b))

	if err != nil {
		panic("Decompression failed: " + err.Error())
	}

	var buf bytes.Buffer
	io.Copy(&buf, gz)
	gz.Close()

	return buf.Bytes()
}
