// Code generated by vfsgen; DO NOT EDIT.

// +build !dev

package assets

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	pathpkg "path"
	"time"
)

// FileSystem contains project assets.
var FileSystem = func() http.FileSystem {
	fs := vfsgen۰FS{
		"/": &vfsgen۰DirInfo{
			name:    "/",
			modTime: time.Date(2019, 11, 14, 12, 3, 4, 361309808, time.UTC),
		},
		"/.vfshash-assets.json": &vfsgen۰FileInfo{
			name:    ".vfshash-assets.json",
			modTime: time.Time{},
			content: []byte("\x7b\x22\x2f\x66\x61\x76\x69\x63\x6f\x6e\x2e\x69\x63\x6f\x22\x3a\x22\x2f\x66\x61\x76\x69\x63\x6f\x6e\x2d\x61\x48\x38\x36\x33\x71\x33\x4a\x46\x43\x2d\x36\x75\x4a\x30\x50\x2e\x69\x63\x6f\x22\x7d"),
		},
		"/favicon-aH863q3JFC-6uJ0P.ico": &vfsgen۰FileInfo{
			name:    "favicon-aH863q3JFC-6uJ0P.ico",
			modTime: time.Date(2019, 11, 14, 10, 21, 59, 186664000, time.UTC),
			content: []byte("\x89\x50\x4e\x47\x0d\x0a\x1a\x0a\x00\x00\x00\x0d\x49\x48\x44\x52\x00\x00\x00\x60\x00\x00\x00\x60\x08\x06\x00\x00\x00\xe2\x98\x77\x38\x00\x00\x00\x06\x62\x4b\x47\x44\x00\xff\x00\xff\x00\xff\xa0\xbd\xa7\x93\x00\x00\x08\x8f\x49\x44\x41\x54\x78\x9c\xed\x9d\x6f\x8c\x15\x57\x15\xc0\x7f\xe7\xbe\xd9\x5d\x54\x28\xa1\x26\xab\x6d\xda\x42\x6a\x8d\x4d\x81\x6a\x4a\x59\xb0\x0b\x71\xed\xb6\x04\x97\x55\xb0\x2d\xd8\x0a\x4d\x6a\x48\x0d\x98\x18\x63\x62\xa2\x7e\xa8\x42\xe4\x4b\x13\x13\x9b\x90\x34\x95\x0a\xad\x29\x58\x6d\xab\x5d\xdb\x2d\x56\x59\xc8\x43\x21\x20\x42\xa5\x02\x95\x10\xd2\x5d\x2a\xad\x42\xb1\x22\x60\x85\xb7\xfb\xee\xf1\xc3\x7b\xf3\x76\xf6\xb1\xef\xdf\xee\x9d\xf7\x98\x71\x7e\xc9\x24\xef\xcd\x9c\x39\xf7\xdc\x7b\xe7\xdc\x7b\xee\x9d\x3b\x33\x90\x90\x90\x90\x90\x90\x90\x90\x90\x90\x90\x90\x90\x50\x57\xa4\xd1\x06\x5c\xf7\x75\x6d\x33\x96\x1f\x01\xb7\x03\xcd\x15\xc4\x33\xc0\x7e\x6b\xf8\xe6\xc9\xf5\xb2\x2f\x7c\xeb\xc2\xa7\xa1\x15\x30\xf5\xab\x7a\x8d\x34\x71\x04\x98\x52\xe3\xa9\xff\xd2\x41\xa6\x9f\xd8\x20\x7f\x0f\xc3\xae\x7a\x62\x1a\x99\xb8\xe7\xf1\x05\xa3\x4c\x31\x0a\x35\x6e\x53\xbc\x14\x9f\x6f\xa4\xed\xae\xf0\x1a\x99\xb8\xc0\x35\xa2\xb9\xdf\x4b\xda\xd8\x79\xf3\xf5\x7c\xa6\x9c\xfc\xd1\xbf\xb1\xb3\x67\x5f\x5e\x46\xb8\x36\x74\x03\xeb\x40\xe8\x15\x30\x73\xb5\x4e\xc9\x58\x3a\x15\x6e\x10\x43\x53\xf0\x98\x5a\xda\xfd\xdf\xfd\x6f\x31\xf1\xc3\x96\xbe\x72\xba\xfa\x4f\xe6\xbc\x25\x77\x32\xed\x9f\x58\xad\xdf\x0e\xc1\x64\x54\xc9\x28\xbc\x35\x38\x81\xed\x03\x8f\xc9\xd9\x30\xd2\xf0\x09\xad\x0f\xf8\xd8\x2a\x6d\x6d\x52\xd6\x01\x0f\xc1\xc8\x82\x8f\x10\x83\x28\x9b\x86\x06\x79\xe4\xf8\x53\xf2\x6e\x18\x09\x84\x52\x01\xb7\x3c\xac\xd3\x11\xb6\x02\x37\x84\xa1\xbf\x01\x9c\x10\xa1\xeb\xc8\x8f\xe5\x0d\xd7\x8a\x9d\x57\xc0\xad\xab\xb4\x55\xb3\xfc\x89\x91\x85\xbf\x17\xd8\x43\x2e\x8c\x8c\x02\xcd\xc0\x1d\xc0\x9c\xc0\xbe\x81\x66\x61\xf6\x81\x0d\x72\xc6\x65\x42\xee\xfb\x80\x41\xd6\x19\x29\x14\xfe\x59\x2b\x7c\xf9\xd0\x93\xf2\x1b\xe7\xe9\xd4\x81\x4f\xad\xd4\x45\x2a\x6c\x01\x26\x03\xd3\x86\x94\x1f\x00\xab\x5d\xa6\xe1\xd4\x03\x66\xae\xd6\x29\x5e\x86\x53\xf8\x6d\xbe\xd2\xf5\xe7\x4d\xd1\x2c\x7c\x9f\xdb\x1e\xd6\x6e\xb5\xbc\x9c\xff\x9b\x31\x29\x5a\x0f\x6c\x90\x7f\xbb\xd2\xef\x74\x1c\xd0\x72\x91\x4e\xa3\x34\x19\x05\x63\xd9\x1b\xf5\xc2\x07\x78\xed\x49\xe9\x35\xca\xbe\xfc\xf8\xa3\x59\x2d\x9d\x2e\xf5\x3b\xad\x00\x11\xa6\x16\x06\x4b\xb0\xdb\xa5\xee\x46\x22\xb0\xcb\xcf\x97\x37\xc4\x34\x97\xba\x9d\xf6\x01\x29\xcb\x24\x3f\x4c\x17\xb8\xe0\x52\x77\x23\xf1\x2c\xe7\x03\xf9\xba\xca\xa9\x6e\x97\xca\x0c\x40\xc0\xd2\xb8\x10\x66\xbe\xdc\x56\x80\x0d\xfc\xd1\x92\x62\x91\x23\xcc\x7c\xb9\xf7\x80\x18\x12\x66\xbe\x92\x26\xa8\x0a\x92\x26\xa8\xc1\x44\xab\x09\x4a\x3c\xa0\x26\x9c\x56\x80\x67\x03\x76\xc6\xc8\x03\xc2\xcc\x97\x73\x0f\xd0\xbc\x81\x12\x33\x0f\x08\x2b\x5f\xce\xfb\x80\x38\x7a\x40\x98\xf9\xba\x62\x3c\x60\xd1\xfd\x3a\x3f\x65\x69\x53\x33\x6c\x93\x2a\x83\x58\xf6\xf4\x3e\x2f\x7b\x8a\xe5\xbb\xef\xd7\x0e\x03\x1d\xe3\xb1\x17\xc0\x42\xba\xf7\xe7\x92\x2e\x27\x13\x6b\x0f\x58\xba\x54\xaf\xce\xc0\x73\xe4\x27\xb9\xc4\x16\x09\x08\x2c\x5e\xa6\xaf\x4a\x86\x07\x7a\x7a\x86\x6f\x0f\x7a\x96\x0e\x85\xef\x8f\xd7\xe6\x7c\x01\xa4\xcb\xc9\xc4\xda\x03\x86\xe0\x39\x03\x9d\x15\xc2\xbb\x85\x34\xf1\x2c\xf0\xb9\xd1\xd2\x1a\x0f\xd5\xd8\x19\x29\x0f\x28\x50\x45\xe1\xdc\x77\x9f\xce\x47\x0b\xd3\xbb\x59\x60\x33\xf0\x8f\x80\xc8\x75\xc0\x03\xe4\xca\x60\xe1\xd2\xa5\xfa\xe9\xe7\xf3\xcd\x91\xb1\xe5\xaf\xda\xaa\xd1\xca\x7a\x62\x3b\x0e\x30\x30\xc7\x97\x17\x78\xe6\x17\xbf\x94\xaf\x14\xcb\x7c\xe9\x5e\xb5\x0a\x0f\x02\x60\xb9\x83\xdc\xad\x4d\x5e\x78\x41\xd2\x54\x68\x3a\x5c\x11\x99\x71\x40\xad\x57\x8a\xb1\x4c\x0c\xfc\x3d\x31\x9a\x8c\x28\x6f\x06\xf2\x3c\x69\x6c\x96\x8d\x8f\x38\x7b\x40\x45\xf9\x52\x32\x2b\x96\x68\x87\x35\xe3\x8f\x82\x8c\x25\xbd\xb9\xa7\x72\x14\x14\x57\x0f\xa8\x28\x5f\x46\xa6\xc3\xd8\xf1\x47\x41\x79\xd2\xe5\x0e\x46\xcb\x03\x1c\xcb\x97\x92\xa9\xe7\xd4\x77\x64\xa6\xa3\x3d\x6a\x0b\xd7\x82\xf2\x40\xfb\x43\x8b\x2f\x5f\x6a\x28\x3a\xbc\x7c\x31\xa8\xd3\x40\xda\x45\xc1\xd8\x2a\x3a\xf2\x5a\xf3\x55\x0b\x0d\x1d\x88\x05\xe5\x81\xbb\xf2\x5b\x49\x82\x3a\x9f\xee\xa9\x63\x14\x14\xd7\x81\x58\xad\x83\xa9\x46\x4d\xf0\x45\x6a\x20\x36\x66\x0f\x50\x76\x63\xd8\x75\x99\x90\x65\x1e\x92\x6b\x86\x82\x3a\x57\x75\x6b\x07\x0e\xe6\x82\x80\xf4\x13\xbd\x15\xa2\xa0\xff\x07\x0f\x10\xa1\xef\xf1\x97\x65\x4d\xb1\xcc\xd7\xba\x75\x8d\xe6\xfb\x81\xa2\x3e\xa0\x43\x75\xfc\x51\x50\x5e\x67\xba\x06\x3b\x9d\x72\xc5\x78\x40\x29\xf9\x52\x32\x45\xfd\xc7\x98\x71\x65\xe7\x58\x69\xe8\x40\x4c\xb2\xec\x4a\x19\x1e\x05\xb0\x76\x94\xe6\xa7\x8c\x8c\x67\x49\x5b\x07\x61\x50\x35\x73\x4a\xb1\x1d\x88\xad\xff\xad\xf4\x41\xf9\xa7\x62\x4a\xc9\x3c\xf6\x6a\x7d\xa3\xa0\x02\x71\xf2\x80\xa8\x10\x5b\x0f\xf8\xd6\x42\xed\xc0\xd6\x10\xc9\x18\xd2\x3f\xcc\x5d\xf9\xb5\x9f\x5b\x85\xce\x92\x22\x71\xf5\x00\x63\xe9\xa0\x96\x48\x26\x57\x10\xe9\x31\x9d\x5b\x85\xce\x52\xc4\xd6\x03\x4c\xf1\xed\xc7\x4a\x14\x45\x41\x4e\x70\x34\x69\x38\x56\x1a\xea\x01\x1e\xa4\x6b\x29\x47\x13\xb8\x52\x6b\x3d\xb7\x1a\x9d\x65\x64\xa2\xe1\x01\xb5\x2e\x60\x5a\xd7\x37\xf6\x48\x66\x3c\xe7\xd6\x4a\xb4\x16\x66\xe5\x7f\xc7\x28\x08\x0a\x35\x5f\x4e\x2b\x40\x08\x5c\x21\x31\xaa\x81\x30\xf3\x95\xac\x8e\xae\x82\xd8\x76\xc2\x51\x21\x32\x9d\x70\xe2\x01\xb5\xe3\xdc\x03\xfc\xb6\x52\x63\xe6\x01\x61\xe5\xcb\xb9\x07\xf8\xf6\xb9\x58\x36\x78\xa5\x10\x66\xbe\x12\x0f\xa8\x82\xc4\x03\x1a\x4c\xe2\x01\x0d\x26\xf1\x80\x06\x13\x19\x0f\xf0\x20\x96\xe3\x80\x30\xf3\x95\x8c\x03\xaa\x20\x52\xe3\x80\x72\x57\xca\xc6\x39\xba\x40\x2d\xb7\x5a\x21\xe5\x32\x5d\x57\x18\x25\x2b\x86\xbf\xac\xfc\xa3\xfc\x6e\xc4\x7e\x88\xbe\x07\x6c\x9c\xa5\x9b\xc9\xb2\x5c\x80\xd4\x95\xec\x1d\x59\xd8\x38\x4b\xb7\xac\x3c\x20\x2b\xfc\x5d\x61\x7a\x80\xd3\x85\xbf\xa6\x68\xf3\x79\x7a\x96\x2e\x30\xc2\xf2\xe2\xe3\x57\xec\x26\x2c\x7f\xea\x76\xbd\xbb\x52\xbe\x5c\x50\xaf\x26\x68\x96\x19\xbe\x72\x76\x43\x61\x7d\xcf\x3c\xc8\xad\x7a\xcb\x7e\x90\x43\x17\xa6\xf1\x5e\x39\xfd\x13\x07\xb8\x3a\xf5\x3e\x33\x47\xd1\xe3\x8a\x82\x3d\x9a\x7b\x99\xf8\x36\x88\x58\x13\x34\x5a\xb8\x96\x52\x5a\x0a\xfb\xa1\xef\xc1\xd7\x72\x4b\x10\x9f\xb9\x4d\xd7\x48\x3e\xc3\x17\x3e\xce\x7b\xef\xac\x2c\xff\xea\xe2\xeb\x37\xb2\x73\xf2\xc1\xcb\xf5\xb8\x22\x68\x8f\x42\x8b\xbf\x3f\x32\x61\x68\xa9\x01\x4b\x35\xfb\x3d\x60\x42\x85\x9b\xbc\x1e\x60\x42\x1c\xe8\xd5\x6a\xbf\x0b\xea\xe2\x01\x9e\xe5\x52\x60\x15\x74\xfb\xb3\x9f\xcc\x3d\x88\x21\x81\x77\x47\x7b\x16\x3e\x50\xa9\x02\xec\x70\x05\x04\xf5\xb8\x22\x68\x8f\x28\x97\xfc\xdf\x91\xf7\x00\x81\xfd\x81\x9b\xd9\xc3\x0f\x62\x04\x2b\x89\xda\x3c\x80\x2a\x1e\xe8\xa8\x99\x60\xe1\x0a\xfb\xfd\x9f\x61\x7a\x80\xd3\x4e\xdd\x53\xce\x17\x5e\x5b\x19\x78\x04\x75\xd9\x41\xd9\x96\x82\x2d\xe5\xbe\x09\xe0\x65\x73\x15\x50\x6e\xf3\xb2\xa5\xcf\x77\xb9\xa5\x60\xcb\xb2\x83\xb2\xcd\xb7\x3f\xa5\x5c\x55\x38\xa6\x9c\x73\x5a\x66\x2e\x95\x19\x65\x20\xf8\xcc\x57\xf0\xd8\xbd\xaf\xcb\x8a\x5f\xcd\xd0\x9f\xaa\x30\x57\xe0\x43\x00\x08\xf3\xc8\xaf\xfd\xf7\xa8\xa2\x09\x22\xe0\x01\xc2\x6e\xd4\x6d\x14\xa4\xf0\x1f\x51\xf6\xde\x73\x78\xb8\xf0\x01\x4c\x96\x76\xbf\x0d\x52\x4b\xbf\xcb\x34\xdd\xbe\x3b\xda\xb0\x3d\x35\xc4\x20\xb9\x57\x17\xcf\xed\x99\xa9\x5d\x4b\x0e\xc9\x56\xff\x70\x3e\x63\x85\xcc\xfd\x7a\x86\xae\xc1\xaf\xa8\xec\x28\x2f\xea\x28\xa2\x29\x3b\xa2\x0f\xe8\x5b\x7c\xd8\x6d\x14\x34\x1a\x2f\xcd\xd0\x6e\x85\xb6\x7c\xf3\x94\xb9\xd4\xc2\x0e\x97\xfa\x9d\x36\x41\x5f\x3c\x28\x67\x53\xb0\xa9\xe0\xae\x96\x9f\xf5\xce\xd0\x45\xe5\x12\xf7\x65\x9b\xc8\x79\x40\xb9\xcd\x0b\x36\x15\x2e\x0d\x2f\xc1\x4b\x33\xb4\x5b\x94\xcd\x81\xe6\xe7\x27\xcb\x0e\xb8\x7b\x6f\x34\x84\xf0\xf6\x74\x73\x91\x47\xb4\x99\x85\xc0\x54\x60\x32\x4a\xef\x2b\xb7\xe8\x3e\x94\xdd\x6a\x38\x15\x94\x0d\x46\x1d\xc6\xe6\xa2\x9c\x72\x0c\x29\xd8\x40\x14\xd4\x3b\x3d\x9c\x2f\x68\x18\xf8\xa8\x2a\xed\x58\x66\x0f\x1b\x4b\xff\xa5\x2c\xdf\x73\x9d\x96\xf3\x0a\xe8\x3a\x2e\xef\x6e\x9d\xae\x5d\x62\x79\x05\x0a\xef\x59\x6e\x03\xda\x28\x53\xc0\xd5\xf4\x01\xff\xb5\xb9\x4a\xc8\x73\x17\xea\x38\x0a\x0a\x50\x14\xec\xf4\x67\x0d\x8b\xee\x39\x2a\xff\x74\x9d\x4e\x28\x9e\xdc\x75\x44\xde\x98\x90\x65\x76\x0a\x9e\x30\x4a\xa6\x9a\xc8\x23\x18\x05\x35\xbd\x0d\xb6\x37\xb7\x79\x27\x87\xf7\xa7\x14\xa9\x47\x14\x14\xd8\x32\xc6\xf2\x78\x93\x32\xbb\xfb\xb0\xfc\x35\x8c\xb2\x0a\xed\x23\x3e\x9f\x3d\x26\x67\x80\xd5\xdb\x6e\xd4\xef\xd0\x4c\x27\x96\xa9\x22\x7c\x24\x28\xa3\x81\xb9\x17\x0f\x98\x30\x04\x17\x76\xc0\xf9\x1d\xf8\xeb\xf6\xb9\xf4\x7b\x98\x74\x27\x4c\xec\x84\xf7\x2d\x64\x03\x73\x4a\xe2\x7e\x2e\x28\x67\x97\x72\x4a\x61\x40\x86\xd8\x71\xf7\x9b\x6e\xdb\xfc\x62\x1a\x7a\xdf\x6a\xfb\xcd\xba\x46\xf2\x0f\x59\xb4\x4c\x65\xaf\x0e\x72\x63\xe6\x1d\x5a\x47\x93\x6d\xbe\x96\xd3\xa6\x89\xfe\x8b\x27\x72\x9f\x15\x51\x61\x6d\xe7\xd1\xf0\xa3\xa0\xb0\x69\xe8\x77\xc4\x8c\xe5\x8c\x7f\x05\x0c\x0e\x30\x17\x46\xb4\x89\x7e\xb8\x77\x27\xc0\xd0\xdb\xb4\x02\xad\xfe\x71\x51\x42\xf9\xaa\x51\xbd\xa9\xe7\x4b\x47\x2e\xa3\x05\x5e\x34\xca\xb9\xa2\x76\x37\x2b\xb0\xf6\xf4\x31\x16\x9c\x3e\xc6\x02\x81\xb5\x46\xc9\x16\xc9\x9c\x4b\x09\x2f\x36\xd2\x76\x57\x34\xfc\xd6\xf9\xae\x9b\x74\x2e\xb0\x1e\x98\x0e\xbc\xae\xf0\xdd\xf9\xc7\x47\x3e\x34\xf7\x87\x9b\xb4\x03\x78\x54\x60\xa6\xc2\x21\x81\x6f\xcc\x3b\x2e\x7b\xeb\x6f\x6d\x42\x42\x42\x42\x42\x42\x42\x42\x42\x42\x42\x42\x42\xe4\xf9\x1f\x4f\x80\xc5\x86\x19\x37\x68\x02\x00\x00\x00\x00\x49\x45\x4e\x44\xae\x42\x60\x82"),
		},
	}
	fs["/"].(*vfsgen۰DirInfo).entries = []os.FileInfo{
		fs["/.vfshash-assets.json"].(os.FileInfo),
		fs["/favicon-aH863q3JFC-6uJ0P.ico"].(os.FileInfo),
	}

	return fs
}()

type vfsgen۰FS map[string]interface{}

func (fs vfsgen۰FS) Open(path string) (http.File, error) {
	path = pathpkg.Clean("/" + path)
	f, ok := fs[path]
	if !ok {
		return nil, &os.PathError{Op: "open", Path: path, Err: os.ErrNotExist}
	}

	switch f := f.(type) {
	case *vfsgen۰FileInfo:
		return &vfsgen۰File{
			vfsgen۰FileInfo: f,
			Reader:          bytes.NewReader(f.content),
		}, nil
	case *vfsgen۰DirInfo:
		return &vfsgen۰Dir{
			vfsgen۰DirInfo: f,
		}, nil
	default:
		// This should never happen because we generate only the above types.
		panic(fmt.Sprintf("unexpected type %T", f))
	}
}

// We already imported "compress/gzip" and "io/ioutil", but ended up not using them. Avoid unused import error.
var _ = gzip.Reader{}
var _ = ioutil.Discard

// vfsgen۰FileInfo is a static definition of an uncompressed file (because it's not worth gzip compressing).
type vfsgen۰FileInfo struct {
	name    string
	modTime time.Time
	content []byte
}

func (f *vfsgen۰FileInfo) Readdir(count int) ([]os.FileInfo, error) {
	return nil, fmt.Errorf("cannot Readdir from file %s", f.name)
}
func (f *vfsgen۰FileInfo) Stat() (os.FileInfo, error) { return f, nil }

func (f *vfsgen۰FileInfo) NotWorthGzipCompressing() {}

func (f *vfsgen۰FileInfo) Name() string       { return f.name }
func (f *vfsgen۰FileInfo) Size() int64        { return int64(len(f.content)) }
func (f *vfsgen۰FileInfo) Mode() os.FileMode  { return 0444 }
func (f *vfsgen۰FileInfo) ModTime() time.Time { return f.modTime }
func (f *vfsgen۰FileInfo) IsDir() bool        { return false }
func (f *vfsgen۰FileInfo) Sys() interface{}   { return nil }

// vfsgen۰File is an opened file instance.
type vfsgen۰File struct {
	*vfsgen۰FileInfo
	*bytes.Reader
}

func (f *vfsgen۰File) Close() error {
	return nil
}

// vfsgen۰DirInfo is a static definition of a directory.
type vfsgen۰DirInfo struct {
	name    string
	modTime time.Time
	entries []os.FileInfo
}

func (d *vfsgen۰DirInfo) Read([]byte) (int, error) {
	return 0, fmt.Errorf("cannot Read from directory %s", d.name)
}
func (d *vfsgen۰DirInfo) Close() error               { return nil }
func (d *vfsgen۰DirInfo) Stat() (os.FileInfo, error) { return d, nil }

func (d *vfsgen۰DirInfo) Name() string       { return d.name }
func (d *vfsgen۰DirInfo) Size() int64        { return 0 }
func (d *vfsgen۰DirInfo) Mode() os.FileMode  { return 0755 | os.ModeDir }
func (d *vfsgen۰DirInfo) ModTime() time.Time { return d.modTime }
func (d *vfsgen۰DirInfo) IsDir() bool        { return true }
func (d *vfsgen۰DirInfo) Sys() interface{}   { return nil }

// vfsgen۰Dir is an opened dir instance.
type vfsgen۰Dir struct {
	*vfsgen۰DirInfo
	pos int // Position within entries for Seek and Readdir.
}

func (d *vfsgen۰Dir) Seek(offset int64, whence int) (int64, error) {
	if offset == 0 && whence == io.SeekStart {
		d.pos = 0
		return 0, nil
	}
	return 0, fmt.Errorf("unsupported Seek in directory %s", d.name)
}

func (d *vfsgen۰Dir) Readdir(count int) ([]os.FileInfo, error) {
	if d.pos >= len(d.entries) && count > 0 {
		return nil, io.EOF
	}
	if count <= 0 || count > len(d.entries)-d.pos {
		count = len(d.entries) - d.pos
	}
	e := d.entries[d.pos : d.pos+count]
	d.pos += count
	return e, nil
}