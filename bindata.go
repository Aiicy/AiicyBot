// Code generated by go-bindata.
// sources:
// locales/locale_en-US.ini
// locales/locale_zh-CN.ini
// DO NOT EDIT!

package main

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

var _localesLocale_enUsIni = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x4c\xcd\x3b\x0e\xc2\x40\x0c\x84\xe1\xde\xa7\x18\x0a\x3a\xae\xb0\x3d\x77\x40\x14\x79\x98\xd8\x92\xd7\xb6\xb2\x2b\x56\xdc\x1e\x25\x34\x54\x23\xcd\x57\xfc\xc2\x66\x81\x82\xfb\xb1\x37\x5c\xdb\x85\x1e\xc2\x96\x4f\x12\x45\xc1\x8f\x47\xec\xb6\x52\xea\x82\x82\x8d\x3b\xba\x30\xf6\xc9\xd7\xa8\xa8\x15\xa9\x0b\xa5\xfa\xf6\x8f\x9c\xf6\x81\x36\x64\xf8\x86\x21\xec\xe7\x3d\x47\x47\xb8\xa9\x33\x1d\x0d\x14\x34\x89\x71\xd2\x34\xc7\x9b\xa1\xfe\x8a\x6f\x00\x00\x00\xff\xff\x08\xb3\xe3\xc4\x92\x00\x00\x00")

func localesLocale_enUsIniBytes() ([]byte, error) {
	return bindataRead(
		_localesLocale_enUsIni,
		"locales/locale_en-US.ini",
	)
}

func localesLocale_enUsIni() (*asset, error) {
	bytes, err := localesLocale_enUsIniBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "locales/locale_en-US.ini", size: 146, mode: os.FileMode(420), modTime: time.Unix(1529846076, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _localesLocale_zhCnIni = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x4c\xcd\x4d\x8a\x83\x30\x18\x80\xe1\x7d\x4e\xf1\x21\xcc\x31\x72\x92\x61\x56\x83\xa8\x20\x9a\x2b\x38\x08\xc3\x88\xe3\xac\xe6\x2f\x06\x69\x41\x8b\xdd\xf4\x87\x2e\x2a\x69\x8c\x97\xc9\x97\xb4\x2b\xaf\x50\x4a\x37\x5d\xbf\xf0\x3e\xa1\x1f\xc7\x29\x50\x30\xe3\x02\xbb\x71\x56\xe5\xd3\xac\x32\xf2\x1c\xfa\x31\x7b\x21\x61\xf4\x58\xcc\xf0\xe3\xbe\x4b\xc2\xa2\x57\xa0\x70\xae\x8e\xa8\x7f\x2f\xfc\xcb\x0a\xe9\x78\xee\x74\x85\xdd\x01\x6b\xed\x3e\xde\x09\x8b\x92\x00\x28\xe0\xea\xcd\x36\xc2\x0a\x89\xff\xbd\x91\x12\x45\xef\xe4\x34\xab\x12\x77\x7b\xa3\xf8\x7d\x80\x75\x83\xed\xa7\xc7\xd2\x24\xf0\xc8\x0d\x05\x0a\xf6\x4f\xbb\x56\x9a\x53\x67\x86\xc2\xf1\x1c\x87\x0d\x16\x6b\x33\x2d\x6d\xb6\xbd\x06\x00\x00\xff\xff\xc1\x88\x01\xb9\xae\x00\x00\x00")

func localesLocale_zhCnIniBytes() ([]byte, error) {
	return bindataRead(
		_localesLocale_zhCnIni,
		"locales/locale_zh-CN.ini",
	)
}

func localesLocale_zhCnIni() (*asset, error) {
	bytes, err := localesLocale_zhCnIniBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "locales/locale_zh-CN.ini", size: 174, mode: os.FileMode(420), modTime: time.Unix(1529846171, 0)}
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
	"locales/locale_en-US.ini": localesLocale_enUsIni,
	"locales/locale_zh-CN.ini": localesLocale_zhCnIni,
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
	"locales": &bintree{nil, map[string]*bintree{
		"locale_en-US.ini": &bintree{localesLocale_enUsIni, map[string]*bintree{}},
		"locale_zh-CN.ini": &bintree{localesLocale_zhCnIni, map[string]*bintree{}},
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

