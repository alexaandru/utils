package utils

import (
	"crypto/md5"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
)

// FileHashes holds a map between file paths and their hashes
type FileHashes map[string]string

// Dump writes file hashes to disk.
func (fh FileHashes) Dump(fname string) (e error) {
	if out, err := os.Create(fname); err == nil {
		defer out.Close()
		for k, v := range fh {
			if _, err2 := out.Write([]byte(k + ":" + v + "\n")); err2 != nil {
				e = err2
				break
			}
		}
	} else {
		e = err
	}

	return
}

// Load reads file hashes from disk.
func (fh FileHashes) Load(fname string) {
	for _, line := range strings.Split(LoadFile(fname), "\n") {
		if line == "" {
			continue
		}

		kv := strings.Split(line, ":")
		k, v := kv[0], kv[1]
		fh[k] = v
	}
}

// Diff computes the difference between two file hashes/maps.
func (fh FileHashes) Diff(other FileHashes) (diff []string) {
	diff = []string{}
	for k, v := range fh {
		if v != other[k] {
			diff = append(diff, k)
		}
	}

	return
}

// Filter keeps only those FileHashes that are in fnames list.
func (fh FileHashes) Filter(fnames []string) (out FileHashes) {
	out = FileHashes{}

	for _, fname := range fnames {
		if hash, ok := fh[fname]; ok {
			out[fname] = hash
		}
	}

	return
}

// Reject returns all the FileHashes that are NOT in the fnames list.
func (fh FileHashes) Reject(fnames []string) (out FileHashes) {
	out = FileHashes{}
	rejects := map[string]struct{}{}
	for _, fname := range fnames {
		rejects[fname] = struct{}{}
	}

	for fname, hash := range fh {
		if _, rejected := rejects[fname]; !rejected {
			out[fname] = hash
		}
	}

	return
}

// FileHashesNew walks all the folders/files starting at "root".
func FileHashesNew(root string) (hashes FileHashes) {
	if cwd, err := os.Getwd(); err == nil {
		os.Chdir(root)
		hashes = FileHashes{}
		filepath.Walk(".", func(path string, f os.FileInfo, err error) (e error) {
			if !f.IsDir() {
				hashes[path] = fileHash(path)
			}

			return
		})
		os.Chdir(cwd)
	} else {
		panic(err)
	}

	return
}

// Compute an md5 hash of a file.
func fileHash(fname string) string {
	hash := md5.New()
	io.WriteString(hash, LoadFile(fname))
	return fmt.Sprintf("%x", hash.Sum(nil))
}
