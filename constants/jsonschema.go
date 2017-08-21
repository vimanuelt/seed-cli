// Code generated by go-bindata.
// sources:
// schema/seed.manifest.schema.json
// schema/seed.metadata.schema.json
// DO NOT EDIT!

package constants

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

var _schemaSeedManifestSchemaJson = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xec\x5a\x4b\x73\xdb\x36\x10\xbe\xf3\x57\x60\xd0\x1c\x24\xbb\x8c\x9d\x99\x5e\xec\x4b\x27\x33\xf5\xa1\x87\x3e\xa6\xe9\xe4\x10\x4b\xcd\x40\xe4\x8a\x82\x43\x02\x2c\x08\xb6\xa3\xd4\xfa\xef\x1d\x3e\x2c\x51\x78\x53\xb2\xd2\x66\xac\x5c\xe2\x59\x2c\x76\x17\xcb\xef\xfb\x00\x91\xf8\x27\x42\x08\xbf\xaa\x92\x15\x14\x04\xdf\x22\xbc\x92\xb2\xbc\xbd\xba\x7a\xa8\x38\x8b\x3b\xeb\x6b\x2e\xb2\xab\x54\x90\xa5\x8c\xaf\xbf\xbb\xea\x6c\xdf\xe0\x6f\x9b\x79\x72\x5d\x42\x33\x89\x2f\x1e\x20\x91\x9d\x8d\xa4\x29\x95\x94\x33\x92\xff\x2a\x78\x09\x42\x52\xa8\xf0\x2d\x5a\x92\xbc\x82\xd6\xa1\x1c\x9a\x9b\xf4\x08\xe1\x0a\x20\x7d\x0f\xa2\xa2\x9c\x6d\x8d\x83\xf8\x95\x14\x94\x65\x6d\xfc\xd6\x5e\x12\x29\x41\x34\xae\xf8\x8f\x77\x77\x77\x3f\x7c\x7c\x7f\xf7\xdb\xbb\x1f\x7f\xf9\xf9\x15\x6e\x3d\x36\x9d\x23\x7e\xe0\x0b\x53\xb4\x41\xb5\xad\xdd\x57\x71\x97\x52\xaf\xba\xb5\x33\x52\xc0\x9e\xc5\x5e\xb6\x56\xfa\x3d\x89\x3f\x5f\xc7\x37\x1f\xe3\xf9\x65\x5f\xf9\xa0\xfa\xae\xb4\x3c\xe3\x82\xca\x55\xa1\x37\x67\x54\xa6\xc9\xf5\xe3\xfd\x9b\xf8\x66\x7e\x7f\x1d\xdf\xcc\x2f\xa6\xb3\xd9\x6b\xaf\x65\x12\xef\x1b\x1e\xbb\xff\x9a\x9a\xdf\xc6\x1f\xe2\xd6\xfa\xf4\xf7\xc5\x74\xa2\x05\x70\xfb\x4f\x2f\xa6\xdf\x4f\x66\xb3\xcb\xa1\xf5\xb2\x09\xb2\x67\x68\xbc\x2c\x9d\x29\x49\xf2\x89\x64\x70\xee\x8b\xd2\x17\x49\x65\xee\x03\xa4\x71\x62\x0a\x55\x22\x68\x29\xfd\xdd\x34\xe7\x25\x59\x65\x9b\x47\x84\x20\xeb\xfd\x87\x40\x25\x14\xaa\xbf\x23\x13\x42\x1b\x63\xd6\x82\x50\x26\x09\x65\x20\x6c\xb9\x15\xb2\xa3\x50\xc2\x23\x07\xe9\x91\x85\xf8\x9e\x25\xec\x95\xde\xfa\x72\x91\x11\x46\x3f\x13\x43\xd3\x47\xc7\x82\x82\xd0\xfc\xd8\x20\xb5\x38\x3a\x44\xb9\xe2\x6c\x7c\x63\x22\x4b\x40\x2c\xe0\xcf\x9a\x0a\x48\xf1\x2d\xba\x37\x3c\x00\x63\x13\x06\xb6\xb9\x85\x24\x05\xf0\x5a\xda\x30\x43\x99\x84\x0c\x84\x19\xe8\x02\x2a\x5e\x8b\x44\x83\xc4\xe9\x11\x57\x25\x24\x27\x2a\xd0\x91\x9b\x68\xc8\x4e\x36\x5f\xd1\x23\x4b\xf7\x2f\xa0\xf7\xb0\x10\x47\x29\xc8\xa0\xdc\xc3\x2c\xca\x2e\xfa\x36\xfe\xa0\xec\xa2\xbb\x7f\x1b\x53\x0c\xfc\x17\xc9\xeb\x80\x32\x58\x5d\x2c\x86\x48\xf0\x85\xa5\xac\xac\xe5\x4f\x75\x2e\x69\x99\x53\x4d\x95\xc6\x25\xd0\x6c\x86\x94\x76\x7e\xf4\xe3\x06\x96\xf4\x23\x5d\x07\xb4\x91\x79\xe4\x49\xea\x4e\xf9\x84\xd2\xc8\x15\x75\xe3\x15\xf6\x86\x82\x62\x49\x12\xeb\x56\x76\x32\x96\x25\xbc\x28\x08\x4b\x8f\x95\xc1\x16\x08\x26\x0e\xf8\x28\x37\x8a\x70\x3e\xba\xe1\x25\xcd\xed\x4c\x74\xcb\x06\x72\x4b\x47\xc8\x5a\x0e\x58\x51\xd8\xba\x7a\x2f\xa7\x94\xa0\x60\x39\x41\x63\x25\x05\xd9\xf8\x8f\x14\x7e\x04\x54\xb6\xe0\x3c\x07\xc2\x5c\xa5\xa5\xb0\x24\x75\xde\x6c\x57\x52\xd4\x30\xba\x9e\x02\x52\x4a\x7e\x5f\x97\xce\x5e\x06\xc1\xa1\x77\x74\x83\x42\x89\x66\x22\xcb\x5e\xdd\xe3\xd7\xd3\xc9\x6b\xe0\x93\x1f\xd5\xdf\x16\x92\xb6\x82\x8c\x76\x4b\x99\x3e\x5d\x46\x5b\xf8\x1a\x07\x55\x11\x36\xe7\x37\xed\x07\xcd\x0f\xf7\x33\xdf\x5f\x34\xdf\xfb\x4c\xcf\xd2\x25\x60\x75\xe1\x80\x70\xeb\xe3\xd3\x0b\x14\xb2\x68\xb4\x3b\x76\x3b\x9d\xfa\xf3\x92\xd3\xc7\x89\xce\xde\xc7\x27\x4b\x26\x06\xa2\xd3\xa9\x80\xfb\x61\x1e\xa3\x11\xea\x71\xce\xfd\x83\xb4\x96\xe7\x63\xcb\x21\x2b\x0a\x5b\x57\xef\xf5\xff\x94\xb1\xed\x31\x61\x54\x69\xa3\xd3\xec\x0a\x3e\x61\x92\x84\xd7\x4c\xfd\x79\xef\x4e\x11\x24\xc7\xf8\x4d\xe8\xa3\x98\xb4\x6f\xf9\x2e\x1f\x67\xb3\x8b\xe9\xd7\xb3\xa5\x7c\x61\x6d\xdb\x21\xce\x8b\x96\xf3\x29\xe9\x6b\x97\x97\x4f\xb0\x3e\x2d\xe7\xcf\xc7\x9e\x53\x1e\x7b\x5e\xb8\x54\x7d\xd9\x63\x58\xd1\x6c\x5f\xce\x53\xd8\xf9\x4d\xaf\xa9\xed\x46\xb1\x2e\x89\x5c\x05\x57\x11\x1c\xb5\xe0\xa9\x63\x6d\x1e\xfd\xc0\x82\xdb\xa1\x26\xfe\xb6\x00\xcd\xd2\xa1\xc1\xf1\x44\xf0\xff\xe2\x45\x72\xdb\x60\x6d\x40\x7b\x8f\xec\x44\x7c\x05\x52\x52\xa6\x7d\x42\x44\x67\xcc\x1f\x84\xf9\x0a\x12\x01\xf6\x03\x70\xa0\x1c\x07\xbc\x18\x7b\x46\x6c\x8d\x45\x50\x64\xfa\x7b\xf8\x0d\x01\x84\xe0\xe2\x19\x3f\x4a\x1b\x51\x34\x02\x41\x6e\xf4\xe0\xc4\xa6\x28\xf6\xef\x92\xda\x9a\x9f\x26\x18\x2e\x01\x28\xb1\xcc\x7a\xa7\x87\xb2\x5f\x0b\x38\x30\x60\x42\x24\x64\x5c\x98\x0f\x83\x7e\x82\x0c\xe5\x6e\x7b\x3b\xc6\xe4\xe7\x90\x60\xe7\xc4\x26\x05\x91\xc4\x3c\x52\xad\x2b\x09\xc5\xb1\x62\xe7\xa0\x43\x87\x02\xcb\x03\x8d\x6c\x29\x07\xf8\x8f\x94\x94\xe6\x64\xaa\x9a\xeb\x17\x8d\x06\x63\xca\x55\x1b\xed\xb2\x89\xe5\x12\x89\xf9\x9a\x86\xe1\x2b\x7c\x34\x5c\x50\xb3\x80\xb6\x78\xbd\xf0\xbd\x5b\x62\x5d\x9c\xf6\x8a\x57\xd4\xcc\xdd\xfc\x1b\x00\x00\xff\xff\xe9\x7d\x52\x0a\xc9\x26\x00\x00")

func schemaSeedManifestSchemaJsonBytes() ([]byte, error) {
	return bindataRead(
		_schemaSeedManifestSchemaJson,
		"schema/seed.manifest.schema.json",
	)
}

func schemaSeedManifestSchemaJson() (*asset, error) {
	bytes, err := schemaSeedManifestSchemaJsonBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "schema/seed.manifest.schema.json", size: 9929, mode: os.FileMode(420), modTime: time.Unix(1503330285, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _schemaSeedMetadataSchemaJson = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xcc\x57\xdf\x4f\xdb\x30\x10\x7e\xcf\x5f\x71\x32\x7d\x00\xd1\x52\x84\x78\x59\x5f\x10\xd2\xd0\xc4\xb4\x51\x44\x27\x1e\x06\xdd\x64\x9a\x4b\x6b\x94\xd8\xc1\x76\x85\x2a\xd6\xff\x7d\xb2\xf3\xa3\x76\x9a\x42\x18\x1d\xe4\x09\xe7\x7c\x77\xfe\xee\xbb\xb3\xfb\xf1\x14\x00\x90\x8e\x9a\xcc\x30\xa1\x64\x00\x64\xa6\x75\x3a\xe8\xf7\xef\x95\xe0\xbd\xcc\x7a\x20\xe4\xb4\x1f\x4a\x1a\xe9\xde\xe1\x71\x3f\xb3\xed\x90\xae\x89\x63\xe1\x33\x21\x53\x14\xc6\x64\xfe\x26\xa8\xe5\xe2\xc0\x7c\xe5\x81\x9a\xe9\x18\x4d\x6c\xb1\x99\x99\x43\x54\x13\xc9\x52\xcd\x04\x37\x9b\x43\x8e\x50\x38\x00\x55\x10\x62\xc4\x38\x86\x70\xb7\x80\x2f\x28\xbe\x8e\x86\x17\x79\xb6\x45\x6a\x93\x89\xbb\x7b\x9c\xe8\xcc\x96\x4a\x91\xa2\xd4\x0c\x15\x19\x80\x29\x12\x80\x28\xc4\xf0\x1a\xa5\xca\xd2\x67\x46\x27\x5c\x69\xc9\xf8\xd4\x86\x5b\x7b\x4a\xb5\x46\x69\x91\xfc\x1a\x9d\x9d\x7d\xfe\x7d\x7d\x76\x35\x3a\x1f\x5e\x74\x88\xf5\x58\x66\x8e\xab\x12\x6a\x52\x3a\x88\xac\x5d\xe2\xc3\x9c\x49\x34\xb4\xdd\xe4\xb6\xc2\xbb\xbb\xfa\x9e\x08\x21\x43\xc6\xa9\x46\x45\x72\xeb\xb8\x4c\x21\x38\x0e\x23\x2f\xfe\xa9\x5c\xb9\xcc\x5e\x0a\xc6\xb5\x93\x15\x80\xd0\x30\x64\x86\x5b\x1a\x5f\xba\xec\x44\x34\x56\xe8\x39\xd6\x90\xe7\x83\xad\x5a\x01\x08\xf2\x79\xe2\xc1\x2a\x77\x32\x24\x15\xfb\xd8\xfb\x5e\x76\xfd\x43\x5c\x06\x6a\xce\xea\x48\x34\x14\x90\x9d\xbe\x1d\x09\x5b\x94\xea\xa7\x42\xd9\x95\x7f\xd4\x32\xa8\x5b\x3b\x07\xd6\xd3\xf7\x7d\x1e\x6b\xd6\x1e\x0e\x1d\x38\xef\x47\xe4\xa9\x94\x74\xb1\x15\x36\xbf\x31\x8e\x23\xff\x7a\x7d\x24\x9b\x0e\x9c\xff\xcf\x66\xbc\xe1\xb0\xb7\x0c\x66\xcb\xf8\xac\x62\xda\x26\xa9\xc5\x53\x4a\xed\x30\x76\xab\xdb\x4c\x63\x52\x17\xf7\xea\x76\xf8\x0d\xf9\xe7\xf6\x5c\x8a\x78\x31\x15\xbc\x15\x6d\x29\xb0\xbc\xc7\x8b\x51\x73\xd2\xdb\x5e\xde\xf6\xd0\xe8\x01\x6a\xf3\x68\xd7\x76\xa1\xe1\x5c\x07\x6e\x3d\x85\xb4\xd1\x2c\xc1\x26\xb2\xa6\x49\x63\x36\x35\x85\x28\x4d\xa5\xae\x94\xb9\x49\x93\x41\x55\x97\xdd\x1c\xf6\x3e\x8d\x9f\x8e\x97\xbd\x6c\x71\xb4\x5a\xfc\x28\x16\x83\xb5\xc5\xee\xed\xed\x81\x5d\xef\xef\x9d\xec\xee\xde\xec\xf7\xc6\x6b\x2e\x7b\x7f\x7e\xee\x9d\x74\x48\xdd\xd4\x12\xe4\x61\xeb\xe1\x06\x15\xd8\x1b\x94\x67\xc6\x7d\xa5\x36\x7f\x14\x82\x3c\xcb\x7a\x06\x4f\x4f\x07\xb9\x46\x25\xce\x40\xae\x94\x77\x29\xcd\x9c\x61\xaa\x88\xfd\x53\x50\x8c\x4f\x63\x84\xd2\xb7\x5b\x1d\x3b\xff\x9e\x90\x84\xf1\xf3\xfc\x8a\x1c\x95\xc6\xe2\xd2\x6c\x92\xc7\x79\x2a\x3e\x4f\xee\x50\xd6\xf6\xb7\x99\xff\x9a\x2e\x5f\xdd\x81\x02\x94\x1d\x7f\xef\x3e\xf9\xc2\xea\x19\x2e\x38\xd8\x5a\x41\x44\x25\x1d\xea\x45\x3e\xd6\xdf\x8b\x86\x42\x79\xe9\x81\x74\x7e\x20\x9b\x21\xd4\x8f\x02\x84\x84\x44\x48\xac\x43\x4b\xe3\xf8\xb9\x7f\x58\x5e\x25\x41\x37\xb5\xc9\x1d\x85\x17\xde\x34\x53\x1e\x95\x57\xcd\xcb\x8b\xc4\x5c\xae\xea\x82\xc7\x19\x4a\x04\x3d\x43\x88\x98\x54\x1a\xf0\x61\x4e\x63\x65\x0d\x31\x55\xfa\xc3\xea\x3e\x7e\xa1\xee\xe2\xc7\xa1\x59\xd1\x19\x4b\x60\xa6\x60\x8b\x83\xe7\x70\xef\x8f\x9e\x79\x64\x82\x65\xf0\x37\x00\x00\xff\xff\xbb\x44\x9e\x3d\x8d\x10\x00\x00")

func schemaSeedMetadataSchemaJsonBytes() ([]byte, error) {
	return bindataRead(
		_schemaSeedMetadataSchemaJson,
		"schema/seed.metadata.schema.json",
	)
}

func schemaSeedMetadataSchemaJson() (*asset, error) {
	bytes, err := schemaSeedMetadataSchemaJsonBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "schema/seed.metadata.schema.json", size: 4237, mode: os.FileMode(420), modTime: time.Unix(1503330285, 0)}
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
	"schema/seed.manifest.schema.json": schemaSeedManifestSchemaJson,
	"schema/seed.metadata.schema.json": schemaSeedMetadataSchemaJson,
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
	"schema": &bintree{nil, map[string]*bintree{
		"seed.manifest.schema.json": &bintree{schemaSeedManifestSchemaJson, map[string]*bintree{}},
		"seed.metadata.schema.json": &bintree{schemaSeedMetadataSchemaJson, map[string]*bintree{}},
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

