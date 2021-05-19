package config

import (
	"fmt"
	"io/ioutil"
	"os"
)

type FileStore struct {
	path string
}

func NewFileStore(path string) (*FileStore, error) {
	fs := &FileStore{
		path: path,
	}
	return fs, nil
}

func (fs *FileStore) Load() ([]byte, error) {
	f, err := os.Open(fs.path)
	if os.IsNotExist(err) {
		return nil, nil
	} else if err != nil {
		return nil, fmt.Errorf("读取配置文件(%s)失败 err: %w", fs.path, err)
	}
	defer f.Close()
	fileBytes, err := ioutil.ReadAll(f)
	if err != nil {
		return nil, err
	}
	return fileBytes, nil
}

func (fs *FileStore) Close() error {
	return nil
}
