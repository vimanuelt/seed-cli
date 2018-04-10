// Code generated by go-bindata.
// sources:
// schema/1.0.0/seed.manifest.example.json
// schema/1.0.0/seed.manifest.schema.json
// schema/1.0.0/seed.metadata.schema.json
// DO NOT EDIT!

package assets

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

var _schema100SeedManifestExampleJson = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xb4\x56\xdf\x6f\x22\x37\x10\x7e\x2e\x7f\xc5\xc8\xba\xa7\x8a\x9f\xb9\x03\x55\x3c\xb5\x3d\x82\x42\x55\x92\xd3\x85\xe4\xa5\x8a\x90\xe3\x9d\x5d\x4c\xd6\xf6\xd6\xf6\x86\x50\xb4\xff\x7b\x65\x2f\xb0\xde\xb0\x34\xb4\x52\xf3\x14\xcf\x7c\x33\xf3\xf9\x9b\x19\x2f\xbb\x16\x00\x31\x88\xd1\x23\x6a\xc3\x95\x24\x63\x20\x83\x6e\xbf\xdb\x27\x6d\xe7\x59\xab\x67\x32\x06\x07\x02\x20\x92\x0a\x74\x7e\xb1\xed\x38\x7b\xbb\xb4\xae\xd5\x73\x63\x2c\x00\xc9\x28\x7b\xa1\x09\x9e\x73\x5b\x6e\x53\x9f\x70\xbe\x85\x98\x6b\x63\x21\x48\x1b\xa1\x61\x9a\x67\x76\x1f\xf8\x1d\x69\x64\x80\x4a\xb8\x99\x4c\x87\x10\xf3\x14\x81\xca\x08\x54\x6e\xb3\xdc\x1a\xb0\x1b\x05\x8b\xd9\x74\x0a\x5c\xd0\x04\x4d\x1b\x28\x7c\xbd\x7f\xf4\x10\x41\x25\x8f\xd1\x58\x60\x4a\x5a\xca\x25\x97\x09\x30\x4c\xd3\x25\x53\xb9\xb4\x47\x2e\x34\x31\x64\x0c\x7f\xf8\x13\x00\x59\x45\xf1\x70\xef\xf3\x4c\xe3\xb8\x3a\x31\xf3\x5a\x1d\x7c\x41\xc8\xb4\x62\x68\x0c\x97\x09\xf1\x8e\xa7\x7d\x5a\x41\xb9\x2f\x8a\xfa\x28\x63\x20\xe4\x6f\x6a\x25\x61\xa2\xb0\xca\xa6\x74\x42\x25\xff\x8b\x1e\xae\x7d\xdd\x61\x4a\x67\x95\x1f\x05\xe5\xa9\x73\xac\x23\x85\x3f\xe3\x1b\x15\x59\x8a\x5d\xa6\x44\x05\xc9\xb5\x07\xac\xac\xcd\xc6\xbd\xde\x66\xb3\xe9\x36\xc2\xb2\x95\x92\x9e\xc4\x68\x34\xea\x0c\x87\xc3\xce\x97\xcf\x57\x83\x92\x7c\x71\xec\x8f\x40\x95\x5b\x32\x86\xcf\xa3\x7e\x7f\x6f\xe4\xd2\xa2\x8e\x29\xc3\xf0\x42\x4c\x09\x41\x65\xe4\xd2\x7d\xda\xcd\x6e\xbf\x3d\x2c\x96\xd3\xd9\xef\xd7\x05\x7c\xda\xdd\x3d\x2c\xdc\x71\x32\xfb\x5e\x04\xa2\x49\xd7\xb5\x20\x03\x00\x71\x3d\x0d\x5b\xe0\xfe\x76\xc1\xff\x81\x6e\x55\x05\xd2\x6e\xfd\xb0\x1f\x35\x6d\x39\x75\x37\xb7\x3a\xc7\x76\x3d\x4e\xe3\x9f\x39\xd7\x18\x35\x7b\x05\x46\x9c\x2e\xb6\xd9\x49\xf5\xaa\xbf\xbd\xb7\x8e\x1b\x88\x8e\x3f\x90\x1a\xe4\x29\x38\x15\xad\xf7\xd6\xa2\xea\x6c\x39\xa9\xff\xf9\xce\x65\xfc\xd2\x45\x2c\xdd\x38\x1a\x72\xee\x1a\x0e\x5d\xb2\xae\x8d\xed\x01\x96\xa7\x96\x67\x7e\xef\x1a\xb4\xc8\xa8\xb5\xa8\xe5\xbe\xa0\x2b\xf6\x63\xd7\xf2\x38\xbc\x72\xd1\xfe\x77\x64\xc3\x6d\x69\xa2\x6a\xf1\xcd\xf6\x1a\x40\x4d\x54\x1c\xac\x59\xef\x2a\x98\xac\x8d\xdf\x9d\x0b\x34\x3d\x79\x06\x8e\x88\x17\xdc\x1e\x00\x5f\x9b\xfc\xf6\xa0\xb3\xb4\x98\xa0\x3e\x43\xea\x64\x08\x84\xcb\x55\x6f\x78\x48\xed\x48\x6c\x7e\xf7\x70\xbb\x58\x7e\xfb\x65\x71\x53\x2b\xec\x34\x59\x39\x7f\xcf\xae\xb0\xb7\x7f\xcf\x50\xf7\xbc\xb9\x06\x14\x2a\xf2\x89\xb4\xaa\xa8\x1d\x88\x1d\xb5\x22\x06\xad\xe5\x32\xb9\x80\xd0\xe4\xd7\xe5\xcd\xdd\xfd\xa2\x5e\xc4\x20\xd3\xe8\x9e\x87\x98\xa6\x06\x4f\xeb\xd4\x1e\x13\x8d\x46\xe5\x9a\x61\xb8\x01\xc4\x30\x9a\x52\xfd\x71\x79\x96\xe5\xf5\xd2\xaf\x34\xcd\x9d\x67\x50\x55\x6d\xff\x73\x0a\x81\xe2\x4c\x8a\xab\x9f\x2e\x4e\x62\x56\x54\x63\x34\x3f\x97\xaa\x7f\x71\xa2\x88\x9b\x97\x33\x74\xfa\x35\xb3\x7f\x29\xe7\xe5\xce\x72\xff\x15\xf9\xf2\x81\xd0\xa8\xb5\xd2\x61\x4b\x83\xf7\x86\x95\x63\x31\x08\xb6\xe5\xf8\x11\xbe\x76\x71\x70\xeb\x08\x06\xee\x77\x9f\xe1\x12\x34\x09\x8c\x01\x96\x51\x8b\x89\xd2\x7e\x73\x22\x6a\x29\x79\x3f\xff\xa7\x4c\xae\xfe\x7f\x26\xee\x37\x45\x2b\x14\xcc\xc9\x55\xb4\x8a\xd6\xdf\x01\x00\x00\xff\xff\xc8\x49\xb1\x90\xfb\x08\x00\x00")

func schema100SeedManifestExampleJsonBytes() ([]byte, error) {
	return bindataRead(
		_schema100SeedManifestExampleJson,
		"schema/1.0.0/seed.manifest.example.json",
	)
}

func schema100SeedManifestExampleJson() (*asset, error) {
	bytes, err := schema100SeedManifestExampleJsonBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "schema/1.0.0/seed.manifest.example.json", size: 2299, mode: os.FileMode(420), modTime: time.Unix(1510606562, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _schema100SeedManifestSchemaJson = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xec\x5a\xcb\x6e\xeb\x36\x10\xdd\xeb\x2b\x08\xf5\x2e\xec\xa4\x8a\x1d\xa0\x9b\x78\x53\xe4\x03\x0a\x74\x51\x74\x11\xdb\x2d\xc6\xd2\xd8\xa6\x23\x91\x2a\x45\xb5\x70\x1a\xff\x7b\xa1\x87\x13\x99\xe2\x43\xf2\x23\x4d\x1b\x5f\xe0\xe2\x1a\xa3\xe1\xcc\x21\x79\xe6\x70\xae\xc4\xbf\x3d\x42\xfc\x6f\x59\xb8\xc6\x04\xfc\x09\xf1\xd7\x52\xa6\x93\xd1\x68\x93\x71\x16\x54\xd6\x3b\x2e\x56\xa3\x48\xc0\x52\x06\xe3\x1f\x46\x95\xed\x3b\xff\xfb\x62\x9c\xdc\xa6\x58\x0c\xe2\x8b\x0d\x86\xb2\xb2\x41\x14\x51\x49\x39\x83\xf8\x67\xc1\x53\x14\x92\x62\xe6\x4f\xc8\x12\xe2\x0c\x4b\x87\xb4\x69\x2e\xd2\x13\xe2\x67\x88\xd1\xaf\x28\x32\xca\xd9\x9b\xb1\x11\x3f\x93\x82\xb2\x55\x19\xbf\xb4\xa7\x20\x25\x8a\xc2\xd5\xff\xed\x7e\x36\xbb\x1b\x17\x7f\xbf\xf9\xe5\xe3\x5d\xe5\xe5\x6f\xf8\x42\x17\xaa\x01\xb5\xb4\xbb\xe0\x56\xf9\xda\x90\x4b\x3b\x83\x04\x0f\x2c\x66\xcc\x2d\xdc\x53\x08\x5e\xc6\xc1\xc3\xef\xc1\xfc\xb6\x46\xde\x40\xbf\x9f\x41\x7b\x4d\x7a\xe5\x18\x8c\x5f\xa7\xf7\xc1\xc3\x7c\x3a\x0e\x1e\xe6\x37\xc3\xd9\xec\xce\x69\x19\x04\x87\x86\xd7\xea\x9f\x02\xed\x63\xf0\x14\x94\xd6\xfd\xef\x9b\xe1\xa0\x15\xc0\xee\x3f\xbc\x19\xfe\x38\x98\xcd\x6e\x9b\xd6\xdb\x22\xc8\x81\xa1\xf0\x32\xac\x49\x0a\xe1\x33\xac\xf0\xba\x2e\xca\xba\x48\x2a\x63\x17\x15\xb5\x03\x23\xcc\x42\x41\x53\xe9\x5e\x4d\x7d\x5e\x58\x65\xa6\x71\x20\x04\x6c\x0f\x37\x81\x4a\x4c\x54\x7f\x4b\x26\x42\x76\xda\xac\x09\x50\x26\x81\x32\x14\xa6\xdc\x4a\x99\x93\xae\xa5\x4e\x2c\xe5\x4e\x0c\x25\xef\x98\xc2\x01\xf4\xd2\x97\x8b\x15\x30\xfa\x02\x9a\x45\xef\x1d\x0b\x13\xa0\xf1\xa9\x41\x72\x71\x72\x88\x74\xcd\x59\xff\x85\xf1\x0c\x01\x7d\x81\x7f\xe4\x54\x60\xe4\x4f\xc8\x54\xb3\x01\xda\x45\x68\xd8\xe6\x86\x22\x49\x90\xe7\xd2\xc4\x19\xca\x24\xae\x50\xe8\x89\x2e\x30\xe3\xb9\x08\x5b\x94\xb8\x3c\xe3\xb2\x10\x62\x50\x89\x4e\xec\x85\x46\xcc\xc5\xe6\x02\xdd\x13\xba\x7b\x02\xb5\x87\xa1\x70\x14\x40\x1a\xe5\x6e\x66\x51\xce\xcf\xc7\xe0\x49\x39\x3f\xdf\xff\xec\x74\x31\xfc\x3f\x21\xce\x3b\xc0\x60\x79\xb2\x68\x32\xc1\x15\x96\xb2\x34\x97\x3f\xe5\xb1\xa4\x69\x4c\x5b\xaa\xd4\x2f\x41\xcb\xa6\x49\x69\xae\x8f\xfa\xb9\xa6\x4a\xea\x27\xd5\x0a\xb4\x9e\xcc\x3d\x47\x52\x7b\xca\x3d\x4b\x3d\x5b\xd4\x9d\x53\xd8\x8b\x12\x14\x4b\x08\x8d\x47\xd9\xc5\xaa\x2c\xe4\x49\x02\x2c\x3a\x55\x06\x4b\x22\xe8\x6a\xc0\x55\x72\xbd\x0a\xce\x55\x6e\xfe\x92\xc6\xe6\x4a\xb4\xcb\x06\xb1\x4b\x47\x97\xb9\x1c\x31\xa3\x6e\xf3\xaa\xbd\xac\x52\x42\x3a\xcb\x09\xe9\x2b\x29\xc4\x54\xff\x44\xa9\x8f\x0e\xc8\x16\x9c\xc7\x08\xcc\x06\x2d\xc2\x25\xe4\x71\x71\x5c\x49\x91\x63\x6f\x3c\x09\x46\x14\x7e\xd9\xa6\xd6\xb5\xec\x44\x87\xda\xd1\x4e\x0a\x25\x9a\xae\x58\x0e\x70\xf7\x9f\x4f\x25\xaf\x1d\x77\xbe\xd7\xfa\x96\x94\xec\x0d\x28\x05\x21\x29\xe8\x3a\xa7\x8b\xe2\xd1\xda\x0d\x28\x5d\xe7\x04\x79\x2b\x27\xed\x43\xf5\x50\xd0\xe7\xd7\x9d\x4f\x9b\x4c\xdb\xda\x92\xab\xfe\xa8\x19\xff\xb7\xfa\x53\x67\x3a\xcb\x2a\x21\xcb\x13\x0b\x85\x4b\x1f\x97\x7e\x91\x2e\x93\x26\xef\xff\x0d\xb0\x3a\xd5\xfd\x9b\xd5\xc7\xca\xce\xda\xc7\x25\x93\xba\x0a\x24\x97\x53\x01\xfb\x66\x9e\xa2\x11\x6a\x7b\x69\xed\xa2\x78\x2e\xaf\x6d\xd4\x31\x33\xea\x36\xaf\xda\xeb\x73\xca\xd8\x5b\xdb\xd2\x0b\xda\x11\x87\xf7\x1e\xf0\x05\x93\x7c\xba\x96\xe5\x5f\x3a\x23\x3e\x58\xac\xf6\x7b\x7b\xed\x69\xfe\xeb\x62\xf0\x8c\xdb\xcb\x56\xe8\xb5\x49\xb9\x64\x93\xf2\xc5\x75\xe8\x63\x9b\xa6\x84\xe7\xcc\xde\x33\x7d\x95\xf7\xc4\xe4\x3c\xef\x8a\x53\x90\xeb\xce\x48\x3a\x47\x4d\x78\x64\x99\x9f\x43\x43\x7c\xc1\xcd\x74\x13\x7f\x19\xc8\x66\x20\xf8\x7b\xcd\x14\x61\x8f\x3e\x16\x8f\x7f\x15\x5d\x2e\x70\xeb\x41\xeb\x4d\xb4\x95\xf5\x19\x4a\x49\x59\xeb\x23\x24\xf9\x8a\xbc\x3f\x03\xe7\x33\x0c\x05\xaa\x5f\xc8\x34\x38\xec\x92\xdc\xa1\x4f\x3d\x23\xb7\xfa\x32\xc8\xd3\xfd\x6e\x7e\x85\x40\x21\xb8\x38\xe3\x67\x6d\x2d\x8b\x7a\x30\xc8\xce\x1e\x3f\x34\x29\x8a\xf9\xcb\x66\x6b\xce\xfb\x01\x9a\x6b\x04\x4a\x2c\xbd\xde\xb5\x43\x99\x2f\x16\x1c\x19\x30\x04\x89\x2b\x2e\xf4\x0d\xa1\xbb\x40\x9a\x72\xb7\xe1\x0b\x9d\x87\x45\x7c\x0d\x43\x8a\xb0\x20\xe1\x54\x11\xb3\xd0\xbc\xda\x5d\xc3\x46\x79\xa6\x94\x0d\x5e\x7b\x4a\x4a\x7d\x32\x55\xa5\x9b\x97\x8f\xcc\xd7\x6f\x5a\x17\x50\x0c\x17\x4b\xf4\x57\x37\x34\x5f\xe6\xbd\xe6\x54\x0a\xe8\x25\xec\x36\xe4\x83\x0b\x63\x55\x9c\x72\x87\xbc\x62\xec\xce\xfb\x27\x00\x00\xff\xff\x0b\x94\xb7\x98\xd5\x26\x00\x00")

func schema100SeedManifestSchemaJsonBytes() ([]byte, error) {
	return bindataRead(
		_schema100SeedManifestSchemaJson,
		"schema/1.0.0/seed.manifest.schema.json",
	)
}

func schema100SeedManifestSchemaJson() (*asset, error) {
	bytes, err := schema100SeedManifestSchemaJsonBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "schema/1.0.0/seed.manifest.schema.json", size: 9941, mode: os.FileMode(420), modTime: time.Unix(1523070963, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _schema100SeedMetadataSchemaJson = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xcc\x57\xc1\x6e\xdb\x38\x10\xbd\xfb\x2b\x06\x4c\x8e\x4e\xb4\xd8\x0d\x76\xb1\xbe\x05\x05\x52\xb4\x68\x9b\xa0\x39\x06\x3e\xd0\xf2\xc8\x66\x40\x91\x0e\x49\x23\x31\x0a\xfd\x7b\x41\x5a\x94\x48\x89\x76\xe4\xc4\x0d\x7c\x93\x46\x33\x9c\x37\xef\x0d\xa9\xe1\xaf\x11\x00\x39\xd7\xf9\x12\x4b\x4a\x26\x40\x96\xc6\xac\x26\x59\xf6\xa8\xa5\xb8\xd8\x5a\x2f\xa5\x5a\x64\x73\x45\x0b\x73\xf1\xd7\x55\xb6\xb5\x9d\x91\xb1\x8d\x63\x73\x1f\xa2\x27\x59\x66\xa4\xe4\xfa\x92\xa1\x29\x5c\xc8\xd2\x94\x3c\x53\x45\xfe\xdf\xff\x57\xff\xd6\xfe\x86\x19\x8e\x36\xe4\x33\x4a\xf8\x7a\x7f\xfb\x03\xe4\xec\x11\x73\xb3\xfd\x3a\x47\x9d\x2b\xb6\x32\x4c\x0a\xeb\x73\xef\x32\x41\x21\x15\x50\x48\x06\x98\xcd\xca\xad\x16\xda\x14\x3e\xad\x99\x42\x0b\xec\xa1\xf6\x80\xa9\xfb\xb2\x52\x72\x85\xca\x30\xd4\x64\x02\xb6\x6c\x00\x32\x9b\xc9\x17\xfb\x06\xe4\x5c\x61\xb1\xa7\xfe\x05\x4a\x6b\xca\x6c\xc0\xa5\x7d\x3a\x23\x50\x8d\x00\x2a\xb7\xb4\x14\x78\x6b\xa3\x1f\xdc\xaa\xc1\x6a\x67\xd9\x1c\x0b\x26\x98\xad\x49\xdb\x45\x4a\x34\x6a\x43\xb6\x61\x03\x3c\x3f\x49\xce\x31\x77\x8c\xbc\x1a\x53\x20\x35\x6b\x85\x83\x1d\xa3\xb5\x47\x50\xb3\x14\x38\xb6\x34\x35\xc0\xbd\x25\x4d\xbe\xb3\x87\x02\xd4\x36\xef\x3d\x6e\xdf\x73\x29\xd5\x9c\x09\x6a\x50\x93\xda\x3a\x6d\x96\x88\xe9\x74\xb5\x34\x4f\x61\x17\xdd\x49\x26\x4c\xb0\x2a\x00\xa1\xf3\xb9\x03\x4f\xf9\x5d\x28\x77\x41\xb9\xc6\xc8\x31\xd1\x0d\x31\xd8\xae\x15\x80\xa0\x58\x97\x11\xac\xe6\xcb\x16\x49\xc7\x3e\x8d\xde\xab\x71\x9c\x24\x64\x20\x91\x2b\x2d\xdd\x4a\x6a\xf7\x14\xa7\xaa\x46\xa9\xe7\x20\x61\x9a\xbe\xef\x6b\x6e\xd8\xe9\x70\x18\xc0\xf9\x38\x22\xaf\x95\xa2\x9b\xa3\xb0\xf9\x8d\x09\xbc\x37\x8a\x89\xc5\x49\xb0\x19\xc0\xf9\xf3\x6c\xf2\x1d\xc9\xde\xd3\x98\x27\xc6\x67\x17\xd3\x31\x49\xf5\x47\x29\x75\xcd\x38\xee\x7e\x66\x06\xcb\x54\xdc\xc1\x72\xc4\x82\xbc\x59\x9e\x3b\xc9\x37\x0b\x29\x4e\x42\x16\x8f\xe5\x23\x4e\x8c\x44\xa6\xf7\x9d\xbc\xa7\x43\x63\x04\xe8\x94\x5b\x3b\xa9\xc2\xc0\xbe\x1e\x85\xf5\xd4\x75\x90\xc4\xa4\x15\x0e\x39\xc1\xc0\xda\x75\x6b\xc6\x95\xce\xdc\x7a\x0d\x79\xe3\x04\xb2\x00\x9f\xa1\x1e\x5d\x75\x7a\x54\x6a\x90\x58\x45\x83\x59\x68\x87\xce\x8d\xc6\xad\xaa\x49\x90\x30\x0d\x15\x0b\x73\x4c\x3a\xbd\xb9\x53\xa8\x56\xa4\x01\x73\x6d\x8f\xed\x2a\x62\xdb\xcf\xa8\x29\x8a\x6f\xea\x6f\xbb\x79\x6d\x2e\x01\xf5\x32\xb0\x6f\xf6\x0c\xa6\xd6\x71\x44\xe3\x1b\xc9\xf5\xf0\xd2\x8c\x6e\xba\x7c\xf6\x67\x58\xd8\xce\xe4\x9e\x68\xb1\xe6\x9c\x74\xb7\xd3\x41\x14\x87\x5b\xb3\xda\x73\x38\x34\x39\x1f\xda\x69\xdd\xe7\x9f\x76\x22\xdd\x8d\x2e\x8a\xd0\xf5\x3f\xd8\x46\x94\x33\x54\x2e\x66\x9f\xb8\xaf\xec\xa4\x9b\x9e\xd7\x01\x82\xe7\x89\xa8\x58\xf4\xda\xf3\xbd\x32\xef\xde\x42\x4d\x82\x23\x6e\xa0\xf6\xea\xb6\x7f\xff\xd8\x5a\xfc\x0d\x20\xaa\xa4\xc7\x9c\x66\x62\xc1\x11\x1a\xef\x71\xaf\xea\x1e\x50\x52\x32\xf1\xa5\x46\xfa\x77\x68\xa6\x2f\xde\xfc\x4f\x60\xf6\x45\x85\x0d\x9e\xfe\x0f\xf9\xc6\x09\x4f\xe5\xf1\x07\x46\x35\xcf\xd3\x00\x7e\xfb\xa3\xf5\xc5\xb9\x7f\xec\xa8\x93\x89\xc4\x17\x85\xbd\xa4\x0b\x70\x94\xda\x33\xdf\x47\xe9\x41\xc4\xa7\xfe\x82\x83\x2f\x80\x55\x0f\x72\x30\xfe\x0d\xc5\x6b\x9e\x25\x48\x05\xa5\x54\x98\xc6\x4e\x39\xef\x1d\x67\x1d\x05\x0e\xbc\x68\xed\x13\x33\xec\xc4\xb4\x92\xc9\xa2\xa9\xfa\x79\x48\xd1\x85\x5c\xab\xb6\x5a\x78\x5e\xa2\x42\x30\x4b\x84\x82\x29\x6d\x00\x9f\xd6\x94\x6b\x67\xe0\x54\x9b\x93\x60\xe3\x6a\x20\x1b\x7e\x4c\x1a\x4a\xc5\x96\x3d\xb0\x3d\x73\xf4\x96\x0d\x94\x49\x9e\x6e\xd5\xa8\xfa\x1d\x00\x00\xff\xff\x08\x83\x9f\x03\x85\x14\x00\x00")

func schema100SeedMetadataSchemaJsonBytes() ([]byte, error) {
	return bindataRead(
		_schema100SeedMetadataSchemaJson,
		"schema/1.0.0/seed.metadata.schema.json",
	)
}

func schema100SeedMetadataSchemaJson() (*asset, error) {
	bytes, err := schema100SeedMetadataSchemaJsonBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "schema/1.0.0/seed.metadata.schema.json", size: 5253, mode: os.FileMode(420), modTime: time.Unix(1523054135, 0)}
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
	"schema/1.0.0/seed.manifest.example.json": schema100SeedManifestExampleJson,
	"schema/1.0.0/seed.manifest.schema.json": schema100SeedManifestSchemaJson,
	"schema/1.0.0/seed.metadata.schema.json": schema100SeedMetadataSchemaJson,
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
		"1.0.0": &bintree{nil, map[string]*bintree{
			"seed.manifest.example.json": &bintree{schema100SeedManifestExampleJson, map[string]*bintree{}},
			"seed.manifest.schema.json": &bintree{schema100SeedManifestSchemaJson, map[string]*bintree{}},
			"seed.metadata.schema.json": &bintree{schema100SeedMetadataSchemaJson, map[string]*bintree{}},
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

