// Code generated by go-bindata.
// sources:
// public/index.html
// public/js/app.js
// public/js/jquery.blockUI.min.js
// DO NOT EDIT!

package main

import (
	"github.com/elazarl/go-bindata-assetfs"
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
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}

	var buf bytes.Buffer
	_, err = io.Copy(&buf, gz)
	clErr := gz.Close()

	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
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

func (fi bindataFileInfo) Name() string {
	return fi.name
}
func (fi bindataFileInfo) Size() int64 {
	return fi.size
}
func (fi bindataFileInfo) Mode() os.FileMode {
	return fi.mode
}
func (fi bindataFileInfo) ModTime() time.Time {
	return fi.modTime
}
func (fi bindataFileInfo) IsDir() bool {
	return false
}
func (fi bindataFileInfo) Sys() interface{} {
	return nil
}

var _publicIndexHtml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x9c\x54\x4d\x53\xdb\x30\x10\xbd\xf3\x2b\x84\xae\xd4\x56\x4c\x42\x3e\x3a\x76\x66\x02\x25\x94\xd0\x82\x13\x13\x12\x7a\x53\x24\x39\x92\x23\x4b\x46\x92\x4d\xfc\xef\x3b\x89\x49\xdb\x01\x3a\x74\x7a\xf2\xee\x4a\xef\x69\xdf\xee\x1b\x87\xc7\x5f\xee\x2e\xee\x1f\xe3\x4b\xc0\x5d\x2e\x87\x47\x61\xf3\x01\x20\xe4\x0c\xd3\x5d\x00\x40\x98\x33\x87\x01\xe1\xd8\x58\xe6\x22\x58\xba\xd4\xeb\xc3\x3f\x8f\x14\xce\x59\x04\x2b\xc1\x9e\x0b\x6d\x1c\x04\x44\x2b\xc7\x94\x8b\xe0\xb3\xa0\x8e\x47\x94\x55\x82\x30\x6f\x9f\x7c\x02\x42\x09\x27\xb0\xf4\x2c\xc1\x92\x45\xc1\x81\xc8\x09\x27\xd9\xf0\x9e\x33\x30\x2b\x73\x10\x4b\x5c\xaf\x8d\x2e\x15\x0d\x51\x73\xd2\xdc\x92\x42\x6d\x80\x61\x32\x82\x96\x6b\xe3\x48\xe9\x80\x20\x5a\x41\xc0\x0d\x4b\x23\x98\xe2\x6a\x97\xfa\x82\x68\x38\x3c\x6a\x20\xc7\x9e\x07\xce\xb5\x76\xd6\x19\x5c\x80\x8b\x24\x01\x9e\xf7\x96\xcd\xd5\x92\x59\xce\x98\x3b\x50\x71\xe7\x0a\xfb\x19\xa1\x1c\x6f\x09\x55\xfe\xea\xc0\xb0\x4b\x88\xce\xd1\xaf\x02\xea\xf8\x2d\xbf\x85\x88\xb5\xbf\x6b\x7e\x2e\x94\x4f\xac\x85\x40\x28\xc7\xd6\x46\xb8\x7a\xd7\x31\x6e\xf7\x3b\xde\x95\x3a\x6b\xf7\x3b\xdb\xa7\x69\x80\xf5\x62\x39\x3a\x69\x9d\xf5\x67\xcb\x78\x1b\xaf\xbb\x69\xdd\xb9\x5e\x54\xf7\xb7\xbc\x75\x79\xda\x6d\x2f\xf3\x31\x99\xc8\x64\xf4\x2c\xae\xd6\xe3\xd1\x02\xd1\x91\x48\xba\x93\x65\x0e\x01\x31\xda\x5a\x6d\xc4\x5a\xa8\x08\x62\xa5\x55\x9d\xeb\xd2\x36\x82\x43\x74\xd8\x5b\xb8\xd2\xb4\xde\x07\x54\x54\x40\xd0\x08\xbe\xac\x05\x0e\x43\x44\x45\xd5\x5c\xb7\xc4\x88\xc2\x01\x6b\x48\x04\x11\x22\x9a\x32\x3f\x7b\x2a\x99\xa9\xf7\x22\x9b\xd0\x0b\xfc\x20\xf0\x83\xbd\xa8\xcc\xee\xe0\x0d\x6a\xf8\x86\x20\xb3\x2f\x10\x7f\x25\x35\xd9\xcc\xaf\x3f\xc6\x1c\xe6\x4c\xa8\xca\xac\x4f\xa4\x2e\x69\x2a\xb1\x61\xfb\xf7\x71\x86\xb7\x48\x8a\x95\x45\x85\x2e\x0a\x66\xfc\xcc\xa2\xc0\x0f\x4e\xfd\x01\x2a\x73\x7a\x28\xbe\xbc\xf1\xce\xac\x47\xc5\xed\x6a\xcd\x07\xe7\x27\x8f\xc1\xf4\xc6\x55\xed\x99\xea\x2d\xda\xf9\x3a\xde\xf2\xf9\xe0\x06\x25\x64\x6a\x47\x71\x8f\xcf\xc5\x6a\xd9\x1e\x64\xbd\x14\x6f\xc6\xb1\xdd\x54\xcb\xd2\x56\x29\x6e\xad\x3a\xd3\xbf\xcf\xfa\x43\x3d\xff\xea\x9b\xec\xb5\x6d\xde\x57\x32\xf9\x31\xeb\x26\x05\xcb\x78\x67\xde\x3a\xa5\xfd\xec\xce\x75\xab\x6f\x97\x5f\x53\x86\x26\xd3\x2b\x31\x9b\x25\xd3\xe9\x36\x49\xc7\x8b\x42\x04\xdf\x9f\xca\x07\x3a\xaa\xb3\x39\x36\x67\x27\xbd\x6e\xfc\x70\x91\x3f\xca\xff\x50\xb2\x6b\x0d\x17\xc5\xab\xf5\xed\x3d\xd6\x58\x2b\x44\xcd\xcf\xe2\x67\x00\x00\x00\xff\xff\x92\xe2\x24\x31\x44\x04\x00\x00")

func publicIndexHtmlBytes() ([]byte, error) {
	return bindataRead(
		_publicIndexHtml,
		"public/index.html",
	)
}

func publicIndexHtml() (*asset, error) {
	bytes, err := publicIndexHtmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "public/index.html", size: 1092, mode: os.FileMode(420), modTime: time.Unix(1518438667, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _publicJsAppJs = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xbc\x58\xe1\x6e\x1b\x39\x0e\xfe\xef\xa7\xe0\xce\x16\xeb\x19\x24\x19\x27\xbb\x6d\x77\x37\xae\x0b\xf4\x7a\xd9\xbd\x02\xdd\xf4\x70\x4e\x51\x1c\xda\xc2\x90\x35\x1c\x8f\x2e\xb2\x34\x27\x71\xec\x1a\x6d\xde\xfd\x40\x69\xc6\x1e\x27\x69\xbc\x87\x43\x2f\x3f\x5a\x5b\x94\xc8\x8f\xe4\x47\x4a\xf4\x4a\x38\x58\xfb\xf1\x60\x50\x36\x46\x92\xb2\x06\x3c\x9a\xe2\x62\x85\x86\x52\xc0\x15\x41\x06\x9f\x07\x00\xa5\x75\x29\x6f\x55\x93\xd3\x31\x28\x78\x06\xc2\x2d\x9a\x25\x1a\xf2\xb9\x46\xb3\xa0\x6a\x0c\xea\xe8\x28\x1b\x00\xff\xad\x7d\xce\x4a\xd2\xed\x9e\xf7\xea\x63\x36\x1e\xec\x04\xc9\x6c\x36\xbb\xb8\xfc\xeb\x9b\xdf\xfe\xb8\x98\x4e\x5f\xfc\x7e\x31\x9b\xcd\x92\x6c\x70\x33\x18\x04\x13\xc5\x14\xbd\x67\x24\x13\x48\x66\x33\x5c\xd6\xb4\x99\xcd\x92\x71\x10\x7a\x12\x84\x3b\xf9\xe9\x18\x46\xa3\x53\x38\x79\x0e\xca\x28\x82\x2f\x70\xc6\x9f\x5d\x63\x8c\x32\x8b\x9e\x4f\xd2\xa1\x20\x7c\x37\x4d\xa3\x37\xac\x49\x5b\x09\x13\x58\x2b\x53\xd8\x75\xae\xad\x14\xbc\x71\xdc\x0a\x1b\xa7\x60\x02\xc3\xb5\x3f\x1f\x8e\x07\x03\x00\x55\x42\xaa\xad\xcc\x6b\x67\xc9\x4a\xab\x61\x32\x99\xc0\xb0\x22\xaa\xfd\xf9\x30\xea\x84\xdd\xa1\x70\x0a\xe0\x66\x10\xd7\x8e\x26\x30\x1c\x8d\x86\x70\xc4\x36\xf3\xca\x7a\x1a\xef\x24\x41\xab\xa0\xca\x88\x25\xc2\x11\x9f\x8e\x16\xd7\x1e\x26\x60\x70\x0d\xef\x70\x3e\xb5\xf2\x1a\x29\x6d\x9c\xca\x5a\x59\x6e\x8d\xad\x91\x23\xd0\xb9\x98\x76\x28\xa4\x35\xde\x6a\xcc\xb5\x5d\xa4\xc3\x97\xd6\x18\x94\x84\xc5\x30\xa6\xe6\x6e\xf8\x20\x00\xed\x94\x2e\xd1\x7b\xb1\xc0\xbe\x5e\x5c\x51\xa7\x9a\xa3\x80\x2b\xca\x0b\x41\x02\x26\x21\x3b\xbf\x5f\x5c\x4d\x2f\xa6\xd3\x57\x6f\x2e\x67\xb3\x04\x7e\xf8\xe1\x96\x85\x09\x9c\x76\x87\x77\xd9\xdf\x26\x38\x70\x82\xff\x1c\x52\xe3\x4c\xfc\x76\xb3\x35\xf5\x80\xaa\x3e\x47\x3a\x44\x9d\xb2\x5b\x3e\x9e\x75\xeb\x8f\xf2\xc6\xcc\xb5\x95\xd7\x6f\x5f\xa5\x0f\x58\xc6\x95\xd0\x5b\x2f\xb3\xf1\x5e\x78\xa4\xb6\xbe\x1f\x1c\x48\xfb\xa1\xd9\x33\xfc\xdd\x1e\x5e\x78\x94\x77\xa6\x3f\x43\x1b\xe4\x73\x18\x3e\xab\xce\x9e\xbf\xb6\x9e\xc0\x96\xe0\xd1\xad\xd0\x71\xfa\x38\x63\xca\x9a\xef\xe0\xef\x1a\x85\x47\x78\x27\x14\xe5\x79\xfe\x6c\x54\x9d\x3d\x1f\xc2\x4d\xd6\x47\x7b\x27\xa1\x71\x15\xe9\x4a\x2d\xd1\x36\x94\xee\xf8\xf1\x19\x3a\x30\x81\x5b\xbb\x92\x68\xf5\x1d\xc3\xd9\xe9\xe9\xe9\x2d\x8f\xd1\x39\xeb\xf6\xe8\xc0\x0b\x9d\x63\x5c\x2a\x0e\x85\xb7\x6d\x04\x47\xa3\x91\xd0\xe8\x28\x45\xee\x1f\xb9\xb4\x05\x66\x9d\x04\xa6\x88\xc0\x15\x73\x3e\x1a\x91\xb5\xda\xe7\x0a\xa9\xcc\xad\x5b\x8c\x2a\x5a\xea\x91\x2b\xe5\xd3\xc7\x4f\x9e\x7c\xef\xa3\xf7\x27\x3f\xe7\x8f\xf3\xb3\x1e\xed\x3a\x85\xcc\x84\x00\x74\x9b\x41\xb6\xcf\xbd\xe2\xd2\xba\xa5\xd0\xc0\x49\x6a\x1c\x1e\xc3\x12\x05\xb7\x01\xa0\x4a\x10\x50\x85\x50\x37\xae\xe6\x04\x96\xd6\xc1\xba\x52\xb2\x0a\xab\xbb\x88\xc3\x5a\x78\x40\x4f\x62\xae\x95\xaf\xb0\x80\x4a\x78\x98\x23\x1a\x28\x1b\x5d\x2a\xad\xb1\xc8\x93\xe8\x0f\x6a\x8f\xa0\xca\xbb\xb8\xce\xee\xe2\x7a\x61\x00\x4d\x51\x5b\x65\x08\x94\x87\x0f\xc9\xc2\x32\x2c\xb1\x16\x9b\x0f\xc9\x31\xf8\x46\x56\x20\x3c\x88\x8e\x02\x51\x5c\xd8\xb5\x01\xeb\x40\xc0\xdc\xd9\xb5\x47\x07\x95\x58\xb1\xc0\x88\x95\x5a\x08\xc2\x22\x68\x80\xd2\xd9\x25\x08\xa8\xc5\x02\x0f\x82\xfb\xf1\x20\x38\x42\xb7\x54\x46\x50\x8c\xdb\x5e\x70\x8a\x06\x81\x2c\x9b\xea\x9a\x60\xe0\xc2\x21\x9b\x3f\xfd\x2f\x36\xe7\x28\x45\xc3\x8a\x29\xe4\xc2\xa1\x44\xb5\x62\xcf\x81\x36\x35\x72\xd9\x84\x5e\xa4\x08\xa4\x30\xc6\x12\x08\x29\xb1\x26\x48\x31\x5f\xe4\xc7\x20\x7a\x86\x02\x0b\x1a\x53\xa0\xf3\x24\x4c\xe1\xc1\x1a\xbd\x01\xc2\x4f\x14\x75\xfc\xf1\xe2\x9f\xe1\xea\x03\xaa\x94\x67\xce\x29\xea\xec\x71\x6e\xe6\xca\x08\xb7\xe9\x6a\x37\x3b\x18\xea\xc7\x77\xdd\xfe\x07\x86\x04\x17\x39\x5c\x55\x08\xbe\x46\xa9\x4a\x25\xb7\x34\x5d\xaa\x45\x45\x30\x47\x28\xb0\x54\x06\x0b\x50\x26\x84\xa3\x6c\xa8\x71\x87\x73\xfb\xe4\xbe\x82\x08\xfd\xa1\xf1\x10\xf6\x31\xbb\x85\xa4\x46\x68\xbd\x81\xda\xa1\x67\x0d\x87\xd4\x3e\xbd\xab\xf6\xea\x6e\xd1\x84\xc6\x58\x80\x98\x9b\x50\x83\x7a\x73\x0c\x31\x03\x6b\x45\x95\x6d\x28\x44\x96\x9d\xb4\xae\x0d\x6a\xa8\x00\x78\x19\x1a\xaa\xb4\x86\x9c\xd5\x50\x3a\xb1\xc4\x43\x80\x7e\xfe\x16\x7c\x0a\x0c\x60\xac\xca\x80\xe8\x92\x1c\x29\xc3\xee\x31\xb3\xf8\x5e\x55\x9e\xd0\x50\xd8\x18\x34\x77\x24\xe4\xcf\xdd\xa1\x96\x7a\xc6\x9a\x93\xb7\x57\xbf\x9d\xfc\x02\xef\x1f\x6e\x7a\x3f\x3d\xfd\xf1\xd7\x8f\xb7\x00\x04\x56\xfe\x69\xaa\xfd\xf2\x6d\x2a\x6c\x2f\x0a\x1f\x92\x95\xb2\x5a\x10\x7a\x50\xe4\xa1\xb6\x5a\xc9\xcd\x87\x84\xa9\xac\x7c\x67\x58\x79\x58\xa8\x15\x1a\x40\x45\x15\x3a\x2e\x22\xfe\x1f\x59\x60\x2c\xd8\xb0\xe8\x1b\x52\x73\x8d\xed\x99\x63\x66\x44\x7f\x9f\x00\x83\x58\x70\xa3\xa9\x54\xd1\xab\x92\x02\x49\x28\xed\x41\xcc\x99\x4e\xa1\x9b\x07\x0c\x07\xa3\xf3\xeb\xff\x21\x3a\xac\xc4\x5a\x98\xab\x45\xb8\x5b\x14\x31\xfe\xda\x59\x89\xde\x1f\x02\x78\x76\x9a\xf1\xed\x78\x69\x09\xbb\x9b\x4a\xf9\xbd\xba\x55\x91\x80\x0d\x57\xd8\x7c\x13\x20\xc6\x9b\xe2\xb8\x8f\x4f\x0a\x03\xa5\x50\x3a\xc8\xb7\x8f\x46\xa8\xb8\xd7\x55\xe2\x1a\x41\x19\x4f\x28\x8a\xfc\xc1\x70\xa4\x52\x2b\x34\x94\xfd\x77\x71\xc1\x4f\x75\x78\x65\xf6\xb0\x71\x04\x0c\x2e\x2c\x29\x41\x08\xd6\x20\x27\x7a\x69\x1d\x02\x7e\x22\x34\xfc\x5e\x39\x86\x79\x9b\xca\xf6\x48\xa1\x0a\x33\xa4\xf6\x4d\xc6\x82\x65\xd7\x03\x1d\xfa\xda\x1a\xbf\xab\xb2\xb6\xe8\xee\xf1\x33\x87\x67\x73\x07\xa3\xe7\x30\x6d\xa9\x13\xdb\x11\xef\xde\x5a\xf6\x31\xd2\xc2\x61\x60\x1b\x27\xd4\xe1\x39\x24\x70\x04\x31\x3d\xfd\x47\xcd\x57\x13\x77\xdf\x55\xdf\xb9\xf2\xa7\xc3\x87\x46\xda\xc6\x10\x3a\x06\x61\xa0\x31\xdb\x58\x4a\x6b\x0a\x15\xb6\x07\xb0\xb5\x0b\x00\xf8\x5e\xa0\x78\xe9\xb7\x8f\x92\x4e\xbf\xc3\x7f\x37\xe8\x0f\x76\xf4\xb3\x7b\x2e\x8a\xaf\x77\xf4\xed\xa5\xcf\xd4\x6a\x5c\xf8\x52\xa3\x2b\xad\xe3\x47\xc7\xd5\xeb\x69\x8f\x60\x6d\xe7\xeb\x25\x54\xa2\xa3\x90\x02\x42\xe6\xe7\x30\x5c\x6f\x2b\x74\xaa\x54\x58\xec\x35\xb6\x3b\x90\xde\x9a\x6b\xc3\x2f\xa0\xb8\x92\x84\xb9\x07\xe0\x51\x9a\x7c\xcf\x3d\x72\xe1\xaf\x2a\x41\x7f\x13\x75\x8d\x06\x8b\x24\xcb\xb9\x9b\xa6\x0f\x4b\x33\x38\x82\x24\x72\x83\x13\xfd\x80\xd3\xa5\xed\x9e\xb5\x91\x13\xf1\x73\x7c\x19\xf3\x3f\x2d\x41\x79\x6c\xbe\x19\x0c\x1e\xa5\x50\x58\x19\x06\x5d\xc8\x98\x39\xc5\x26\xbd\x35\x95\xf1\x0b\x36\xd9\x72\x35\x61\x56\xc7\xc1\x33\x8b\xef\xe8\x7b\x1e\xe4\x37\x31\x81\x51\x1e\x1f\xd5\xac\xc1\x47\xb6\xb7\x2d\xc1\x37\x75\x6d\x1d\xc5\xbe\xb0\xb1\x8d\xeb\xde\x8a\x49\xfb\xec\xde\x8d\x37\x37\x03\x1e\x1c\x76\x13\xb1\xf2\x53\xe9\x54\x4d\xaf\xad\x28\xb0\x48\x1b\xa7\x77\xa3\xb1\x0f\x12\xc6\xd4\x39\x96\x2f\x90\x2e\x34\x86\x61\xfe\x2f\x9b\x2b\xb1\xb8\x14\x4b\x4c\x87\x71\xe3\x30\x18\xe3\xa0\xc5\x9f\x08\x60\xd2\x69\xd8\xfd\x36\x70\x72\x32\xde\x1b\x95\xa2\xfc\xbd\xfa\x98\x7b\x27\x99\x98\x01\x40\x57\xf9\xae\xc1\x6e\x80\x6e\x97\x4a\xa1\x3d\x86\x70\x6f\x3d\xd0\x56\x14\xd1\x87\x94\x4f\x1f\x03\x57\xfa\x5c\xc8\x6b\xb8\xed\x49\xdf\x91\x18\xe5\xd6\x97\x14\x92\xb8\x23\x81\x76\x3e\x0e\xdf\xf2\x70\xaf\x4f\x20\xe1\x7b\x78\xf4\x2f\xb1\x12\xed\xae\x71\xc8\x65\x8b\x3e\xa6\x7a\xca\x03\x58\x06\x9f\x79\xc4\x79\x75\xd1\x57\x62\x4d\xd8\x10\x26\x34\x59\x09\xb3\x3f\x56\xf7\x66\xd9\x12\x52\xb8\xa3\x32\xfc\xc2\x90\xe8\x90\x9d\x04\xbe\x7c\xf9\xda\x0e\x69\x97\xb5\x46\xc2\x04\x7a\xd3\xe6\x83\x10\x4c\xa3\xf5\x78\xbb\xb3\x0b\xda\x6e\x24\x8e\x83\xe5\x4d\x9f\x84\xec\xdc\x1b\xbe\xa2\xfd\xbe\x83\x0c\xef\x7e\xa7\x6e\xab\xbd\xd9\x8e\x95\xed\xe1\x90\x76\xce\x1b\x0b\x1e\xa6\x19\x24\x15\x8a\x22\x81\xec\xfd\xe9\xc7\x3c\x14\x75\xf1\xb2\x52\xba\xe8\xa2\x06\x59\x20\xc6\x7f\x02\x00\x00\xff\xff\x99\xb4\xea\x70\xca\x12\x00\x00")

func publicJsAppJsBytes() ([]byte, error) {
	return bindataRead(
		_publicJsAppJs,
		"public/js/app.js",
	)
}

func publicJsAppJs() (*asset, error) {
	bytes, err := publicJsAppJsBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "public/js/app.js", size: 4810, mode: os.FileMode(420), modTime: time.Unix(1518116806, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _publicJsJqueryBlockuiMinJs = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xdc\x5a\x5f\x73\x1b\x39\x72\x7f\xf7\xa7\x18\xe2\xea\x46\x80\x09\x42\x43\xc9\x7b\x97\x0c\x17\x62\xad\x75\x4e\x56\x75\xa7\x78\x13\xc9\xb5\x49\x54\x4a\x0a\x9c\x01\x49\x98\xc3\xc1\xdc\x00\x24\x45\x8b\xfc\xee\xa9\x06\xe6\x1f\x29\xca\xf6\x5d\xf6\xe9\x5e\x44\x0c\xba\x01\x34\x1a\x8d\xee\x5f\x37\x74\xfe\xf6\xcd\xdb\xe0\xf3\xbf\xaf\x64\xb9\x0d\xde\x67\x3a\x59\x7c\xba\x19\x05\xeb\x8b\x68\xf8\x6e\x38\xbc\xb8\x7c\xf3\x36\x98\x5b\x5b\xc4\xe7\xe7\x9f\xff\x0a\x2c\x6c\x29\x32\xb3\x2a\x58\xa2\x97\xe7\x13\xe0\x3e\x7f\xf3\x36\xb8\xd6\xc5\xb6\x54\xb3\xb9\x0d\x70\x42\x02\x18\x1a\xdc\xb2\xe0\x27\x60\x1c\x05\x7f\x5a\x89\x2c\xc8\x54\x22\x73\x23\xd3\x38\xb8\xbd\xb9\x3f\xff\xd7\x5f\xfe\xf2\xe6\xed\xf9\x1b\x3c\x5d\xe5\x89\x55\x3a\xc7\xe4\x19\xad\x8c\x0c\x8c\x2d\x55\x62\xd1\xa8\xee\x0f\x24\x96\xe4\xb9\xf9\xd2\x58\x53\x45\x9e\xd7\xa2\x0c\x0c\x9d\xd3\x05\xd7\x9c\x6f\x54\x9e\xea\x0d\x5d\x73\x15\x86\x6b\xad\xd2\x20\xea\x71\xae\xd8\x52\x1a\x23\x66\x72\xdc\xb4\x62\x4f\x1c\xa9\x29\x56\x5c\x32\xf9\x64\x65\x9e\xe2\xe7\x3d\x95\x6c\xe2\x37\xcd\x52\x39\x15\xab\xcc\x1a\xaa\x76\xbb\xe7\x3d\xa1\x3d\xc5\xd4\x2c\xd7\xa5\xbc\x99\x3a\xbd\xc8\x74\xb7\xeb\x49\xac\x09\x4b\x85\x15\x18\xd5\xe3\x94\xa9\xc8\x88\x90\x67\x98\x9f\xe9\xb5\x2c\x33\xb1\xbd\xbe\xbb\xfb\xc6\x52\x1d\x4e\xda\x1d\xe6\x05\x30\xdf\x1a\x9d\x18\x43\x15\xfc\xf5\xfc\x8a\xe9\xfc\x63\x35\x47\xa6\x92\x45\x18\x1e\xc8\xc2\x92\x55\x69\x74\xc9\x51\xa1\x55\x6e\x65\x89\x08\x9d\x7f\x6b\x05\x3b\x97\x4b\x99\x7a\xf1\x9a\xb6\x5f\x6d\xcd\xbd\x4a\x39\xe7\xeb\xae\x9e\xe9\x22\x0c\x27\x61\x68\x71\x75\x34\xcf\x53\x91\xca\x8f\x2b\x1b\x47\x30\x28\x0c\x11\x9c\x72\x3e\x43\x3d\x6e\xb7\x85\xd4\xd3\x60\x1d\x86\x78\xcd\x0a\x51\xca\xdc\xfe\x9b\x4e\xe5\x6e\xb7\x66\xde\xd8\x88\x3f\xec\x2d\xaf\x3b\xc6\xeb\x87\xe8\x31\x5e\xd3\x25\x7f\xde\x8f\x4e\x1c\xc5\x5c\x19\xab\xcb\x2d\xa2\x4b\x42\x97\x4c\x66\x7c\x4b\x97\xd5\xcc\x7c\xdb\x59\x82\x2e\x59\xaa\x4c\x91\x89\x2d\xdf\x32\x63\xb7\x99\xac\xbf\x81\x5f\x1b\x05\xf6\xd6\x90\xea\x8e\x66\xae\x30\xac\x5b\xac\x94\x4b\xbd\x96\xd7\x73\x95\xa5\x78\x4b\xf6\x27\x64\xd2\xf9\xa7\xdc\xb5\x91\x3b\xa0\xea\x83\x8c\x60\x67\x33\x7a\x43\x37\xf4\x13\x7d\xe2\x8a\x4d\x84\x91\xff\x3d\x9a\xf1\x72\xb7\x53\x6c\xaa\xcb\x44\xde\x4c\x4b\xb1\x94\x63\x89\xcf\x7e\x54\xae\x19\x24\x99\x30\x86\xd7\x53\xa3\xc0\xc9\xc7\xd1\x97\x81\xca\x53\xf9\x14\x9f\xf5\x9f\xfa\xfd\xa0\x7f\x36\xaa\x36\x13\xe7\x3a\x97\xa3\x89\x2e\x53\x59\xfa\xf6\x52\x94\x33\x95\xc7\xd1\xa8\x10\x69\xaa\xf2\x19\xb4\xaa\xdd\xc5\x62\x62\x74\xb6\xb2\x72\xb4\x51\xa9\x9d\xc7\xc3\x28\xfa\xfd\x68\x2e\xe1\x5a\xfb\xb6\xd5\x45\x1c\x8d\x32\x39\xb5\x71\x84\x02\x53\x26\x1c\x9d\xf5\x15\xf3\xa2\xdd\x95\x49\xff\x0c\x5d\xfd\x78\xee\x3f\xaf\xce\x48\x0c\x72\xa7\x6a\xfd\x9a\xd0\x5d\x19\x61\x60\xaa\xd6\x57\x67\x84\xde\xf0\xca\xd2\xc6\xa7\xc7\x07\xee\xb7\x32\xf3\x60\xa5\x06\x1b\x95\xce\xa4\x1d\x54\x76\xfe\x7d\x3a\x69\xd7\x7b\x45\xca\x83\x55\x7e\x1b\x3d\x7f\x8f\x56\x5b\x35\x54\x4a\x08\xc3\xc5\x18\x7f\xe2\x27\x45\x04\xe5\xbb\xf6\xad\x99\x5d\x03\xa9\x7f\xe6\xc5\xfe\x45\xcc\x24\x68\x26\x55\x22\xd3\xb3\x56\x47\xd0\x4a\x74\x99\xcb\x72\x20\xb2\xec\xc4\xae\xf0\x53\x7f\x18\x91\xe3\x7d\x35\x06\x32\x55\x4f\x32\x45\x57\x67\x20\x9c\xb2\x99\x0c\x43\xfc\xa9\x7f\x28\x5b\x7b\x1e\x73\x29\x52\x59\xb6\x62\x0c\xdc\x90\x89\x28\x0f\xa5\xf0\x12\xdf\x03\x0d\x5d\x9d\xf5\x71\x35\xf5\x6e\x87\xc2\x7c\x62\x8a\x11\x22\x7d\xe4\x95\x82\x08\x7d\x7d\xb5\x44\xe7\x56\xe6\xb6\xb3\x5c\xd5\xd3\xa8\x14\x06\x37\x33\xc5\xb5\x8d\xfd\xcd\xba\xfd\x90\xc9\xe5\xc1\x42\xbf\x9d\x7a\xeb\xfb\xf7\x8f\xa1\xe1\x4f\x7c\x31\xfe\xdb\xcd\xf6\xef\x36\xca\x4a\x86\xf8\xef\x3a\xcd\xff\xcf\x59\xd5\x9b\xdf\x70\x89\x3f\xb9\x00\x87\x1b\xe3\xda\x40\x6c\xc6\x73\x42\x37\x4c\xa4\xa9\x5b\x19\xbf\x54\x29\x22\x24\xf6\x9c\x86\x34\x17\x7f\xb7\xbb\x71\x5d\xdd\x00\x4e\xa8\xef\x43\xb5\x18\x88\x2e\xc6\xc8\x2b\x20\x46\x8d\x48\x84\xe2\xe3\x18\x42\xc2\x70\xe6\x87\xea\x42\x24\xca\x6e\x11\x8d\x7c\x04\xba\xe6\x0f\x2e\x08\x3d\xd2\x3b\xbe\x18\x4b\x8c\x26\x3a\xdd\x22\x70\x8b\x9a\x8c\x24\x93\x22\x99\xe3\x6b\xda\x01\x6a\x76\xae\x0c\x13\x45\x21\xf3\xf4\x5e\xe3\x3b\xb2\xef\xf8\x2a\xc5\xd2\x52\xcc\x66\x62\x02\x96\x2b\xd9\x34\xef\x7e\x6f\xda\x0f\xfc\x3c\x17\x79\x9a\xc9\x18\xb1\x97\xd6\x8b\x68\x22\xf2\x44\x66\x31\xca\x14\xda\x7b\x29\x3f\xf2\x69\x18\xe2\x9e\x64\x66\x55\x14\xba\xb4\x6c\xa2\x9f\x6e\x75\x2a\xb3\xdd\x4e\x62\xa4\x27\x9f\x65\x62\xa9\x5c\x4e\x64\x0a\x2a\xc9\x57\x59\x16\x6b\xc2\x32\x99\xcf\xec\xfc\x2a\x22\x80\xfa\x56\xbb\xdd\x47\x07\xcf\x16\x20\xa7\xc8\x32\xbd\x79\xaf\xd3\xed\x9d\x2d\xa5\x4d\xe6\x20\xee\xf1\xdc\x61\x28\x31\x9a\xdb\x65\x46\xbd\x4e\xbc\x02\xbd\xef\x46\x14\x81\xf7\x06\x5d\xaf\x00\x15\xbe\x18\x4c\xc2\xb0\xb7\x20\x20\xfb\x07\x9e\x62\x4d\x91\x8f\x0e\xf7\xba\xf8\x15\xe2\x00\x22\xf4\xbe\xdb\xff\x17\x39\xb5\x35\xe1\x96\x7f\x18\x23\x1c\x05\x83\x00\xf5\x3f\xf4\x11\x41\x71\x44\xdf\xf3\xfb\xa6\xef\xde\xf7\x9d\x38\x1d\x49\xb5\x47\x4c\x96\xeb\x87\xe8\xd1\xc3\x17\xd8\xbc\x6d\x41\x4d\x6b\x27\xf4\xe2\x4a\x92\xc5\xd8\x32\x23\xed\x87\xa7\xa2\x94\xc6\xc0\x24\xed\x0e\x6f\x85\x9d\xb3\xa5\x78\xc2\xa9\x4e\x56\x70\x4f\x18\x28\x82\x99\xa4\xd4\x59\xf6\xb3\xe3\xa2\xc1\x21\x4d\x4f\xa7\x46\x5a\x4f\x23\xc1\x20\xc0\x3e\xb1\x78\xa1\x9e\x71\x14\xa3\xbe\x62\x7f\x5d\xa9\x72\x61\x96\x3a\x95\x1f\xfd\x40\x91\x2c\xfa\x67\x24\xe8\x07\xa8\x78\x42\x67\x24\x7e\x55\xb8\x33\x67\x87\x2d\xa2\x3b\x58\xb9\x19\x4f\x4f\x6c\xcf\xc5\x61\x44\xcf\x5e\x91\x2c\x08\xc3\x76\x4f\x75\xa3\xf2\x13\x2c\xc9\x94\xcc\xfd\x41\x05\xbb\xdd\xd1\xde\xbb\xc4\xd7\x37\x50\x2f\x7f\x5a\xfe\xc3\xd1\xf4\x7d\x18\xbe\x98\x00\xa0\x02\xa2\xef\x09\xbd\x3d\x41\xb4\xba\x40\xf4\x96\x8c\x64\x66\x64\xe0\x72\x91\x44\x02\xd6\xff\x2f\xb2\x78\x8d\xfb\x0c\x7f\x63\xbb\x95\x4e\x5f\xd9\x6f\x7d\xd6\xe7\xc1\x05\x9c\xb7\xdb\xd6\xc1\x59\x9c\x07\x17\x70\xa0\x78\x92\x89\x79\xc0\x5f\xd7\xad\x37\xab\x7b\x5d\x04\xe3\xef\x61\x8a\x83\x53\x56\x79\xaf\x8b\xd6\x7a\xa8\x65\x1e\x87\xdd\xeb\x82\x47\x8d\x4e\x7a\x8d\x52\xc2\x70\xe1\xef\x4b\xce\x5d\x1a\x05\x9e\x21\x31\x86\x59\x5d\x8c\x0b\x51\x1a\x79\x93\x5b\xdc\x74\xd1\x61\x44\xe2\x88\x1a\x8e\xf0\xeb\x2a\xfb\x4d\x77\xd1\xcf\x3b\x97\x61\x74\xfa\xf8\x0c\xd9\xef\xc9\x5e\x4d\x71\x37\xf2\x6c\xd8\x54\xe5\x29\x76\xee\xf5\x38\xd8\x54\xce\x1b\xaf\x21\xea\x34\x6d\x8a\xeb\x1c\x0b\xd2\xaf\x5c\xa7\xf2\x7e\x5b\x40\xdc\x90\x78\x4d\x98\x99\xeb\x0d\x26\xa7\x03\x8b\x72\xd4\x0a\x2a\x43\x9c\xf1\xcc\x54\x31\x48\xfc\x6e\x72\xaf\xe1\xcf\x1c\x92\x1f\x97\x28\x8f\x9b\x56\x9c\xd0\x9f\xf9\xd1\xf8\xde\x7a\xfc\x39\x4e\xe8\x17\xee\x7e\x47\x47\xd4\x1b\xf6\xbf\x7e\x56\x5c\x4f\x4f\x7f\x76\x41\x77\xf3\x92\xf0\x85\xec\xfd\x91\x1f\x4f\x51\x09\xe8\x46\x35\xc2\x56\x22\xc1\x7e\xaa\x26\x9b\x80\x0e\x37\x04\xbb\xf8\x91\xe3\x21\xd5\x54\x81\x4b\xc1\x13\xbe\x79\x88\x1e\x69\xc1\x25\xac\xa6\x93\x95\x81\xa8\x56\x9d\xae\xa1\x13\xb7\x79\xe8\xbe\xc9\x8b\x95\x0d\x43\x23\xed\xbd\x5a\x4a\xbd\xb2\x38\xa3\x17\x11\x21\xb1\xc0\x6e\x86\xda\x12\xff\x93\xb6\x17\xd5\x61\x3f\xc7\xec\x35\xf7\x2b\xef\x0c\xef\x04\xe2\xc5\x58\xb2\x55\x5e\x41\x1c\xac\x7c\xc8\xae\x7b\xb0\x22\xfb\xce\x44\xa7\x72\xe5\x8a\x86\xe8\xaf\x64\xbf\xdf\x37\x45\x16\x8b\x35\xb5\x75\x91\x25\xeb\x94\x58\x04\x87\x49\x68\xca\xc5\x2b\x49\x37\xa1\xc9\x0b\x5a\xbd\x08\x19\x25\x61\x88\x93\x4c\x8a\xb2\xde\x4a\x42\xa8\xa8\xd2\xe7\x3f\x9d\x1e\x43\xa8\xfd\x56\xa9\xc6\xfa\x5a\x44\x8e\x23\x0a\x62\x53\x08\xfd\x9c\x73\xdb\x26\xda\x61\x88\x3b\x5f\x2f\x04\x6c\xb3\xf3\xd7\xc4\xe9\x70\x78\x30\x52\x8e\x4a\x9e\xb5\x50\x89\x25\x90\xfc\x97\x32\xc7\x84\x4d\x55\x66\x65\x89\x51\x2d\x28\xdc\xb6\x34\xf5\x9c\xc1\x55\xd0\x76\xc7\xa2\xba\xa1\x57\x6d\x1f\xb5\x55\x85\xe6\x3f\xa4\x91\x36\x0c\x71\x59\x23\x98\x21\x7c\x3c\x0c\xab\x80\x5e\xd7\x71\x0e\xd8\x09\x6d\xb8\x2f\x1c\xf7\xc5\x57\xb9\x09\xcd\xc2\x10\x4f\x78\xc1\x41\x63\xb0\x74\x55\xa5\x19\x63\xc3\xeb\x99\x68\xc9\x8c\xd5\x05\xec\xcb\x13\x71\xc3\xd6\x45\x84\x11\xe7\x7c\x30\x00\xe7\x89\x4b\x9a\x52\x4b\x35\xd9\x13\x12\x77\xbe\x1a\xe3\x52\x60\x5c\x54\xd1\xca\x29\x18\x2e\x71\xee\x6e\x57\xcf\x7c\xad\xa8\xa6\x3d\xd2\x39\x46\xa1\x6d\xf4\x0c\xc3\xe3\x70\xda\xad\xcb\x00\x0d\x70\xaa\x85\x10\x08\xb8\x0e\xc3\xcf\x61\xdd\x87\xdb\xa6\x02\xd4\x21\x36\xa8\xa9\x05\x50\x5d\x72\x5d\x50\xab\xcc\x11\x51\xdb\x14\x88\xea\x56\xe5\x61\x2b\x41\x98\xcc\x08\x35\x27\xcd\xac\xb9\x45\xc0\x70\xa8\x0b\x63\x85\x55\x09\x22\x61\x68\x8e\xf1\x3f\xaa\x69\x14\xd5\xda\x41\xbc\x2e\xac\xa9\xee\x2d\xe8\x7c\xe0\x9c\x2a\x6f\xca\x19\x97\x87\xe0\x8e\x50\xc1\x33\xe6\x30\x0a\x86\xbb\x9e\x35\x30\xd2\x77\x8e\x6a\xa2\x18\x0c\x49\xdd\x24\xf4\x98\x8d\xa7\x1d\x8f\x92\xfb\x43\xaf\x23\xad\x6d\x3c\x0a\xac\x6e\xdd\xf1\x63\xbd\xdb\xe1\x5e\xbe\xdb\x4d\x48\x18\xe2\x7c\xb7\xcb\xbe\x62\x0e\xc0\xf2\x3a\x9d\x6a\x42\x73\xd8\x2e\x78\xef\x0f\x6b\x70\xc8\x90\x3a\x68\x08\x5c\x9d\x38\x40\xaa\xda\xa2\xe0\x68\xa9\x57\x46\xa6\x7a\x93\x07\xae\xb5\x2a\x82\x85\xdc\xba\xef\x85\xdc\xba\x68\x0b\x8d\x55\x11\x58\xbd\x4a\xe6\xc6\x8a\xd2\xfa\xa6\xcc\x53\xdf\x80\xd3\x44\x23\x3d\x6e\x95\x49\x7c\xf0\x10\x54\x51\x03\x9e\xb9\xed\x5f\xe5\x15\x05\x02\x77\xa3\x23\x83\xb5\x4b\x4e\x50\xb5\x34\xe2\x9c\x6b\x06\xe7\x18\x86\x9a\x2d\xe4\xf6\xda\x59\xf9\x3f\x43\x6f\xf3\x35\x01\x1a\xe8\x81\x25\x3a\x37\xb6\x14\x2a\xbf\x17\x93\x3f\xcb\x6d\x9d\x04\x14\x54\xf1\x9e\x66\x66\xae\xa6\xf6\xcf\x72\x0b\xec\x56\x94\x33\x69\xc1\x45\x3e\xd8\xea\x9a\x0f\x86\x8f\x34\xe7\xaf\xb2\x45\x8f\xae\x68\xbe\xdb\xe5\xa4\x94\x76\x55\xe6\xc1\xe9\x78\x94\xe1\x9c\xec\x01\x23\xd1\xde\x70\xef\x6f\xb7\x97\xce\x47\x8e\x6a\x4a\x32\xaa\x26\x11\x6c\x2e\x4c\x95\x12\x1f\x14\xdb\x9c\x99\x1f\x17\xb2\x8f\x7b\x20\x12\x89\xea\x8a\x19\x8c\x52\xb5\x66\xa8\x6f\x0e\x93\xfc\x36\x01\x1c\xf7\xa2\x18\xbc\x54\x3b\xe2\xa4\xd7\x86\x59\x5a\xc7\xed\x07\xb7\x47\x94\x61\xe9\x8e\xa8\xf0\xda\xd5\xbc\x78\x90\x9c\xf3\x5e\x34\x2e\x1a\x45\xc6\xd1\xe3\x48\x83\xfe\x5c\xf4\xc7\xdd\x13\x16\x90\x9d\xd5\x91\x55\x71\xd9\xad\x46\xe7\x5c\xfa\xdb\x43\x0d\x87\x8c\xbf\x4d\x05\x06\x07\x89\x01\x39\xbf\x18\xa4\x58\x9d\x4a\x1a\xb3\x76\xa0\xc7\xdd\x83\xc3\x94\xe8\x68\x68\x9b\x87\x82\xbc\x38\x67\x90\x55\x70\x73\x15\x8d\x4d\x1f\xd0\x66\x8c\x22\xe4\x1c\x26\xce\x01\xfa\xf2\xec\x2a\x1a\x67\x2d\xa5\xdd\x56\xea\xe1\x42\x75\xa8\x0d\x68\x96\xce\x51\xb9\x90\x3c\x8c\xc8\x6e\x17\xed\x5d\x39\xa0\xc2\x68\xdc\x7d\xf8\xb6\xf3\x43\x09\x97\x2c\xd7\xba\xd8\xed\x3a\x06\xb5\xa7\x25\x3f\xbf\xbd\xbb\xf9\x70\xce\xac\x34\x16\xe7\x62\xad\x66\xc2\xea\x92\xad\x8c\x2c\x7f\x9a\xc1\x75\xa2\x2b\xcf\x12\xfc\x81\x45\x5f\x61\x0b\xc3\x9e\x67\xfb\xa7\xaf\xb2\x8d\x5e\x20\xf6\x5b\xf7\x1e\x11\x39\x21\xa7\x5c\x32\x65\xfe\xa5\x16\xb0\xe1\x4d\x4a\x29\x6c\x8d\xfe\x9c\x15\x21\x52\x39\xc3\x03\xc0\x4e\x46\x0d\x80\xe1\x6d\xca\x4e\x9e\x75\xfd\x52\x22\x09\x60\x9c\x06\xd3\x1d\x30\xd9\x43\xa6\x59\xa9\x37\x59\x97\xe5\x45\x5c\x3d\x28\x84\x55\xec\x6d\x7d\x19\x4e\xbd\x2e\xe3\x60\xf4\xe3\x7c\x78\x85\xfa\xba\x8f\x7e\x3c\x87\x96\x3b\xf8\x2e\xf5\xe2\x0a\xf5\xad\xa3\x5e\x00\xb5\x79\xf7\x51\x90\x72\xf0\x4b\x79\x59\x47\x93\x56\x1c\xf2\xac\xb9\x06\x70\xd6\xa2\x36\xfc\x5c\xbf\x11\x19\xea\xcf\x3e\xd6\x95\x11\xf4\x78\xf5\x98\x34\xae\x7b\xe2\x3f\x46\x11\xad\x5f\x8e\x74\x0d\x3b\x8e\xf9\x80\x38\x94\x97\xb4\x82\x8c\xb1\xae\xc1\x63\x97\xb1\x26\x2a\x5a\x41\xec\xb8\x37\xa4\x9d\x50\x00\x9f\x4d\x6c\x8c\x73\x9a\x18\x13\x9f\x78\x0a\x73\x3a\xbc\xbe\xbb\xdb\x93\xfd\x28\xc3\x10\xa5\x0f\x0a\x6a\xd0\xe1\xe2\x87\x5e\xcb\xf2\xd0\x2f\x3e\x57\x5b\x8a\x1a\x39\x2f\xe5\xbb\xaa\xb2\xa5\xb9\xac\xe1\xe2\xad\x99\xc1\x75\xac\x30\x17\xf5\x3b\xbc\xd7\xf8\x32\x8a\xe8\x90\xec\x49\x35\xfd\xa1\xd7\x3d\x1c\xdd\x00\xb5\xa1\xbc\x24\x7b\x67\x29\xd3\xdc\xd3\xdb\xb3\xb1\xce\x8d\x01\x28\x7a\x88\x1e\x79\x1d\x8d\x6b\xe7\xde\x1e\x97\x25\x14\x98\x46\xb5\xc3\xfa\x0e\x20\x5e\x3b\x77\x87\xc6\x8e\x31\x5b\xbd\x5d\x07\xc7\x46\x2f\x1e\x57\xeb\x60\x76\x32\xe8\xef\x76\xba\x49\x6d\xba\x0f\x8a\x7b\x2f\xe3\x8b\xb5\x6a\x5c\xc4\xb9\xf7\x44\xc0\x44\x5b\xdc\x04\x10\xc2\x8d\x3b\x82\x79\xa8\x94\x99\xb0\x6a\x2d\x11\xad\xe4\x3c\x0d\xc5\x68\x2f\x22\xd5\xca\x7e\x86\x2f\x5a\x2f\xf9\x90\x6a\xbf\x90\x6d\x75\x5f\x09\x7d\x70\x33\x3a\x4a\xea\x9e\xc0\x18\x77\x13\x3a\xed\xe7\x27\xf1\xc9\xfd\x59\xbf\x90\xae\x16\xaa\xc5\x5b\xcb\x12\x3c\x0d\xbf\x60\x7f\x3c\x71\x48\xbc\xb9\x80\xee\xc6\xff\x92\x49\x61\x64\xb0\x11\xca\x32\xc6\xfc\xd5\xa7\xae\x0a\x1b\x43\x52\x40\x9b\x72\x6d\xdc\x8b\xa8\xab\x2b\xc0\x4d\x81\xdb\xf1\xdc\xbc\x6a\xd1\xfa\xa1\x8b\xfa\xe7\x2d\x74\x19\xfd\x1e\x51\xab\x8b\x18\xbd\x83\x96\x7b\xd7\x42\x97\x3f\x40\xa7\x7c\xb2\x3f\x65\x6a\x96\xc7\xc8\xdf\x43\x44\x13\x9d\xe9\x32\x46\xbf\x8b\xa2\x08\xd1\xea\xf9\x0c\x5d\x16\x4f\x81\xd1\x99\x4a\x83\xdf\x09\x21\x10\x9d\x88\x64\x31\x2b\xf5\x2a\x4f\xaf\x2b\xf6\xe9\x74\x8a\xa8\xc7\xde\x31\x02\xf1\xd1\x9e\x36\x8f\xd3\xf1\xf3\xb7\x24\xd9\xd3\xb6\xa4\x1e\x3f\xbf\x9c\xde\x49\x53\x5d\xeb\x98\xfd\xe1\x78\xa5\x4e\x1a\x15\xb7\xc8\xbf\xf6\x0e\xed\xf2\x3f\x44\xc5\x53\x25\xc0\xd0\x35\xbd\x04\x88\xba\xff\x94\xa8\x3b\xeb\x5d\xbb\x57\x49\x5a\xeb\x15\xfd\x00\xb4\x13\x32\x34\xeb\x25\x5d\x5d\xbc\xb2\x07\x34\xd8\xc8\xc9\x42\xd9\x81\x5f\x64\x50\x8a\x54\xad\x0c\xaa\x97\x46\x83\xa5\xfe\xf2\x1a\xed\x64\xf7\x9e\x36\x8f\xbd\xf1\xf9\xff\xcc\xad\x2d\xcc\xb9\xf2\xe1\xd4\x9b\x30\xcb\x74\x22\xc0\x44\xd9\xbc\x94\xd3\xdd\x0e\x21\x32\x46\x9f\xc5\x5a\x98\xa4\x54\x85\x8d\xa7\x22\x33\xd2\xbd\x55\x80\x17\x9c\x64\x22\x5f\x20\xda\x29\x27\x81\x81\xb9\x77\x70\xe7\xd5\xab\xda\x08\x98\x5f\xe3\xb9\x23\x7a\x5c\xba\x87\xbe\x16\xeb\x3b\xe6\x43\x34\x0c\x5d\x95\x07\xbe\xe8\x04\x95\x77\x51\xeb\x8f\xa3\xc3\x78\x10\xd1\xb6\x7a\xd3\x7c\x75\x4b\x3c\x31\x8a\x95\x23\xca\x1c\xba\xd3\x78\xad\x8c\x9a\x64\x12\xd1\xba\xb2\xe5\x6e\x50\x27\xaa\xf8\xcf\x2e\x84\xf5\x7d\xa7\x2a\xe0\xf1\x3b\x7a\x00\x64\x63\xd4\x78\x78\x7a\xe4\x34\xe3\xde\x70\xef\x1c\xf4\xc4\x25\xf2\xb4\xe0\x0f\x8f\xfb\x13\xa9\x60\x2a\xa7\x2a\x97\x61\xe8\x7f\x99\x58\xa6\xdd\x36\xf3\xb5\xf0\xb1\xef\xc1\x0f\xc8\x57\x02\xd1\x23\x95\x90\xbc\x78\x2a\xd9\x13\x4c\x46\xff\x17\x00\x00\xff\xff\x63\xbf\xf7\x6e\x2e\x24\x00\x00")

func publicJsJqueryBlockuiMinJsBytes() ([]byte, error) {
	return bindataRead(
		_publicJsJqueryBlockuiMinJs,
		"public/js/jquery.blockUI.min.js",
	)
}

func publicJsJqueryBlockuiMinJs() (*asset, error) {
	bytes, err := publicJsJqueryBlockuiMinJsBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "public/js/jquery.blockUI.min.js", size: 9262, mode: os.FileMode(420), modTime: time.Unix(1518116806, 0)}
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
	"public/index.html": publicIndexHtml,
	"public/js/app.js": publicJsAppJs,
	"public/js/jquery.blockUI.min.js": publicJsJqueryBlockuiMinJs,
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
	"public": &bintree{nil, map[string]*bintree{
		"index.html": &bintree{publicIndexHtml, map[string]*bintree{}},
		"js": &bintree{nil, map[string]*bintree{
			"app.js": &bintree{publicJsAppJs, map[string]*bintree{}},
			"jquery.blockUI.min.js": &bintree{publicJsJqueryBlockuiMinJs, map[string]*bintree{}},
		}},
	}},
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


func assetFS() *assetfs.AssetFS {
	assetInfo := func(path string) (os.FileInfo, error) {
		return os.Stat(path)
	}
	for k := range _bintree.Children {
		return &assetfs.AssetFS{Asset: Asset, AssetDir: AssetDir, AssetInfo: assetInfo, Prefix: k}
	}
	panic("unreachable")
}
