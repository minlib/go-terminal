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
	fw, err := os.Create(tarPath)
	if err != nil {
		panic(err)
	}
	defer fw.Close()
	// gzip write
	gw := gzip.NewWriter(fw)
	defer gw.Close()
	// tar write
	tw := tar.NewWriter(gw)
	defer tw.Close()
	// traverse the file or directory
	for _, srcPath := range paths {
		// remove the trailing path separator if path is a directory
		srcPath = strings.TrimSuffix(srcPath, string(os.PathSeparator))
		// visit all the files or directories in the tree
		err = filepath.Walk(srcPath, func(path string, fi fs.FileInfo, err error) error {
			if err != nil {
				return err
			}
			if fi.IsDir() {
				return nil
			}
			name, err := filepath.Rel(filepath.Dir(srcPath), path)
			if err != nil {
				return err
			}
			name = strings.ReplaceAll(name, string(os.PathSeparator), "/")
			// fmt.Println(name)
			// 信息头
			th := new(tar.Header)
			th.Name = name
			th.Format = tar.FormatGNU
			th.Size = fi.Size()
			th.Mode = int64(fi.Mode())
			th.ModTime = fi.ModTime()
			err = tw.WriteHeader(th)
			if err != nil {
				panic(err)
			}
			f, err := os.Open(path)
			if err != nil {
				return err
			}
			defer f.Close()
			// 写文件
			_, err = io.Copy(tw, f)
			if err != nil {
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

// UnTar decompresses a tar.gz file to specified directory.
func UnTar(tarFile, destDir string) error {
	if !strings.HasSuffix(destDir, string(os.PathSeparator)) {
		destDir += string(os.PathSeparator)
	}
	srcFile, err := os.Open(tarFile)
	if err != nil {
		return err
	}
	defer srcFile.Close()
	gr, err := gzip.NewReader(srcFile)
	if err != nil {
		return err
	}
	defer gr.Close()
	tr := tar.NewReader(gr)
	for {
		hdr, err := tr.Next()
		if err != nil {
			if err == io.EOF {
				break
			} else {
				return err
			}
		}
		filename := destDir + hdr.Name
		if err := os.MkdirAll(filepath.Dir(filename), os.ModePerm); err != nil {
			return err
		}
		file, err := os.Create(filename)
		if err != nil {
			return err
		}
		io.Copy(file, tr)
	}
	return nil
}
