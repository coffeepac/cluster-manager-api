// Code generated by go-bindata.
// sources:
// api/swagger/api.swagger.json
// DO NOT EDIT!

package swaggerjson

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

var _apiSwaggerJson = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xec\x5b\x5f\x6f\xdc\xb8\x11\x7f\xf7\xa7\x18\xa8\x05\x9a\x03\x62\x6f\x2e\x7d\xcb\xcb\xd5\xf0\x01\x77\xc6\xc5\xa9\xd1\x0d\xe2\x87\x26\x30\xb8\xd2\xac\xc4\x33\x45\x2a\xe4\x68\xb7\xdb\xc2\xdf\xbd\x20\xf5\x9f\x2b\xd9\xbb\xd2\xd6\x5e\xd7\x77\x2f\x17\x8b\xe4\xfc\xfd\x71\x38\x1c\xce\xfe\xe7\x04\x20\x30\x6b\x16\xc7\xa8\x83\x0f\x10\xbc\x3f\x7b\x17\xbc\xb5\xdf\xb8\x5c\xaa\xe0\x03\xd8\x71\x80\x80\x38\x09\xb4\xe3\x17\x22\x37\x84\x1a\xae\x98\x64\x31\x6a\x38\xbf\xbe\x74\xf3\x01\x82\x15\x6a\xc3\x95\xb4\xb3\x56\xef\xce\x7e\x2c\x09\x01\x04\xa1\x92\xc4\x42\xaa\xa9\x01\x04\x92\xa5\x8e\xdc\x15\x0f\x13\x86\x02\xbe\xa0\xc4\x7f\x73\x56\xae\x00\x08\x72\x2d\xec\x78\x42\x94\x99\x0f\xb3\x59\xcc\x29\xc9\x17\x67\xa1\x4a\x67\x2b\x7f\x2a\xa6\x8c\xbb\xc9\x69\x39\xf4\xb7\xd8\x7e\xb1\x93\x03\x37\xe7\xfe\x04\xe0\xde\x29\x65\xc2\x04\x53\x34\xc1\x07\xf8\x67\x21\x9a\xa5\x5f\x89\xe9\x78\xd9\x15\xdf\xdc\xdc\x50\x49\x93\x77\x26\xb3\x2c\x13\x3c\x64\xc4\x95\x9c\xfd\x6e\x94\x6c\xe6\x66\x5a\x45\x79\xb8\xe3\x5c\x46\x89\x69\x2c\x3b\x63\x19\x9f\xad\x7e\x9c\x85\x85\x61\xdb\x46\x8a\xb1\x6d\x33\x2b\x7e\x9e\xa6\x4c\x6f\xac\xae\x37\x5c\x08\xd0\x48\x9a\xe3\x0a\x81\x12\x04\x43\x8c\x72\x03\x6a\x09\x0c\x4a\x62\xc0\x64\x04\x9c\x0c\xdc\xe5\x0b\x0c\x95\x5c\xf2\x18\x96\x4a\x43\xa8\xa4\xc4\x90\xf8\x8a\xd3\xa6\xb6\x23\x40\xa0\x32\xd4\x4e\xe4\xcb\xc8\xf2\xf8\x05\xa9\x74\x77\x7b\x92\x46\x93\x29\x69\xd0\x74\x64\x03\x08\xde\xbf\x7b\xe7\x7d\x02\x08\x22\x34\xa1\xe6\x19\x95\xc0\x68\x11\x2a\x34\xb2\x0e\x61\x5b\xcb\x00\x82\x3f\x6b\x5c\xda\x15\x7f\x9a\x45\xb8\xe4\x92\x5b\x0a\xa6\xb2\xd2\x6d\x5a\xc0\xef\x96\x65\xbc\x91\xf2\x1f\x98\x89\x4d\xd0\x21\x74\x7f\xd2\xf7\xef\xfb\x96\x3a\x19\xd3\x2c\x45\x42\xdd\x38\xaf\xf8\xcf\x53\xa4\x42\xac\xfb\xff\xdb\x07\x95\xfc\xc4\x52\xb4\x7e\xb0\x5e\xa9\x3c\x41\x0a\x16\x08\x42\xa9\x3b\x8c\x20\xcf\xce\x7c\x12\xdc\xad\xfc\x9e\xa3\xde\xf8\x43\x1a\xbf\xe7\x5c\xa3\x75\xc9\x92\x09\x83\xde\x30\x6d\x32\x27\x98\x21\xcd\x65\x1c\xf4\x2a\xfc\xad\xa5\x30\xb1\xd8\x57\xb5\xda\xd5\xcd\xe2\x6f\x27\x9e\xa5\x82\x08\x05\x12\x3e\x8c\xc7\x62\x4e\x83\xbf\x07\xb0\xf5\xb3\x9b\xfa\x02\xe0\xd5\x11\xf4\x58\x10\x76\x93\x30\x02\x6e\xda\x08\xfb\x8b\x01\xbb\xd0\x02\x2d\x42\x43\x5a\x6d\x5e\x1e\xc6\xb2\xfc\x91\x80\x97\x69\xb5\xe2\xf6\x90\xd9\x09\x63\x17\x1a\xd9\x8b\xc0\x58\x47\xd0\x27\xc1\xd8\x42\x45\x5b\x18\x28\xe0\xd1\x37\xd2\x42\x07\xe9\xdc\x07\xc7\xa1\x0d\x70\x65\xe2\x5d\xd4\x1f\x8f\xb7\x93\x96\xf5\xfc\xf3\x77\x26\xb8\xa1\x71\x87\x30\x03\xbb\xd6\x06\xfe\x92\x96\xd9\xe9\x6c\xfd\x68\x19\x1e\x31\x38\xbb\x92\x8e\x42\xe7\x01\xdc\x13\x23\x65\x2a\x0a\x55\x2e\x3b\xde\xc9\x94\x19\x76\xcf\x9c\xa7\x99\x40\xc8\x54\x04\x6e\x21\x54\xa6\x3d\x03\x28\x5d\x67\xff\x8e\x60\xcd\x29\x71\xb1\x54\xe6\xe9\x02\xb5\x75\x61\xa6\x22\x03\x5c\x16\x5f\x59\x8a\x26\x63\x21\x16\xf1\x27\xc2\xe8\x61\xc7\x5e\xab\xe8\xc2\x09\x7a\xdc\x5e\xad\xc4\x7c\x8d\x01\xa7\xa5\xfe\xf3\x84\x9b\x04\x85\x50\x6b\xa5\x45\xb4\x2f\x9c\x7f\xb5\x2b\xe1\xc6\x2e\x85\x39\xea\x15\x0f\x11\x4e\x61\x5d\xe0\x39\x43\x46\xc0\x20\xd6\x88\xc4\x65\x6c\xb3\x81\x0a\xc1\xbb\x80\xd7\xd1\x76\xa4\x8f\x19\xbb\x8d\x94\xaf\x11\xba\x8d\xf6\xcf\x86\xdc\xb4\x8d\xd9\x3d\x2e\x06\x11\x66\x42\x6d\x30\x02\x4b\x03\xc2\x84\x69\x7a\xf4\x8a\xf0\x2b\x8a\xf4\xc2\x9f\x79\x6c\x88\xf4\x44\x7d\x12\x58\x6e\xe7\xbf\x7d\xca\x56\xf5\x1a\xe2\x42\xa0\x06\x93\xa8\x5c\x44\xf6\x2a\xca\xa5\x21\x26\x04\x46\xa0\xe4\xd3\xdc\x14\xde\x3e\xae\x53\x21\xe5\x59\x7d\xe0\xee\x71\x0d\x2a\x15\xac\x97\x1e\x9b\x4e\x55\x75\x6c\x77\x8d\xca\x15\xae\x92\xe0\x68\xb8\xad\x77\x34\x7a\xf9\xdb\xb7\x4f\x1d\xb7\x1b\xe0\x13\x4b\x9f\xc8\x1d\x07\x88\x79\xbb\x5f\x46\xcb\x1d\x04\x6f\x94\x06\x8d\xe5\x5f\x3f\xb4\x62\xdb\x57\xf9\x39\xe1\xa6\x38\x9a\x17\x58\x5c\xcb\xf3\x90\xf8\x0a\x81\xbb\x2a\x9d\xb3\x4e\xc2\x0c\x30\xa1\x91\x45\x1b\x58\x20\xca\x26\x4a\xd6\x69\xa9\xb1\xc7\xb7\x57\x16\xf0\xa3\xe5\x65\xc1\xff\x45\x84\x4b\x5f\xd6\xd7\x78\x8c\xfb\x36\x78\x9e\xc3\xbc\x08\x2c\x3b\xa7\xa0\x0f\xe1\xbe\x24\x35\x8c\xd0\xeb\xaa\x78\xf3\x79\x6b\xe6\xb1\x01\xd4\x13\xf5\x35\xe2\xd3\x33\xc1\x13\xc1\xb3\x7e\xa9\x69\x49\xd7\xbc\x95\xf4\x55\x8e\x9a\xd2\xc4\x25\x61\x27\x31\xad\xce\x09\xb5\xf8\x1d\xc3\x26\x1e\x06\x99\xb6\xc0\x24\xee\xa1\xad\x72\x44\x07\x7f\xde\x59\xf3\xb6\x33\x56\xbd\x89\xf5\xd4\xfb\x83\x5e\x48\x14\x4f\x34\xa3\x78\xb4\x13\x83\xe6\xa5\xa7\x9f\x63\xef\x8e\x7f\xac\xea\x76\x7e\x33\x9f\x67\x18\x4e\x31\xa0\xc6\xb8\xd8\x92\xfb\xab\xf7\x39\x41\x28\x96\xbb\x07\xaa\xf3\x9b\xf9\x80\x05\x31\xd4\x48\xb7\x77\xb8\xb9\xe5\xd1\x68\x4e\x73\x47\xe5\x37\xdc\x5c\x46\x05\xbb\xeb\x4b\x38\x0f\x43\x34\xe6\x41\xae\xcc\x4d\xb1\xcc\x27\x72\x2e\x78\xfd\x86\x9b\x9a\x3b\xf3\xb9\x8f\xf2\xe1\x15\x63\x93\x9d\x88\x32\xca\x14\x97\x34\x5a\x45\x2b\x84\xd3\xa9\xa6\xd4\x6b\xd3\xdc\xa0\x1e\xbd\xe3\x2c\x9f\x8a\x40\x55\xa5\xab\xf8\xf6\xb3\x53\x2c\xa7\x64\x92\xeb\xfe\x7e\x9e\x53\x02\x77\xa5\xd3\x2c\xc7\x1e\x05\xc7\xb9\xcd\xc4\xc7\x12\xb7\xca\x77\xca\xfa\x9d\x05\xa3\x7e\x73\x96\xb5\x24\xed\xf3\x1d\x57\xf1\xbf\x2e\xa9\x39\xec\x0e\x7b\xa0\x62\x0a\x26\xc3\x90\x2f\xcb\xe7\xf5\xa9\xd6\xef\x30\x7f\x0e\x37\xb4\x43\x7b\xad\xe2\x29\x84\xb9\xd6\x28\x49\x6c\x80\xec\x2d\x82\x1b\x60\x6b\x03\x4a\x43\xca\xd8\x40\x9c\x72\x23\x07\xf1\x48\x1d\x49\x1e\xd9\xe7\x43\x9e\x68\x09\xc5\xd6\x07\x92\xa9\x3a\xa1\x86\x45\x3a\xbf\x19\x96\x68\x14\x36\x8a\xc4\x6f\x02\x28\xd4\xdd\x10\x24\x16\x4a\x09\x64\xdd\x7a\x40\xb0\x54\x3a\x65\x34\x38\xdc\x82\x0c\x52\x82\xda\xa2\x41\x2a\xea\x6c\xe0\x35\x33\xed\xed\x0b\x8b\x12\x3f\x36\x61\x44\x33\x10\x8c\x27\x24\x26\x9f\xcb\x84\x64\x2b\x92\xc4\x28\xcb\xf4\x7f\x84\x17\x7a\x5e\xdd\x8f\xcf\x0b\x17\xae\x96\xd6\xd6\xb9\xba\xe5\xdb\xeb\xfb\xc1\x2d\x3d\xef\xa4\x7d\xdb\xfe\xdc\xcf\xb6\xde\xe5\xfb\xf8\xac\x7b\xc3\x8a\x90\x58\xdf\x22\xc1\xe4\x2e\x4b\x5a\xe6\x62\x20\xfc\xa1\x31\x2c\x9e\x10\x84\x4b\x02\x06\xd6\xa8\x11\x62\xbe\xc2\x31\xd8\xfd\xc5\x02\x9f\x87\x4d\x25\xe6\xb9\x0f\x15\xd9\x3a\xe7\x8b\xe2\x52\x8a\x43\x69\x59\x53\x7a\x3d\x04\xd7\xe2\xd5\xd4\xb5\xa4\x58\xb6\xee\x7b\xab\x33\x0e\x48\xf5\x4b\xa1\x31\x53\x93\x05\x28\xca\x6b\x96\x94\xe1\xa4\xf4\xa6\x9f\x55\xe8\xb9\x68\x0a\x2f\xe7\xad\x5e\x2e\x2b\x26\xf2\xed\xea\xc6\xee\x6c\x98\xc6\x72\x33\xb8\xb3\x12\xde\x48\x34\x84\x11\x6c\x58\x2a\xe0\xd4\x0d\x7d\x71\x2c\xce\xdc\x97\x50\x49\x42\x49\xe6\x87\xf1\xf0\x2d\xae\xfd\x73\x24\xb2\xa2\x4d\x84\xf0\x61\x10\xe5\x97\xf8\x07\x4c\x5d\xb7\xa1\x4e\xe3\xd6\x5f\x7e\x1f\x65\xd0\xed\x06\x8a\xe3\x0b\xb7\x97\x9d\x6e\xb2\xea\x52\x65\x36\x86\x30\x1d\xd8\x39\x55\x97\xcb\x80\x38\x4c\x6b\xd6\xad\x5b\x05\x9c\x30\xf5\xe7\xef\x95\x0f\x7a\xd5\x9e\x16\x99\xfb\x7e\xb5\x3e\xfa\x2d\x39\x53\xfc\xf7\xff\xe3\xbb\x27\xaa\x44\xb5\x38\x36\x2d\xc8\x93\xb9\xb6\xba\x99\x49\x55\xcd\xcc\x55\xab\xc3\xf8\x72\x98\xd7\x13\xf2\x6c\x21\xef\x53\xfb\xe0\x5c\x22\x85\xc5\x23\x50\xdd\xc6\xf4\x55\x7e\x44\xb6\x42\xc0\x34\xa3\x8d\x9d\xe3\x1e\xd1\x80\x09\xd1\x84\xc6\x91\x38\xef\xb6\x04\x4d\xb0\x40\xa6\xa2\x41\x74\x71\x49\x18\x7b\x0f\xd7\x2d\xa8\x73\x49\x7f\x7d\x3f\x64\x9a\x47\xfa\xb3\xde\x28\xed\xd9\x61\xcc\x09\xd8\xed\xb0\xf8\xdf\x66\x6f\xdb\xd2\xd5\x54\x3b\x57\xac\x76\xd7\x4f\x27\xf5\xdf\x49\x8f\xc9\xfe\xdc\x2d\xb5\xde\x55\x99\xea\xb9\xc7\x82\xb7\xa5\xd8\xa3\x0a\xf5\xbd\x99\x4d\x50\x6a\xfb\x27\x17\x7d\x4a\x0d\xdc\xfd\x76\xe9\xab\xe8\x8f\x84\x5b\x6f\x6e\xb0\xd7\x19\xd8\x9b\x9f\x0d\xdc\xcd\x4b\xe1\x8a\x49\x03\x45\xa3\xde\x04\x78\x5f\x69\xfa\x9e\x9d\x3b\xe6\x72\xf9\xf1\x7c\x5b\x90\x5d\xf7\x64\xff\x93\xf1\xf1\x1d\xc5\x2f\xf6\xd6\xda\xf3\xde\xf7\xf2\xb6\xd6\xc1\x0e\xde\x61\xd6\x7c\x80\xf5\x94\x6b\xc7\x17\xbb\x56\x2a\xef\xaa\x61\xa3\x63\xf5\xd6\x3e\x83\x3c\x8b\x35\x8b\x70\xf0\xca\x5c\xb9\x74\xcd\xa3\x41\xed\x0f\x91\x64\x96\xa6\xa9\x7f\x7f\x71\x6a\x19\x96\x5f\x7f\xfa\x2a\xe7\x85\xc9\x38\x41\x62\x73\x94\xce\x24\x16\xa5\x5c\x42\xa6\xf9\x8a\x0b\x8c\xd1\xfc\x34\x50\x30\xb6\xd3\x6e\x5b\xa9\xcc\xe4\xfb\xc5\x8e\x8d\x4e\x3d\x48\xb0\x3a\xdb\xc4\x93\x5a\x58\x60\x0b\xe1\x4e\xae\x42\x9d\x51\xe5\xcd\xde\xfe\x82\x3f\x82\xd9\x9e\xc1\xac\x6e\x16\xc0\x7f\x11\x6a\xc9\xc4\xcf\x2a\x6c\x75\x0b\x78\x6d\x22\x57\x4a\x5b\xdf\xa9\x9c\xe0\x81\x1f\xb1\x0e\xff\xee\xd4\xb0\xd4\xe4\x32\x3e\x0d\x65\x48\xd5\x49\x78\x5a\xfa\xf4\x94\x65\xdc\x8a\x79\x7f\x72\x7f\xf2\xdf\x00\x00\x00\xff\xff\x40\xcc\x03\xd3\x52\x3b\x00\x00")

func apiSwaggerJsonBytes() ([]byte, error) {
	return bindataRead(
		_apiSwaggerJson,
		"api.swagger.json",
	)
}

func apiSwaggerJson() (*asset, error) {
	bytes, err := apiSwaggerJsonBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "api.swagger.json", size: 15186, mode: os.FileMode(420), modTime: time.Unix(1526658472, 0)}
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
	"api.swagger.json": apiSwaggerJson,
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
	"api.swagger.json": &bintree{apiSwaggerJson, map[string]*bintree{}},
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
