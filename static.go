// Code generated for package dis_doi by go-bindata DO NOT EDIT. (@generated)
// sources:
// sensitive_words
// sensitive_words_test
// whitelist
package dis_doi

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

// Mode return file modify time
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


func sensitive_wordsBytes() ([]byte, error) {
	return bindataRead(
		_sensitive_words,
		"sensitive_words",
	)
}

func sensitive_words() (*asset, error) {
	bytes, err := sensitive_wordsBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "sensitive_words", size: 652564, mode: os.FileMode(384), modTime: time.Unix(1710826342, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _sensitive_words_test = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x44\x8c\xc1\x4e\xc2\x40\x10\x86\xef\xf3\x96\x03\x6c\xb6\xd3\x94\xd9\x60\xd9\x94\x7a\x82\xc4\x44\x14\x23\x20\x07\x49\x00\xe3\x85\x44\x8d\x86\x6a\x0f\x74\x0f\xc5\x97\xe9\xac\xf0\x16\xa6\x11\xd3\xd3\xcc\x37\xf3\xfd\x3f\x72\x4b\x51\x48\x1c\x23\x43\x4b\x71\x84\x1d\xc5\x1a\xd0\x74\x51\x5d\x40\x3b\x40\x8a\x88\x35\x74\x30\x42\x82\x90\x90\x75\xaf\xe6\xcb\x00\x59\xb7\x03\xcb\x3d\x42\x03\x29\x9a\x44\x71\x6a\x91\x21\x41\xd6\x81\x61\x9d\x28\x86\x8e\x61\xdd\x27\x34\x29\xb1\x0e\x09\x06\xd4\x57\x91\x82\xae\x89\xad\x89\x88\x09\x74\xdd\x61\x79\x40\x36\x0e\x2c\x9c\x96\x5f\xe0\x37\xf7\x27\x37\x05\x19\x8f\x8e\xdf\x2b\x71\x39\x54\xee\x59\xb2\xf2\xb8\xbb\x3a\xdf\x9a\x4d\x5c\x2e\xb3\x85\xbf\x7d\x6f\x1c\x99\x2d\x64\x9a\x35\x66\x55\x4c\x6a\x3e\x8f\xff\xef\x5f\xc8\xef\x56\x3e\x2f\xab\x62\x0d\x32\x7f\xf1\xa3\x57\xbf\xd9\x82\x8c\x97\x92\xcd\x65\x7b\x0d\x7e\xfd\x29\x87\x3d\x48\xf9\x21\x87\xbd\x1f\x0e\xc1\x3f\x3d\x4a\xf9\x06\x52\xdc\xfd\xdc\x38\x99\x3c\xfc\x06\x00\x00\xff\xff\x12\x1b\xaa\xa6\x36\x01\x00\x00")

func sensitive_words_testBytes() ([]byte, error) {
	return bindataRead(
		_sensitive_words_test,
		"sensitive_words_test",
	)
}

func sensitive_words_test() (*asset, error) {
	bytes, err := sensitive_words_testBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "sensitive_words_test", size: 310, mode: os.FileMode(420), modTime: time.Unix(1710903120, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _whitelist = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xaa\xca\xc8\xcc\x2e\x05\x04\x00\x00\xff\xff\xef\x17\xb2\x6a\x05\x00\x00\x00")

func whitelistBytes() ([]byte, error) {
	return bindataRead(
		_whitelist,
		"whitelist",
	)
}

func whitelist() (*asset, error) {
	bytes, err := whitelistBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "whitelist", size: 5, mode: os.FileMode(420), modTime: time.Unix(1711078202, 0)}
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
	"sensitive_words":      sensitive_words,
	"sensitive_words_test": sensitive_words_test,
	"whitelist":            whitelist,
}

// AssetDir returns the file names below a certain
// directory embedded in the file by go-bindata.
// For example if you run go-bindata on data/... and data contains the
// following hierarchy:
//
//	data/
//	  foo.txt
//	  img/
//	    a.png
//	    b.png
//
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
	"sensitive_words":      &bintree{sensitive_words, map[string]*bintree{}},
	"sensitive_words_test": &bintree{sensitive_words_test, map[string]*bintree{}},
	"whitelist":            &bintree{whitelist, map[string]*bintree{}},
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