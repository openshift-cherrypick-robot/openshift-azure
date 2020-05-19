// Package arm Code generated by go-bindata. (@generated) DO NOT EDIT.
// sources:
// data/master-startup.sh
// data/node-startup.sh
package arm

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func bindataRead(data []byte, name string) ([]byte, error) {
	gz, err := gzip.NewReader(bytes.NewBuffer(data))
	if err != nil {
		return nil, fmt.Errorf("read %q: %v", name, err)
	}

	var buf bytes.Buffer
	_, err = io.Copy(&buf, gz)
	clErr := gz.Close()

	if err != nil {
		return nil, fmt.Errorf("read %q: %v", name, err)
	}
	if clErr != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

type asset struct {
	bytes []byte
	info  os.FileInfo
}

type bindataFileInfo struct {
	name    string
	size    int64
	mode    os.FileMode
	modTime time.Time
}

// Name return file name
func (fi bindataFileInfo) Name() string {
	return fi.name
}

// Size return file size
func (fi bindataFileInfo) Size() int64 {
	return fi.size
}

// Mode return file mode
func (fi bindataFileInfo) Mode() os.FileMode {
	return fi.mode
}

// ModTime return file modify time
func (fi bindataFileInfo) ModTime() time.Time {
	return fi.modTime
}

// IsDir return file whether a directory
func (fi bindataFileInfo) IsDir() bool {
	return fi.mode&os.ModeDir != 0
}

// Sys return file is sys mode
func (fi bindataFileInfo) Sys() interface{} {
	return nil
}

var _masterStartupSh = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xbc\x59\x5f\x77\x1b\xb5\x12\x7f\x46\x9f\x62\x58\x07\x4a\x4b\xe4\x4d\x0a\xdc\x0b\x86\xf4\x9c\x34\x49\x39\xbd\x94\x24\x37\x69\x2f\x0f\xa5\xa7\x47\x5e\xcd\xda\xaa\xb5\xd2\x22\x69\xed\x98\xd4\xdf\xfd\x9e\x91\xd6\x8e\x1d\x3b\x49\x4b\x03\x7d\x48\x6d\x69\x34\x33\x9a\xf9\xcd\x3f\xb9\xf3\x79\xde\x57\x26\xef\x0b\x3f\x04\x8e\x17\x8c\x75\xe0\x99\x75\x10\xd0\x07\x65\x06\x3d\xd0\x76\x00\xc2\x48\x90\xce\xd6\x20\xb4\x86\xe0\x44\x59\xaa\x02\xc2\x50\x04\x98\xd8\x46\x4b\x70\xb6\x09\x08\x63\x25\x20\x0c\x11\x2a\xe1\x03\x3a\x38\x7a\xf1\x94\x75\xe0\xec\xe8\xfc\xe4\xd5\xd9\xc1\xd1\xcf\x67\x27\xaf\x4e\xf7\xb2\xa9\x6d\x1c\x77\xe8\x6d\xe3\x0a\xe4\x03\x67\x9b\x3a\x63\x1d\x38\x39\x7f\xfb\xec\xbf\x87\xc7\x7b\x99\xad\xd1\xf8\xa1\x2a\x43\x77\x6b\xe5\x64\xd7\x7a\x21\x71\xdc\x2d\xb4\x6d\x64\xc6\x3a\xac\x03\xaa\x0e\xa2\xaf\xd1\x03\x7f\x0e\xcf\x8f\x4f\x5f\xbd\x04\xee\x61\xeb\x2b\xa9\x06\xf0\xb5\x1f\x5a\x17\x20\xdb\x6a\xf9\x66\xf0\x1e\x82\x50\x1a\xf8\xee\x43\xe0\xef\xe0\xc5\xc9\xcf\xc0\xb9\xb6\x03\x5e\x3b\x2c\xd5\x05\x64\xbf\xbc\x7a\x7a\x04\x44\x0a\x87\x67\x27\xa7\xbd\xec\xd3\xf8\x13\x0f\xc6\x2e\x2f\x41\x95\xd0\x3d\xb0\xa6\x54\x83\xee\x39\x16\x8d\x53\x61\x7a\x2a\x42\x31\x3c\x15\xc5\x48\x0c\xd0\xc3\x6c\xc6\xb4\x1d\x0c\xd0\x01\x0f\xad\xe1\xb8\x0f\xc2\x85\xa6\xee\xfa\x21\x64\xca\xf8\x20\xb4\x56\x66\x00\x0e\x25\x90\xc9\x0b\x69\xa0\x88\x3c\x1b\x27\x82\xb2\x06\xac\x81\xad\xaf\x86\xd6\x07\x23\x2a\x7c\x98\xb1\x42\x04\x78\x92\x8f\x85\xcb\xb5\xea\xe7\xd3\xa6\xca\x0b\xad\xd0\x04\x5e\xa0\x0b\xdd\x1a\x2b\xf8\xe9\xa7\x07\x47\x27\xcf\x1e\x90\x8a\x07\xe8\xc2\xbe\x7f\x3a\x0d\xe8\x17\xba\xd2\x9a\x2a\x55\x21\x02\xfa\x6e\xab\xeb\x19\xd6\xd6\xab\x60\xdd\x34\x6e\xc3\x7b\x38\x0f\x8e\xf4\x9a\xcd\xd8\xd1\xc9\xb3\x9b\x85\x8e\x70\x7a\x5d\xe6\xa9\x53\x63\x11\xf0\x17\x9c\x7e\xa4\xe4\x5f\x70\xba\x26\xf8\x83\x0d\xb8\x7f\x76\x02\xbe\xf5\x02\x34\xb5\x24\x19\xf0\xfa\xf2\xb2\xe5\xe7\xff\x63\x95\xb9\xc3\x5d\xd9\x36\x64\x30\x9b\xbd\x59\x33\x79\x69\x1d\x88\x10\xb0\xaa\x03\x28\x03\x97\xbb\xdd\xee\x77\xb3\x1f\x41\x5a\x06\x30\x6d\x2a\x68\xd5\x00\x3e\x05\xfe\x07\x7c\x9c\xcc\x28\x12\xbe\xfc\x12\xfa\x0e\xc5\x88\x01\xdc\x7a\xe1\xd7\x73\x35\xb6\x2e\xdb\x4f\xb3\x37\x9b\xaf\xde\xea\x94\x30\x54\x0a\xa5\x51\x66\x0c\x08\xb3\xaf\x5f\x2f\x9d\x06\xae\x03\x7c\x07\x6f\xde\xfc\x48\xd1\x6d\xc0\x6b\xc4\x1a\x76\x7f\x04\xd4\x1e\x01\x2f\x54\xa0\x2f\xa5\x62\xd2\x1a\xbc\xc3\x1b\x0e\x2b\x3b\xfe\x38\x30\x93\xf5\x0a\x8d\xc2\x50\xf2\x61\xae\x02\xee\x4a\xb8\x15\xdc\xb7\x80\x90\x5d\x5e\xa2\x91\xb3\x19\x65\xb9\xc2\xa1\x08\x48\xd2\x83\x50\x06\x1d\xd4\x8d\xd6\x64\x25\x87\x81\x55\x23\xa9\x1c\xf0\xfa\x8a\x99\x75\x6a\xa0\x4c\xde\x95\xb6\x18\xa1\xbb\x06\xf7\xd5\xcd\x3c\xdd\xa8\xfb\xce\x5b\xb3\x0c\xfb\xee\x21\x3a\x35\x46\xd9\x3d\xb0\x55\x5f\x19\x94\xcf\x2b\x31\xc0\xd3\x46\xeb\xf3\x28\x75\x0e\x84\x35\x88\x6b\x43\xb9\xe7\x06\x69\x90\x3b\x6b\x43\x4e\x57\x7a\x79\x72\x78\xd2\x03\x89\x1a\x03\xc6\x54\x5c\x5a\xad\xed\x84\x38\xc5\x54\x9b\xee\x4c\x56\x16\x25\xa5\x68\x15\x40\x79\xe8\x8b\x11\x4a\x50\x26\x58\xb0\x8d\x83\xff\xfd\x0a\x8a\xf4\xf2\x2c\x9e\x11\x52\x02\x2f\xa1\xbd\x36\x53\x25\x7c\x0e\x03\x87\x4b\x96\x99\xab\x81\xa1\xc8\x4b\x1f\x44\x3f\x01\x85\x01\xf8\xa9\x0f\x58\x15\x41\x83\x0f\xb6\x6e\x79\xf0\xe8\xcd\xa6\xee\x06\x55\xa1\xbb\x93\xca\xa3\x1b\xab\x02\x6f\xa2\x5b\xda\xaf\x46\xa5\xef\x5e\x94\x9e\xd4\xcd\x25\x8e\x73\xa9\xfc\x28\x17\x7f\x36\x0e\xf3\x45\xc9\xa9\x85\x0b\xbb\x0c\x00\x8b\xa1\x85\x07\xb7\x93\xc1\xda\x1d\x81\xd8\xc3\xc0\xd5\x7f\x34\x36\x08\x80\x1d\xd8\x79\x00\x4f\x9e\x5c\x5d\x9d\xd4\xb0\x8d\x09\xd7\x4f\x32\x00\x87\x3e\x58\x87\x85\x35\xc0\xcf\x36\xec\x27\x44\x11\xa7\x16\x45\x52\x60\x65\xcd\x35\x14\x31\x80\x8c\x0a\x97\x24\x24\xb9\xac\x07\xd9\x3b\xdb\x38\x23\xb4\xcc\xb6\x69\x4f\x2a\x4f\x55\x8b\x6b\x1c\x88\x62\xca\x1d\x0e\x94\x0f\x6e\x9a\xf5\x20\xb8\x06\x59\xc2\xd3\xaa\x2d\x85\x0b\xeb\xc6\xdc\x4c\x70\xcd\x77\xa5\x62\xac\xb5\x4c\x0c\x1e\xc2\x78\x9b\xcb\x22\xb4\x7d\xf7\xd8\x4a\x8c\xd9\xeb\x49\x34\xb5\x21\xaa\x2f\x37\xa2\x08\x43\x21\x37\x61\x68\xe1\xd5\xeb\xbe\xf2\x85\x57\xbb\xb9\x6e\xcc\x0e\xbc\x7f\x9f\x6e\x77\x93\x5b\x97\x48\xaf\x09\x4c\x0e\x95\x58\x8a\x46\x07\xff\x41\x0e\xa5\x73\x37\xbb\x33\xee\x92\x5d\x3a\x20\x8a\x02\x6b\x6a\xa2\xe0\xfb\x6f\xbf\xfd\x06\xa8\x44\x50\x4c\x0a\x59\x29\xef\x29\x08\x29\xf5\x38\xab\x35\x59\xd2\x3a\x90\x3e\xd6\x8e\x50\xd4\xdb\xf1\x40\xfb\xe1\xdb\x58\x46\x3e\xab\x9d\x0d\x76\x6f\xeb\x52\xfa\xf0\xc5\x17\xdb\x8f\x66\xec\xb3\xda\xba\x90\x16\x3a\x9d\x47\xdb\x33\xf6\xd9\x55\xc7\xb2\x1f\x3b\xaa\xe7\x67\x47\xbf\xed\xbf\x78\xf1\x76\xff\xc5\x8b\x93\xdf\x28\x99\x6d\x45\x26\xc0\x2b\x72\x6a\x40\xe0\x3c\xfd\x7f\x7c\xf4\x1b\x2d\xce\xb7\xb9\x24\xd6\xb0\x15\xff\xf2\x77\xb0\x7f\x70\x70\x74\xfa\x12\xf8\xa4\x4d\xf1\x73\x39\xdc\x8b\x31\xb6\x98\xf5\x53\x9f\xb2\x5e\x3e\xdf\xbd\xa3\x14\x10\x60\xc8\x36\xeb\x98\x39\x4f\x54\x30\x9b\xdd\x5e\x57\x6f\x47\xde\x15\x97\x4f\x2b\x9d\x1f\x2c\xe5\x63\x2b\xe8\xbf\x76\x6e\x2a\xa1\xd4\xdd\x1e\x9f\xbc\x3c\xea\xc1\x73\x03\x65\x13\x1a\x87\xdb\x50\xd9\x31\xa6\x9e\x5b\x99\xd2\xba\xaa\xad\x96\x4d\xf0\x4a\x22\xd8\x12\xd0\x8c\x95\xb3\xa6\x42\x13\x60\x2c\x9c\x4a\x4e\xe8\x30\x8f\x01\xbe\xbe\x60\x78\x11\xdd\x79\xbe\x7f\xfe\xea\xec\xf9\xde\x83\xa5\xab\xfc\x1a\x2d\xd1\xde\x24\xed\xc3\x6c\xf6\x20\x1e\xe4\x71\x20\x70\x8d\x89\xd0\x6d\x8d\x05\x9c\x2b\xa3\x02\x04\x0b\x7d\x6b\x83\x0f\x4e\xd4\x70\x78\x7c\x0e\x1e\x43\x53\xcf\x33\x02\x1d\xe2\xbc\x76\x6a\xac\x34\x0e\x50\x02\xe7\x54\xbd\xb9\xc1\x30\xb1\x6e\x04\x54\xe3\x81\x8f\x21\xef\xe5\xf4\xb1\xf7\x27\x70\x6c\xd5\xbb\xdd\xd0\xab\x5a\xcc\xf9\x45\x3d\x31\x25\xab\xb9\x88\x4a\x18\x41\xee\x0e\x16\x6a\x55\x8c\xa0\xa9\xc1\xe0\x24\x49\xf6\x18\x28\x36\x7d\x0c\xb9\x21\x82\x1c\xa6\x56\x81\xc5\x71\xe8\x2a\xfd\xcd\x79\x1e\x27\x9e\xbf\x26\x96\xac\x43\x0c\xe2\xc1\x08\x7f\x2a\x1c\x7a\xdc\xa5\x10\x60\x71\x61\x95\x9c\x92\x51\x4d\x5d\x1d\xba\xae\xcc\x7f\xf8\x81\xa7\x12\xce\xa5\xf1\x5d\x3f\x5c\xd2\x5c\x1a\x5f\x09\xff\x07\x69\x3c\x68\xf9\x93\xc6\xcb\xca\xa6\x12\x91\xe8\x5a\x81\x9b\x35\x9e\xd3\xcc\x93\x3a\xeb\xc0\x24\xc2\x8f\x76\x29\xf4\x0c\x25\xe7\x89\x10\x03\x82\x0c\xcd\x78\x73\xbb\x35\x41\x69\x15\x14\x7a\x18\xd8\xd8\x28\x06\x0b\x4e\x14\xb1\x59\x92\x8a\x80\xd7\xa5\x01\xa9\x5c\x1c\x76\x8d\xf1\xd0\xc7\xd2\x3a\x24\xb1\xd4\x55\x8c\x8c\x9d\x18\xba\x47\xbc\x43\x94\x84\x80\x46\x92\x13\x26\x2a\x0c\x81\x42\x63\x0a\x3e\x36\x3a\x6c\x32\x54\x1a\x63\xd4\x2c\x7a\x3f\xe0\xf2\x21\xec\xed\x41\x96\xc5\xc8\x91\xf6\xaa\xf3\x4c\x61\xf2\x0f\xe1\x8c\x35\x86\x5c\x9d\x4e\x30\x96\xba\x67\x5e\x08\x1e\x5c\xe3\x03\x4b\x48\x88\xf6\x6c\x6a\x18\xa0\xc1\xb1\x88\x69\x86\x56\x7c\x10\xc5\x08\x84\x07\x6f\xa9\xe9\xf2\xd1\x18\xab\xfd\xae\xf2\xa0\x85\x92\x14\xc9\xd0\x9f\xb2\xce\x4a\xa0\x2d\x9a\xd3\xed\x74\x52\x5b\x4f\x70\x1e\xaa\x68\xe2\xd6\xb8\x37\x10\x57\xd6\x21\xeb\x90\x2a\x1e\x4a\x67\xab\x15\xda\xda\xd9\x02\xbd\x27\x9f\x4c\x14\xb5\xbd\x43\x55\x27\xd0\x91\xfe\x2c\xa9\xe1\x11\xfc\x30\x0d\xf8\x0d\x35\xe6\x05\x82\x00\x29\xa6\x60\x8d\x9e\xd2\x6d\x6a\x4c\x25\x4d\xda\xc2\xb3\xbc\xf1\x2e\xd7\xb6\x10\x3a\xe2\x51\xfc\xe9\xb1\x90\xed\x65\xa9\x7d\xed\x0b\x8f\x5a\x19\xf2\x2b\x9c\xee\x1e\xde\x49\xef\x6d\x19\x26\xc2\x7d\x30\x7d\xa1\x45\x25\xc6\x73\x6a\xd6\x01\x34\x94\x02\x23\xb0\x53\x40\xac\x7a\xa5\x0d\x0b\xcf\xae\x22\xa7\x31\x95\xf0\x23\xa8\xa4\x97\xf3\xa8\x81\x24\x67\xf5\x6b\x65\xcd\xd5\x4a\xa9\x1b\x34\x61\xf1\x7d\x89\x5d\xab\xc0\x7d\xb1\x4b\x97\xf8\x34\x6e\xac\x03\xa7\xca\xc0\xa8\xe9\x63\xb2\x5c\x44\x51\xe3\x11\xa2\x65\x41\xd4\x8a\x13\x2d\x3a\xe6\x29\x94\x14\x70\x87\x90\xf9\xce\x57\xf0\x28\xad\xf7\xe0\x61\xf7\x51\xe7\xf7\xdd\x61\x08\xb5\xef\xe5\xf9\xd2\xb0\xd6\xc9\x52\x72\x6a\xe7\x93\x54\x61\x73\xea\x77\x4c\xf7\x4a\xe2\xbd\x31\x5e\xbc\x15\xf1\xb4\x70\xaf\x32\x28\x63\xc5\x3f\xf7\xcf\xd5\xcb\x7b\x30\x47\x9c\xf9\x22\x9b\x76\xda\x64\xec\xf2\x92\x53\x4e\x36\x08\x5b\xdd\xa7\xa2\x18\x35\xf5\x53\x6d\xfb\xc7\x94\x4a\xb3\xec\xce\x97\xa6\x45\x55\xa0\x9e\x76\x8c\x6e\xba\x36\x89\x53\xa6\x0b\x94\x80\x17\xa5\xa9\x1f\xa5\xc4\xa1\xfc\xac\x5c\xed\x81\xf3\x47\x8c\x9a\x1f\xd2\xe3\x50\xb9\xbd\xd5\xbd\xf6\x5c\x1a\xb2\xb7\x96\xe8\xfe\x72\xcf\x78\x14\x0a\x99\xee\xfc\x89\x6d\xe3\x0a\xa3\xbf\xb3\x73\x5c\x15\x74\x7f\xcd\xe3\xad\x7a\x4a\x3b\x31\xda\x0a\x49\x46\x4c\x4e\xc8\x56\xcb\xe8\x7a\xe5\xfc\x9d\x41\xac\x9e\x6b\xf1\xd7\x5b\x5f\xda\x44\x1c\x5f\x6c\x6b\x67\xc7\x4a\xa2\xcb\x7b\xf9\x5b\x29\x82\xc8\xdf\x52\xb9\x6b\xa9\x97\x01\xd0\xcb\x6d\x43\x25\x9a\xb6\xee\xb2\x19\x41\x29\x5d\x22\x71\xe2\xfd\x16\xee\x7b\x74\xf2\x5a\x04\xcc\x66\x2d\x91\x8c\x0f\xdb\xb1\xf6\xee\x91\xb0\x16\x8c\x5d\xd9\x6f\x09\x44\x11\xf7\xe6\xa6\xba\xdd\xa0\xad\xfc\x39\x31\xb9\x70\x1e\x26\x8f\xe7\x23\xe2\x5f\xc5\x74\xea\xcb\xe9\xce\x9f\x88\xe9\x15\x46\x7f\x27\xa6\x57\x05\xfd\x43\x98\x4e\x56\x8e\x75\xdd\x88\xda\x0f\x6d\xf8\x28\x4c\x13\x8a\x7a\x8b\x4f\x8b\xad\xe5\x7c\xd5\x5b\xfd\x96\xd0\xc9\x11\x8e\x5e\x1e\x1c\x1e\xbc\x7c\xf1\x76\xff\xf4\xf9\x5e\xf6\x4d\x76\x03\x68\x57\x8d\x42\x34\xc4\x25\x16\xf4\x56\xdf\x39\x50\x56\x22\x61\x0d\x97\x14\x37\x9c\x12\xe6\x6a\x2e\xa5\xf1\x20\x11\xc4\xa6\x79\x29\x63\xb7\xcb\x34\x24\x29\xa1\x79\xa1\x9b\x18\xa3\x59\x6b\xc3\x9d\xf8\x6f\x6f\x5e\x5f\x56\x56\x7b\x8f\xbf\xf9\x7e\x67\x7b\x79\x69\x77\x23\xe1\xee\x3a\xe1\xe3\x8d\x84\x8f\x23\x61\xb6\x59\x25\x1e\xec\x08\x4d\x34\x0b\x2f\xad\xe3\xf1\xd1\xe6\x1a\xa9\x90\x63\x74\x41\x79\xe4\x35\xa2\xe3\x8d\xd3\x1e\x36\x94\xc6\x28\x86\xb1\x6a\xbc\x6e\xa5\xfc\xd1\xb5\xb5\xb5\xc7\xe4\x85\x3d\x57\x4a\xd2\xca\x43\xcf\x35\xbe\x1f\x82\x4c\x8c\xe3\x4a\x16\xcb\x33\xcd\x3f\xb3\x19\x63\xa1\x31\x28\xb9\x90\x15\x35\xe2\x25\x8d\x3e\x57\xcd\x4c\xfb\x2a\xc4\x6b\x2d\x4c\x1a\xda\x10\x84\xf6\x16\x0c\xa2\xbc\xa2\xeb\xc6\x86\xad\x3b\xb6\xba\xa9\xd0\x03\x01\x23\xbd\x68\xcb\xf9\x20\x76\x51\x7a\x48\xef\x94\x05\x8d\x5f\x34\xa3\xcd\xdf\xb5\x2b\xd8\xf9\xf7\x77\x3b\x9b\xde\xb7\x6f\xe0\x4f\x7a\xa4\xa7\xc5\xd8\x22\xf8\xa9\xd7\x76\x00\x5e\xd1\x4c\x30\xc1\x76\xd0\x06\xa4\xbe\x21\x0c\x89\x24\x0c\x9d\x6d\x06\x43\x98\xbf\x4e\x2e\xf5\xb1\xed\x13\xe5\x9c\xcb\xc6\x4e\xd7\xd6\x6b\xdb\xac\x03\xc6\x06\xec\x81\x08\xb6\x52\x05\xbf\xb2\x58\x9c\x2e\x0b\x27\xfc\x10\xb4\xb5\xb5\x87\xc6\x04\xa5\xe7\xbf\x43\x2a\x0f\x4d\xbd\xde\x95\x6f\xe4\xb2\x10\x76\x1f\xbf\xdd\xf9\x62\x88\xb2\x89\x06\x5b\x8e\x4a\x87\x7d\x6b\xe3\xe3\x49\x61\xab\x3a\x3e\xd5\x6f\xfa\x79\x26\x63\x7e\xd8\x04\x2a\x2c\x94\xc2\xd2\x99\xaf\x1f\xb3\xcb\x4b\x4a\x91\xb3\xd9\xda\x5c\x70\xeb\x7d\x16\x2f\xa4\xf3\xdf\x3f\xfe\x1f\x00\x00\xff\xff\x94\x26\x66\x09\xf6\x1d\x00\x00")

func masterStartupShBytes() ([]byte, error) {
	return bindataRead(
		_masterStartupSh,
		"master-startup.sh",
	)
}

func masterStartupSh() (*asset, error) {
	bytes, err := masterStartupShBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "master-startup.sh", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _nodeStartupSh = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xbc\x58\x51\x6f\xdb\x38\x12\x7e\xe7\xaf\x98\xb5\x17\x9b\x5b\x6c\x29\xa5\x07\xf4\x0e\x9b\x6e\x0b\xf4\xda\x2e\xd0\x5b\x6c\x13\x24\x7b\x77\x0f\x45\x1f\x68\x71\x24\x71\x4d\x91\x2a\x87\xb4\xe3\xba\xfe\xef\x87\xa1\xe4\xc4\x8e\x13\xb7\x45\x0f\xf7\x16\x4b\xc3\x99\x6f\x66\xbe\xf9\x38\xca\xf4\xbb\x72\x66\x5c\x39\x53\xd4\x82\xc4\x6b\x21\xd6\x6b\x30\x35\x14\x2f\xbd\xab\x4d\x53\x5c\x61\x95\x82\x89\xab\x0b\x15\xab\xf6\x42\x55\x73\xd5\x20\xc1\x66\x23\xac\x6f\x1a\x0c\x20\x23\x38\xaf\x51\x52\x54\x21\xa6\xbe\xa0\x16\x26\xc6\x51\x54\xd6\x1a\xd7\x40\x40\x0d\xad\x8a\x50\x69\x07\x55\xf6\x98\x82\x8a\xc6\x3b\xf0\x0e\xbe\xff\x4b\xeb\x29\x3a\xd5\xe1\x8f\x13\x51\xa9\x08\xcf\xcb\x85\x0a\xa5\x35\xb3\x72\x95\xba\xb2\xb2\x06\x5d\x94\x15\x86\x58\xf4\xd8\xc1\x2f\xbf\x9c\xbc\x3e\xff\xf5\x84\x01\xbe\xc4\x10\x5f\xd0\x3f\x56\x11\xe9\x06\x29\x3f\x33\xb5\xa9\x54\x44\x2a\x46\xa4\x97\xd8\x7b\x32\xd1\x87\x55\x7e\x0d\x9f\xe0\x2a\x06\xc6\xb5\xd9\x88\xd7\xe7\xbf\x3e\x1c\x74\x8e\xab\xbb\x31\x2f\x82\x59\xa8\x88\xbf\xe1\xea\x2b\x23\xff\x86\xab\x83\xc0\x5f\x58\xbe\x17\x97\xe7\x40\x63\x07\x20\xf5\x9a\x23\xc0\xbb\xf5\x7a\xf4\x46\xff\xf4\xc6\x7d\xa6\x55\x93\x47\x30\x81\xcd\xe6\xfd\x41\xc1\x6b\x1f\x40\xc5\x88\x5d\x1f\xc1\x38\x58\x3f\x2e\x8a\x27\x9b\xa7\xa0\xbd\x00\x58\xa5\x0e\x46\x18\x20\x57\x20\x3f\xc0\xd7\xc5\xcc\x21\xe1\x87\x1f\x60\x16\x50\xcd\x05\xc0\x91\x74\xdf\x6d\x41\x7c\xbf\x1e\xff\xda\xbc\xbf\x3f\xf1\x11\xd1\xc0\x9f\x5a\x19\x8b\x7a\x22\x80\xd9\xfa\xee\xdd\xce\x69\x90\x36\xc2\x13\x78\xff\xfe\x29\xc4\x16\x1d\x90\x45\xec\xe1\xf1\x53\x40\x4b\x08\x78\x6d\x22\xff\xa8\x8d\xd0\xde\xe1\xd1\x4e\x04\xec\xfc\xe2\xeb\x68\xcc\x95\xab\x2c\x2a\x07\xca\x5a\x11\x3a\x90\xa1\x86\xa3\xb4\x3e\x42\x3f\xb1\x5e\xa3\xd3\x9b\x8d\x10\x53\xa8\x02\xaa\x88\x1c\x3d\x2a\xe3\x30\x40\x9f\xac\xe5\x1a\x05\x8c\xa2\x9b\x6b\x13\x40\xf6\xb7\xce\x7c\x30\x8d\x71\x65\xa1\x7d\x35\xc7\x70\x87\xe8\xfb\x2f\xcb\x21\xa3\xe2\x4f\xf2\x6e\x97\xf0\xc5\x2b\x0c\x66\x81\xba\x78\xe9\xbb\x99\x71\xa8\xdf\x74\xaa\xc1\x8b\x64\xed\x55\x8e\xba\x25\xc1\x01\xb9\xad\x03\x49\x0f\x41\x81\x32\x78\x1f\x4b\x4e\xe9\x8f\xf3\x57\xe7\x67\xa0\xd1\x62\x44\x6e\x15\xd4\xde\x5a\xbf\x64\x4f\x4d\xf0\xa9\x1f\x72\xe6\x2a\xab\x3a\x62\x00\x13\xc1\x10\xcc\xd4\x1c\x35\x18\x17\x3d\xf8\x14\xe0\xdf\xbf\x83\x61\x5c\x24\xf2\x19\xa5\x35\xc8\x1a\xc6\xb4\x85\xa9\xe1\x3b\x68\x02\xee\x54\x66\x0b\x03\x63\x55\xd6\x14\xd5\x6c\xa0\x89\x00\xa0\x15\x45\xec\xaa\x68\x81\xa2\xef\x47\x1f\x32\x77\x33\xf5\x45\x34\x1d\x86\xcf\x5a\x11\x86\x85\xa9\xf0\x21\xbb\x9d\xf7\xdd\xbc\xa6\xe2\xba\x26\x86\x5b\x6a\x5c\x94\xda\xd0\xbc\x54\x1f\x53\xc0\x32\x20\xf9\x14\x2a\x94\xbd\x0a\xf1\xb1\x00\xc0\xaa\xf5\x70\x72\xdc\x0c\x0e\x72\x04\x76\x0f\x4d\xe8\x3f\x24\x1f\x15\xc0\x29\x9c\x9e\xc0\xf3\xe7\xb7\xa9\x33\x0c\x9f\x5c\xbc\x7b\x52\x00\x04\xa4\xe8\x03\x56\xde\x81\xbc\x3c\x78\xbf\x5e\x4b\x9e\x3b\xfc\x00\xc5\xa5\xb7\xc8\xa2\x55\x07\xc5\x53\x2f\x00\x06\xb2\x71\x90\x91\x60\x5a\x61\xe7\xdd\x1d\x82\x09\x80\x89\xf5\x8d\xd4\x4c\xb2\x30\x39\x83\xc9\x9f\x3e\x05\xa7\xac\x9e\x3c\xe2\x77\xda\x90\x9a\x59\x94\x16\x1b\x55\xad\x64\xc0\xc6\x50\x0c\xab\xc9\x19\xc4\x90\x50\x0c\x54\x63\x1c\xe8\xf4\x10\x77\xb7\xe2\x2a\xc4\xc3\x92\xdf\x6f\x70\xa7\xc3\xb5\x11\x62\xac\x5f\x1e\x31\x9e\x84\x51\xed\xf2\x00\x50\xf1\xd6\x6b\xcc\xfa\xf6\x3c\x37\xc4\xb1\xd5\x0f\x47\x85\x84\x1d\x31\xad\x0f\x7d\x5d\x0d\x56\xb0\xd9\x1c\x57\xe4\xe3\x88\x6e\xbd\x7c\x8b\xe8\x7e\x71\x8c\xaf\xd5\xde\xbf\x9d\x3e\x24\xbe\x53\x31\x85\xb7\xe7\x7f\xbc\x3e\x83\x37\x0e\xea\x14\x53\xc0\x47\xd0\xf9\x05\xeb\x81\xe2\x2a\xd4\x3e\x74\xa3\xd2\xa6\x48\x46\x23\xf8\x1a\xd0\x2d\x4c\xf0\xae\x43\x17\x61\xa1\x82\x61\x9e\x90\x98\x0a\xc2\x08\x3f\x5d\x0b\xbc\xee\x7d\x88\x70\xf5\xe2\xea\x5f\x97\x6f\x9e\x9d\xec\xa4\xf2\x1f\x1f\xe6\x18\xc6\x4c\x86\xf7\xb0\xd9\x9c\xe4\x83\xf2\x9a\x35\x29\x24\x97\xa5\x68\x2c\x16\x48\x69\x9c\x89\x10\x3d\xcc\xbc\x8f\x14\x83\xea\xe1\xd5\xdb\x2b\x20\x8c\xa9\xdf\xf2\x84\x0f\x49\xd9\x07\xb3\x30\x16\x1b\xd4\x20\x25\x2b\xbf\x74\x18\x97\x3e\xcc\x81\xef\x07\x90\x0b\x28\xcf\x4a\xfe\xf3\xec\x23\x48\x1c\xe1\x1d\x2f\xf4\x3e\x8a\xad\xbf\x8c\x13\x07\x0a\x6f\x43\x74\xca\x29\x6e\x76\xf4\xd0\x9b\x6a\x0e\xa9\x07\x87\xcb\x21\x32\x61\x8c\x7c\x73\x33\xa9\x38\x37\xdd\x0e\xd7\x8c\xc8\x9b\xdf\xed\x50\x6c\x7d\xbe\x1d\x7c\xfe\x3e\xb8\x14\x53\x76\x90\x0f\xe6\x99\x66\xd1\xb1\x8b\x82\x2f\x0d\x91\x1f\xec\x9b\xb3\x3e\xf5\xbc\x0d\x60\x28\x74\xf9\xf3\xcf\x72\x90\x7f\xa9\x1d\x15\xd4\xee\x20\xd7\x8e\x3a\x45\x1f\x18\x71\x33\xfa\x67\xc4\xbb\x60\x07\x0d\x19\xec\xc6\x80\xf7\x23\xde\xda\x6c\x47\x5d\x4c\x61\x99\xe9\xc7\x6f\x79\xf0\x78\x00\x60\xa9\x54\xc3\x94\x51\x4e\xdf\xd4\x2d\x45\x63\x4d\x34\x48\xd0\xf8\xbc\x62\x44\x0f\x41\x55\xf9\xa2\xd5\x86\x89\x57\x88\x29\xd3\x7c\x7b\x38\x24\x47\x30\xc3\xda\x07\xe4\xb0\x7c\x23\xcd\x9d\x5f\x3a\xce\x23\xe7\x90\x23\x61\x56\xa5\xd4\xc3\xd2\xc4\x16\x78\x34\x56\x40\xf9\x92\x14\xcb\xd6\x58\xcc\x53\x73\xb3\x37\x80\xd4\x3f\xc2\xb3\x67\x30\x99\xe4\xc9\xd1\xfe\x76\x67\x19\xc6\xe4\xff\xc4\x33\x91\x1c\xb7\x7a\x38\x21\xc4\xb0\x77\xc9\x4a\xc9\x18\x12\x45\x31\x30\x21\xd7\x33\xf5\xd0\xa0\xc3\x85\xca\x22\xc3\x4f\x28\xaa\x6a\x0e\x8a\x80\x3c\x5f\xd8\x94\x8b\xb1\xbf\x2b\x19\x02\xab\x8c\xe6\x49\x86\xd9\x4a\x4c\xf7\x06\xed\x66\xb1\x79\x34\x9c\xb4\x9e\x98\xce\xad\xc9\x25\x1e\x8b\xfb\x80\x71\xe7\x03\x8a\x29\x43\x21\xa8\x83\xef\xf6\x6c\xfb\xe0\x2b\x24\xe2\x9e\x2c\x0d\xaf\x4c\xad\xe9\x07\xd2\x31\x7e\x31\xc0\x20\x04\x6a\x7d\xb2\x3a\xd7\xd8\xbb\x0a\x41\x81\x56\x2b\xf0\xce\xae\x38\x9b\x3e\x83\x41\xd6\x48\x12\x65\xa2\x50\x5a\x5f\x29\x9b\xf9\xa8\x3e\x12\x56\x7a\x4c\x96\x57\x9f\x99\x22\xb4\xc6\x71\x5f\xe1\xe2\xf1\xab\xcf\xda\x93\xaf\xe3\x52\x85\x2f\xb6\xaf\xac\xea\xd4\x62\x6b\x2d\xa6\x80\x8e\x25\x30\x13\x7b\x18\x88\xfd\xae\x8c\x63\x41\xe2\x76\x72\x92\xeb\x14\xcd\xa1\xd3\xa4\xb7\x53\x03\x43\x9c\xfd\x9f\x9d\x77\xb7\x4f\x6a\x9b\xd0\xc5\x9b\xdf\x3b\xee\x46\x00\xff\x2b\x77\x43\x12\xdf\xe6\x4d\x1c\x5b\x52\x62\x72\xa8\xa5\xd2\x1d\xd3\xa3\xe6\x81\xf4\x3d\x3a\x6a\x4d\x1d\x25\x93\x2b\x78\x2b\x7b\xab\x1c\x0e\x1b\x06\xdf\x5d\x9f\x39\xc5\x73\xbf\xbb\x8e\xb0\x04\x21\x28\x4b\x1e\x1c\xa2\xbe\xb5\x2c\x72\x63\x8b\x85\xb7\xa9\x43\x02\x5e\xdc\x87\xdd\x5e\x6f\x65\x85\xb7\xb6\x61\x63\xab\x58\x4c\x58\x71\xb6\x1b\x7e\x07\xa7\x7f\x7f\x72\x7a\xdf\xa6\xff\x80\x7f\xc6\x31\x6c\x52\xf9\x3b\x86\x56\x64\x7d\x03\x64\x98\xe1\x4b\x1c\xaf\x0d\xc0\x05\x86\x55\x6c\xd9\x24\xb6\xc1\xa7\xa6\x85\xed\x32\xb6\xd3\x95\x71\x23\xdb\x7a\xb9\xb7\x6f\xbe\x3f\x78\x2d\xa6\xe0\x7c\xc4\x33\x50\xd1\x77\xa6\x92\xfb\x35\x83\x2a\x28\x6a\xc1\x7a\xdf\x13\x24\x17\x8d\x85\x4e\x51\x5e\xf7\x09\x52\x7f\xc8\xb1\x7b\xbd\xdc\x04\xfb\xf6\xff\x5e\x50\xd5\xa2\x4e\xb9\x5c\x3b\x5f\x76\x10\x90\x17\x00\x16\x8e\xca\x77\x7d\xfe\x64\xb9\xef\x23\x75\x22\xa8\x4d\x51\xf3\x85\x20\xe5\x78\xe6\xa7\xbf\xf2\x87\x9c\x25\xdc\x6c\x0e\x38\x7e\x34\x1b\xf8\xf4\x69\x58\x76\xb7\xdf\x81\xff\x0d\x00\x00\xff\xff\xd9\x9d\x8e\x5d\xad\x11\x00\x00")

func nodeStartupShBytes() ([]byte, error) {
	return bindataRead(
		_nodeStartupSh,
		"node-startup.sh",
	)
}

func nodeStartupSh() (*asset, error) {
	bytes, err := nodeStartupShBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "node-startup.sh", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

// Asset loads and returns the asset for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func Asset(name string) ([]byte, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("Asset %s can't read by error: %v", name, err)
		}
		return a.bytes, nil
	}
	return nil, fmt.Errorf("Asset %s not found", name)
}

// MustAsset is like Asset but panics when Asset would return an error.
// It simplifies safe initialization of global variables.
func MustAsset(name string) []byte {
	a, err := Asset(name)
	if err != nil {
		panic("asset: Asset(" + name + "): " + err.Error())
	}

	return a
}

// AssetInfo loads and returns the asset info for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func AssetInfo(name string) (os.FileInfo, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("AssetInfo %s can't read by error: %v", name, err)
		}
		return a.info, nil
	}
	return nil, fmt.Errorf("AssetInfo %s not found", name)
}

// AssetNames returns the names of the assets.
func AssetNames() []string {
	names := make([]string, 0, len(_bindata))
	for name := range _bindata {
		names = append(names, name)
	}
	return names
}

// _bindata is a table, holding each asset generator, mapped to its name.
var _bindata = map[string]func() (*asset, error){
	"master-startup.sh": masterStartupSh,
	"node-startup.sh":   nodeStartupSh,
}

// AssetDir returns the file names below a certain
// directory embedded in the file by go-bindata.
// For example if you run go-bindata on data/... and data contains the
// following hierarchy:
//     data/
//       foo.txt
//       img/
//         a.png
//         b.png
// then AssetDir("data") would return []string{"foo.txt", "img"}
// AssetDir("data/img") would return []string{"a.png", "b.png"}
// AssetDir("foo.txt") and AssetDir("notexist") would return an error
// AssetDir("") will return []string{"data"}.
func AssetDir(name string) ([]string, error) {
	node := _bintree
	if len(name) != 0 {
		cannonicalName := strings.Replace(name, "\\", "/", -1)
		pathList := strings.Split(cannonicalName, "/")
		for _, p := range pathList {
			node = node.Children[p]
			if node == nil {
				return nil, fmt.Errorf("Asset %s not found", name)
			}
		}
	}
	if node.Func != nil {
		return nil, fmt.Errorf("Asset %s not found", name)
	}
	rv := make([]string, 0, len(node.Children))
	for childName := range node.Children {
		rv = append(rv, childName)
	}
	return rv, nil
}

type bintree struct {
	Func     func() (*asset, error)
	Children map[string]*bintree
}

var _bintree = &bintree{nil, map[string]*bintree{
	"master-startup.sh": {masterStartupSh, map[string]*bintree{}},
	"node-startup.sh":   {nodeStartupSh, map[string]*bintree{}},
}}

// RestoreAsset restores an asset under the given directory
func RestoreAsset(dir, name string) error {
	data, err := Asset(name)
	if err != nil {
		return err
	}
	info, err := AssetInfo(name)
	if err != nil {
		return err
	}
	err = os.MkdirAll(_filePath(dir, filepath.Dir(name)), os.FileMode(0755))
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(_filePath(dir, name), data, info.Mode())
	if err != nil {
		return err
	}
	err = os.Chtimes(_filePath(dir, name), info.ModTime(), info.ModTime())
	if err != nil {
		return err
	}
	return nil
}

// RestoreAssets restores an asset under the given directory recursively
func RestoreAssets(dir, name string) error {
	children, err := AssetDir(name)
	// File
	if err != nil {
		return RestoreAsset(dir, name)
	}
	// Dir
	for _, child := range children {
		err = RestoreAssets(dir, filepath.Join(name, child))
		if err != nil {
			return err
		}
	}
	return nil
}

func _filePath(dir, name string) string {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	return filepath.Join(append([]string{dir}, strings.Split(cannonicalName, "/")...)...)
}