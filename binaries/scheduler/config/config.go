// Code generated by go-bindata.
// sources:
// config/config.go
// config/local.local
// config/local.memory
// DO NOT EDIT!

package config

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

var _configConfigGo = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x01\x00\x00\xff\xff\x00\x00\x00\x00\x00\x00\x00\x00")

func configConfigGoBytes() ([]byte, error) {
	return bindataRead(
		_configConfigGo,
		"config/config.go",
	)
}

func configConfigGo() (*asset, error) {
	bytes, err := configConfigGoBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "config/config.go", size: 0, mode: os.FileMode(420), modTime: time.Unix(1486507161, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _configLocalLocal = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x6c\x90\x41\x6b\x32\x31\x10\x86\xef\xfe\x8a\x90\xb3\x7c\x9f\xb5\xb4\x94\xbd\x6a\x2f\xa5\x8b\x8b\x2e\xf4\x3c\x66\x27\x31\x98\xcd\x2c\x93\x89\x54\x8a\xff\xbd\xc4\xba\x68\xb1\xb9\x65\x9e\x87\x99\x77\xe6\x6b\xa2\x94\x5e\x84\x9c\x04\x59\x57\xaa\x7c\x95\xd2\xed\x71\x40\x5d\x29\x1d\xc8\x40\xd0\x13\xa5\x4e\xd3\xe2\x7d\x10\xef\x91\xd3\xbd\xc7\x83\xd1\xd3\x9f\x52\x43\x21\xf8\xe8\x1a\x64\x4f\x5d\x61\xf3\xa7\x59\x9f\x46\xfa\x1a\x2d\xb1\xc1\x16\xd2\xbe\xf5\x3d\x52\x16\x5d\x29\xe1\x8c\x17\xfe\x1b\xe8\xc7\x59\x7f\x9d\xbe\x31\x3b\xec\x72\x40\x5e\x50\xb4\xde\xdd\xa7\x48\x02\x82\x36\x87\x71\x58\x0d\x9f\x6b\x14\xf6\x98\x1a\xe4\xd2\x59\xab\x4a\xcd\x2f\x70\x89\xdb\xec\x6a\xea\xb0\x14\x2d\x84\x34\x46\x58\xa3\xa1\x03\xf2\x1b\x6d\xd3\x2a\x6e\x04\x58\xf2\x50\x9c\x9b\x94\x4b\xb4\x90\x83\xdc\x84\xad\x53\x51\x1e\x5e\x66\xe5\x8d\x8d\x72\x8c\xc8\xab\x03\xf2\x0e\xa1\xab\xcb\xd9\x9e\xcf\xfc\xba\x11\x38\x78\xa7\x3f\x36\xb1\x3e\xe0\xb8\xc5\xd2\x33\x1a\x21\x3e\x16\xf0\x2f\x19\x22\xe9\x40\xe0\x7f\x71\x12\x38\x08\xe4\xce\x37\x9a\x9c\xbe\x03\x00\x00\xff\xff\x29\xbc\x57\x88\xcc\x01\x00\x00")

func configLocalLocalBytes() ([]byte, error) {
	return bindataRead(
		_configLocalLocal,
		"config/local.local",
	)
}

func configLocalLocal() (*asset, error) {
	bytes, err := configLocalLocalBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "config/local.local", size: 460, mode: os.FileMode(420), modTime: time.Unix(1486507145, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _configLocalMemory = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x6c\x90\xc1\x6a\xc3\x30\x10\x44\xef\xfe\x8a\x65\xcf\x3d\x24\x3d\x94\xa2\xab\x73\x2a\x35\x29\x89\x7f\x40\x89\xc7\x8e\xa9\x2c\x85\xd5\x6e\x68\x28\xf9\xf7\xa2\xd6\xa6\xd0\x46\x37\xcd\x3c\x1e\xcc\x7e\x56\x44\x5c\x07\xcb\x0a\x61\x47\xe5\x4b\xc4\xed\xf5\x0c\x76\xc4\x13\xa6\x24\x57\x7e\xf8\x49\xeb\x64\x51\xd9\xd1\x7a\x55\x11\xdd\x4a\xc8\xfb\xe3\x09\x9d\x05\x48\x9d\x62\x3f\x0e\xff\x0d\x59\xbd\xa2\xb7\xb0\x38\x1a\xff\xb1\x83\xca\x88\xfc\x06\x69\x7d\x7e\x67\x72\xf4\x38\x97\x1b\x1c\x6c\x68\x52\x87\x12\xf6\x3e\x64\xcc\xc5\x0e\xc7\x74\x81\xbc\xa4\x43\xde\xc6\xbd\x7a\x51\x3b\xff\x65\x36\xe8\xbd\x05\x2d\xce\x76\x9c\x90\x4c\x9b\x5c\x98\xf5\xf3\xaa\xbc\xc5\x64\x31\x42\xb6\x17\xc8\x09\xbe\x6b\x32\x3b\x7a\xfa\xee\x7f\x27\xf9\xc1\xbf\xa6\x3b\x53\xe6\x63\x14\xb0\xba\x7d\x05\x00\x00\xff\xff\x0a\xdb\x3f\x04\x38\x01\x00\x00")

func configLocalMemoryBytes() ([]byte, error) {
	return bindataRead(
		_configLocalMemory,
		"config/local.memory",
	)
}

func configLocalMemory() (*asset, error) {
	bytes, err := configLocalMemoryBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "config/local.memory", size: 312, mode: os.FileMode(420), modTime: time.Unix(1486507155, 0)}
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
	"config/config.go": configConfigGo,
	"config/local.local": configLocalLocal,
	"config/local.memory": configLocalMemory,
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
	"config": &bintree{nil, map[string]*bintree{
		"config.go": &bintree{configConfigGo, map[string]*bintree{}},
		"local.local": &bintree{configLocalLocal, map[string]*bintree{}},
		"local.memory": &bintree{configLocalMemory, map[string]*bintree{}},
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

