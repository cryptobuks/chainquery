// Code generated by go-bindata.
// sources:
// migration/000_init_schema.sql
// DO NOT EDIT!

package migration

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

var _migration000_init_schemaSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xec\x5b\x5b\x73\xdb\xba\x11\x7e\xf7\xaf\xc0\x9c\x97\x50\x53\x7b\xc6\xca\x49\x4f\xd3\xe9\xe4\x81\xa6\x60\x9b\x27\x12\xe5\x52\x54\x9a\xf4\x05\xa2\x48\x48\x42\x4d\x81\x0a\x01\xc6\x76\x7f\x7d\x87\x00\x2f\xe0\x4d\xa2\x24\xa6\x27\xce\x9c\xb7\x58\x5c\x2c\x6e\xbb\xdf\x7e\xbb\x8b\x5c\x5d\x81\xbf\x6c\xc9\x3a\x72\x39\x06\xf3\xdd\xc5\x85\xfa\xf7\x8c\xbb\x1c\x6f\x31\xe5\x37\x78\x4d\xe8\x85\x61\x43\xdd\x81\xc0\xd1\x6f\xc6\x10\x98\xb7\xc0\x9a\x3a\x00\x7e\x36\x67\xce\x0c\x2c\x96\x41\xe8\x3d\x2e\x2e\xb4\x0b\x00\x00\x58\x10\x7f\x01\x66\xd0\x36\xf5\xf1\xe5\x85\xfc\x65\x49\x38\x5b\x80\x4f\xba\x6d\xdc\xeb\xb6\xf6\xf6\x7a\x20\x46\x5b\xf3\xf1\xf8\x52\x0a\x78\x1b\x97\xd0\xa7\x30\x7a\x2c\xa4\xfe\x76\x3d\x00\xc9\x3f\x74\xc3\x81\x36\x98\x41\x07\x04\x2e\x27\x74\x08\x8c\xe9\x78\x9c\xac\x44\xfe\x89\xd6\x98\xe2\xc8\x0d\x90\x47\x6a\x4a\x43\xba\x22\xd1\xd6\xe5\x24\xa4\x6c\x01\x4c\xcb\x81\x77\xd0\x06\x73\x6b\x66\xde\x59\x70\x54\x15\xf7\xc9\x6a\x45\xbc\x38\xe0\x2f\x0b\x30\x9a\xce\x6f\xc6\x50\x1b\xbe\xbf\x7c\x5f\x5b\xeb\xc6\x65\x9b\x5e\x96\x09\x32\x85\x98\xac\x37\x7c\x01\x6e\xcc\x3b\xd3\x72\x5a\xd7\xb7\xc5\x3e\x71\x29\xe2\x64\x8b\x3b\xc8\x46\x8f\x01\x46\x51\x18\xf2\x3e\x4f\x94\xba\x5b\x8c\xbc\xc0\x25\xdb\xfe\x55\x87\xd4\x3b\xbc\xaf\x5d\x84\xbf\x91\x30\x66\x48\x58\x1c\x3a\xf7\x2a\xb2\xb9\xf1\x33\xef\x57\xa3\x54\xc6\xc8\x7f\x0f\x6f\x89\xbb\xd1\x1a\xf7\x7a\x94\x72\xf2\x4e\x76\xf2\x0d\x47\x8c\x84\xb4\xab\x1c\xda\xe0\xe7\x62\xa5\xc3\xb3\x57\xca\x23\x97\x32\xd7\xe3\x42\xb5\xcb\x36\x98\x2d\x80\x03\x3f\x3b\xf5\xcf\x0c\xed\xa2\xd0\xc3\x8c\x61\x7f\x01\x1c\xd3\xfa\x62\x5a\x8e\x36\x1c\x80\x11\xbc\xd5\xe7\x63\x07\x5c\x2b\xaa\x53\xef\x8f\xb0\xcb\x13\xe9\x91\xee\x40\xc7\x9c\xc0\xc2\xef\xb2\x41\xc6\xdc\xb6\xa1\xe5\xa0\xe4\xeb\xcc\xd1\x27\x0f\x99\xf7\x84\x3e\x59\x91\x23\xc7\x82\xa9\x05\xe6\x0f\xc9\x80\x26\xbd\x42\xf1\x83\x6d\x4e\x74\xfb\x0b\xf8\x08\xbf\x80\xc5\xc3\x47\x74\x23\x60\x13\x68\x09\x62\x0e\xe4\xdc\x73\xcb\xfc\xe7\x1c\x4a\x09\xd3\x7f\x96\x22\xf7\xc2\x28\x35\x89\x3c\xa9\xa0\x31\xb5\x66\x8e\xad\x27\x97\xb6\x30\x28\x47\x4e\x71\x56\xf7\xe2\x24\x3f\xb9\x01\xf1\x7f\x67\xc9\xe5\x1a\xf7\xd0\xf8\xa8\x35\x9e\xb6\x39\x93\xdb\x9a\xda\xe0\xf7\xd9\xd4\x42\x9f\xf4\xb1\x39\x6a\x14\x1d\xa4\x13\x9b\xd6\x08\x7e\x56\x17\x97\xa2\x97\x96\xe1\x58\x8b\x9c\x23\xec\x51\x53\xad\xb3\x2e\x39\x11\x18\x97\x89\xaa\x88\x57\x97\x7d\x48\xb1\xa0\x74\x44\x4d\x00\xd1\xb2\x20\x23\x33\x10\x2d\xb7\x95\x16\xc9\x49\x6e\x0f\x5a\x61\x1b\x83\x8b\x01\x80\xd6\x9d\x69\xc1\x0f\x26\xa5\xe1\xe8\xa6\x30\x8d\x7b\xdd\x9e\x41\xe7\x43\xcc\x57\xef\xb7\xcb\x77\xb9\x3b\xa4\x7f\xa3\x98\x12\x2f\xf4\x71\xe2\x0f\xf6\xf4\x5f\xe8\x76\x6a\x4f\x74\xe7\x83\x31\x9d\x3c\xd8\x70\x36\x83\xa3\xe4\xf6\xd1\xcd\x78\x6a\x7c\x44\x33\xf3\xdf\xf0\xc3\xbb\x7f\x34\xc7\x65\x48\xfd\x93\x23\xb6\x72\xc1\x8d\x71\x5b\x01\x92\xe5\x8b\x38\x46\x94\x7c\x3e\x1b\x18\x09\xdd\xc5\x1c\x79\x61\x4c\xf9\xe1\x88\x1c\xc6\xfc\x08\xe9\x6f\x6e\x10\xe3\xfd\xa1\x7b\x85\xab\x02\x4d\xf8\x51\x83\xa6\x46\x28\x6d\x90\xeb\x84\xf7\xbd\xd1\x87\x2a\x84\x67\xe7\x53\xf9\xae\x04\x83\x43\x27\x18\xb9\x4f\x25\xfc\xfd\xb1\x20\x54\x5d\x52\xba\xa1\xa3\xd7\x55\x45\x60\x05\x35\xcb\x38\x7c\x3b\xb5\xa1\x79\x67\x49\xc1\xdb\x92\x60\x09\x6f\x6a\x2e\x32\x00\x36\xbc\x85\x36\xb4\x0c\x98\x13\xe3\x1c\xba\x93\x0d\x8e\xe0\x18\x26\x1b\xd4\x67\x86\x3e\x82\xca\x96\xad\x29\xd0\x0d\xc7\x9c\x5a\xa0\x39\x14\x54\x10\xbe\x1a\x10\x14\xd0\x52\x24\x33\x28\xad\x99\xf3\xde\x51\x29\x34\x66\x83\x4b\x87\xde\x65\xe0\x7e\x4c\x55\xe4\x7f\x42\x64\x75\x7d\x3f\xc2\x8c\xb5\xa3\x6a\x26\x90\x63\xc0\xbb\xe3\x31\x20\xb5\x8d\x2a\xbc\x91\x88\x71\xc4\x30\xa6\x85\x6b\xe4\x2c\x73\x5d\x4c\xf8\xeb\xf5\x20\xd5\x50\x7c\x45\x71\x14\xa8\x29\xda\xf5\xe0\x87\x05\x81\xaa\x13\xeb\xd9\x81\xee\x23\x52\xa9\x90\x22\x9b\xdd\x43\x66\xa0\xd4\xc7\xcf\x25\x59\x27\x39\x33\x4d\x1c\x5d\xdd\x88\x53\x99\x4e\x06\x9f\xca\xfe\x84\xc6\x2e\xc2\x79\xbb\xa9\xab\xa8\x93\x7c\x3c\x94\x0a\x55\x58\x67\x9f\x61\x52\x12\x8f\xf4\xce\x9b\x16\x93\xc9\x31\xe4\x85\x84\x2e\x5d\x86\x0f\x26\x1a\x69\x95\x21\x93\x3e\x9b\x1c\x25\xe4\x35\x8c\x79\x5f\x49\x68\xa6\x8e\xd6\x23\x7f\x45\x82\xed\x30\xf5\x51\xbc\xf3\xa5\x2d\x77\xd9\x36\xc3\x5f\x63\x2c\x52\xf6\x13\x78\x59\xa6\xc3\x8b\xc8\x8e\x23\x46\xd6\xc8\x65\x5b\x49\x3d\x4e\xdb\xaa\xa2\x49\xa4\xa8\x27\x69\x7a\x2d\x78\x67\x0a\xa7\xdb\x4b\x57\x84\x88\x8a\x8b\x55\xeb\x2f\xd3\x14\xb7\x8c\xa0\xdd\x68\xca\x9e\x99\x2b\xb4\xaa\x02\x03\xe5\xb9\x79\x9d\x82\x1d\x33\xbf\x02\xb3\x62\xea\x4f\xd2\xda\xb4\xd4\xec\x9a\xd3\xc6\x30\xe6\x6a\xc2\x98\xfb\x5c\x5d\x5a\xe8\xec\x04\xf2\x42\xf2\x67\x85\x78\x54\x63\x35\xe2\xd7\x0e\x98\xbe\x07\x71\x2b\xa2\x8d\x66\x5e\xb3\x61\xe2\x2f\x2e\x4b\x5a\x9b\x3c\x20\xbf\x8f\x2c\xde\x63\x56\x38\x4d\xa6\xa6\x6c\x86\x44\x75\xaa\x93\x1d\xa0\x61\xe2\x3a\xe3\xe8\xc9\xff\x5e\xb5\x55\xc9\xd4\xfe\x55\x30\x87\xf6\xf8\x95\x00\xc7\xe1\xf8\xc7\x5f\x76\xb8\xdc\xfc\x38\x27\xc4\xed\xe2\x25\x7a\xc4\x2f\x3d\x05\xcc\x4c\xdb\xe9\x41\x53\x16\x0e\xf0\xd7\x98\x44\xd8\x4f\x02\x30\x75\x79\x1c\xe1\x86\x7e\x8b\x52\xff\x18\xfe\x76\x5d\x9c\xc8\x5f\x4f\x3d\x91\xcc\x9d\x02\xc2\xf8\x39\xab\x27\x4c\x10\x20\xde\x91\xf8\x24\xa2\x49\xc6\xdf\x0a\x81\xaf\x26\x71\x9a\x4a\x27\x6c\xcd\x9b\xe6\xaa\x48\xcd\xc5\x2e\xa5\xfd\xb7\x30\x10\x39\xf0\xff\x48\x04\x9a\x17\x30\x4b\x6e\xeb\xe6\x25\x47\xff\xfa\xed\xf5\x14\x06\xaa\xf5\xf8\x1c\xff\xeb\x75\xf8\xb2\xdd\xb6\x54\xe0\x4b\x42\x0d\xb5\x77\xb9\xbd\x83\x6c\x67\x1a\xef\x42\x42\xc5\x8e\xc4\x6d\x5d\xd6\xef\x71\x00\x8c\xe9\x64\x02\x2d\x07\xbc\xa1\x18\xfb\xd8\x07\xab\x30\x02\x11\x5e\xe1\x28\x21\xf8\x0c\x10\x0a\xf8\x86\x30\xe0\x85\x41\xbc\xa5\x20\x8c\x7c\x1c\xbd\x69\x9a\x89\x77\xa4\x4a\x72\xed\x3f\x21\x57\x4a\x0b\xd6\x55\xb2\x94\xfe\xfc\x5d\xd9\x92\x3c\x53\x85\x6c\x14\x93\x76\xe3\x4b\x52\x41\xc1\x5b\x0a\xcf\x2f\x34\x95\x7d\x25\x2c\xc1\xc7\x59\x9c\xa9\x3a\xf9\x9f\xa4\xe9\x40\xd3\xa6\x66\x63\x47\x12\xa6\xee\x86\xb6\xf0\xf1\x92\x70\xe4\x6e\x65\x17\xe6\x40\xf7\xa4\xc0\x92\x59\xbc\x05\xe1\x0a\xf0\x0d\x06\x02\x54\x19\xe0\xa1\xc4\x91\x74\x6e\x01\x33\xc9\x67\xfe\xfc\xa6\x08\x99\xfe\x99\x73\x49\xa3\xec\x32\x59\xe0\x72\xcc\x38\xaa\x37\x7a\xfa\xec\x2a\x28\x86\x5c\xb9\xa1\x6e\x5e\xa9\xa8\x2a\xbc\xe3\x8f\x8a\xaa\xad\x0b\xfa\x6e\xee\xba\xaf\x77\x50\x4c\x3f\x16\x37\xd9\xd0\xf5\x68\xbb\xe2\xbd\x3d\x89\x42\xef\x28\xb1\xfc\x44\x4f\xc9\x05\x3a\x0e\x36\x84\x2d\xa7\x81\x50\xb1\xea\xd7\x1d\xe1\xc4\xe3\x9f\x6e\x69\x5b\xaf\x7d\xe3\x6e\xa9\x16\x75\xb7\x58\x7d\x9f\xf2\xf6\x5d\xfd\xa5\x99\x78\xbd\x94\xac\xe9\xd4\xde\x4b\xa3\x42\x99\xe4\x29\xf9\x43\x2e\x05\xae\xae\xc0\x10\x5c\x01\x03\x47\x9c\xac\x88\xe7\x72\xec\xbc\xec\xf0\x25\x78\x0b\xae\xc0\x8c\x47\xd8\xdd\x26\x7f\x03\xb6\x09\xe3\xc0\x07\x34\xe4\x60\x89\x01\x27\xf4\x85\x50\x0e\x02\xb2\x25\x9c\x01\x97\x81\x65\x18\x06\x80\x7d\x0d\x96\x21\x09\x70\x94\xd6\x6f\xe3\x65\x40\xd8\x06\x47\x67\x6d\x28\x87\x51\x85\xf3\xb9\xc0\x48\xf6\x65\xfa\xe0\x89\xf0\x4d\x75\xf1\x19\x84\x16\xf3\x33\xb2\x2e\x77\x8f\x4e\xbb\x66\xaf\x98\xa7\xd4\x89\x66\x7e\x5f\x45\xf1\xae\x7d\xfd\xbc\xad\xde\xdf\x6b\x27\xc1\xd1\x91\xcb\x64\xb6\x3d\x81\x23\x73\x3e\x11\x59\x6b\x9b\xdc\x7f\x44\xd6\x50\x08\xa6\x25\xea\xab\x2b\xa0\xfb\x3e\x49\xb6\xe0\x06\x60\x45\x70\xe0\xcb\xf8\x86\x5d\xf6\x02\x08\xf5\xf1\x33\xa1\xeb\x24\x20\x32\x61\x5e\x20\x31\x4e\x96\xc6\xfd\x98\x6f\xc2\x48\x49\xbe\x87\x6f\x07\x79\x9c\x97\x45\x01\x52\x9d\x34\x6d\x73\x50\x9e\xa4\x4e\xe5\x6a\xc6\xf0\xb7\xb7\xa7\x3e\x04\x61\x88\x22\x86\x56\xe8\xa9\x5b\xda\x1d\xb8\x74\x1d\xbb\xeb\x3e\x2a\x29\x7c\x13\x6f\x97\xd4\x25\x81\xec\x79\xaa\x0f\xce\x08\x0f\xca\x86\x77\xc4\x8b\x91\x15\xc6\xc8\x8b\xa3\xc4\x81\x5e\x52\x67\xfc\x75\x50\xec\x76\x45\x02\x8e\xa3\xae\xed\x95\x25\xf1\x11\xe3\xc2\x0d\xd4\xfd\x66\xf2\xbf\xe8\x9e\x87\x77\x1c\xfb\xbf\xfc\xe0\xef\xde\x9a\xe8\x91\x40\x96\xfd\xfd\x0b\x21\xb2\x87\xe6\xb4\x3e\xb8\xa8\xd0\x9d\x63\x9e\x5d\xec\x59\xc9\x43\x86\x73\xa2\x63\xa0\x82\x6e\x79\x76\x2f\xdb\x59\x1e\x65\x8e\x9a\xbb\xda\xb3\x16\x73\xcf\x29\xf9\x1a\xd7\xde\x70\x28\x07\x70\x99\x46\xc7\x4b\x25\xb8\xb5\x3c\x12\x14\x0a\x0d\x15\x63\xd3\x92\x44\x09\x77\x5b\x2a\x12\xaa\xcc\x60\x20\xe2\x9a\xf2\x93\x00\x99\x2a\x31\x32\xea\xe7\x51\x23\x4f\xd5\xab\x3e\xe2\xc5\x8a\xdc\x4e\x97\x8a\x83\x90\x6c\x2e\x38\xa4\x7e\x53\x15\xd6\x53\xa0\xd4\x32\xc8\xd4\x86\x7f\x1f\x36\xd4\x61\xe4\x1a\x24\x38\x3a\x02\x1b\xb5\x32\x56\xb6\x0c\x19\xe7\x78\xa6\x15\xd8\xd6\x76\x3a\x12\x96\xb4\x14\x9f\xe4\x4a\x5e\x35\x8d\x64\xf1\x6e\x17\x46\x7b\xea\xff\xa9\x00\xf6\x51\xef\x84\x2d\x55\xdd\x35\xcb\x3c\x03\x90\x7b\x6f\x4d\x74\x25\xc3\x8d\xb0\x9b\x9d\x79\x3b\xf0\xae\x1e\x51\x7e\xee\x39\x94\x35\xdc\x44\xaf\xa0\xa7\x18\x7b\x7a\xbe\x9a\x72\xd8\x75\x8f\x28\xaf\x90\xf8\x6d\x6b\xac\x0d\xe4\xed\xe1\xa4\xa5\xf9\x2b\xcf\x59\x2b\xd5\xb9\x95\xcf\x61\xcc\xf3\xe2\x6a\x6b\x75\xfc\x55\xbb\x69\x4c\x1f\x69\xf8\x44\xd1\x81\xac\xef\xa8\xb4\xab\xa0\x8e\xf5\xff\x06\x44\x58\xfa\x08\xa5\x23\x49\xea\xf3\xbf\x6c\xfc\x11\x7e\x7a\x64\x75\xb8\x87\x04\xa2\x09\x18\xd2\x5b\xf6\xea\xb4\xac\xca\x1f\x44\xbb\xa1\xde\xd0\xa8\xcc\xd7\xc2\x1f\xca\x52\x83\x16\xe6\xa7\x2e\x26\xfc\x6e\x25\x68\xc5\x8b\xe7\x62\x42\x11\x67\x1b\x5e\xf6\x16\xc8\xd0\x3c\xa4\xb1\x4a\xde\x2e\xee\x3c\x67\xea\xeb\xf0\xf3\x43\x20\xc5\xff\x02\x00\x00\xff\xff\xbe\x12\x6e\x4a\x11\x38\x00\x00")

func migration000_init_schemaSqlBytes() ([]byte, error) {
	return bindataRead(
		_migration000_init_schemaSql,
		"migration/000_init_schema.sql",
	)
}

func migration000_init_schemaSql() (*asset, error) {
	bytes, err := migration000_init_schemaSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "migration/000_init_schema.sql", size: 14353, mode: os.FileMode(420), modTime: time.Unix(1522719345, 0)}
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
	"migration/000_init_schema.sql": migration000_init_schemaSql,
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
	"migration": &bintree{nil, map[string]*bintree{
		"000_init_schema.sql": &bintree{migration000_init_schemaSql, map[string]*bintree{}},
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

