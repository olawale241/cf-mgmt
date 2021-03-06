// Code generated by go-bindata.
// sources:
// files/cf-mgmt.sh
// files/cf-mgmt.yml
// files/pipeline.yml
// files/security-group.json
// files/vars.yml
// DO NOT EDIT!

package generated

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

var _filesCfMgmtSh = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x52\x56\xd4\x2f\x2d\x2e\xd2\x4f\xca\xcc\xd3\x4f\xcd\x2b\x53\x48\x4a\x2c\xce\xe0\x2a\x4e\x2d\x51\xd0\x4d\xe5\x4a\x4e\x51\x48\xce\xcf\x4b\xcb\x4c\xd7\x2d\x4a\x2d\xc8\xe7\x4a\x4e\xd3\xcd\x4d\xcf\x2d\x51\x50\x71\x76\x8b\xf7\x75\xf7\x0d\x89\x77\xf6\xf7\xf5\x75\xf4\x73\xe1\x02\x04\x00\x00\xff\xff\x4b\xdb\x02\xbb\x43\x00\x00\x00")

func filesCfMgmtShBytes() ([]byte, error) {
	return bindataRead(
		_filesCfMgmtSh,
		"files/cf-mgmt.sh",
	)
}

func filesCfMgmtSh() (*asset, error) {
	bytes, err := filesCfMgmtShBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "files/cf-mgmt.sh", size: 67, mode: os.FileMode(493), modTime: time.Unix(1518199076, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _filesCfMgmtYml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x54\xcc\x4b\x6a\xc4\x30\x0c\x00\xd0\xbd\x4e\x21\xb2\xae\x9b\xbd\x2f\x53\x84\xab\x38\x22\xfe\x21\xc9\xa1\xa1\xf4\xee\x25\xc3\x30\x30\xeb\x07\x2f\x84\x00\xa3\x90\x6f\x5d\x6b\xc4\x22\x6d\xfe\x00\x48\xa5\xcc\x5f\xca\xd6\xa7\x26\x8e\x80\xe8\xd7\xe0\x88\xdf\x3d\x1d\xac\xe1\xc1\x80\xf8\x64\xfc\x55\x1e\xdd\xc4\xbb\x5e\x11\x87\x9c\xdd\xa9\x18\xeb\x29\x89\x6d\x4d\x5b\xa8\xb9\xfa\x07\x3a\xe5\x88\x4b\x21\x67\xf3\xe5\x0f\x01\xa4\x8d\xe9\x76\xef\x01\x1b\x55\x8e\x98\x7a\xdb\x24\x87\xbb\x03\xd0\xd9\x6e\x1b\xe4\xfb\x9b\xac\x49\x56\x27\x3b\x5e\xf5\xa7\xed\xf0\x1f\x00\x00\xff\xff\x06\xd9\x38\x16\xc7\x00\x00\x00")

func filesCfMgmtYmlBytes() ([]byte, error) {
	return bindataRead(
		_filesCfMgmtYml,
		"files/cf-mgmt.yml",
	)
}

func filesCfMgmtYml() (*asset, error) {
	bytes, err := filesCfMgmtYmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "files/cf-mgmt.yml", size: 199, mode: os.FileMode(420), modTime: time.Unix(1554829763, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _filesPipelineYml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xec\x98\xdf\x6f\xda\x30\x10\xc7\xdf\xf3\x57\xdc\x63\x91\x9a\xf5\x3d\x6f\xa8\xa4\x15\x12\x3f\x2a\x60\x9b\xa6\xa9\xb2\xdc\xe4\x48\xbd\x39\x3f\x66\x3b\x4c\xfc\xf7\x93\x09\x0d\xc4\x31\x05\x02\x55\x27\x25\x6f\x55\x7c\xb9\xbb\x7c\x3f\xbd\xf3\x71\x02\x65\x9a\x8b\x00\xa5\xe7\xb8\x90\xd0\x18\x3d\x08\xd2\x64\xc9\x22\x57\x60\x96\x3a\x00\x6a\x9d\xa1\x07\x11\x53\x0e\x40\x61\xea\x39\x00\x00\xb9\x60\x1e\xdc\xdc\x44\x4c\x11\x6d\x49\x72\xc1\x7a\xbd\xcd\xc9\x8b\xa0\x49\xf0\x5a\x39\x2c\x1e\xf5\x7a\x65\x0c\xc5\x62\x74\x95\x60\x51\x84\xa2\x0c\xa2\x1f\x1a\x51\x58\xa2\x50\xac\x28\xd7\xde\xf6\xdf\xe9\xf5\x1c\xe7\x57\xfa\xb2\x9f\xb5\x40\xaa\xd0\x4d\x45\x24\x1d\x80\x8c\xd3\x44\x7b\x70\x21\x42\x65\x7e\x12\xc0\xd6\x8b\x07\x4a\xe4\xb8\x31\x53\x54\xfe\x36\x9d\x00\x2c\x19\xaf\x0a\x72\x17\xb0\x3b\x6d\x2a\xef\x82\xa5\x1b\x47\xb1\xfa\xb2\x8e\xf9\xc6\x34\xa3\x82\xc6\xb2\xc8\x1a\x60\xfe\x63\xbe\xf0\xc7\x64\x30\x1d\xf7\x87\x13\x9d\xbc\x5c\x4b\x85\x31\x09\xd3\x98\xb2\x64\x2b\x14\xc0\xd7\xb9\x3f\x23\xc3\x81\x36\xc8\x25\x0a\xc2\xc2\xf2\xe8\xa9\x3f\x9f\x7f\x9f\xce\x36\x67\x19\x95\xf2\x6f\x2a\x76\x87\xf7\xd3\xc9\xc3\xf0\x91\x0c\x86\x33\x7d\x5c\xe4\x47\x42\x26\x76\x06\xa3\xa1\x3f\x59\x90\xb9\x7f\x3f\xf3\x17\x1b\x1b\xce\x30\x51\x44\x62\x20\x50\x95\x66\xa3\xe9\x23\x19\xf9\xdf\xfc\x91\x36\xe1\x69\x44\x38\xae\x90\xef\xbc\x3c\x90\xf1\xe3\x78\x41\xee\xa7\xe3\x71\x7f\x32\xa8\xea\x53\x2a\x2f\x5f\xa9\xc0\xd0\x2d\x3e\xed\x02\xf1\x6b\x7e\x3a\xfd\x4d\xfd\x0d\x89\xcc\x7f\x7e\x89\x41\x2e\x98\x5a\xbb\x91\x48\xf3\xec\xf2\x3a\xa8\xfb\xeb\x90\x1c\x28\x09\x53\xaa\x12\x0d\x95\x92\x45\x89\x1b\xe2\x92\xe6\x5c\x5d\x8e\x48\xcb\x2a\x25\x86\x1e\xfc\xb4\x87\x7e\xde\xe3\x78\x34\x78\xc7\xd3\xe4\x79\x44\xb2\x92\x6b\x88\x1c\x2f\xb8\x6f\x36\x66\xc6\x2d\x78\xb8\x1e\xab\xc1\x3a\x68\x26\xb4\x7d\x7d\x2c\x13\x81\x9b\x09\xb6\xd2\x7f\x9f\x7c\x47\x99\x35\xa6\x3d\x3f\x9f\x38\x3a\x58\xa2\x75\xc4\x0e\x4f\x12\x35\xb9\x76\x83\x05\x8a\x15\x0b\xd0\xa5\x41\x80\xf2\x43\xa0\xd5\x22\x74\xa0\x6a\x23\x47\x55\xa2\xea\xd4\x77\xb5\xe2\x32\x9d\x1c\xc1\xf6\x4e\xec\x8e\xa0\x75\x68\x7c\xbf\xd2\xde\x26\x89\x8c\x06\xd8\xb0\xd0\x6e\x21\xcf\xc2\x37\x9a\x5a\x97\xd3\xfa\x65\x19\xb2\x05\xdc\x06\xfd\x27\xb2\x1f\x89\x87\x34\x23\xb5\x70\x4d\x07\xd0\x42\x48\x73\x3e\xb9\x04\xe9\x21\x80\x8d\x66\x97\x36\x81\x6e\x36\xbd\x98\x04\xb7\xf5\xd4\x94\x60\xf1\xde\x75\x18\x9a\xa9\x74\x0c\x4d\x86\x55\x85\xac\x0c\x8b\xae\xf8\xff\x80\x2c\xf3\x69\x01\xcd\x8f\x69\xbd\x16\x35\xed\xe4\xff\xe4\xa9\xa2\x57\x47\x6f\x41\x5a\x06\x6a\x01\xd3\x4b\x60\x6d\x75\xb2\xd3\x3a\x7b\x47\x63\xc7\x76\x7b\x78\x3f\x73\x22\xcd\x6e\x5f\x73\x26\xd6\x83\xdb\x1a\x73\x36\xfd\xcc\x81\xc8\x92\x4b\x0b\xb0\x7e\x68\x07\xde\x69\x69\x23\xde\xb4\xfb\x1e\xdf\x1f\xd8\x82\xb4\x80\x65\x73\x48\x66\xdf\x65\x32\xe5\x54\xb1\x34\x71\x25\x46\x31\x26\xea\x2a\x1b\xf1\xe2\xf7\x68\xed\x0e\x3d\xb7\x50\xad\xc9\x7d\x2a\xde\x13\x01\x1d\x05\x7d\x2e\x41\x8b\x12\xbb\xdd\x01\x47\x9a\xe4\x59\xa3\xde\x6a\x96\xef\x6d\x7d\xa2\xba\x4e\xcf\xb5\x25\xd9\x15\x6a\x6d\x95\x50\x53\xe9\x5f\x00\x00\x00\xff\xff\x19\x1a\xd5\x2c\x1f\x20\x00\x00")

func filesPipelineYmlBytes() ([]byte, error) {
	return bindataRead(
		_filesPipelineYml,
		"files/pipeline.yml",
	)
}

func filesPipelineYml() (*asset, error) {
	bytes, err := filesPipelineYmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "files/pipeline.yml", size: 8223, mode: os.FileMode(420), modTime: time.Unix(1555521269, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _filesSecurityGroupJson = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x8a\x8e\xe5\x02\x04\x00\x00\xff\xff\x44\xd2\x68\x70\x03\x00\x00\x00")

func filesSecurityGroupJsonBytes() ([]byte, error) {
	return bindataRead(
		_filesSecurityGroupJson,
		"files/security-group.json",
	)
}

func filesSecurityGroupJson() (*asset, error) {
	bytes, err := filesSecurityGroupJsonBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "files/security-group.json", size: 3, mode: os.FileMode(420), modTime: time.Unix(1510432475, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _filesVarsYml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x9c\x90\x41\xab\x14\x31\x10\x84\xef\xf9\x15\xc5\xdb\xf3\xb8\x78\xf0\x32\xb7\x87\xbb\x82\x20\xfa\x10\x3d\x0f\xd9\xa4\x27\xdb\x92\xa4\x87\x4e\xcf\x5b\xe6\xdf\x4b\x76\x56\x45\x8f\x9e\x92\x6a\xaa\xba\x3e\xfa\xb0\xc9\xaa\x48\x6c\x50\x5a\x04\xab\xb2\x4b\x6c\x53\x17\xd3\xaa\x3c\xfe\x51\x17\xf5\x35\x5c\x47\x14\xdf\x8c\xd4\xed\xc1\x30\xa3\x6d\xcd\xa8\x20\x4a\xf1\x5c\xdd\xae\xa6\x5d\x8d\xee\x80\xb5\x91\xc2\x87\x20\x6b\x35\xdc\xd8\xae\x58\x48\x0b\xb7\xc6\x52\x61\x82\xa0\xe4\x8d\x20\x9a\xda\xb1\x2d\x3e\x50\x73\x3d\x32\x71\xec\xe9\xd3\xf9\xe5\xeb\xf9\xfd\xf3\xb7\xf3\x09\x03\xbe\x37\x42\xc8\x4c\xd5\xa6\x46\x41\xc9\x30\x60\xf1\xad\xdd\x44\x23\x64\xfe\x9f\xaa\x5f\xf1\x11\x4f\x4f\xee\xf0\xd8\x8e\xc7\xf6\x59\x14\xab\xf7\x61\xff\xec\x50\xee\x2f\x80\x8e\xf8\x9b\xc0\x04\x17\xae\xf7\x37\x47\xbf\x60\x80\xd4\xbc\xa1\x12\x45\x8a\xe0\xce\xc7\x35\xe1\xd3\xe9\xf9\xc5\x75\xc3\xf4\x4f\x79\x96\x94\xba\x21\xd3\x2b\xe5\x7b\x67\x98\x87\x92\x8a\x21\x48\x29\xbe\xc6\x06\xae\xb0\x2b\x61\xe1\x85\x32\x57\x72\x59\xd2\x74\xb7\x8f\xf8\xf8\xf9\xc3\x17\x77\x80\x71\x21\x70\x35\xd2\x57\x9f\x3b\x8a\x29\xa7\x44\x8a\x75\x89\xde\xe8\x18\x29\x93\x11\x7e\xc8\xa5\x41\xaa\xeb\xf6\xe1\x61\x19\xf1\xf6\x5d\xe9\x47\x90\x3a\x73\x5a\xd5\x5b\xbf\x5b\x64\xa5\x60\xa2\x9b\xdb\xe7\x53\x64\x1d\xf1\xe6\xb8\x2b\xf7\x33\x00\x00\xff\xff\xfb\x82\x67\xfb\x40\x02\x00\x00")

func filesVarsYmlBytes() ([]byte, error) {
	return bindataRead(
		_filesVarsYml,
		"files/vars.yml",
	)
}

func filesVarsYml() (*asset, error) {
	bytes, err := filesVarsYmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "files/vars.yml", size: 576, mode: os.FileMode(420), modTime: time.Unix(1555521290, 0)}
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
	"files/cf-mgmt.sh": filesCfMgmtSh,
	"files/cf-mgmt.yml": filesCfMgmtYml,
	"files/pipeline.yml": filesPipelineYml,
	"files/security-group.json": filesSecurityGroupJson,
	"files/vars.yml": filesVarsYml,
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
	"files": &bintree{nil, map[string]*bintree{
		"cf-mgmt.sh": &bintree{filesCfMgmtSh, map[string]*bintree{}},
		"cf-mgmt.yml": &bintree{filesCfMgmtYml, map[string]*bintree{}},
		"pipeline.yml": &bintree{filesPipelineYml, map[string]*bintree{}},
		"security-group.json": &bintree{filesSecurityGroupJson, map[string]*bintree{}},
		"vars.yml": &bintree{filesVarsYml, map[string]*bintree{}},
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

