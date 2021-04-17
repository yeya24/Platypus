// Code generated for package resource by go-bindata DO NOT EDIT. (@generated)
// sources:
// lib/runtime/config.example.yml
// lib/runtime/template/rsh/bash.tpl
// lib/runtime/template/rsh/go.tpl
// lib/runtime/template/rsh/lua.tpl
// lib/runtime/template/rsh/nc.tpl
// lib/runtime/template/rsh/perl.tpl
// lib/runtime/template/rsh/php.tpl
// lib/runtime/template/rsh/python.tpl
// lib/runtime/template/rsh/python2.tpl
// lib/runtime/template/rsh/python3.tpl
// lib/runtime/template/rsh/ruby.tpl
package resource

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

var _libRuntimeConfigExampleYml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x8c\x92\xcb\x8e\xda\x4c\x10\x85\xf7\x48\xbc\x43\x89\x91\xa5\xff\x97\x92\x99\x10\x16\x13\xb1\x9b\x90\xdb\x2c\x26\xa0\xc0\xec\x5d\xd8\xc5\xb8\x45\x5f\x9c\xae\x6a\x18\xde\xa6\x9f\xa5\x9f\x2c\x6a\xdb\x10\xb2\x89\x22\xaf\x5c\xd5\x7d\xbe\xd3\xa7\x8a\xc9\x1f\xc8\xf3\x1c\xc6\x23\x80\xb7\xd0\x38\x96\x39\x4c\xde\xdd\x76\xdf\x24\x17\x01\x5a\xe7\x65\x0e\xd3\xd9\xec\xbe\xff\xbf\x49\x71\xa5\x51\x4e\x6d\xe0\x14\x15\xa7\x88\x5b\x4d\x29\x8a\x4b\x31\x30\xa5\xc8\x74\x20\x8f\x3a\xc5\xd6\xbb\x96\xbc\x28\xca\x67\x38\xc5\x60\xd5\xcf\x40\x29\xaa\x9a\xac\xa8\x9d\x22\x9f\xe2\x7f\xad\x57\x46\x79\xf4\xa7\x14\xf7\x74\xfa\x3f\x45\xb7\x4b\x11\x53\x64\x65\x5f\xb2\x6c\xa5\x15\x59\xb9\xbd\xa0\x1f\xb4\x4e\x11\x0f\xa8\x74\x8f\xfd\x03\xe2\x29\x45\xad\x58\xa8\x4e\x71\x4b\xda\x1d\xe7\x97\x7b\x65\xa1\xca\x14\x1f\x57\x57\x85\x50\xa6\xf8\xcc\xe4\x2d\x1a\xba\x2a\x9b\x32\xc5\xa7\x87\x45\x8a\x58\xd7\x9e\x98\xaf\x5a\xae\x4c\x71\xd9\x92\x47\x51\xf6\x25\xc5\xf5\x89\x85\xcc\x55\x5f\x32\xc2\x56\xce\x10\x6c\x94\xa1\xb5\xa0\x69\xfb\x76\x83\xdc\x7c\x71\xde\x60\xce\xb7\x50\x50\x04\x28\x0c\x14\x6e\xf2\x2f\xc1\x7f\x18\x10\xf0\x9c\x53\xf9\x2d\x0d\xa8\xb5\x3b\x32\x04\x06\x71\x20\x1e\xab\x7d\x2e\x41\xe5\xac\xa5\x4a\x94\xb3\x0c\x3b\xef\x0c\x48\x43\xc0\x68\x08\x1e\x57\x70\x07\xe7\x37\xc3\x1d\x2c\xd7\x80\xb6\x86\xa7\x87\xc5\xed\xdf\x7d\x42\x21\x93\xf1\xc8\x13\xcb\x2e\xe8\x2e\xd5\x1b\xf8\x48\x15\x06\x26\xf8\xf1\x79\xbd\xd9\x05\x0d\x9f\x96\xf0\x7d\xb9\x01\x0e\x6d\x76\x0e\x68\x4f\x80\x41\x9a\x3c\xec\x0a\xb3\x1b\x30\x54\x35\x68\x15\x9b\x37\xbd\xc2\x70\x83\x5e\x5b\xc7\xd4\xb9\x1c\x08\xd0\x2f\x26\x28\x2b\xae\x13\xa2\x57\xc9\xa6\x35\x58\x92\xa3\xf3\xfb\xce\xee\x90\xda\xf4\xfd\x7d\x97\xdb\xb4\xcb\xad\x4f\xed\x7e\x36\x9b\xf6\x8c\x92\x6c\xde\x94\x39\x88\x0f\x54\x82\x21\xb4\x0c\x2c\xe8\xf3\x0c\x2f\xde\xd7\x3d\xef\xd8\x90\x85\xf3\x76\xf7\xa7\xb8\x43\x5d\x8b\x8c\x47\x37\xb0\x68\xa8\xda\x83\xa5\x23\x78\xd2\x84\x4c\x43\xd2\x5f\x95\x7c\x0b\xdb\x5e\xe7\x02\x39\x0b\x8e\x47\xa1\xad\x51\x06\x99\x5f\x01\x00\x00\xff\xff\x3c\x4f\x27\x33\x7c\x03\x00\x00")

func libRuntimeConfigExampleYmlBytes() ([]byte, error) {
	return bindataRead(
		_libRuntimeConfigExampleYml,
		"lib/runtime/config.example.yml",
	)
}

func libRuntimeConfigExampleYml() (*asset, error) {
	bytes, err := libRuntimeConfigExampleYmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "lib/runtime/config.example.yml", size: 892, mode: os.FileMode(438), modTime: time.Unix(1618208039, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _libRuntimeTemplateRshBashTpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xd2\x2f\x2d\x2e\xd2\x4f\xca\xcc\xd3\xcf\xcb\xcf\x28\x2d\x50\x00\x33\x93\x12\x8b\x33\x14\x74\x93\x15\xd4\x91\x78\x99\x0a\x76\xfa\x29\xa9\x65\xfa\x25\xc9\x05\xfa\xf1\xf1\x1e\xfe\xc1\x21\xf1\xf1\xfa\xf1\xf1\x01\xfe\x41\x21\xf1\xf1\x0a\x06\x76\x6a\x86\xea\x50\x15\x79\xa5\x39\x39\x0a\x6a\x80\x00\x00\x00\xff\xff\x8f\x47\x46\xc1\x58\x00\x00\x00")

func libRuntimeTemplateRshBashTplBytes() ([]byte, error) {
	return bindataRead(
		_libRuntimeTemplateRshBashTpl,
		"lib/runtime/template/rsh/bash.tpl",
	)
}

func libRuntimeTemplateRshBashTpl() (*asset, error) {
	bytes, err := libRuntimeTemplateRshBashTplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "lib/runtime/template/rsh/bash.tpl", size: 88, mode: os.FileMode(438), modTime: time.Unix(1616234053, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _libRuntimeTemplateRshGoTpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x6c\x8e\x4f\x4b\xc4\x30\x10\xc5\xbf\xca\x30\x87\x6e\x0b\x6b\x73\x6f\xe9\x5e\xf4\xe0\x6d\x65\x77\x8f\x81\x90\x9d\xc6\x36\xd8\x4c\x42\xfe\x88\x22\x7e\x77\xb1\x54\x3d\xb8\xb7\xf7\x7e\xbf\x61\x78\xa2\xa4\x28\xae\x96\x05\xfb\xb9\x04\x58\xe3\x55\xa7\x19\xee\x08\xd0\xd0\xec\x61\x17\x34\xbd\xe8\xc9\x80\xd3\x96\x7b\xeb\x82\x8f\x59\xa2\x4f\xc2\xbc\x19\x92\xf8\x4b\xd8\x64\x89\xfd\x73\x61\x5a\x2f\xeb\xe6\x83\xf6\xaa\x1b\xd8\xe4\xf6\xc1\xea\xa5\x96\x98\x29\x48\xdc\x4b\x54\xea\xf1\x78\xbe\x28\xd5\x29\xf5\x74\x3c\x5d\x94\x92\xd8\xf4\xe4\xc6\x6e\xf8\x7e\xd9\xde\x7b\xe7\x34\x8f\xb5\xc4\x75\x4d\x9a\x37\xdd\x9e\xf3\x68\x79\xa0\x9f\xec\x4b\xfe\x2b\x26\xc6\xad\x9c\x0a\xd7\xcd\xe7\x0e\x0e\x20\xb2\x0b\x22\x2c\x3a\xbf\x87\x92\xda\xc9\x43\x55\xc1\xe4\x21\x16\xbe\xa9\xa2\xfb\x87\x11\x0e\x62\x34\xaf\x82\xcb\xb2\x40\xf5\x15\x00\x00\xff\xff\xd8\x37\xb8\xf7\x2c\x01\x00\x00")

func libRuntimeTemplateRshGoTplBytes() ([]byte, error) {
	return bindataRead(
		_libRuntimeTemplateRshGoTpl,
		"lib/runtime/template/rsh/go.tpl",
	)
}

func libRuntimeTemplateRshGoTpl() (*asset, error) {
	bytes, err := libRuntimeTemplateRshGoTplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "lib/runtime/template/rsh/go.tpl", size: 300, mode: os.FileMode(438), modTime: time.Unix(1616234259, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _libRuntimeTemplateRshLuaTpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x4c\x8e\xc1\xaa\x83\x30\x10\x45\x7f\x65\x70\x61\x14\x9e\x66\xe1\xee\xd9\xba\xee\xce\xd2\x76\x59\x08\x3a\x1d\x50\x1a\x12\x9b\xcc\x94\x7e\x7e\xd1\x74\xe1\xee\x70\x38\x5c\xae\x96\x18\xf4\x38\x3b\xed\xfc\x24\x0b\x6c\x38\x0e\x71\x82\x0a\x41\x59\x19\xa0\x22\xc8\x02\xbd\x64\x0e\x54\xa8\xbb\x52\xd1\xe3\x93\x78\xa5\xb2\xdd\x7b\x1f\x93\xe3\x63\x2a\x6a\xc6\xa5\x28\x5b\xfe\x47\xef\x1c\x21\x6f\x91\x31\xa7\xfe\x7a\x33\x66\xe5\xbf\x24\xce\xfd\xe5\x27\xca\xd6\xc7\x9a\x3e\x84\xc2\x69\x72\x77\x66\x86\x43\xde\x40\x97\x37\x29\xcc\x14\x74\xfa\x41\x6f\xed\xc4\x5a\xc8\xbf\x01\x00\x00\xff\xff\x8c\xf3\xba\xa6\xc5\x00\x00\x00")

func libRuntimeTemplateRshLuaTplBytes() ([]byte, error) {
	return bindataRead(
		_libRuntimeTemplateRshLuaTpl,
		"lib/runtime/template/rsh/lua.tpl",
	)
}

func libRuntimeTemplateRshLuaTpl() (*asset, error) {
	bytes, err := libRuntimeTemplateRshLuaTplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "lib/runtime/template/rsh/lua.tpl", size: 197, mode: os.FileMode(438), modTime: time.Unix(1616234223, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _libRuntimeTemplateRshNcTpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xd2\x2f\x2d\x2e\xd2\x4f\xca\xcc\xd3\xcf\xcb\xcf\x28\x2d\x50\x00\x33\x93\x12\x8b\x33\x14\x74\x93\x15\x94\x72\xb3\xd3\x32\xd3\xf2\x15\xf4\x4b\x72\x0b\xf4\xf5\x0a\x72\x12\x4b\x2a\x0b\x4a\x8b\xad\xf3\x92\x15\xe2\xe3\x3d\xfc\x83\x43\xe2\xe3\x15\xe2\xe3\x03\xfc\x83\x40\x0c\x03\x1b\x54\x55\x0a\x35\x48\x66\xd5\x28\x94\xa4\xa6\xa2\x19\xa3\xa4\x60\xa7\x9f\x92\x5a\xa6\x9f\x57\x9a\x93\xa3\xa0\x06\x08\x00\x00\xff\xff\x5e\x90\x17\x06\x87\x00\x00\x00")

func libRuntimeTemplateRshNcTplBytes() ([]byte, error) {
	return bindataRead(
		_libRuntimeTemplateRshNcTpl,
		"lib/runtime/template/rsh/nc.tpl",
	)
}

func libRuntimeTemplateRshNcTpl() (*asset, error) {
	bytes, err := libRuntimeTemplateRshNcTplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "lib/runtime/template/rsh/nc.tpl", size: 135, mode: os.FileMode(438), modTime: time.Unix(1616234260, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _libRuntimeTemplateRshPerlTpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x4c\x8f\xe1\x6a\xf3\x20\x18\x85\x6f\x45\x24\x54\x05\x8b\x17\x20\x2d\x7c\x7c\xeb\x58\x19\x6b\x4a\x75\xff\x06\x2f\x89\x79\xb7\xca\xd2\x57\x89\x66\x50\xc6\xee\x7d\xa4\x50\xd8\xbf\xc3\xc3\x39\x07\x1e\x33\x97\xc9\xf4\x91\x0c\xa5\xf3\x9c\xd9\x2d\xf6\x5d\x39\xb3\x75\x60\x22\xe3\x34\xb2\x35\x32\xf1\x26\xc4\x5c\x90\xb9\x14\x3e\xb1\xda\x26\x6e\x38\xc0\x53\xeb\x3c\x00\xb7\x4d\xde\x00\x1c\xdb\x93\x07\xb0\xe5\x56\x90\x4e\x1f\x1f\x61\x7f\xd8\x79\xed\xda\xff\xcf\xe0\xfc\x69\xf7\xef\x45\x7f\x60\xcd\x53\xaa\xa9\xbf\x52\x77\x41\xc9\x6b\xc8\x5c\x29\x1b\xdf\x65\x48\x44\x18\x96\xdd\x72\xd0\x0d\xc3\x04\x91\x64\x93\x75\x24\xac\xd0\xd5\x44\xb2\x89\x4a\x29\xf5\x9d\x32\x92\x74\xfe\x61\x7f\xd0\x7c\xbb\x72\x5c\xd9\x3b\x69\x5f\xfd\x1d\x95\x6b\xa9\x78\x91\xfc\x8f\x4c\xe4\xca\xfe\xd8\xc5\x43\xb0\xad\x19\xf0\xcb\xd0\x3c\x8e\x6c\xf5\x1b\x00\x00\xff\xff\xf4\xfb\xfa\x74\xff\x00\x00\x00")

func libRuntimeTemplateRshPerlTplBytes() ([]byte, error) {
	return bindataRead(
		_libRuntimeTemplateRshPerlTpl,
		"lib/runtime/template/rsh/perl.tpl",
	)
}

func libRuntimeTemplateRshPerlTpl() (*asset, error) {
	bytes, err := libRuntimeTemplateRshPerlTplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "lib/runtime/template/rsh/perl.tpl", size: 255, mode: os.FileMode(438), modTime: time.Unix(1616234223, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _libRuntimeTemplateRshPhpTpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x4c\xcc\xb1\x0a\xc2\x30\x10\x87\xf1\x57\x39\x82\xf4\x5a\xb0\xdc\xd0\xb1\xda\xd9\xad\xa2\x1d\x85\x3f\x36\x9e\xa4\x18\x92\xd0\x10\xf1\xf1\xa5\x4e\x2e\x1f\xdf\xf4\x93\x92\x57\x99\x97\x20\x21\xba\x92\xe8\xb7\xf3\x3d\x3b\x6a\x2d\x71\x72\x89\xda\x95\xf8\xc6\xbc\xcb\xd1\xbe\x8e\xcf\xad\x31\x69\xa8\x0d\x70\x1a\xaf\x13\x60\xf6\xc0\x79\xbc\x4c\x40\xd3\x67\xa7\xde\x43\x3f\x6a\x6b\xf3\x27\x2d\x74\xa8\x3a\x1a\xaa\xce\x34\xfd\x66\x31\x0d\xf2\xd0\xb7\x84\xe2\x3d\x55\xdf\x00\x00\x00\xff\xff\x2c\xe7\x7f\x38\x82\x00\x00\x00")

func libRuntimeTemplateRshPhpTplBytes() ([]byte, error) {
	return bindataRead(
		_libRuntimeTemplateRshPhpTpl,
		"lib/runtime/template/rsh/php.tpl",
	)
}

func libRuntimeTemplateRshPhpTpl() (*asset, error) {
	bytes, err := libRuntimeTemplateRshPhpTplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "lib/runtime/template/rsh/php.tpl", size: 130, mode: os.FileMode(438), modTime: time.Unix(1616234223, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _libRuntimeTemplateRshPythonTpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x6c\x8f\x41\x4b\xc4\x30\x10\x85\xff\x4a\xc8\xc1\x24\x10\x13\xf5\x1a\x14\x16\x59\x51\x44\x2b\xdb\x1c\x85\xc1\x66\x23\x2d\xb6\x99\xd0\x49\x84\xfe\x7b\xb1\x2d\x9e\x3c\xcd\x7b\x1f\x73\x78\x9f\xad\x34\xdb\x6e\x48\x36\x61\x5f\x33\x5b\x63\xf7\x41\x3d\xbb\x0c\x4c\xe4\xa5\xf4\x98\xd6\xf8\x2e\xc4\x30\x65\x9c\x0b\x23\x0c\x5f\xb1\x68\xaa\x5d\x9e\x31\x44\x22\x8d\xe4\xe8\x76\xc3\x66\x3b\x72\x6f\x87\x07\x78\x7a\x3d\x7a\xbd\xd7\xb6\xb9\x7f\x86\xd6\x9f\x8e\x87\x17\xe5\xc8\x04\x4c\x29\x86\x22\x25\x07\x78\x6c\x5a\x0f\xc0\x35\xc0\x5b\x73\xf2\x00\x4a\x39\x24\x73\xae\xf9\x46\x92\xf9\x1c\xc6\x98\x50\x2a\x7d\xa5\x1c\xfb\x07\x5f\x2b\xb7\x8f\x43\x5a\x1f\x68\xa1\x12\x27\xc9\xff\x74\xb8\xfa\x35\x10\xec\xce\x9e\xe3\xb7\x4d\x75\x1c\xd9\xc5\x4f\x00\x00\x00\xff\xff\x7b\xea\xca\x30\xfb\x00\x00\x00")

func libRuntimeTemplateRshPythonTplBytes() ([]byte, error) {
	return bindataRead(
		_libRuntimeTemplateRshPythonTpl,
		"lib/runtime/template/rsh/python.tpl",
	)
}

func libRuntimeTemplateRshPythonTpl() (*asset, error) {
	bytes, err := libRuntimeTemplateRshPythonTplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "lib/runtime/template/rsh/python.tpl", size: 251, mode: os.FileMode(438), modTime: time.Unix(1616234223, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _libRuntimeTemplateRshPython2Tpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x6c\x8f\x41\x4b\xc4\x30\x10\x85\xff\x4a\xc8\xc1\x24\x10\x13\xdd\x6b\x50\x58\x64\x45\x11\xad\x6c\x73\x14\x06\x9b\x8d\xb4\xd8\xcd\x84\x4e\x22\xf4\xdf\x8b\x6d\xf1\xb4\xa7\x79\xef\x63\x0e\xef\xb3\x95\x26\xdb\x0d\xc9\x26\xec\x6b\x66\x4b\xec\x3e\xa9\x67\xd7\x81\x89\x3c\x97\x1e\xd3\x6e\xc9\x1f\x42\x0c\xe7\x8c\x53\x61\x84\xe1\x3b\x16\x4d\xb5\xcb\x13\x86\x48\xa4\x91\x1c\xdd\xad\xd8\xac\x47\x6e\x6d\xff\x08\xcf\x6f\x07\xaf\xb7\xda\x36\x0f\x2f\xd0\xfa\xe3\x61\xff\xaa\x1c\x99\x80\x29\xc5\x50\xa4\xe4\x00\x4f\x4d\xeb\x01\xb8\x06\x78\x6f\x8e\x1e\x40\x29\x87\x64\x4e\x35\xef\x24\x99\xaf\x61\x8c\x09\xa5\xd2\x37\xca\xb1\x0b\xf8\x56\xb9\x6d\x1c\xd2\xf2\x40\x33\x95\x78\x96\xfc\xdf\x87\xab\x3f\x03\xc1\xee\xed\x29\xfe\xd8\x54\xc7\x91\x5d\xfd\x06\x00\x00\xff\xff\x89\x2f\x51\xfc\xfc\x00\x00\x00")

func libRuntimeTemplateRshPython2TplBytes() ([]byte, error) {
	return bindataRead(
		_libRuntimeTemplateRshPython2Tpl,
		"lib/runtime/template/rsh/python2.tpl",
	)
}

func libRuntimeTemplateRshPython2Tpl() (*asset, error) {
	bytes, err := libRuntimeTemplateRshPython2TplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "lib/runtime/template/rsh/python2.tpl", size: 252, mode: os.FileMode(438), modTime: time.Unix(1616234223, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _libRuntimeTemplateRshPython3Tpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x6c\x8f\x4f\x4b\xc4\x30\x10\xc5\xbf\x4a\xc8\xc1\x24\x10\x13\xff\x1c\x83\xc2\x22\x2b\x8a\x68\x65\x9b\xa3\x30\xd8\x6c\xa4\xc5\x6e\x26\x74\x12\xa1\xdf\x5e\x6c\x8b\xa7\x3d\xcd\x7b\x3f\xe6\xf0\x7e\xb6\xd2\x64\xbb\x21\xd9\x84\x7d\xcd\x6c\x89\xdd\x27\xf5\xec\x32\x30\x91\xe7\xd2\x63\xba\x5d\xf2\x87\x10\xc3\x29\xe3\x54\x18\x61\xf8\x8e\x45\x53\xed\xf2\x84\x21\x12\x69\x24\x47\x77\x2b\x36\xeb\x91\x5b\xdb\x3d\xc2\xf3\xdb\xde\xeb\xad\xb6\xcd\xc3\x0b\xb4\xfe\xb0\xdf\xbd\x2a\x47\x26\x60\x4a\x31\x14\x29\x39\xc0\x53\xd3\x7a\x00\xae\x01\xde\x9b\x83\x07\x50\xca\x21\x99\x63\xcd\x37\x92\xcc\xd7\x30\xc6\x84\x52\xe9\x2b\xe5\xd8\x19\x7c\xad\xdc\x36\x0e\x69\x79\xa0\x99\x4a\x3c\x49\xfe\xef\xc3\xd5\x9f\x81\x60\xf7\xf6\x18\x7f\x6c\xaa\xe3\xc8\x2e\x7e\x03\x00\x00\xff\xff\xc3\x60\x7f\x49\xfc\x00\x00\x00")

func libRuntimeTemplateRshPython3TplBytes() ([]byte, error) {
	return bindataRead(
		_libRuntimeTemplateRshPython3Tpl,
		"lib/runtime/template/rsh/python3.tpl",
	)
}

func libRuntimeTemplateRshPython3Tpl() (*asset, error) {
	bytes, err := libRuntimeTemplateRshPython3TplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "lib/runtime/template/rsh/python3.tpl", size: 252, mode: os.FileMode(438), modTime: time.Unix(1616234223, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _libRuntimeTemplateRshRubyTpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xd2\x2f\x2d\x2e\xd2\x4f\xca\xcc\xd3\xcf\xcb\xcf\x28\x2d\x50\x00\x33\x93\x12\x8b\x33\x14\x74\x93\x15\x94\x8a\x4a\x93\x2a\x15\x74\x8b\x8a\xf3\x93\xb3\x53\x4b\x14\x74\x53\x15\xd4\x53\x2b\x52\x93\x35\x62\x94\xe0\xca\x62\x94\x74\x62\x94\x74\x93\xc1\x14\x92\xde\x4c\x05\x3b\xfd\x94\xd4\x32\xfd\x92\xe4\x02\xfd\xf8\x78\x0f\xff\xe0\x90\xf8\x78\xfd\xf8\xf8\x00\xff\xa0\x90\xf8\x78\x05\x03\x3b\x35\xc3\x18\x25\x4d\x6b\x75\x25\xa8\xb2\xbc\xd2\x9c\x1c\x05\x35\x40\x00\x00\x00\xff\xff\x2c\x10\xbd\x97\x8b\x00\x00\x00")

func libRuntimeTemplateRshRubyTplBytes() ([]byte, error) {
	return bindataRead(
		_libRuntimeTemplateRshRubyTpl,
		"lib/runtime/template/rsh/ruby.tpl",
	)
}

func libRuntimeTemplateRshRubyTpl() (*asset, error) {
	bytes, err := libRuntimeTemplateRshRubyTplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "lib/runtime/template/rsh/ruby.tpl", size: 139, mode: os.FileMode(438), modTime: time.Unix(1616234261, 0)}
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
	"lib/runtime/config.example.yml":       libRuntimeConfigExampleYml,
	"lib/runtime/template/rsh/bash.tpl":    libRuntimeTemplateRshBashTpl,
	"lib/runtime/template/rsh/go.tpl":      libRuntimeTemplateRshGoTpl,
	"lib/runtime/template/rsh/lua.tpl":     libRuntimeTemplateRshLuaTpl,
	"lib/runtime/template/rsh/nc.tpl":      libRuntimeTemplateRshNcTpl,
	"lib/runtime/template/rsh/perl.tpl":    libRuntimeTemplateRshPerlTpl,
	"lib/runtime/template/rsh/php.tpl":     libRuntimeTemplateRshPhpTpl,
	"lib/runtime/template/rsh/python.tpl":  libRuntimeTemplateRshPythonTpl,
	"lib/runtime/template/rsh/python2.tpl": libRuntimeTemplateRshPython2Tpl,
	"lib/runtime/template/rsh/python3.tpl": libRuntimeTemplateRshPython3Tpl,
	"lib/runtime/template/rsh/ruby.tpl":    libRuntimeTemplateRshRubyTpl,
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
	"lib": &bintree{nil, map[string]*bintree{
		"runtime": &bintree{nil, map[string]*bintree{
			"config.example.yml": &bintree{libRuntimeConfigExampleYml, map[string]*bintree{}},
			"template": &bintree{nil, map[string]*bintree{
				"rsh": &bintree{nil, map[string]*bintree{
					"bash.tpl":    &bintree{libRuntimeTemplateRshBashTpl, map[string]*bintree{}},
					"go.tpl":      &bintree{libRuntimeTemplateRshGoTpl, map[string]*bintree{}},
					"lua.tpl":     &bintree{libRuntimeTemplateRshLuaTpl, map[string]*bintree{}},
					"nc.tpl":      &bintree{libRuntimeTemplateRshNcTpl, map[string]*bintree{}},
					"perl.tpl":    &bintree{libRuntimeTemplateRshPerlTpl, map[string]*bintree{}},
					"php.tpl":     &bintree{libRuntimeTemplateRshPhpTpl, map[string]*bintree{}},
					"python.tpl":  &bintree{libRuntimeTemplateRshPythonTpl, map[string]*bintree{}},
					"python2.tpl": &bintree{libRuntimeTemplateRshPython2Tpl, map[string]*bintree{}},
					"python3.tpl": &bintree{libRuntimeTemplateRshPython3Tpl, map[string]*bintree{}},
					"ruby.tpl":    &bintree{libRuntimeTemplateRshRubyTpl, map[string]*bintree{}},
				}},
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