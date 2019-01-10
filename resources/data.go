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

var _goCentrifugeBuildConfigsDefault_configYaml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xb4\x58\x59\x73\xdb\xba\x15\x7e\xd7\xaf\x38\x23\xbf\xb4\x33\xa5\xcc\x7d\xd1\xcc\x9d\x8e\xbc\x25\xb9\x71\x5c\xd9\x96\xe3\x1b\xbf\x34\x20\x70\x28\x22\xa6\x00\x06\x00\xb5\xe4\xd7\x77\x00\x52\x8e\x1d\x2f\xb7\xcb\x34\x2f\xa6\x00\x9c\xfd\x3b\x1f\x0e\x72\x00\x27\x58\x91\xae\x31\xc0\x70\x8d\x8d\x6c\x57\x28\x0c\x18\xd4\x46\xa0\x01\xb2\x24\x5c\x68\x03\x8a\x8b\x7b\x2c\x77\x23\x8a\xc2\x28\x5e\x75\x4b\xbc\x40\xb3\x91\xea\x7e\x0a\xaa\xd3\x9a\x13\x51\xf3\xa6\x19\x39\x65\x5c\x20\x98\x1a\x81\x0d\x7a\x45\x7f\x52\x83\xa9\x89\x81\xe3\x07\x0d\xb0\x22\x5c\x18\xab\x7f\xb4\x3f\x32\x1d\x01\x1c\xc0\xb9\xa4\xa4\x71\x2e\x70\xb1\x04\x2a\x85\x51\x84\x1a\x20\x8c\x29\xd4\x1a\x35\x08\x44\x06\x46\x42\x89\xa0\xd1\xc0\x86\x9b\x1a\x50\xac\x61\x4d\x14\x27\x65\x83\x7a\x32\x82\xbd\xbc\x55\x09\xc0\xd9\x14\xa2\x28\x72\xdf\x68\x6a\x54\xd8\xad\x86\x08\x3e\xb0\x29\xe4\x51\xde\xef\x95\x52\x1a\x6d\x14\x69\xe7\x88\x4a\xf7\xb2\x1e\x8c\x0f\x79\x1b\x1f\x06\x61\x36\xf1\x27\xfe\x24\x38\x34\xb4\x3d\x8c\xf2\xd0\x0f\x0f\x79\x5b\xe9\xc3\xcb\xd5\xe2\x72\x5b\x6e\xee\xbb\xbb\x2f\x5f\x4e\xaa\xee\xc7\xa2\xdc\x9e\xce\xae\x70\x71\x71\x7c\x2e\x7f\xec\x76\x49\x92\xaf\x2f\xc5\xf2\xf3\x7a\xfe\xe9\xdb\xf9\x97\xfb\xf1\x9f\x28\x8d\xf6\x4a\x3f\x57\xe9\xe9\x45\xba\xba\xff\x7e\x8b\xdf\x6e\x3f\xde\x86\xdf\xe7\x5d\x90\xfe\xd1\xb2\x77\xd1\xfd\xef\x32\x58\x44\xab\x9a\xd4\xf3\xa3\xe4\x1a\x13\x11\xf4\x4a\xf7\xa9\x9a\xed\x33\xd5\x07\x60\xc3\x47\x61\xb8\xd9\x9d\x11\x6a\xa4\xda\x4d\x61\x3c\xfe\x65\xe7\x0a\x97\x5c\x9b\x27\x5b\x44\xd0\x5a\xaa\x2b\x6c\xa5\xe6\xbf\x48\xb5\x64\x67\x61\xf2\x8f\xb2\xe1\x4b\x62\xb8\x14\x6e\xcf\x15\xef\x13\xe1\xe2\x45\x28\x0d\x35\x1e\xc1\x63\xc4\xf4\x0e\x1e\xc0\x45\xb7\x42\xc5\x29\x7c\x38\x01\x59\x39\xf4\x3c\xc2\xc9\x4f\xc9\xbe\x90\x49\x30\x48\x1d\xed\xab\x05\x0d\xd7\xc6\x4a\x0a\xc9\xf0\x39\xd0\x5a\x25\xd7\xdc\x6d\x48\xa7\xfb\x91\x03\x7b\xf7\xfe\xb4\xfa\x51\x32\x09\xc3\x64\x12\xfa\xfe\x24\x0e\x7f\x45\x40\x10\x9e\x44\x1f\xa5\xbc\x3d\xe7\x9c\x5e\x7e\xde\x2c\xea\xc5\xd1\x97\x74\xfb\x91\xce\xe5\x79\x95\x5e\x5d\x7e\xf9\xfd\xac\xdd\x54\x81\xca\x92\xcd\xf9\x36\xbc\xbb\x8a\xda\x63\x16\x8c\x5f\x52\x9f\xa7\x93\x30\xf0\x5f\x53\x7f\x79\xf7\x69\x96\xbf\x9b\xbf\x57\xeb\xd3\xbb\xa3\x62\xc3\xee\xe5\x0d\x9d\xcd\x56\xc7\x77\xef\xdb\x02\x77\xbb\xbb\xf8\xfa\x34\x5f\x9e\xa9\xa8\x5e\x5c\xfc\x31\x1e\x72\x74\x3a\xa0\x7d\x9f\x45\x9b\x62\x0f\xae\x86\x7e\x7e\xa5\x1f\xe2\x41\xf8\x9c\xd8\xf4\x00\xc3\xb6\x91\x3b\x64\x70\xbd\x22\xca\xc0\xf1\x00\x33\x0d\x95\x54\x2e\xa1\x4b\xbe\x46\xf1\x24\x95\xcf\xa1\x08\xaf\x62\xd1\xdf\x16\x3e\x0b\x8b\x38\xc9\x02\xcc\xa2\x3c\x0e\xd3\x22\x23\x69\x5a\x66\xa4\x28\x88\x5f\x30\x96\xd2\x2c\x62\x51\x92\xb2\x37\x50\xeb\x6f\x8b\x34\xf5\xa9\x1f\x15\x2c\x0a\x82\x38\x89\x48\xe5\xb3\x24\xa7\x49\x9a\xa6\x59\x18\xb1\x82\x86\x15\xc9\x58\x8a\xf4\x0d\x7c\xfb\xdb\xac\xca\x93\x98\x55\xa4\xc8\xfd\x20\x64\x59\x45\x92\x84\xe6\x7e\x54\x96\x24\x0c\x53\xbf\xa4\x0c\x31\x2e\x13\x64\x6f\x75\x82\xbf\x65\xa5\x9f\xe4\xc1\xac\x88\xc2\x3c\x4d\xe3\x3c\x49\xa2\x30\x9f\xb1\x93\xd2\x3f\x0d\x93\x20\xc8\xe3\x34\xf6\xab\x02\x93\x13\xd7\x33\x25\x2a\x41\x9a\x1a\xf9\xb2\x36\x03\xe8\x0e\x0e\x0e\x86\x0a\x7c\x94\x6b\x22\xe0\x6c\x76\x39\xfc\xf6\xe0\xd6\xb2\x1d\x17\x55\xa7\x08\xec\x64\x07\x4b\x4b\xd3\x02\x50\x29\xa9\x2c\x9c\x16\x35\xd7\xa0\xf0\x7b\x67\x2b\xc7\x35\x08\x69\x40\x77\x6d\x2b\x95\x41\x06\x25\x52\xd2\x69\xb4\x92\xca\x75\x8b\x3d\xa2\x3a\x21\x2c\xd5\x3a\x22\xd5\x86\x18\xdb\x32\x9d\x5d\x9a\xc0\x55\x27\xfa\x75\xcf\x1b\xd6\x7e\x23\x8a\xd6\x7c\x8d\x93\xf1\xdf\x06\xa7\x00\x36\xb6\xe3\x8c\x04\x26\xff\xee\x24\x08\x34\x8e\xc4\x5b\xa2\xb8\xd9\xf5\x86\x9c\x96\x7b\x17\x0f\x2e\xa7\xfd\xcf\xaf\xc3\x01\xcf\xa3\x35\xe1\xe2\xb7\x7e\xdb\xf3\xac\xb7\xbf\x45\x7e\xe4\xc7\xe0\x79\x1b\xa2\xda\xe1\x8f\x57\x12\xa5\x38\x2a\x48\xd2\xdc\xf7\x7d\x1f\x3c\x4f\x48\x8f\x08\xca\x51\x18\xaf\x6c\x24\xbd\xd7\xfd\x9a\x46\xb5\x46\xaf\xb1\x49\x05\xcf\x5b\x91\xad\xd7\xda\xa6\x86\x30\xb1\x42\x5a\x90\x56\xd7\xd2\x0c\x8b\x6e\x6d\xc5\xc5\x93\x9f\xd6\x67\x42\x0d\x5f\x23\x78\x9e\x05\xb3\x4d\x91\xac\xaa\xe7\x99\x00\xcf\x63\xa5\x47\xe5\xaa\xb5\xe7\xa5\x00\xad\x99\x0d\x89\xd0\x1a\x3d\xcd\x7f\x20\xc4\x7e\x91\x82\xe7\x7d\xd3\x52\xa8\x96\x7a\xb5\xd4\x46\x03\x69\x9a\x47\x6b\x5c\x18\x54\x15\xa1\x68\xd7\xbf\x3e\x2d\xf7\xf3\x64\xbe\x54\xf9\x23\x1b\x3e\x32\xdb\x7b\x02\x7b\x47\x8c\x84\x5b\x2c\xaf\xed\xba\xd1\xe0\x72\xa2\xa0\x52\x72\x05\x9d\x30\xaa\xd3\x16\x12\x52\xf1\x25\x17\x53\x98\x4c\xc6\xaf\xd6\xd3\x36\xf9\xb3\x5a\x7e\xf5\xbc\x4e\x68\x52\xa1\x87\xdb\x56\x6a\xfc\x0a\x55\x43\x96\xbf\x00\xf8\x3f\x63\xf6\xf0\x7f\x64\xf6\x27\xbd\xf4\x6f\x73\x7b\xe0\xc7\x93\x20\x89\x27\x41\x3e\x49\x9e\xdd\xee\x7b\xf2\x9d\xeb\x94\x13\xbc\xe9\xce\xee\x2e\xba\xe0\xdd\x76\xad\x77\x47\x8b\x6b\xb5\xd0\xc5\xda\x1c\xa5\xa5\xf9\x34\x13\xef\xcf\xe4\xf9\xb7\xf2\xfe\xc7\x31\x19\xbf\xa0\x3e\x99\x04\x79\x32\x09\xa3\xec\x55\x03\xc7\xef\xe8\x86\x2f\xbe\xc9\x8f\xb7\xef\xab\x23\x12\xe7\xe1\xcd\xdc\x10\xbc\xd9\x5e\x9c\x6f\x58\xfe\xa3\x14\x47\xc1\x75\xb6\xc1\xd9\xdd\xcd\xf6\xee\x6d\x76\x77\xa4\xf1\x2a\xb7\x87\xff\x07\x72\x7f\x83\xdb\xf3\xa4\x8c\xc2\x2a\x23\x51\x15\xfb\x71\x1e\x54\x41\x18\x45\xb1\x1f\x07\x69\xe6\xd3\x9c\x96\xe8\x67\x55\xc6\xb2\x82\xbe\xc9\xed\x49\x4c\x30\xca\xa2\xca\x2f\xd2\x8a\x54\x21\x2b\xd3\x32\x27\x71\x9a\x05\x19\xf5\xcb\x22\x47\x5a\x11\x3f\x4b\x18\x7b\x93\xdb\xe3\x38\xae\xd2\xb8\xc0\xc8\xcf\xe2\x38\xc4\x2c\xa5\xb4\xca\xa2\x2c\x4e\x53\x4c\xc2\x2a\x48\xfd\xa2\x2c\xf2\x30\xf5\xdf\xe6\x76\x3f\x0e\x32\x2c\xa3\xac\x88\x83\x20\x8d\xa3\x34\x8f\xfd\xe0\x24\x4d\xd3\x22\x8f\xe9\xe9\x49\x96\x16\xf1\xec\x88\x1e\x95\xc1\x78\x64\xc7\x61\x62\x08\x5c\x1b\xa9\xc8\x12\x47\xba\xff\xdb\x0f\xb9\x73\x62\x6a\x97\xe2\xc6\xce\x4a\x27\x47\x50\xf1\x06\x47\xd6\xa8\xa9\xa7\x70\x68\x56\xed\xe1\xcf\x61\xfb\x9f\x8c\x18\x32\x71\x27\x59\x69\xf5\x1e\x4b\x51\xf1\x65\xa7\x9c\x5b\x0f\x06\xa8\x5b\xbd\xfe\xef\xcd\xf4\x0a\x9e\x59\xfb\xf0\x40\x4d\x1b\x0b\x27\x87\x8a\xd9\xfc\x03\x10\xc1\x60\x1e\xce\xe1\xba\xe7\x15\xdb\xaa\x28\x6c\x2f\x8e\x6c\xb7\xbe\x97\xda\x08\xb2\xc2\x29\xf8\x6e\xc4\xf5\x47\x07\x30\x97\xca\x0c\x4a\xac\x82\x97\x05\xed\xa1\x29\xe4\x7e\x1e\x5a\xe3\xb6\x5f\x3d\x23\x1d\x35\x03\x7d\x1c\xb7\x1e\xb5\x61\xdb\x87\x79\xdd\x22\xe5\xd5\x0e\x4e\xb7\xc6\x31\x00\x7c\x98\x3f\xf2\xd5\x51\x16\x25\xc2\x3e\x18\x14\x5a\x56\x66\x40\x0c\xf0\x0a\x4a\xac\xb9\x60\x70\x31\x5b\x58\x35\x38\x48\x7f\x98\x4f\x61\x33\xd9\x4e\x76\x93\x1f\x7d\x12\xad\xd7\x9d\x46\xf6\xd0\x13\x36\xea\x86\xec\x50\xd9\x54\x3a\x77\x5d\x43\xbb\xd3\x0b\xbe\x42\xd9\xb9\x30\x05\xc8\x16\xc5\xf0\x8a\x19\x38\xd9\x71\x95\xbb\x67\x46\xb0\x5f\x1e\x44\xa6\x30\x8e\x7c\xed\x80\x73\xd9\x61\x87\xbf\x84\xeb\xac\x13\xbd\x13\xb4\x56\x52\xc8\x4e\x5b\xfa\xa3\xa8\x35\x17\xcb\xd1\x77\x2b\xd0\x27\xa3\x7f\x83\xe9\x3e\xf4\x6e\x55\xa2\xb2\x04\x6a\xfb\x1f\x95\x3e\xa4\x52\x68\xcb\xc9\x03\x99\x6e\xec\x10\x5c\xba\x4b\x47\x52\x62\xfa\xcc\x68\x43\x94\xe9\xda\x11\x58\xf9\xdb\x5e\x70\x0a\x7d\x78\x67\x0a\x51\x43\xd7\xc2\xf1\xfc\x06\xe8\x8e\x36\xa8\xfb\x50\x7b\x03\x76\x9e\xd8\x10\xee\x9e\x6e\xd6\x5f\x5c\xa3\x30\x36\xd4\x7e\xfb\x96\x70\x17\xed\xa7\xeb\x29\x04\x36\xd0\x07\x06\xd3\xae\x84\x9c\x3e\x0d\x7a\xb4\x67\xb0\xa1\xce\xd8\xa0\xe5\xa6\x4d\xcd\x69\xfd\xc0\x6e\x40\x28\x95\x9d\x70\x37\x96\x1d\x6e\x86\x8b\x46\xda\x24\x0c\x37\x04\x03\xde\xdf\x62\xb4\xd3\x46\xae\x06\x23\xfb\x6e\x18\x9e\xaa\xb3\x5e\xcd\x85\x03\xed\xd8\x3e\x4f\xc7\x0f\x0f\x52\xd7\x68\x83\xe2\x07\xbb\xb4\xb1\x73\x47\x8f\xaf\xbf\x6c\xd0\x8d\x5d\x5c\x21\x6c\x34\x48\x05\xbc\xa5\xc3\x2b\xd5\x3e\x4a\xed\x27\x25\xc6\xba\xed\x52\xf2\x57\x9b\x5d\xc9\xf0\xe6\xea\x7c\x0a\xb5\x31\xed\xf4\xf0\xd0\xdd\xf3\x76\x38\x98\x16\x49\x9c\xec\x8b\xe9\x5e\xd1\x4b\x62\x63\xe1\xd4\xba\xbb\x24\x7a\x6e\x3f\xa7\x10\xf8\xfb\x7f\xcf\x0e\x37\x7c\xc5\x4d\x7f\xf8\xdc\x7e\x4e\x21\xce\x82\x30\xca\xf3\x27\x20\x35\xd2\x55\xab\x87\x96\xf8\x19\x99\x51\x44\x68\xf2\x30\x44\xd8\x18\x18\xeb\x5f\xdd\x04\xdc\x9c\xe5\xba\xbf\x0f\x05\x8c\xe2\xcb\x25\x2a\x64\x3d\xa4\x0d\x6e\xcd\xbe\xd0\x3d\xac\x53\xdf\xe2\xfa\x35\xc3\x0a\x09\x03\x29\x9a\x9d\x6d\x97\x3d\xd8\xf7\xff\xf5\xb0\x77\xe9\xa7\xea\x2b\x24\xec\xa9\xfa\x20\x19\xb4\x5f\xd8\x4a\x3c\xf6\xbd\x95\xb2\x81\x15\xd9\x82\x42\xa3\x78\x3f\x28\x68\x14\x0c\xc8\x93\x63\x72\xed\x5a\x79\x45\xb6\x57\xfd\xb9\x29\x84\x43\x4e\x5f\x56\xe9\xa6\xb5\x35\x69\x9c\xde\x5d\xdf\x00\xc4\x3a\x48\x3b\xa5\xdc\xb3\xf7\x91\x44\x4d\x34\x94\x88\xf6\x5d\x6c\x90\x1a\x97\xa6\xbd\x02\x6b\xcf\x5e\x4f\xe1\x10\xc1\x09\xd7\x0e\x2d\x4e\xa3\x96\xab\x67\x68\xd3\xc0\xe4\xe3\xa1\x1e\xcc\xd6\x79\x44\x5a\x3e\x02\x30\xdb\xb9\x94\xcd\x8c\x5a\x5a\x38\x15\x56\x13\x9b\x82\x51\x1d\xda\x5e\x23\x62\x07\x0c\xcb\x6e\xb9\x1c\x28\xc9\xb6\x80\x23\x80\xa5\x04\x6b\x64\xe4\x76\xfb\x56\x6b\x5b\x25\x2b\x57\x9e\x07\x11\x4b\x76\x76\x75\x0a\x15\x69\x34\x8e\xfe\x15\x00\x00\xff\xff\x54\x43\x46\xfe\x3c\x12\x00\x00")

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

	info := bindataFileInfo{name: "go-centrifuge/build/configs/default_config.yaml", size: 4668, mode: os.FileMode(420), modTime: time.Unix(1545075267, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _goCentrifugeBuildConfigsTesting_configYaml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x94\x92\xcd\x6e\x1c\x37\x0c\xc7\xef\xf3\x14\x02\x2f\xbe\xec\xae\xf5\xfd\xf5\x06\x45\xd0\x5e\x5a\x20\x67\x4a\xa2\x6c\x61\xbd\xb3\x53\x8d\xc6\xce\x22\xc8\xbb\x17\xb3\xb1\x9b\xab\x83\xb9\x90\x04\xff\xfc\x93\xa3\x5f\xa6\x79\xf4\x56\xb7\x27\xfa\x8b\xc6\xdb\xb5\x9f\x23\x1b\xb4\x8e\x36\x3f\x4d\x34\x9e\xa9\xd3\x76\x89\x13\x63\x98\xf3\x75\x9b\xc7\xba\xc7\x8c\x5d\xb0\xcd\x91\xdd\x43\xc6\xce\x74\x8b\xec\xe1\x3b\x60\x29\x9d\xd6\x15\x22\xf8\x90\x38\x7a\x6b\xbc\xca\x5a\x6b\x8d\xb9\x16\x27\x92\xb6\x8a\x78\x51\xd9\x18\x24\xa1\x85\x44\x03\x07\xc8\xfd\xb6\x8c\x2b\xc4\xef\x90\xdb\xf2\x4c\x1d\x22\x20\xad\x47\x21\xfd\x31\x8f\xbe\x37\xdc\xcb\x83\xbe\x0d\x88\x90\x9d\x0b\xd5\x2b\x17\x8a\x73\xbc\x04\x99\x6b\x16\xa5\x14\x8d\xbe\x2a\x51\x0c\x72\x2c\xd9\x57\x89\x3c\x49\x14\x9a\x0b\xe5\x78\x51\x56\xf1\xaa\x7c\xe6\xd9\xe3\xff\xf3\x16\xec\x78\x59\x77\xdb\xf6\x0a\x11\x94\xcd\xc2\x7a\x72\x2a\xd5\xe0\x79\x25\x67\x12\x77\xd2\x55\x1f\x38\x3a\x81\x05\x7e\x1c\xe0\x5c\x2a\x44\x58\xef\x0b\xc3\x3d\xfd\x35\xa4\x9c\x5f\x68\x86\xa8\xe4\x01\x66\x88\xd2\x4a\xa1\xf5\x01\x16\x88\xe2\x00\x1d\xa2\x3f\xc0\x8a\x2f\xfb\x01\x85\x44\x22\x61\x49\xe5\xe0\x45\xd0\xba\x08\xca\x28\x93\x4f\xd2\x91\x26\x4b\x3c\x99\x54\x93\x56\x89\xb8\x72\x16\x4d\xf1\xde\x87\x8a\xd6\x05\x94\x5e\x48\xb9\x2f\x72\xc1\xbc\xff\x8a\x2c\xa4\x4f\x5e\x18\x63\x4c\x42\x41\x58\x5c\x46\x0a\xdc\x72\xf2\x5e\x4b\xac\x19\xbd\x32\xb6\x70\xab\x8d\x49\x25\xa0\x71\x46\x26\xb4\x35\x67\x1e\x24\xd5\x7d\x52\x2b\x10\x41\x1b\xe2\x96\xa3\x3d\x16\x89\x74\xd4\x2a\xf9\x63\x90\xb2\x1e\xb5\xf6\x32\xe8\x10\x8a\x72\x05\x0e\xf0\x4a\x7d\x6d\xd7\xfd\xc8\x1f\x0f\xef\x0f\xbf\xe0\xba\xbe\x5d\x7b\x89\xec\xe1\xa3\xf4\xce\x40\x64\x9f\x45\x60\x9a\x5a\xa1\x79\xb4\x71\xfb\xa3\x44\x06\xfc\x1b\x17\xbf\x3e\x98\xa6\x7f\x37\xda\x68\x87\x6e\xde\x2e\x5f\xaf\xfd\x4c\x7d\x8d\x4c\x4e\x8c\xbd\xdd\x93\xaf\xd8\xc6\x3f\xed\x42\x7f\xfe\x1d\x99\x98\xa6\x33\xdd\xee\x84\xae\xed\x69\x6e\xf3\xd3\x4f\x58\x97\x2d\xbd\xb4\xfc\x65\xa7\xf4\x74\x7a\x3c\x9d\x1e\xd3\xd6\x5e\xca\x63\xa7\xf5\xba\xf5\x4c\xeb\xe3\x7b\xf7\x17\xba\x9d\x96\x2d\x9d\x16\xba\xfc\xd4\xf5\xf6\x8a\x83\x3e\x27\x3c\xef\xe2\xbb\x90\xc6\x33\x6e\xe3\xf9\x93\xde\xef\xdd\xbf\x69\xfc\xa1\xfa\x70\xfd\x2f\x00\x00\xff\xff\x56\xed\xff\x60\xc5\x03\x00\x00")

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

	info := bindataFileInfo{name: "go-centrifuge/build/configs/testing_config.yaml", size: 965, mode: os.FileMode(420), modTime: time.Unix(1545556716, 0)}
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
