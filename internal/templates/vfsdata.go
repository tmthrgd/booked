// Code generated by vfsgen; DO NOT EDIT.

// +build !dev

package templates

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

// FileSystem contains project templates.
var FileSystem = func() http.FileSystem {
	fs := vfsgen۰FS{
		"/": &vfsgen۰DirInfo{
			name:    "/",
			modTime: time.Date(2019, 11, 18, 15, 29, 24, 373593514, time.UTC),
		},
		"/error.tmpl": &vfsgen۰CompressedFileInfo{
			name:             "error.tmpl",
			modTime:          time.Date(2019, 11, 17, 3, 46, 24, 770694000, time.UTC),
			uncompressedSize: 563,

			compressedContent: []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x9c\x91\xbf\x4e\xc4\x30\x0c\x87\xf7\x3e\x45\xe8\xc0\xc4\xb5\xba\x0d\x41\x52\x06\x66\x24\x04\xbc\x80\x95\xba\x4d\x74\xce\x1f\xc5\xbe\x2b\xa5\xea\xbb\xa3\xde\x89\x81\x8d\x63\x8a\xf2\x59\x9f\x2d\xfb\xa7\x6f\xfa\x64\x65\xce\xa8\x9c\x04\xea\x2a\xbd\x3d\x8a\x20\x8e\x06\x63\x57\xe9\x80\x02\xca\x3a\x28\x8c\x62\x8e\x32\xec\xee\x7f\x60\x84\x80\xe6\xe4\x71\xca\xa9\x88\xb2\x29\x0a\x46\x31\xf5\xe4\x7b\x71\xa6\xc7\x93\xb7\xb8\x3b\x7f\xee\x7c\xf4\xe2\x81\x76\x6c\x81\xd0\xec\xeb\xae\xd2\xe2\x85\xb0\x5b\x96\xe6\x5d\x40\x8e\xfc\x9c\x7a\x5c\x57\xb5\x2c\x4e\x24\x5f\xd0\x07\x7e\x8a\xfa\x55\xd6\xed\xc5\xaa\x34\xf9\x78\x50\x05\xc9\xb0\xcc\x84\xec\x10\x45\xb9\x82\x83\xa9\x97\x05\x98\x51\x5e\x41\x9c\xaa\xdb\x98\x4a\x00\xf2\x5f\xd8\x04\x1f\x1b\xcb\x5c\xaf\x6b\x7d\x85\xcf\x07\x24\x94\x14\xff\xa9\x63\x29\xa9\xfc\xc9\xdb\xd6\xe6\x87\xb6\x1d\x52\x14\x6e\xc6\x94\x46\x42\xc8\x9e\x1b\x9b\x42\x6b\x99\x9f\x06\x08\x9e\x66\xf3\x06\x84\x13\xcc\xb7\x10\xf2\x63\xef\x39\x13\xcc\x86\x27\xc8\x5b\xf7\x00\x3e\x2a\x4b\xc0\x6c\xb6\x30\xc0\x47\x2c\x5b\x9e\xfb\xab\xcf\xec\xf6\x5d\xa5\xf3\xa6\xbd\x20\x33\x8c\x67\x98\xbb\x4a\xb7\xdb\x8c\xee\x3b\x00\x00\xff\xff\xc0\x80\xef\x61\x33\x02\x00\x00"),
		},
		"/message.tmpl": &vfsgen۰CompressedFileInfo{
			name:             "message.tmpl",
			modTime:          time.Date(2019, 11, 18, 15, 31, 4, 339210225, time.UTC),
			uncompressedSize: 1218,

			compressedContent: []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x9c\x93\x41\x8f\xd3\x30\x10\x85\xef\xf9\x15\xa6\x07\x04\x07\x12\xed\x72\x41\x6c\x62\x04\x2b\x2d\xe2\xb0\x52\xc5\x22\xee\x83\x3d\xad\x4d\x1d\x3b\xf2\x4c\x5b\x05\x2b\xff\x1d\x39\x69\xa1\x69\x2e\x2c\xa7\xb6\x6f\xbe\xf7\xaa\x19\xcf\xd4\x2f\x74\x50\xdc\x77\x28\x0c\xb7\x4e\x16\x75\xfe\x10\x0e\xfc\xb6\x41\x2f\x8b\xba\x45\x06\xa1\x0c\x44\x42\x6e\xf6\xbc\x79\xf3\x4e\x16\x35\x5b\x76\x28\x53\x2a\xbf\xe5\x2f\xc3\x50\x57\x93\x52\xd4\xce\xfa\x9d\x88\xe8\x1a\xe2\xde\x21\x19\x44\x16\x26\xe2\xa6\x59\xa5\x04\x44\xc8\x6b\x60\x23\x56\x95\x0f\xb1\x05\x67\x7f\x61\xd9\x5a\x5f\x2a\xa2\xd5\x30\xac\x9e\xe1\xa7\x1d\x3a\xe4\xe0\xff\xd3\xde\x22\x11\x6c\xf1\x9f\x9c\x86\xb9\xa3\xf7\x55\xb5\x09\x9e\xa9\xdc\x86\xb0\x75\x08\x9d\xa5\x52\x85\xb6\x52\x44\x1f\x36\xd0\x5a\xd7\x37\x5f\xc1\xe1\x11\xfa\x97\xd0\x76\x77\xda\x52\xe7\xa0\x6f\xe8\x08\x5d\x4e\x6f\xc1\x7a\xa1\x1c\x10\x35\x2a\x78\x06\xeb\x31\xe6\x59\xdf\xcc\x86\x68\x6e\x64\x51\x13\x2a\xb6\x21\x8f\xde\xdc\xca\x4f\x40\x56\x89\x27\x06\xa6\xba\x32\xb7\xb2\xa8\x75\x7e\x23\xcd\x72\x0d\x91\xad\xb2\x1d\xf8\x5c\xd2\x9c\x55\x2d\x53\x72\xe8\x45\x79\x59\xcc\xc1\x5a\x4f\xa6\xc7\xa9\xed\xa5\xe1\x5c\xb8\x84\xd7\x26\x70\x98\xa1\xe5\x24\xdd\x87\xbd\xe7\x4b\xf2\xbb\xd5\x78\x45\x4e\xd2\x82\xfc\xb8\xd7\x36\xcc\xc0\x51\x59\x70\x9f\xbf\x3c\xcc\xf3\xb2\xb0\xa0\x1e\xac\x9b\xf7\x52\x8e\xca\x82\x7b\x62\xab\x76\x18\xe7\xe8\x59\x5c\xd0\x6b\x07\xfe\xaa\xed\xac\x2c\x53\x0d\xc4\xab\xbf\x9f\xa4\x05\x79\x0f\xce\xcd\xc1\x51\x39\x71\x29\x1d\x2d\x1b\x51\x3e\x5a\x22\xd4\x97\x15\xf1\x2a\xa5\x72\x18\x44\x3b\x56\x84\xca\xa5\xd7\x29\xa1\xd7\x7f\xd2\xab\x71\x1f\xaa\xbf\x4b\x33\x5b\x9f\xf3\xb3\x8a\x1f\xbd\x78\x42\xaf\xc7\x19\x4c\x6b\x64\x0f\xa7\x7d\xa4\x49\x97\x75\xa5\xed\x61\x9e\x55\xe5\xb5\xcd\x99\x2a\xda\x8e\xe5\x01\xa2\x38\xd1\xa2\x11\xb9\xdd\xe9\xc7\x30\xdc\xd5\xd5\x89\x39\xc3\x42\xe3\x06\xa3\xa0\xa8\xae\x4f\x4f\xbf\x1d\x6f\xf6\xe7\x74\x78\xcf\x30\x9e\x6f\xf6\xda\xf9\x3b\x00\x00\xff\xff\x1d\x83\x71\xad\xc2\x04\x00\x00"),
		},
		"/threads.tmpl": &vfsgen۰CompressedFileInfo{
			name:             "threads.tmpl",
			modTime:          time.Date(2019, 11, 18, 15, 19, 7, 926790519, time.UTC),
			uncompressedSize: 783,

			compressedContent: []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x9c\x91\x4f\x6f\xd4\x30\x10\xc5\xef\xfe\x14\x26\x07\x6e\x1b\xab\x3d\x21\x70\x8c\x38\x72\x40\x42\x15\x77\x34\x38\x93\xd8\xea\xd8\x0e\x9e\x29\x25\x44\xfb\xdd\x91\x37\xbb\x50\xa8\x54\xfe\x9c\x46\xb6\xe7\xf7\x66\xfc\x9e\x7d\x36\x16\x2f\xeb\x82\x3a\x48\x22\xa7\x6c\x2b\x9a\x20\xcf\x03\x66\xa7\x6c\x42\x01\xed\x03\x54\x46\x19\xee\x64\x3a\xbc\x70\xca\x4a\x14\x42\xf7\x0e\x99\x61\x46\xfd\x21\x54\x84\x91\xad\xd9\xaf\x95\xa5\x98\x6f\x75\x45\x1a\x58\x56\x42\x0e\x88\xa2\x43\xc5\x69\xe8\xb6\x0d\x98\x51\xde\x83\x04\xdd\x99\x5c\x6a\x02\x8a\xdf\xb0\x4f\x31\xf7\x9e\xb9\x3b\x1e\xbb\x7f\xe0\xf9\x16\x09\xa5\xe4\xff\xc4\x65\xdf\xfb\xaf\xc8\x20\xb2\xf0\x4b\x63\xa6\x92\x85\xfb\xb9\x94\x99\x10\x96\xc8\xbd\x2f\xc9\x78\xe6\xd7\x13\xa4\x48\xeb\x70\x03\x84\xf7\xb0\x3e\x87\xb4\xbc\x1a\x23\x2f\x04\xeb\xc0\xf7\xb0\x34\xf5\x04\x31\x6b\x4f\xc0\x3c\xf8\x92\x05\x62\xc6\xda\x0c\xbf\x7a\xec\x64\xb8\x72\xca\x32\x7a\x89\xa5\x85\x10\xae\xdd\xdb\xfc\xa9\x7c\xb5\x26\x5c\x3b\xb5\x6d\x82\x69\x21\x10\xd4\xdd\xf9\x0f\x1f\x29\xb2\x74\xba\x3f\x75\x1d\x8f\xca\x9a\x9f\xf0\x2f\x32\x97\x49\x37\xf8\xf9\x0e\x59\xf8\x8f\x8a\x97\xc6\xa7\x44\xdf\x54\x1f\xe2\x17\x1c\x1f\xec\xff\xb4\xe8\x05\xf8\x4d\xd4\x34\x87\x9c\x52\xdb\x76\xd0\x23\x4e\x31\x3f\x22\x0f\x8d\x28\x74\x76\xf1\xfc\xe6\x4e\x40\x85\x3c\xa3\xee\x5b\x03\x45\x67\xe1\x47\xe4\x7d\xcb\xf6\x54\xac\x01\x67\x0d\xc5\x1d\xc0\xbc\xcf\x2f\xf4\xe0\xfc\x3d\x00\x00\xff\xff\xad\xf4\xcf\xf9\x0f\x03\x00\x00"),
		},
	}
	fs["/"].(*vfsgen۰DirInfo).entries = []os.FileInfo{
		fs["/error.tmpl"].(os.FileInfo),
		fs["/message.tmpl"].(os.FileInfo),
		fs["/threads.tmpl"].(os.FileInfo),
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
	case *vfsgen۰CompressedFileInfo:
		gr, err := gzip.NewReader(bytes.NewReader(f.compressedContent))
		if err != nil {
			// This should never happen because we generate the gzip bytes such that they are always valid.
			panic("unexpected error reading own gzip compressed bytes: " + err.Error())
		}
		return &vfsgen۰CompressedFile{
			vfsgen۰CompressedFileInfo: f,
			gr:                        gr,
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

// vfsgen۰CompressedFileInfo is a static definition of a gzip compressed file.
type vfsgen۰CompressedFileInfo struct {
	name              string
	modTime           time.Time
	compressedContent []byte
	uncompressedSize  int64
}

func (f *vfsgen۰CompressedFileInfo) Readdir(count int) ([]os.FileInfo, error) {
	return nil, fmt.Errorf("cannot Readdir from file %s", f.name)
}
func (f *vfsgen۰CompressedFileInfo) Stat() (os.FileInfo, error) { return f, nil }

func (f *vfsgen۰CompressedFileInfo) GzipBytes() []byte {
	return f.compressedContent
}

func (f *vfsgen۰CompressedFileInfo) Name() string       { return f.name }
func (f *vfsgen۰CompressedFileInfo) Size() int64        { return f.uncompressedSize }
func (f *vfsgen۰CompressedFileInfo) Mode() os.FileMode  { return 0444 }
func (f *vfsgen۰CompressedFileInfo) ModTime() time.Time { return f.modTime }
func (f *vfsgen۰CompressedFileInfo) IsDir() bool        { return false }
func (f *vfsgen۰CompressedFileInfo) Sys() interface{}   { return nil }

// vfsgen۰CompressedFile is an opened compressedFile instance.
type vfsgen۰CompressedFile struct {
	*vfsgen۰CompressedFileInfo
	gr      *gzip.Reader
	grPos   int64 // Actual gr uncompressed position.
	seekPos int64 // Seek uncompressed position.
}

func (f *vfsgen۰CompressedFile) Read(p []byte) (n int, err error) {
	if f.grPos > f.seekPos {
		// Rewind to beginning.
		err = f.gr.Reset(bytes.NewReader(f.compressedContent))
		if err != nil {
			return 0, err
		}
		f.grPos = 0
	}
	if f.grPos < f.seekPos {
		// Fast-forward.
		_, err = io.CopyN(ioutil.Discard, f.gr, f.seekPos-f.grPos)
		if err != nil {
			return 0, err
		}
		f.grPos = f.seekPos
	}
	n, err = f.gr.Read(p)
	f.grPos += int64(n)
	f.seekPos = f.grPos
	return n, err
}
func (f *vfsgen۰CompressedFile) Seek(offset int64, whence int) (int64, error) {
	switch whence {
	case io.SeekStart:
		f.seekPos = 0 + offset
	case io.SeekCurrent:
		f.seekPos += offset
	case io.SeekEnd:
		f.seekPos = f.uncompressedSize + offset
	default:
		panic(fmt.Errorf("invalid whence value: %v", whence))
	}
	return f.seekPos, nil
}
func (f *vfsgen۰CompressedFile) Close() error {
	return f.gr.Close()
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
