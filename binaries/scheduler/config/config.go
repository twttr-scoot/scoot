// Code generated by go-bindata.
// sources:
// config/config.go
// config/inMemory.json
// config/local.json
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

	info := bindataFileInfo{name: "config/config.go", size: 0, mode: os.FileMode(420), modTime: time.Unix(1476107478, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _configInmemoryJson = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xaa\xe6\x52\x50\x50\x72\xce\x29\x2d\x2e\x49\x2d\x52\xb2\x52\x00\x71\x15\x14\x94\x42\x2a\x0b\x52\x95\xac\x14\x94\x72\x53\x73\xf3\x8b\x2a\x95\x74\x20\xa2\xce\xf9\xa5\x79\x25\x4a\x56\x0a\x86\x06\x5c\x0a\x0a\xb5\x5c\xb5\x80\x00\x00\x00\xff\xff\xfe\xbe\xda\x8f\x3c\x00\x00\x00")

func configInmemoryJsonBytes() ([]byte, error) {
	return bindataRead(
		_configInmemoryJson,
		"config/inMemory.json",
	)
}

func configInmemoryJson() (*asset, error) {
	bytes, err := configInmemoryJsonBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "config/inMemory.json", size: 60, mode: os.FileMode(420), modTime: time.Unix(1475870485, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _configLocalJson = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xaa\xe6\x52\x50\x50\x72\xce\x29\x2d\x2e\x49\x2d\x52\xb2\x52\x00\x71\x15\x14\x94\x42\x2a\x0b\x52\x95\xac\x14\x94\x72\xf2\x93\x13\x73\x94\xb8\x14\x14\x6a\x75\x40\xea\xc2\xf3\x8b\xb2\x53\x8b\x8a\x31\xd5\x15\x15\x24\x83\x55\x71\xd5\x02\x02\x00\x00\xff\xff\x03\xdc\x1c\x93\x50\x00\x00\x00")

func configLocalJsonBytes() ([]byte, error) {
	return bindataRead(
		_configLocalJson,
		"config/local.json",
	)
}

func configLocalJson() (*asset, error) {
	bytes, err := configLocalJsonBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "config/local.json", size: 80, mode: os.FileMode(420), modTime: time.Unix(1475870485, 0)}
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
	"config/config.go":     configConfigGo,
	"config/inMemory.json": configInmemoryJson,
	"config/local.json":    configLocalJson,
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
		"config.go":     &bintree{configConfigGo, map[string]*bintree{}},
		"inMemory.json": &bintree{configInmemoryJson, map[string]*bintree{}},
		"local.json":    &bintree{configLocalJson, map[string]*bintree{}},
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
