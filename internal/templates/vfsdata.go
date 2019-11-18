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
			modTime: time.Date(2019, 11, 18, 15, 17, 51, 787320794, time.UTC),
		},
		"/error.tmpl": &vfsgen۰CompressedFileInfo{
			name:             "error.tmpl",
			modTime:          time.Date(2019, 11, 17, 3, 46, 24, 770694000, time.UTC),
			uncompressedSize: 563,

			compressedContent: []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x9c\x91\xbf\x4e\xc4\x30\x0c\x87\xf7\x3e\x45\xe8\xc0\xc4\xb5\xba\x0d\x41\x52\x06\x66\x24\x04\xbc\x80\x95\xba\x4d\x74\xce\x1f\xc5\xbe\x2b\xa5\xea\xbb\xa3\xde\x89\x81\x8d\x63\x8a\xf2\x59\x9f\x2d\xfb\xa7\x6f\xfa\x64\x65\xce\xa8\x9c\x04\xea\x2a\xbd\x3d\x8a\x20\x8e\x06\x63\x57\xe9\x80\x02\xca\x3a\x28\x8c\x62\x8e\x32\xec\xee\x7f\x60\x84\x80\xe6\xe4\x71\xca\xa9\x88\xb2\x29\x0a\x46\x31\xf5\xe4\x7b\x71\xa6\xc7\x93\xb7\xb8\x3b\x7f\xee\x7c\xf4\xe2\x81\x76\x6c\x81\xd0\xec\xeb\xae\xd2\xe2\x85\xb0\x5b\x96\xe6\x5d\x40\x8e\xfc\x9c\x7a\x5c\x57\xb5\x2c\x4e\x24\x5f\xd0\x07\x7e\x8a\xfa\x55\xd6\xed\xc5\xaa\x34\xf9\x78\x50\x05\xc9\xb0\xcc\x84\xec\x10\x45\xb9\x82\x83\xa9\x97\x05\x98\x51\x5e\x41\x9c\xaa\xdb\x98\x4a\x00\xf2\x5f\xd8\x04\x1f\x1b\xcb\x5c\xaf\x6b\x7d\x85\xcf\x07\x24\x94\x14\xff\xa9\x63\x29\xa9\xfc\xc9\xdb\xd6\xe6\x87\xb6\x1d\x52\x14\x6e\xc6\x94\x46\x42\xc8\x9e\x1b\x9b\x42\x6b\x99\x9f\x06\x08\x9e\x66\xf3\x06\x84\x13\xcc\xb7\x10\xf2\x63\xef\x39\x13\xcc\x86\x27\xc8\x5b\xf7\x00\x3e\x2a\x4b\xc0\x6c\xb6\x30\xc0\x47\x2c\x5b\x9e\xfb\xab\xcf\xec\xf6\x5d\xa5\xf3\xa6\xbd\x20\x33\x8c\x67\x98\xbb\x4a\xb7\xdb\x8c\xee\x3b\x00\x00\xff\xff\xc0\x80\xef\x61\x33\x02\x00\x00"),
		},
		"/message.tmpl": &vfsgen۰CompressedFileInfo{
			name:             "message.tmpl",
			modTime:          time.Date(2019, 11, 18, 15, 24, 25, 667750733, time.UTC),
			uncompressedSize: 982,

			compressedContent: []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x9c\x92\xcf\x8a\xd4\x40\x10\x87\xef\x79\x8a\x76\x0e\xa2\x07\x3b\xec\x9e\x44\x93\x16\x5d\x58\xf1\xb0\x30\xb8\xe2\xbd\xe8\xae\x99\x2e\xb6\xff\x84\x54\x0d\xc3\xd8\xe4\xdd\xa5\x27\xae\x4e\x26\x17\xf5\x94\xf0\xfd\xbe\xaa\x50\x95\xea\x5e\xb8\x6c\xe5\x34\xa0\xf2\x12\x83\x69\xba\xfa\x50\x01\xd2\xbe\xc7\x64\x9a\x2e\xa2\x80\xb2\x1e\x46\x46\xe9\x0f\xb2\x7b\xf3\xd6\x34\x9d\x90\x04\x34\xa5\xe8\x6f\xf5\x65\x9a\xba\x76\x26\x4d\x17\x28\x3d\xa9\x11\x43\xcf\x72\x0a\xc8\x1e\x51\x94\x1f\x71\xd7\x6f\x4a\x01\x66\x94\x2d\x88\x57\x9b\x36\xe5\x31\x42\xa0\x1f\xa8\x23\x25\x6d\x99\x37\xd3\xb4\xf9\x87\x7a\x7e\xc2\x80\x92\xd3\x7f\x96\x47\x64\x86\x3d\xfe\x55\xa5\x17\x19\xf8\x5d\xdb\xee\x72\x12\xd6\xfb\x9c\xf7\x01\x61\x20\xd6\x36\xc7\xd6\x32\x7f\xd8\x41\xa4\x70\xea\xbf\x42\xc0\x23\x9c\x5e\x42\x1c\xde\x3b\xe2\x21\xc0\xa9\xe7\x23\x0c\xb5\x7b\x04\x4a\xca\x06\x60\xee\x6d\x4e\x02\x94\x70\xac\xbb\xbe\x59\x2c\xd1\xdf\x98\xa6\x63\xb4\x42\xb9\xae\xde\xdf\x9a\x4f\xc0\x64\xd5\xa3\x80\x70\xd7\xfa\x5b\xd3\x74\xae\xfe\x23\x27\x66\x0b\xa3\x90\xa5\x01\x52\x8d\x9c\x54\xea\x4c\x29\x01\x93\xd2\x97\x61\x6d\xec\xdc\x5c\xf4\x30\x8f\xbd\x2e\x78\x0e\x2e\xe5\xad\xcf\x92\x17\xaa\x9e\xd1\x5d\x3e\x24\xb9\x34\xbf\x93\xc3\x2b\x73\x46\x2b\xf3\xe3\xc1\x51\x5e\x88\x67\xb2\xf2\x3e\x7f\xb9\x5f\xf6\xab\x60\x65\xdd\x53\x58\xce\xa2\xcf\x64\xe5\x3d\x0a\xd9\x27\x1c\x97\xea\x33\x5c\xd9\xdb\x00\xe9\x6a\xec\x4a\xd6\x5d\x3d\x8c\x57\x9f\x9f\xd1\xca\xbc\x83\x10\x96\xe2\x99\xfc\xf2\x4a\x39\x92\x78\xa5\x1f\x88\x19\xdd\x65\xa2\x5e\x95\xa2\xa7\x49\xc5\x73\xa2\x6c\x8d\x5e\x97\x82\xc9\xfd\xee\xde\x9e\xef\xa1\xfd\x73\x34\x6d\x3d\x35\xf3\x33\x00\x00\xff\xff\x77\x4f\x59\xe1\xd6\x03\x00\x00"),
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
