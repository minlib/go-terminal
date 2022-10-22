package tar

import (
	"archive/tar"
	"compress/gzip"
	"io"
	"io/fs"
	"os"
	"path/filepath"
	"strings"
)

// Tar compresses the specified files or dirs to tar.gz archive.
func Tar(tarPath string, paths ...string) error {
	// file write
	file, err := os.Create(tarPath)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	// gzip write
	gzipWriter := gzip.NewWriter(file)
	defer gzipWriter.Close()
	// tar write
	tarWrite := tar.NewWriter(gzipWriter)
	defer tarWrite.Close()
	// traverse the file or directory
	for _, srcPath := range paths {
		// remove the trailing path separator if path is a directory
		srcPath = strings.TrimSuffix(srcPath, string(os.PathSeparator))
		// visit all the files or directories in the tree
		err = filepath.Walk(srcPath, func(path string, fileInfo fs.FileInfo, err error) error {
			if err != nil {
				return err
			}
			if fileInfo.IsDir() {
				return nil
			}
			name, err := filepath.Rel(filepath.Dir(srcPath), path)
			if err != nil {
				return err
			}
			name = strings.ReplaceAll(name, string(os.PathSeparator), "/")
			header := new(tar.Header)
			header.Name = name
			header.Format = tar.FormatGNU
			header.Size = fileInfo.Size()
			header.Mode = int64(fileInfo.Mode())
			header.ModTime = fileInfo.ModTime()
			if err = tarWrite.WriteHeader(header); err != nil {
				panic(err)
			}
			f, err := os.Open(path)
			if err != nil {
				return err
			}
			defer f.Close()
			if _, err = io.Copy(tarWrite, f); err != nil {
				panic(err)
			}
			return err
		})
		if err != nil {
			return err
		}
	}
	return nil
}

// Untar decompresses a tar.gz file to specified directory.
func Untar(tarFile, destDir string) error {
	if !strings.HasSuffix(destDir, string(os.PathSeparator)) {
		destDir += string(os.PathSeparator)
	}
	srcFile, err := os.Open(tarFile)
	if err != nil {
		return err
	}
	defer srcFile.Close()
	gzipReader, err := gzip.NewReader(srcFile)
	if err != nil {
		return err
	}
	defer gzipReader.Close()
	tarReader := tar.NewReader(gzipReader)
	for {
		header, err := tarReader.Next()
		if err != nil {
			if err == io.EOF {
				break
			} else {
				return err
			}
		}
		filename := destDir + header.Name
		if err := os.MkdirAll(filepath.Dir(filename), os.ModePerm); err != nil {
			return err
		}
		file, err := os.Create(filename)
		if err != nil {
			return err
		}
		io.Copy(file, tarReader)
	}
	return nil
}
