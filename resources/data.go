// Code generated by go-bindata.
// sources:
// ../build/configs/default_config.yaml
// ../build/configs/testing_config.yaml
// DO NOT EDIT!

package resources

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

var _goCentrifugeBuildConfigsDefault_configYaml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x8c\x56\xc9\x76\xdb\x3a\x0f\xde\xeb\x29\x70\x9c\xcd\xff\x2f\x62\x6b\xb0\x64\x5b\xbb\x8c\x9d\x9c\x5c\xc7\x71\x9b\x26\x3b\x8a\x84\x2c\xd6\x12\xa9\x90\x94\x87\x3e\xfd\x3d\xa4\xa4\x34\x43\xd3\x5e\xaf\x78\x0c\xe0\xc3\xf4\x01\xd0\x11\x9c\x63\x4e\x9a\xd2\x00\xc3\x2d\x96\xb2\xae\x50\x18\x30\xa8\x8d\x40\x03\x64\x4d\xb8\xd0\x06\x14\x17\x1b\xcc\x0e\x1e\x45\x61\x14\xcf\x9b\x35\x5e\xa3\xd9\x49\xb5\x49\x41\x35\x5a\x73\x22\x0a\x5e\x96\x9e\x03\xe3\x02\xc1\x14\x08\xac\xc3\x15\xad\xa6\x06\x53\x10\x03\x67\x4f\x08\x50\x11\x2e\x8c\xc5\xf7\x7a\x95\xd4\x03\x38\x82\xb9\xa4\xa4\x74\x21\x70\xb1\x06\x2a\x85\x51\x84\x1a\x20\x8c\x29\xd4\x1a\x35\x08\x44\x06\x46\x42\x86\xa0\xd1\xc0\x8e\x9b\x02\x50\x6c\x61\x4b\x14\x27\x59\x89\x7a\xe8\x41\x6f\x6f\x21\x01\x38\x4b\x21\x8a\x22\xf7\x46\x53\xa0\xc2\xa6\xea\x32\xf8\xc4\x52\x98\x46\xd3\x56\x96\x49\x69\xb4\x51\xa4\x5e\x20\x2a\xdd\xda\x02\x1c\xc3\x60\xc4\xeb\xf1\x28\x08\x27\x43\x7f\xe8\x0f\x83\x91\xa1\xf5\x28\x9a\x86\x7e\x38\xe2\x75\xae\x47\x37\xd5\xea\x66\x9f\xed\x36\xcd\xc3\xfd\xfd\x79\xde\xfc\x5c\x65\xfb\x8b\x93\x25\xae\xae\xcf\xe6\xf2\xe7\xe1\x10\xc7\xd3\xed\x8d\x58\x7f\xdb\x2e\xae\x7e\xcc\xef\x37\x83\xbf\xc2\x46\x3d\xec\xb7\x3c\xb9\xb8\x4e\xaa\xcd\xe3\x1d\xfe\xb8\xfb\x72\x17\x3e\x2e\x9a\x20\xf9\x5e\xb3\x0f\xd1\xe6\xb3\x0c\x56\x51\x55\x90\x62\x71\x1a\xdf\x62\x2c\x82\x16\xb6\x2f\xd7\x49\x5f\xad\x3e\x09\xce\x50\x18\x6e\x0e\x97\x84\x1a\xa9\x0e\x29\x0c\x06\xaf\x24\x4b\x5c\x73\x6d\x5e\x88\x88\xa0\x85\x54\x4b\xac\xa5\xe6\xaf\xac\x6a\x72\xb0\x54\xf9\x27\x2b\xf9\x9a\x18\x2e\x85\x93\xb9\x06\x5e\x11\x2e\x7e\x4b\xa7\xae\xcf\x1e\x3c\x67\x4d\x1b\xe0\x11\x5c\x37\x15\x2a\x4e\xe1\xd3\x39\xc8\xdc\x31\xe8\x19\x57\x7e\x59\xb6\xcd\x8c\x83\xce\xea\xb4\xef\x18\x94\x5c\x1b\x6b\x29\x24\xc3\xb7\x64\xab\x95\xdc\x72\x27\x90\x0e\xfb\x59\x00\x7d\x78\xff\x81\x01\x51\x3c\x0c\xc3\x78\x18\xfa\xfe\x70\x1c\xbe\x66\x41\x10\x9e\x47\x5f\xa4\xbc\x9b\x73\x4e\x6f\xbe\xed\x56\xc5\xea\xf4\x3e\xd9\x7f\xa1\x0b\x39\xcf\x93\xe5\xcd\xfd\xe7\xcb\x7a\x97\x07\x6a\x12\xef\xe6\xfb\xf0\x61\x19\xd5\x67\x2c\x78\xcd\x85\xce\xc1\x34\x19\x86\x81\xff\x9e\x83\x9b\x87\xab\x93\xe9\x87\xc5\x47\xb5\xbd\x78\x38\x9d\xed\xd8\x46\x7e\xa5\x27\x27\xd5\xd9\xc3\xc7\x7a\x86\x87\xc3\xc3\xf8\xf6\x62\xba\xbe\x54\x51\xb1\xba\xfe\x3e\xe8\xea\x74\xd1\xb1\xbe\xaf\xa4\x2d\xf3\x31\x2c\xbb\xb9\x7e\x67\x2e\xc6\x9d\xf1\x9c\xd8\x12\x01\xc3\xba\x94\x07\x64\x70\x5b\x11\x65\xe0\xac\xa3\x9a\x86\x5c\x2a\x57\xd4\x35\xdf\xa2\x78\x51\xce\xb7\x74\x84\x77\xf9\xe8\xef\x67\x3e\x0b\x67\xe3\x78\x12\xe0\x24\x9a\x8e\xc3\x64\x36\x21\x49\x92\x4d\xc8\x6c\x46\xfc\x19\x63\x09\x9d\x44\x2c\x8a\x13\xf6\x07\xe6\xfa\xfb\x59\x92\xf8\xd4\x8f\x66\x2c\x0a\x82\x71\x1c\x91\xdc\x67\xf1\x94\xc6\x49\x92\x4c\xc2\x88\xcd\x68\x98\x93\x09\x4b\x90\xfe\x81\xe3\xfe\x7e\x92\x4f\xe3\x31\xcb\xc9\x6c\xea\x07\x21\x9b\xe4\x24\x8e\xe9\xd4\x8f\xb2\x8c\x84\x61\xe2\x67\x94\x21\x8e\xb3\x18\xd9\x5f\xa6\xe1\x08\x2c\x8d\x8e\x8d\x3c\xae\x11\x95\xad\x46\xce\xd7\x8d\x72\x1a\xda\xab\xc3\xba\xdd\x78\x2b\x5e\xa1\x6c\x0c\xec\x0a\x14\x20\x6b\x14\xdd\xe2\x13\x48\x9d\xa6\x25\xad\x05\xd0\x1e\xf4\x7f\x77\x26\x29\x0c\x22\x5f\x3b\x4f\x37\x0d\x36\xf8\xca\x85\xeb\x0c\xd1\x07\x41\x0b\x25\x85\x6c\xb4\x9d\x03\x8a\x5a\x73\xb1\xf6\x1e\xad\x41\x1b\x40\xbb\xb6\xb5\x6b\xa2\x68\xaa\x0c\x95\x9d\x24\x4b\x03\x54\x7a\x44\xa5\xd0\x76\x38\xbb\xa9\xda\xd9\x99\xc9\x10\x48\x59\x4a\x4a\x0c\x32\x20\x06\xb4\x21\xca\x34\xb5\x07\xd6\xfe\xae\x35\x4c\x21\x74\xe8\x97\x0a\x51\x43\x53\xc3\xd9\xe2\x2b\xd0\x03\x2d\x51\xb7\xa9\xb6\x0e\x80\x6b\xd8\x11\xee\xb6\xbd\x8d\x17\xb7\x28\x8c\x4d\xb5\x15\xdf\x11\xee\xb2\xbd\xba\x4d\x21\xb0\x89\x3e\x51\x59\xd7\x48\x79\xce\xe9\xcb\xa4\xbd\x9e\xc8\x6d\x6a\xb7\x58\xa2\xe5\xe8\xae\xe0\xb4\x78\x22\x39\x10\x4a\x65\x63\x57\x93\x84\x46\x63\xbf\x71\xa4\x2d\x42\xb7\x2a\x18\x70\xe1\xfe\xa4\x8d\x36\xb2\xea\x9c\x40\xce\x4b\xf4\xa0\xbf\x6e\x27\x2d\xcc\x35\xa9\x30\x85\x81\xbd\x68\x83\xa7\x1b\x66\x83\xe9\x81\x9f\xfc\xd2\x92\xdb\x8d\x68\x97\x14\xfc\x6f\x87\xa0\xf0\xb1\xe1\x0a\x61\xa7\x41\x2a\xe0\x35\xed\x0e\x9b\xbd\x63\xf6\x49\x89\xb1\x61\xbb\x92\xfc\xdf\x56\x57\x32\xfc\xba\x9c\xa7\xb0\xd3\xe9\x68\x64\x1b\x50\x16\x52\x9b\x74\x16\x8f\x93\xbe\x95\xee\xec\xae\x89\xcd\x84\x53\x1b\xec\x9a\xe8\x85\x7d\xa6\x10\xf8\xfd\xef\x8d\x72\xc9\x2b\x6e\x5a\xe5\xb9\x7d\xa6\x30\x9e\x04\x61\x34\x9d\xbe\xa0\xa8\x91\xae\x57\x2d\xb1\xc4\xaf\xbc\x8c\x22\x42\x13\x47\xd7\x3e\x03\xc6\xda\x33\x4d\x20\x2b\x25\xdd\x00\x11\xac\x4b\x04\x8c\xe2\xeb\x35\x2a\x64\x2d\xa1\x0d\xee\x4d\xdf\xe6\x96\xd4\x89\x6f\x59\xfd\x9e\x63\x85\x84\x81\x14\xe5\xc1\x0e\x4b\x4f\xf5\xfe\x5b\xa5\x0f\xe9\x17\xf4\x12\x09\x7b\x09\x1f\xc4\x1d\xfa\xb5\xed\xc3\xf3\xd8\x6b\x29\x4b\xa8\xc8\x1e\x14\x1a\xc5\xdb\x7b\xa1\x51\x30\x20\x2f\xd4\xe4\x16\x95\x07\x56\x71\xd9\xea\xa5\x10\x76\x35\xfd\x3d\x24\x17\x06\xd5\x96\x94\x0e\xf7\xd0\xd2\x9f\xd8\x00\x69\xa3\x94\xbb\x91\xcf\x2c\x0a\xa2\x21\x43\xb4\x47\xd4\x20\x35\xae\x4c\x3d\x80\xf5\x67\x77\x54\xd8\x65\x70\xce\xb5\xe3\x8a\x43\xd4\xb2\x7a\xc3\x35\x0d\x4c\x82\x90\x06\x74\x53\xd7\x52\x19\x30\x7b\x17\x11\xa9\xb9\xfd\x4c\xda\x2f\xa4\x2c\x4f\xa8\x5d\x0a\x17\xc2\x22\xb1\x14\x8c\x6a\xd0\xf3\xfe\x0d\x00\x00\xff\xff\x0a\x16\x18\x68\x1b\x0a\x00\x00")

func goCentrifugeBuildConfigsDefault_configYamlBytes() ([]byte, error) {
	return bindataRead(
		_goCentrifugeBuildConfigsDefault_configYaml,
		"go-centrifuge/build/configs/default_config.yaml",
	)
}

func goCentrifugeBuildConfigsDefault_configYaml() (*asset, error) {
	bytes, err := goCentrifugeBuildConfigsDefault_configYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "go-centrifuge/build/configs/default_config.yaml", size: 2587, mode: os.FileMode(420), modTime: time.Unix(1540892585, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _goCentrifugeBuildConfigsTesting_configYaml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x94\x92\xc9\x6e\xe3\x48\x0c\x86\xef\x7a\x0a\x81\x97\x5c\xbc\xd4\xbe\xbd\xc1\x20\x98\xd3\x0c\x90\x33\xab\xc8\x8a\x05\xdb\xb2\x5a\x4b\x12\x23\xc8\xbb\x37\xe4\x38\x9d\x6b\x1a\xba\x90\x04\x7f\xfe\xa4\xea\xe3\xf9\xc0\x23\x2f\xe7\xd4\xb4\x2d\x96\x72\x59\xfa\x79\x5a\xe3\xb6\x3d\x63\xd7\xa7\xf6\x16\xb6\xed\x91\xaf\xa9\x7d\x78\x07\x24\x1a\x79\x9a\x20\x41\x88\x59\x60\x70\x36\xe8\x62\x8c\x31\x58\x2a\x79\x99\x8d\xd3\x2c\x48\x17\x6b\x91\xa5\x91\x0a\x2d\x6c\xa0\x8c\xd7\x61\xbe\x40\x7a\x87\xd2\x0d\x07\x1e\x21\x01\xf2\xb4\x95\x2a\x6c\xcb\x3c\xae\x0d\xb7\xf2\xcc\x6f\x33\x24\x28\xde\xc7\x1a\xb4\x8f\xe4\xbd\xa0\xa8\x4a\x2d\x92\x88\x0c\x86\xaa\x25\x59\x14\x48\x25\x54\x85\x22\x2b\x94\x46\x48\xed\x05\x69\xa7\x45\xd5\xa1\x88\x12\xf0\xcf\xbc\x01\x47\x3c\x4f\xab\x6d\xf7\x02\x09\xb4\x2b\xd2\x05\xf6\x3a\xd7\x18\x44\x65\x6f\xb3\xf0\xca\xd7\x10\x05\x7a\x89\x04\x1f\x1b\x38\x52\x85\x04\xd3\x6d\x61\xb8\xa5\xdf\x43\xe8\x78\xe2\x1e\x92\x56\x1b\xe8\x21\x29\xa7\xa4\x31\x1b\x18\x20\xc9\x0d\x8c\x90\xc2\x06\x26\x3c\xad\x07\x10\xcb\xcc\xd2\xb1\x2e\x31\xc8\x68\x0c\x49\x2e\xa8\x72\xc8\xca\xb3\x61\xc7\x22\xdb\x5c\xb3\xd1\x99\x85\xf6\x0e\x2d\x85\x10\x62\x45\xe7\x23\xaa\x20\x95\x5a\x17\x39\x63\x59\x7f\x45\x91\x2a\xe4\x20\xad\xb5\x36\xa3\x64\x24\x5f\x90\xa3\x70\x82\x43\x30\x0a\x6b\xc1\xa0\xad\x23\xe1\x8c\xb5\x99\x22\x5a\x6f\x55\x46\x57\x4b\x11\x51\x71\x5d\x27\x75\x04\x09\x8c\x65\xe1\x04\xba\x2d\x29\xe4\xad\xd1\x39\x6c\xa3\x52\x75\x6b\x4c\x50\xd1\xc4\x48\xda\x13\x6c\xe0\x85\xc7\xa9\xbb\xac\x47\x7e\x3c\xdc\x1f\x7e\xc0\x69\x7a\xbd\x8c\x94\xda\x87\xaf\xd2\x9d\x81\xd4\xfe\x14\x81\xa6\xe9\x88\xfb\xb9\x9b\xaf\xff\x50\x6a\x41\xbc\x09\xf9\xfd\x41\xd3\xfc\x5a\x78\xe1\x15\xba\x7e\x39\x3f\x5d\xc6\x23\x8f\x53\x6a\x55\xd3\xb6\xaf\xb7\xe4\x09\xbb\xf9\xff\xee\xcc\xff\xfe\x97\x5a\xd9\x34\x47\xbe\xde\x08\x9d\xba\xe7\xbe\xeb\x9f\x3f\x61\x1d\x96\x7c\xea\xca\xe3\x4a\xe9\x6e\xb7\xdf\xed\xf6\x79\xe9\x4e\xb4\x1f\x79\xba\x2c\x63\xe1\x69\x7f\xef\x7e\xe4\xeb\x6e\x58\xf2\x6e\xe0\xf3\xa7\x6e\xec\x5e\x70\xe6\x9f\x09\x8f\xab\xf8\x26\xe4\xf9\x80\xcb\x7c\xf8\xa1\xf7\xbd\xfb\x2f\x8d\xbf\x54\x5f\xae\xbf\x03\x00\x00\xff\xff\xb0\x1c\xaf\x3f\xaa\x03\x00\x00")

func goCentrifugeBuildConfigsTesting_configYamlBytes() ([]byte, error) {
	return bindataRead(
		_goCentrifugeBuildConfigsTesting_configYaml,
		"go-centrifuge/build/configs/testing_config.yaml",
	)
}

func goCentrifugeBuildConfigsTesting_configYaml() (*asset, error) {
	bytes, err := goCentrifugeBuildConfigsTesting_configYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "go-centrifuge/build/configs/testing_config.yaml", size: 938, mode: os.FileMode(420), modTime: time.Unix(1540808141, 0)}
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
	"go-centrifuge/build/configs/default_config.yaml": goCentrifugeBuildConfigsDefault_configYaml,
	"go-centrifuge/build/configs/testing_config.yaml": goCentrifugeBuildConfigsTesting_configYaml,
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
	"go-centrifuge": &bintree{nil, map[string]*bintree{
		"build": &bintree{nil, map[string]*bintree{
			"configs": &bintree{nil, map[string]*bintree{
				"default_config.yaml": &bintree{goCentrifugeBuildConfigsDefault_configYaml, map[string]*bintree{}},
				"testing_config.yaml": &bintree{goCentrifugeBuildConfigsTesting_configYaml, map[string]*bintree{}},
			}},
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
